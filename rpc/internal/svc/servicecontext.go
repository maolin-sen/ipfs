package svc

import "github.com/maolin-sen/go-filecoin-market/apps/ipfs/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
