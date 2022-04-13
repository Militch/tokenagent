package service

import (
	"tokenagent/service/chainservice"
	"tokenagent/service/collectionservice"
	"tokenagent/service/nftservice"
	"tokenagent/service/tokenservice"
)

type ServiceGroup struct {
	ChainServiceProxy chainservice.ChainProxy
	CollectionProxy   collectionservice.CollectionProxy
	NFTProxy          nftservice.NFTProxy
	TokenProxy        tokenservice.TokenProxy
}

var ServiceGroupApp = new(ServiceGroup)
