package nftservice

import (
	"fmt"
	"math/big"
	"strings"
	"tokenagent/contracts/erc721"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils/chaincli"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
)

type ETHNft struct {
}

//@author: [libofeng]
//@function: NftPreMintCreate
//@description: 组织铸造NFT的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreMintRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreMintCreate(args model.NftPreMintRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("mint", tokenId, args.TokenUri)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ToAddress)
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
//@function: NftPreApproveCreate
//@description: 组织铸造NFT的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreApproveRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreApproveCreate(args model.NftPreApproveRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("approve", common.HexToAddress(args.ToAddress), tokenId)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: ", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	to := common.HexToAddress(args.ToAddress)
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
//@function: NftPreBurnCreate
//@description: 组织销毁NFT的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreBurnRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreBurnCreate(args model.NftPreBurnRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("burn", tokenId)
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
//@function: NftPreSetApprovalForAllCreate
//@description: 组织授权NFT的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreSetApprovalForAllRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreSetApprovalForAllCreate(args model.NftPreSetApprovalForAllRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	packed, err := abi.Pack("setApprovalForAll", common.HexToAddress(args.To), args.Approved)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	response := &model.ContractCodeResponse{
		Code: hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: NftPreTransferFromCreate
//@description: 组织转移NFT交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreTransferFromRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("transferFrom", common.HexToAddress(args.FromAddress), common.HexToAddress(args.ToAddress), tokenId)
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
//@function: NftPreSafeTransferFromCreate
//@description: 组织转移NFT交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreTransferFromRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreSafeTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("safeTransferFrom", common.HexToAddress(args.FromAddress), common.HexToAddress(args.ToAddress), tokenId)
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
//@function: NftPreSafeTransferFromExCreate
//@description: 组织转移NFT交易的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreTransferFromRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreSafeTransferFromExCreate(args model.NftPreTransferFromRequestEx) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	packed, err := abi.Pack("safeTransferFrom0", common.HexToAddress(args.FromAddress), common.HexToAddress(args.ToAddress), tokenId, common.Hex2Bytes(args.Data))
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
//@function: NftPreTransferOwnershipCreate
//@description: 组织授权NFT的code码，同时通过TxPreExtra返回Nonce|gasprice|gaslimit信息给前端参考,前端进行签名
//@param: args ethmodel.NftPreTransferOwnershipRequest
//@return: *ethmodel.ContractCodeResponse
func (nft *ETHNft) NftPreTransferOwnershipCreate(args model.NftPreTransferOwnershipRequest) (*model.ContractCodeResponse, error) {
	abi, err := abi.JSON(strings.NewReader(erc721.NftcontractMetaData.ABI))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	packed, err := abi.Pack("transferOwnership", common.HexToAddress(args.NewOwner))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to abi.Pack: %v", zap.Error(err))
		return nil, global.NewRPCError(global.InternalErrorCode, "An exception occurred of contract argument")
	}

	response := &model.ContractCodeResponse{
		Code: hexutil.Encode(packed),
	}

	return response, nil
}

//@author: [libofeng]
//@function: OwnerOf
//@description: 查询NFT的拥有者，直接调用，不用签名
//@param: args ethmodel.NftOwnerOfRequest
//@return: string
func (nft *ETHNft) OwnerOf(args model.NftOwnerOfRequest) (string, error) {
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	contractAddress := common.HexToAddress(args.ContractAddress)
	nftToken, err := erc721.NewNftcontract(contractAddress, cli.Client)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: %v", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}
	addr, err := nftToken.OwnerOf(nil, tokenId)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to OwnerOf: %v", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return addr.Hex(), nil
}

//@author: [libofeng]
//@function: BalanceOf
//@description: 查询地址拥有NFT余额，直接调用，不用签名
//@param: args ethmodel.NftBalanceOfRequest
//@return: uint64
func (nft *ETHNft) BalanceOf(args model.NftBalanceOfRequest) (uint64, error) {
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return 0, global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	contractAddress := common.HexToAddress(args.ContractAddress)
	nftToken, err := erc721.NewNftcontract(contractAddress, cli.Client)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: %v", zap.Error(err))
		return uint64(0), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	balance, err := nftToken.BalanceOf(nil, common.HexToAddress(args.Owner))
	if err != nil {
		global.MARKET_LOG.Debug("Failed to BalanceOf: %v", zap.Error(err))
		return uint64(0), global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}

	return balance.Uint64(), nil
}

//@author: [libofeng]
//@function: TokenUri
//@description: 获取NFT Token的URI，直接调用，不用签名
//@param: args ethmodel.NftTokenUriRequest
//@return: string
func (nft *ETHNft) TokenUri(args model.NftTokenInfoRequest) (string, error) {
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	contractAddress := common.HexToAddress(args.ContractAddress)
	nftToken, err := erc721.NewNftcontract(contractAddress, cli.Client)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: %v", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}

	uri, err := nftToken.TokenURI(nil, tokenId)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to TokenURI: %v", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	return uri, nil
}

//@author: [libofeng]
//@function: GetApproved
//@description: 获取NFT Token的URI，直接调用，不用签名
//@param: args ethmodel.NftTokenUriRequest
//@return: string
func (nft *ETHNft) GetApproved(args model.NftTokenInfoRequest) (string, error) {
	cli := chaincli.GetEthClient(args.BlockChain)
	if cli == nil {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("blockchain:`%v` No connection established in service", args.BlockChain))
	}
	contractAddress := common.HexToAddress(args.ContractAddress)
	nftToken, err := erc721.NewNftcontract(contractAddress, cli.Client)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to NewContracts: %v", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	tokenId, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("Invalid TokenId on the %v Chain", args.BlockChain))
	}

	address, err := nftToken.GetApproved(nil, tokenId)
	if err != nil {
		global.MARKET_LOG.Debug("Failed to TokenURI: %v", zap.Error(err))
		return "", global.NewRPCError(global.InternalErrorCode, fmt.Sprintf("%v on the %v Chain", err.Error(), args.BlockChain))
	}
	return address.Hex(), nil
}

// 通过私钥生成NFT token
// func (ce *CustomEthClient) NftMintWithPrikey(args model.CollectionCreateWithPrikeyRequest) (*model.EthTxResponse, error) {
// 	// 签名
// 	nonce, err := ce.NonceAt(context.Background(), common.HexToAddress(args.From), nil)
// 	if err != nil {
// 		global.MARKET_LOG.Info("Failed to NonceAt: %v", zap.Error(err))
// 		return nil, err
// 	}

// 	gasPrice, err := ce.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		global.MARKET_LOG.Info("Failed to SuggestGasPrice: %v", zap.Error(err))
// 		return nil, err
// 	}

// 	signer := types.HomesteadSigner{}
// 	auth := &bind.TransactOpts{
// 		From:     common.HexToAddress(args.From),
// 		Nonce:    new(big.Int).SetUint64(nonce),
// 		GasPrice: gasPrice,
// 		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
// 			pKey, err := crypto.HexToECDSA(args.Prikey)
// 			if err != nil {
// 				global.MARKET_LOG.Info("Failed to HexToECDSA: %v", zap.Error(err))
// 				return nil, err
// 			}
// 			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), pKey)
// 			if err != nil {
// 				global.MARKET_LOG.Info("Failed to Sign: %v", zap.Error(err))
// 				return nil, err
// 			}
// 			return tx.WithSignature(signer, signature)
// 		},
// 	}

// 	_, tx, _, err := contracts.DeployNft721(auth, ce, args.Name, args.Symbol, common.HexToAddress(args.SingerAddress), args.ContractURI, args.TokenURIPrefix)
// 	if err != nil {
// 		global.MARKET_LOG.Info("Failed to DeployNft721: %v", zap.Error(err))
// 		return nil, err
// 	}

// 	response := &model.EthTxResponse{
// 		TxHash: tx.Hash().String(),
// 	}

// 	return response, nil
// }
