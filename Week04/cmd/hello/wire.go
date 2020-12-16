// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Go-000/Week04/internal/biz"
	"Go-000/Week04/internal/data"
	"Go-000/Week04/internal/service"

	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// InitializeEvent 声明injector的函数签名
func InitializeHello(db *gorm.DB, cache *redis.Client) *service.HelloService {
	wire.Build(service.NewHelloService, biz.NewHelloBIZ, data.NewHelloRepository)
	return &service.HelloService{}
}
