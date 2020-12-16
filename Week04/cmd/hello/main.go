package main

import (
	"Go-000/Week04/api/hello"
	"Go-000/Week04/configs"
	"Go-000/Week04/pkg/model"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func main() {
	// 初始化配置信息
	conf, err := configs.InitConfig()
	if err != nil {
		log.Fatalf("Init config failed. err: %s\n", err.Error())
	}

	// 初始化数据库
	db, err := initDB(conf)
	if err != nil {
		log.Fatalf("Init database failed. err: %s\n", err.Error())
	}

	// 初始化缓存
	cache := initCache(conf)

	// 依赖注入
	service := InitializeHello(db, cache)

	listen, err := net.Listen("tcp", "127.0.0.1:"+conf.Server.Port)
	if err != nil {
		os.Exit(1)
	}

	server := grpc.NewServer()
	hello.RegisterHelloServer(server, service)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("RPC server listen failed. err: %s\n", err.Error())
	}
}

func initDB(conf *configs.Config) (*gorm.DB, error) {
	db, err := model.NewDBEngine(conf)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initCache(conf *configs.Config) *redis.Client {
	return nil
}
