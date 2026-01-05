package svc

import (
	"user/rpc/internal/config"
	"user/rpc/internal/model/postgres"
)

type ServiceContext struct {
	Config config.Config
	UserRepo *postgres.UserRepo
}

func NewServiceContext(c config.Config, userRepo *postgres.UserRepo) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRepo: userRepo,
	}
}
