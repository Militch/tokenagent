package nftservice

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func Test_PreNft_Singer(t *testing.T) {
	codeData := "0x77097fc800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000002e516d5a63474e47586e52634d46674c6b5755617735664a653376746a77706336796e7651574c315457537869536a000000000000000000000000000000000000"
	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("3cd8dc06b4117ac25cfda6a6ba1b94ee24a01c3e735c77bf1a3250bedd912561")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 合约地址
	toAddress := common.HexToAddress("0xA1590e293AB1D15CF3699b1dFAD20C39901426b1")
	nonce, err := cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    44,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       &toAddress,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("Nonce:%v\n", nonce)
	log.Printf("MintPreCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func Test_NftPreTransferFromCreate_Singer(t *testing.T) {
	codeData := "0x23b872dd000000000000000000000000d5d96202f7e7a86c752bc10fb70e4fc8f3e89fdb000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb922660000000000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("3cd8dc06b4117ac25cfda6a6ba1b94ee24a01c3e735c77bf1a3250bedd912561")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 合约地址
	toAddress := common.HexToAddress("0xA1590e293AB1D15CF3699b1dFAD20C39901426b1")

	_, err = cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    45,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       &toAddress,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreTransferFromCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func Test_NftPreSafeTransferFromCreate_Singer(t *testing.T) {
	codeData := "0x42842e0e000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000d5d96202f7e7a86c752bc10fb70e4fc8f3e89fdb0000000000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 合约地址
	toAddress := common.HexToAddress("0xA1590e293AB1D15CF3699b1dFAD20C39901426b1")

	_, err = cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    16392,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       &toAddress,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreTransferFromCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func Test_NftPreSafeTransferFromExCreate_Singer(t *testing.T) {
	codeData := "0xb88d4fde000000000000000000000000d5d96202f7e7a86c752bc10fb70e4fc8f3e89fdb000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("3cd8dc06b4117ac25cfda6a6ba1b94ee24a01c3e735c77bf1a3250bedd912561")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 合约地址
	toAddress := common.HexToAddress("0xA1590e293AB1D15CF3699b1dFAD20C39901426b1")

	_, err = cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    46,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       &toAddress,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreTransferFromCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func TestCustomEthClient_NftPreApproveCreate_singer(t *testing.T) {
	codeData := "0x095ea7b3000000000000000000000000e92b5c51e04f4912601877a98c460aba7cfdc4ab0000000000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	// 地址0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266私钥ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
	// 地址0xD5d96202f7e7a86c752bc10Fb70e4fC8f3E89fDB私钥3cd8dc06b4117ac25cfda6a6ba1b94ee24a01c3e735c77bf1a3250bedd912561
	priKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 签名
	_, err = cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 接收地址
	toAddress := common.HexToAddress("0xA1590e293AB1D15CF3699b1dFAD20C39901426b1")
	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    16393,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       &toAddress,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreApproveCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func TestCustomEthClient_NftPreBurnCreate_singer(t *testing.T) {
	codeData := "0x42966c680000000000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)

	toAddress := common.HexToAddress("0xA1590e293AB1D15CF3699b1dFAD20C39901426b1")

	_, err = cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    16394,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       &toAddress,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreBurnCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func TestCustomEthClient_NftPreSetApprovalForAllCreate_singer(t *testing.T) {
	codeData := "0x77097fc80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000033132330000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("d01cf24b251ece252cac27ba04942b9b6c082e6564e923bef751d85c9fa9c054")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 签名
	nonce, err := cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       nil,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreSetApprovalForAllCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func TestCustomEthClient_NftPreSetContractURICreate_Singer(t *testing.T) {
	codeData := "0x77097fc80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000033132330000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("d01cf24b251ece252cac27ba04942b9b6c082e6564e923bef751d85c9fa9c054")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 签名
	nonce, err := cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       nil,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreSetContractURICreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func TestCustomEthClient_NftPreSetTokenURIPrefixCreate_Singer(t *testing.T) {
	codeData := "0x77097fc80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000033132330000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("d01cf24b251ece252cac27ba04942b9b6c082e6564e923bef751d85c9fa9c054")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 签名
	nonce, err := cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       nil,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreSetTokenURIPrefixCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}

func TestCustomEthClient_NftPreTransferOwnershipCreate_Singer(t *testing.T) {
	codeData := "0x77097fc80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000033132330000000000000000000000000000000000000000000000000000000000"

	cli, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	// 私钥
	priKey, err := crypto.HexToECDSA("d01cf24b251ece252cac27ba04942b9b6c082e6564e923bef751d85c9fa9c054")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 地址
	account := crypto.PubkeyToAddress(priKey.PublicKey)
	// 签名
	nonce, err := cli.NonceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	byteData, err := hexutil.Decode(codeData)
	if err != nil {
		log.Fatal(err)
		return
	}
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      5000000,
		To:       nil,
		Value:    new(big.Int).SetUint64(0),
		Data:     byteData,
	}

	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), priKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	rawTxBytes, _ := rlp.EncodeToBytes(signedTx)
	log.Printf("NftPreTransferOwnershipCreate singnature:%v\n", hex.EncodeToString(rawTxBytes))
}
