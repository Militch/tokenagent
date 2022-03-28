package service

import (
	"tokenagent/service/chainservice"
	"tokenagent/service/collectionservice"
	"tokenagent/service/nftservice"
)

type ServiceGroup struct {
	ChainServiceProxy chainservice.ChainProxy
	CollectionProxy   collectionservice.CollectionProxy
	NFTProxy          nftservice.NFTProxy
}

var ServiceGroupApp = new(ServiceGroup)
