package global

import (
	"fmt"
	"tokenagent/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	MARKET_CONFIG config.ServerConfig
	MARKET_VP     *viper.Viper
	MARKET_LOG    *zap.Logger
)

var (
	JsonRPCVersion = "2.0"

	ParseErrorCode          = -32700 //Parse error语法解析错误	服务端接收到无效的json。该错误发送于服务器尝试解析json文本
	InvalidRequestErrorCode = -32600 //Invalid Request无效请求	发送的json不是一个有效的请求对象。
	MethodNotFoundErrorCode = -32601 //Method not found找不到方法	该方法不存在或无效
	InvalidParamsErrorCode  = -32602 //Invalid params无效的参数	无效的方法参数。
	InternalErrorCode       = -32603 //Internal error内部错误	JSON-RPC内部错误。
)

type RPCError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewRPCError(code int, message string) *RPCError {
	return &RPCError{
		Code:    code,
		Message: message,
	}
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
