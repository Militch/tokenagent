package collectionservice

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"tokenagent/contracts/erc721"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils/chaincli"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

type ETHCollection struct {
}

//@author: [libofeng]
//@function: CollectionPreCreate
//@description: 生成合集的的Code码,同时通过TxPreExtra获取Nonce|gasprice|gaslimit,等信息发给前端进行签名，在通过SendRawTransaction上链
//@param: args ethmodel.CollectionPreCreateRequest
//@return: *ethmodel.ContractCodeResponse
func (collection *ETHCollection) CollectionPreCreate(args model.CollectionPreCreateRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	packed, err := abi.Pack("", args.Name, args.Symbol, args.TokenURIPrefix)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	bin, _ := hexutil.Decode(erc721.NftcontractMetaData.Bin)
	data := append(bin, packed...)

	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		Data: data,
	}
	txPre, err := chaincli.GetEthClient(args.BlockChain).TxPreExtra(msg)
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
//@function: CollectionPreSetContractURICreate
//@description: 组织设置合约URI的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.CollectionPreSetURI
//@return: *ethmodel.ContractCodeResponse
func (collection *ETHCollection) CollectionPreSetContractURICreate(args model.CollectionPreSetURI) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	packed, err := abi.Pack("setContractURI", args.TokenURIPrefix)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ToAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	txPre, err := chaincli.GetEthClient(args.BlockChain).TxPreExtra(msg)
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
//@function: NftPreSetTokenURIPrefixCreate
//@description: 组织设置NFT的URI前缀的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.CollectionPreSetURI
//@return: *ethmodel.ContractCodeResponse
func (collection *ETHCollection) CollectionPreSetTokenURIPrefixCreate(args model.CollectionPreSetURI) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	packed, err := abi.Pack("setTokenURIPrefix", args.TokenURIPrefix)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ToAddress)
	msg := ethereum.CallMsg{
		From: common.HexToAddress(args.FromAddress),
		To:   &to,
		Data: packed,
	}
	txPre, err := chaincli.GetEthClient(args.BlockChain).TxPreExtra(msg)
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
//@function: Name
//@description: Call调用返回合集名称，不上链不用进行签名
//@param: args ethmodel.CollectionNameRequest
//@return: *ethmodel.CollectionNameResponse
func (collection *ETHCollection) Name(args model.CollectionRequest) (string, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	nft721, err := erc721.NewNftcontract(contractAddress, chaincli.GetEthClient(args.BlockChain))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts:", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	name, err := nft721.Name(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to Name:", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return name, nil
}

//@author: [libofeng]
//@function: TokenURIPrefix
//@description: Call调用返回合集的下的NFT对应的URI，不上链不用进行签名
//@param: args ethmodel.CollectionRequest
//@return: *ethmodel.CollectionTokenURIPrefixResponse
func (collection *ETHCollection) TokenURIPrefix(args model.CollectionRequest) (string, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	nft721, err := erc721.NewNftcontract(contractAddress, chaincli.GetEthClient(args.BlockChain))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	tokenPrefix, err := nft721.GetTokenURIPrefix(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to TokenURIPrefix: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return tokenPrefix, nil
}

//@author: [libofeng]
//@function: Owner
//@description: Call调用返回创建合集的地址，不上链不用进行签名
//@param: args ethmodel.CollectionRequest
//@return: *ethmodel.CollectionOwnerResponse
func (collection *ETHCollection) Owner(args model.CollectionRequest) (string, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	nft721, err := erc721.NewNftcontract(contractAddress, chaincli.GetEthClient(args.BlockChain))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	owner, err := nft721.Owner(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to Owner: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return owner.Hex(), nil
}

//@author: [libofeng]
//@function: Symbol
//@description: Call调用返回合集符号，不上链不用进行签名
//@param: args ethmodel.CollectionNameRequest
//@return: *ethmodel.CollectionNameResponse
func (collection *ETHCollection) Symbol(args model.CollectionRequest) (string, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	nft721, err := erc721.NewNftcontract(contractAddress, chaincli.GetEthClient(args.BlockChain))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	symbol, err := nft721.Symbol(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to Symbol: ", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return symbol, nil
}

//@author: [libofeng]
//@function: Owner
//@description: Call调用返回NFTToken的总量，不上链不用进行签名
//@param: args ethmodel.CollectionRequest
//@return: *ethmodel.CollectionTotalSupplyResponse
func (collection *ETHCollection) TotalSupply(args model.CollectionRequest) (uint64, error) {
	contractAddress := common.HexToAddress(args.ContractAddress)
	nftToken, err := erc721.NewNftcontract(contractAddress, chaincli.GetEthClient(args.BlockChain))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewNft721: %v", zap.Error(err))
		return uint64(0), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	totalsupply, err := nftToken.TotalSupply(nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to TotalSupply: %v", zap.Error(err))
		return uint64(0), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return totalsupply.Uint64(), nil
}

// 通过私钥部署合约
func (collection *ETHCollection) CollectionCreateWithPrikey(args model.CollectionCreateWithPrikeyRequest) (*model.TxResponse, error) {
	// 签名
	nonce, err := chaincli.GetEthClient(args.BlockChain).NonceAt(context.Background(), common.HexToAddress(args.From), nil)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NonceAt: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	gasPrice, err := chaincli.GetEthClient(args.BlockChain).SuggestGasPrice(context.Background())
	if err != nil {
		global.MARKET_LOG.Debug("Failed to SuggestGasPrice: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	signer := types.HomesteadSigner{}
	auth := &bind.TransactOpts{
		From:     common.HexToAddress(args.From),
		Nonce:    new(big.Int).SetUint64(nonce),
		GasPrice: gasPrice,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			pKey, err := crypto.HexToECDSA(args.Prikey)
			signature, _ := crypto.Sign(signer.Hash(tx).Bytes(), pKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	contractAddress, tx, _, err := erc721.DeployNftcontract(auth, chaincli.GetEthClient(args.BlockChain), args.Name, args.Symbol, args.TokenURIPrefix)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to DeployContracts: %v", zap.Error(err))
		return nil, err
	}

	response := &model.TxResponse{
		TxHash:          tx.Hash().String(),
		ContractAddress: contractAddress.Hex(),
	}

	return response, err
}
