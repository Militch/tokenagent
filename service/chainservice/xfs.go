package chainservice

import "tokenagent/model/model"

type XFSCommand struct {
}

func (command *XFSCommand) SendRawTransaction(args model.TxRequest) (*model.TxResponse, error) {
	return nil, nil
}

func (command *XFSCommand) GetTransactionReceipt(args model.TxRecRequest) (*model.Receipt, error) {

	return nil, nil
}
