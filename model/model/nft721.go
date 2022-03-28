package model

type NftPreMintRequest struct {
	BlockChain  string `json:"blockchain"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	TokenId     string `json:"token_id"`
	TokenUri    string `json:"token_uri"`
}

type NftPreApproveRequest struct {
	BlockChain  string `json:"blockchain"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	TokenId     string `json:"token_id"`
}

type NftPreBurnRequest struct {
	BlockChain  string `json:"blockchain"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	TokenId     string `json:"token_id"`
}

type NftPreSetApprovalForAllRequest struct {
	BlockChain string `json:"blockchain"`
	To         string `json:"to"`
	Approved   bool   `json:"approved"`
}

type NftPreTransferFromRequest struct {
	BlockChain  string `json:"blockchain"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	TokenId     string `json:"token_id"`
}

type NftPreTransferFromRequestEx struct {
	BlockChain  string `json:"blockchain"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	TokenId     string `json:"token_id"`
	Data        string `json:"data"`
}

type NftPreTransferOwnershipRequest struct {
	BlockChain string `json:"blockchain"`
	NewOwner   string `json:"new_owner"`
}

type NftOwnerOfRequest struct {
	BlockChain      string `json:"blockchain"`
	ContractAddress string `json:"contract_address"`
	TokenId         string `json:"token_id"`
}

type NftBalanceOfRequest struct {
	BlockChain      string `json:"blockchain"`
	ContractAddress string `json:"contract_address"`
	Owner           string `json:"owner"`
}

type NftTokenInfoRequest struct {
	BlockChain      string `json:"blockchain"`
	ContractAddress string `json:"contract_address"`
	TokenId         string `json:"token_id"`
}
