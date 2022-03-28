package chainapi

import (
	"tokenagent/model/model"
	"tokenagent/service"
)

type ChainAPIHandler struct {
}

//@author: [libofeng]
//@function: SendRawTransaction
//@description: 发送交易上链,并且这笔交易是经过签名的
//@param: args model.TxRequest
//@return: resp **model.TxResponse

func (handler *ChainAPIHandler) SendRawTransaction(args model.TxRequest, resp **model.TxResponse) error {
	response, err := service.ServiceGroupApp.ChainServiceProxy.SendRawTransaction(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: GetTransactionReceipt
//@description: 获取交易回执
//@param: args model.TxRecRequest
//@return: resp **model.Receipt

func (handler *ChainAPIHandler) GetTransactionReceipt(args model.TxRecRequest, resp **model.Receipt) error {
	response, err := service.ServiceGroupApp.ChainServiceProxy.GetTransactionReceipt(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}
