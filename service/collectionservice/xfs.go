package collectionservice

import (
	"math/big"
	"tokenagent/model/model"
)

type XFSCollection struct {
}

func (collection *XFSCollection) CollectionPreCreate(args model.CollectionPreCreateRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (collection *XFSCollection) CollectionPreSetContractURICreate(args model.CollectionPreSetURI) (*model.ContractCodeResponse, error) {

	return nil, nil
}

func (collection *XFSCollection) CollectionPreSetTokenURIPrefixCreate(args model.CollectionPreSetURI) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (collection *XFSCollection) Name(args model.CollectionRequest) (string, error) {
	return "", nil
}

func (collection *XFSCollection) Symbol(args model.CollectionRequest) (string, error) {
	return "", nil
}

func (collection *XFSCollection) TokenURIPrefix(args model.CollectionRequest) (string, error) {
	return "", nil
}

func (collection *XFSCollection) Owner(args model.CollectionRequest) (string, error) {
	return "", nil
}

func (collection *XFSCollection) TotalSupply(args model.CollectionRequest) (*big.Int, error) {
	return *new(*big.Int), nil
}

func (collection *XFSCollection) CollectionCreateWithPrikey(args model.CollectionCreateWithPrikeyRequest) (*model.TxResponse, error) {
	return nil, nil
}
