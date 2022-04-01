package chainservice

import (
	"context"
	"fmt"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils/chaincli"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"go.uber.org/zap"
)

type ETHCommand struct {
}

//@author: [libofeng]
//@function: SendRawTransaction
//@description: 以太坊发送签名的交易上链
//@param: args model.TxRequest
//@return: resp **model.TxResponse
func (command *ETHCommand) SendRawTransaction(args model.TxRequest) (*model.TxResponse, error) {
	var tx types.Transaction
	rawTxBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		global.MARKET_LOG.Debug("SendRawTransaction->DecodeString(): ", zap.Error(err))
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("%v of param `data`", err))
	}

	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		global.MARKET_LOG.Debug("SendRawTransaction->DecodeBytes(): ", zap.Error(err))
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("%v of param `data`", err))
	}

	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	chinID, err := cli.ChainID(context.Background())
	if err != nil {
		global.MARKET_LOG.Debug("SendRawTransaction->ChainID(): ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v blockchain:`%v`", err.Error(), args.BlockChain))
	}

	signer := types.NewEIP155Signer(chinID)
	sender, err := types.Sender(signer, &tx)
	if err != nil {
		global.MARKET_LOG.Debug("SendRawTransaction->types.Sender: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v in param `data`", err))
	}

	var contractAddress common.Address
	if tx.To() == nil {
		contractAddress = crypto.CreateAddress(sender, tx.Nonce())
	} else {
		contractAddress = common.Address{}
	}

	response := &model.TxResponse{
		TxHash:          tx.Hash().String(),
		ContractAddress: contractAddress.Hex(),
	}
	err = cli.RpcCli.CallContext(context.Background(), nil, "eth_sendRawTransaction", args.Data)
	if err != nil {
		global.MARKET_LOG.Debug("SendRawTransaction: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v blockchain:`%v`", err.Error(), args.BlockChain))
	}

	return response, nil
}

//@author: [libofeng]
//@function: SendRawTransaction
//@description: 以太坊发送签名的交易上链
//@param: args model.TxRequest
//@return: resp **model.TxResponse

func (command *ETHCommand) GetTransactionReceipt(args model.TxRecRequest) (*model.Receipt, error) {
	var rec *types.Receipt
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	err := cli.RpcCli.CallContext(context.Background(), &rec, "eth_getTransactionReceipt", args.TxHash)
	if err != nil {
		global.MARKET_LOG.Debug("eth_getTransactionReceipt: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v blockchain:`%v`", err.Error(), args.BlockChain))
	}

	if rec == nil {
		return nil, err
	}
	response := &model.Receipt{
		Type:              rec.Type,
		Status:            rec.Status,
		CumulativeGasUsed: rec.CumulativeGasUsed,
		TxHash:            rec.TxHash.Hex(),
		ContractAddress:   rec.ContractAddress.Hex(),
		GasUsed:           rec.GasUsed,
		BlockHash:         rec.BlockHash.Hex(),
		BlockNumber:       rec.BlockNumber,
		TransactionIndex:  rec.TransactionIndex,
	}
	return response, nil
}
