package main

import (
	"net"
	"tracker/counter/service"
	manager_rpc "tracker/manager_grpc/rpc"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("COUNTER_ADDR", "127.0.0.1:6381")
	viper.SetDefault("REDIS_ADDR", "127.0.0.1:6379")

	redisPool := newPool()
	lis, err := net.Listen("tcp", viper.GetString("COUNTER_ADDR"))
	if err != nil {
		log.Fatalf("failed to initialize TCP listen: %v", err)
	}
	defer lis.Close()

	counterService := service.CounterService{RedisPool: redisPool, GetToken: service.NewUUID}
	grpcServer := grpc.NewServer()
	manager_rpc.RegisterCounterServiceServer(grpcServer, counterService)

	log.Printf("Starting gRPC counter on %s", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("gRPC counter server error: %v", err)
	}
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", viper.GetString("REDIS_ADDR"))
			if err != nil {
				log.Fatalf("can not connect redis err:%v", err)
				panic(err.Error())
			}
			return conn, err
		},
	}
}
