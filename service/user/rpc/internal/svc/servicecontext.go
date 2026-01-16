package svc

import (
	"user/rpc/internal/config"
	"user/rpc/internal/model/postgres"
	"common/utils"
)

type ServiceContext struct {
	Config config.Config
	UserRepo *postgres.UserRepo
}

func NewServiceContext(c config.Config, userRepo *postgres.UserRepo) *ServiceContext {
	snowflake.Init()
	return &ServiceContext{
		Config: c,
		UserRepo: userRepo,
	}
}
