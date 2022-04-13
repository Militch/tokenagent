package tokenservice

import (
	"fmt"
	"math/big"
	"strings"
	"tokenagent/contracts/erc20"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils/chaincli"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
)

type ETHToken struct {
}

//@author: [libofeng]
//@function: TokenPreCreate
//@description: 生成合集的的Code码,同时通过TxPreExtra获取Nonce|gasprice|gaslimit,等信息发给前端进行签名，在通过SendRawTransaction上链
//@param: args ethmodel.TokenPreCreateRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreCreate(args model.TokenPreCreateRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	packed, err := abi.Pack("", args.Name, args.Symbol)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	bin, _ := hexutil.Decode(erc20.Erc20MetaData.Bin)
	data := append(bin, packed...)

	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		Data: data,
	}

	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}

	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(data),
	}

	return response, nil
}

//@author: [libofeng]
//@function: Name
//@description: Call调用返回代币名称，不上链不用进行签名
//@param: args ethmodel.TokenRequest
//@return: 代币名称 string
func (ethToken *ETHToken) Name(args model.TokenCallRequest) (string, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	token, err := erc20.NewErc20(contractAddress, cli)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts:", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	name, err := token.Name(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to Name:", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return name, nil
}

//@author: [libofeng]
//@function: Symbol
//@description: Call调用返回代币符号，不上链不用进行签名
//@param: args ethmodel.TokenRequest
//@return: 代币符号 string
func (ethToken *ETHToken) Symbol(args model.TokenCallRequest) (string, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	token, err := erc20.NewErc20(contractAddress, cli)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	symbol, err := token.Symbol(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to Symbol: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return symbol, nil
}

//@author: [libofeng]
//@function: Owner
//@description: Call调用返回NFTToken的总量，不上链不用进行签名
//@param: args ethmodel.TokenRequest
//@return: 代币总量
func (ethToken *ETHToken) TotalSupply(args model.TokenCallRequest) (*big.Int, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return *new(*big.Int), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	token, err := erc20.NewErc20(contractAddress, cli)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewNft721: %v", zap.Error(err))
		return *new(*big.Int), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	totalsupply, err := token.TotalSupply(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to TotalSupply: %v", zap.Error(err))
		return *new(*big.Int), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return totalsupply, nil
}

//@author: [libofeng]
//@function: BalanceOf
//@description: Call调用返回指定账户Token的数量，不上链不用进行签名
//@param: args ethmodel.TokenBalanceOfCallRequest
//@return: 代币总量
func (ethToken *ETHToken) BalanceOf(args model.TokenBalanceOfCallRequest) (*big.Int, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return *new(*big.Int), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	token, err := erc20.NewErc20(contractAddress, cli)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewNft721: %v", zap.Error(err))
		return *new(*big.Int), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	totalsupply, err := token.BalanceOf(nil, common.HexToAddress(args.Account))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to TotalSupply: %v", zap.Error(err))
		return *new(*big.Int), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return totalsupply, nil
}

//@author: [libofeng]
//@function: TokenPreMintCreate
//@description: 组织发行代币的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreMintRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreMintCreate(args model.TokenPreMintRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	tokenId, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}

	packed, err := abi.Pack("mint", common.HexToAddress(args.ToAddress), tokenId)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	contractAddress := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &contractAddress,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: TokenPreApproveCreate
//@description: 组织铸造Token的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreApproveRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreApproveCreate(args model.TokenPreApproveRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("approve", common.HexToAddress(args.SpenderAddress), amount)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: TokenPreBurnFromCreate
//@description: 组织销毁Token的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreBurnRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreBurnFromCreate(args model.TokenPreBurnRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("burn", amount)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: TokenPreIncreaseAllowance
//@description: 组织转移Token交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreAllowanceRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreIncreaseAllowance(args model.TokenPreAllowanceRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	value, ok := new(big.Int).SetString(args.Value, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("increaseAllowance", common.HexToAddress(args.SpenderAddress), value)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: TokenPreDecreaseAllowance
//@description: 组织转移Token交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreAllowanceRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreDecreaseAllowance(args model.TokenPreAllowanceRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	value, ok := new(big.Int).SetString(args.Value, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("decreaseAllowance", common.HexToAddress(args.SpenderAddress), value)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: TokenPreSafeTransferFromCreate
//@description: 组织转移Token交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreTransferRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreTransferFromCreate(args model.TokenPreTransferRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("transferFrom", common.HexToAddress(args.FromAddress), common.HexToAddress(args.ToAddress), amount)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: TokenPrePauseCreate
//@description: 组织暂停Token交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPrePauseCreate(args model.TokenPreRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	packed, err := abi.Pack("pause")
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: TokenPreUnPauseCreate
//@description: 组织暂停Token交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.TokenPreRequest
//@return: *ethmodel.ContractCodeResponse
func (ethToken *ETHToken) TokenPreUnPauseCreate(args model.TokenPreRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	packed, err := abi.Pack("unpause")
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ContractAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	txPre, err := cli.TxPreExtra(msg)
	if err != nil {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	response := &model.ContractCodeResponse{
		TxPreExtra: *txPre,
		Code:       hexutil.Encode(packed),
	}

	return response, nil
}
