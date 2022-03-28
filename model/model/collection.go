package model

type EmptyArgs = interface{}

// Collection PreCreate Request
type CollectionPreCreateRequest struct {
	BlockChain     string `json:"blockchain"`
	FromAddress    string `json:"from_address"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	TokenURIPrefix string `json:"token_uri_prefix"`
}

// Collection Create Request
type CollectionCreateWithPrikeyRequest struct {
	BlockChain     string `json:"blockchain"`
	From           string `json:"from"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	TokenURIPrefix string `json:"token_uriprefix"`
	Prikey         string `json:"prikey"`
}

// Collection PreCreate Response
type ContractCodeResponse struct {
	TxPreExtra
	Code string `json:"code"`
}

// ERC721合集的基本信息请求
type CollectionRequest struct {
	BlockChain      string `json:"blockchain"`
	ContractAddress string `json:"contract_address"`
}

type CollectionPreSetURI struct {
	BlockChain     string `json:"blockchain"`
	FromAddress    string `json:"from_address"`
	ToAddress      string `json:"to_address"`
	TokenURIPrefix string `json:"token_uri_prefix"`
}
