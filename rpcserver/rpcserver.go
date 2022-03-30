package rpcserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/token"
	"io"
	"io/ioutil"
	"net"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"tokenagent/global"

	"github.com/gorilla/websocket"
)

var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint
}

type jsonRPCObj struct {
	jsonrpc string
	id      *int
	method  string
	params  interface{}
}
type jsonRPCRespErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type service struct {
	name    string
	rcvr    reflect.Value
	typ     reflect.Type
	methods map[string]*methodType
}

type RPCServer struct {
	ginEngine  *gin.Engine
	upgrader   websocket.Upgrader
	serviceMap map[string]*service
}

func NewRPCServer(router *gin.Engine) *RPCServer {

	rpcserver := &RPCServer{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		serviceMap: make(map[string]*service),
	}

	rpcserver.ginEngine = router
	return rpcserver
}

//Start starts rpc server.
func (server *RPCServer) Start() error {
	server.ginEngine.Any("/", func(c *gin.Context) {
		//handle websocket request
		// if isWebsocketRequest(c) {
		// 	if err := server.handleWebsocket(c); err != nil {
		// 		global.MARKET_LOG.Warn("ws connect err")
		// 	}
		// 	c.Abort()
		// 	return
		// }
		if "POST" != c.Request.Method {
			httperr(c, 404, errors.New("method not allowed"))
			return
		}
		contentType := c.ContentType()
		if contentType != "application/json" {
			httperr(c, 404, errors.New("not acceptable"))
			return
		}
		if nil == c.Request.Body {
			httperr(c, 404, errors.New("body not be empty"))
			return
		}
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			httperr(c, 500, fmt.Errorf("read body err: %s", err))
			return
		}
		c.Status(200)
		c.Header("Content-Type", "application/json; charset=utf-8")
		var rpcId *int = nil

		if err = server.jsonRPCCall(body, &rpcId, c.Writer); err != nil {
			writeRPCError(err, rpcId, c.Writer)
			return
		}
		c.Abort()
	})

	ln, err := net.Listen("tcp", global.MARKET_CONFIG.Rpcserver.Listen)
	if err != nil {
		return err
	}
	global.MARKET_LOG.Info("RPC Service listen on: %s", zap.String("address", ln.Addr().String()))
	return server.ginEngine.RunListener(ln)
}

func (server *RPCServer) RegisterName(name string, rcvr interface{}) error {
	return server.register(rcvr, name, true)
}

func writeRPCError(err error, reqId *int, w io.Writer) {
	// rpcErr, isRPCErr := err.(*global.RPCError)
	// e := jsonRPCRespErr{}
	// if !isRPCErr {
	// 	e.Code = -32603
	// 	e.Message = rpcErr.Message
	// } else {
	// 	e.Code = rpcErr.Code
	// 	e.Message = rpcErr.Message
	// }

	outMap := make(map[string]interface{})
	outMap["jsonrpc"] = global.JsonRPCVersion
	outMap["id"] = reqId
	outMap["error"] = err
	outBytes, _ := json.Marshal(outMap)
	_, _ = w.Write(outBytes)
}

func (server *RPCServer) jsonRPCCall(data []byte, rpcId **int, w io.Writer) error {
	var personFromJSON interface{}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	_ = decoder.Decode(&personFromJSON)

	jsonObjMap := personFromJSON.(map[string]interface{})
	rpcObj := &jsonRPCObj{}
	if err := server.parseJsonRPCObj(jsonObjMap, rpcObj); err != nil {
		*rpcId = *&rpcObj.id
		global.MARKET_LOG.Debug(rpcObj.method, zap.Error(err))
		return err
	}
	global.MARKET_LOG.Debug(rpcObj.method, zap.Any("Request-->Body", jsonObjMap))
	*rpcId = *&rpcObj.id
	s, t, err := server.getServiceAndMethodType(rpcObj.method)
	if err != nil {
		global.MARKET_LOG.Debug(rpcObj.method, zap.Error(err))
		return err
	}

	_, ok := rpcObj.params.(map[string]interface{})
	if ok {
		if len(rpcObj.params.(map[string]interface{})) == 0 {
			rpcObj.params = nil
		}
	}

	rec, err := s.callMethod(t, rpcObj.params)
	if err != nil {
		global.MARKET_LOG.Error(rpcObj.method, zap.Error(err))
		return err
	}
	outMap := make(map[string]interface{})
	outMap["jsonrpc"] = global.JsonRPCVersion
	outMap["id"] = rpcObj.id
	outMap["result"] = rec

	global.MARKET_LOG.Debug(rpcObj.method, zap.Any("Response-->Body", outMap))
	outBytes, _ := json.Marshal(outMap)
	_, _ = w.Write(outBytes)
	return nil
}

func (server *RPCServer) parseJsonRPCObj(jsonObjMap map[string]interface{}, obj *jsonRPCObj) error {
	idNumber, ok := jsonObjMap["id"].(json.Number)
	if !ok {
		return global.NewRPCError(global.InvalidRequestErrorCode, "jsonrpc id is invalid request")
	}
	idNumberStr := idNumber.String()
	id, err := strconv.Atoi(idNumberStr)
	if err != nil {
		return global.NewRPCError(-32600, err.Error())
	}

	obj.id = &id
	version, ok := jsonObjMap["jsonrpc"].(string)
	if !ok || version != "2.0" {
		return global.NewRPCError(global.InvalidRequestErrorCode, "jsonrpc version is invalid request")
	}
	obj.jsonrpc = version
	methodPack, ok := jsonObjMap["method"].(string)
	if !ok {
		return global.NewRPCError(global.InvalidRequestErrorCode, "jsonrpc method is invalid request")
	}
	obj.method = methodPack
	obj.params = jsonObjMap["params"]
	return nil
}

func (s *service) callMethod(mtype *methodType, params interface{}) (interface{}, error) {
	function := mtype.method.Func
	argIsValue := false
	var argv reflect.Value
	if mtype.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(mtype.ArgType.Elem())
	} else {
		argv = reflect.New(mtype.ArgType)
		argIsValue = true
	}

	if argIsValue {
		argv = argv.Elem()
	}
	if params != nil {
		tk := reflect.TypeOf(params).Kind()
		switch tk {
		case reflect.Slice:
			paramsArr, _ := params.([]interface{})
			if len(paramsArr) != argv.NumField() {
				return nil, global.NewRPCError(-32602, "Invalid params")
			}
			for i := 0; i < argv.NumField(); i++ {
				argv.Field(i).Set(reflect.ValueOf(paramsArr[i]))
			}
		case reflect.Map:
			paramsMap := params.(map[string]interface{})
			for i := 0; i < argv.NumField(); i++ {
				fieldInfo := argv.Type().Field(i) // a reflect.StructField
				tag := fieldInfo.Tag              // a reflect.StructTag
				name := tag.Get("json")
				if name == "" {
					name = strings.ToLower(fieldInfo.Name)
				}
				name = strings.Split(name, ",")[0] //json Tag

				// Support basic types
				if value, ok := paramsMap[name]; ok {
					if reflect.ValueOf(value).Type() == argv.FieldByName(fieldInfo.Name).Type() {
						argv.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
					} else {
						if val, ok := reflect.ValueOf(value).Interface().(json.Number); ok {
							value, err := val.Int64()
							if err != nil {
								return nil, err
							}
							data := int(value)
							if argv.FieldByName(fieldInfo.Name).Type() == reflect.TypeOf(data) {
								argv.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(data))
							}
						}
					}
				}
			}
		}
	}
	replyv := reflect.New(mtype.ReplyType.Elem())
	switch mtype.ReplyType.Elem().Kind() {
	case reflect.Map:
		replyv.Elem().Set(reflect.MakeMap(mtype.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(mtype.ReplyType.Elem(), 0, 0))
	}
	returnValues := function.Call([]reflect.Value{s.rcvr, argv, replyv})
	errInter := returnValues[0].Interface()
	if errInter != nil {
		e := errInter.(error)
		return nil, e
	}
	return replyv.Interface(), nil
}

func isWebsocketRequest(c *gin.Context) bool {
	connection := c.GetHeader("Connection")
	upgrade := c.GetHeader("Upgrade")
	return connection == "Upgrade" && upgrade == "websocket"
}

func (server *RPCServer) getServiceAndMethodType(pack string) (*service, *methodType, error) {
	mpake := strings.Split(pack, ".")
	if len(mpake) != 2 {
		return nil, nil, global.NewRPCError(global.MethodNotFoundErrorCode, "method not found")
	}
	mService := server.serviceMap[mpake[0]]
	if mService == nil {
		return nil, nil, global.NewRPCError(global.MethodNotFoundErrorCode, "method not found")
	}
	mtype := mService.methods[mpake[1]]
	if mtype == nil {
		return nil, nil, global.NewRPCError(global.MethodNotFoundErrorCode, "method not found")
	}
	return mService, mtype, nil
}

func (server *RPCServer) handleWebsocket(c *gin.Context) error {
	conn, err := server.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			continue
		}
		//msgText := string(msg)

		bs := bytes.NewBuffer(nil)
		var rpcId *int
		if err = server.jsonRPCCall(msg, &rpcId, bs); err != nil {
			writeRPCError(err, nil, bs)
		}
		if err = conn.WriteMessage(t, bs.Bytes()); err != nil {
			continue
		}
	}
	return nil
}

func (server *RPCServer) register(rcvr interface{}, name string, useName bool) error {
	s := new(service)
	s.typ = reflect.TypeOf(rcvr)
	s.rcvr = reflect.ValueOf(rcvr)
	sname := reflect.Indirect(s.rcvr).Type().Name()
	if useName {
		sname = name
	}
	if sname == "" {
		return fmt.Errorf("rpc.Register: no service name for type %s", s.typ.String())
	}
	if !token.IsExported(sname) && !useName {
		return fmt.Errorf("rpc.Register: type %s is not exported", sname)
	}
	s.name = sname
	s.methods = suitableMethods(s.typ)
	server.serviceMap[sname] = s

	return nil
}

func suitableMethods(typ reflect.Type) map[string]*methodType {
	methods := make(map[string]*methodType)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := method.Name
		if method.PkgPath != "" {
			continue
		}
		if mtype.NumIn() != 3 {
			continue
		}
		argType := mtype.In(1)
		if !isExportedOrBuiltinType(argType) {
			continue
		}
		replyType := mtype.In(2)
		if replyType.Kind() != reflect.Ptr {
			continue
		}
		if mtype.NumOut() != 1 {
			continue
		}
		if returnType := mtype.Out(0); returnType != typeOfError {
			continue
		}
		methods[mname] = &methodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
	}
	return methods
}

func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return token.IsExported(t.Name()) || t.PkgPath() == ""
}

func httperr(c *gin.Context, status int, err error) {
	c.String(status, "%s", err)
	c.Abort()
}
