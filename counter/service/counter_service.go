package service

import (
	"context"
	"fmt"
	manager_rpc "tracker/manager_grpc/rpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gomodule/redigo/redis"
	"github.com/satori/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type TokenGetter func() (string, error)

type CounterService struct {
	RedisPool *redis.Pool
	GetToken  TokenGetter
}

func (ms CounterService) NewToken(ctx context.Context, in *empty.Empty) (*manager_rpc.NewTokenResponse, error) {
	resp := &manager_rpc.NewTokenResponse{}

	token, err := ms.GetToken()
	if err != nil {
		return resp, err
	}

	conn := ms.RedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("SET", token, "true")

	if err != nil {
		log.Errorf("can not save token: %v", err)
		return resp, fmt.Errorf("can not save token: %v", err)
	}

	resp.Token = token
	return resp, nil
}

func (ms CounterService) UpdateCounter(ctx context.Context, in *manager_rpc.UpdateCounterRequest) (*manager_rpc.UpdateCounterResponse, error) {
	resp := &manager_rpc.UpdateCounterResponse{}

	path := in.Path
	if len(path) == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "path cannot be empty")
	}

	token := in.Token
	if len(token) == 0 {
		resp.Success = false
		return resp, nil
	}

	conn := ms.RedisPool.Get()
	defer conn.Close()
	exists, err := conn.Do("EXISTS", token)
	if err != nil {
		return nil, fmt.Errorf("redis error: %v", err)
	}
	if exists.(int64) == 0 {
		resp.Success = false
		return resp, nil
	}

	key := fmt.Sprintf("%s_%s", token, path)
	connIncr := ms.RedisPool.Get()
	defer connIncr.Close()
	counter, err := connIncr.Do("INCR", key)
	if err != nil {
		return nil, fmt.Errorf("redis error: %v", err)
	}

	resp.Success = true
	resp.Counter = counter.(int64)

	return resp, nil
}

func NewUUID() (string, error) {
	token, err := uuid.NewV4()
	if err != nil {
		log.Errorf("can not create token: %v", err)
		return "", fmt.Errorf("can not create token: %v", err)
	}

	return token.String(), nil
}
