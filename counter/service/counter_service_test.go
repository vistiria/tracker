package service

import (
	"context"
	"errors"
	"testing"

	manager_rpc "tracker/manager_grpc/rpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

const validToken = "3b39ac51-56fd-46c6-8fa7-bf2aa9cbbc1a"

func getToken() (string, error) { return validToken, nil }

func getTokenError() (string, error) { return "", errors.New("can not create token: err") }

func getPool(conn *redigomock.Conn) *redis.Pool {
	return &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}
}

func TestNewToken(t *testing.T) {
	var in *empty.Empty
	conn := redigomock.NewConn()
	ctx := context.Background()

	testCases := []struct {
		cmdToken       string
		tokenGetter    TokenGetter
		resultErr      error
		resultCmdCount int
		resultToken    string
		redisError     bool
	}{
		{
			validToken,
			getToken,
			nil,
			1,
			validToken,
			false,
		},
		{
			"",

			getTokenError,
			errors.New("can not create token: err"),
			0,
			"",
			false,
		},
		{
			validToken,
			getToken,
			errors.New("can not save token: Low level error!"),
			1,
			"",
			true,
		},
	}

	for _, testCase := range testCases {
		conn.Clear()
		cs := CounterService{RedisPool: getPool(conn), GetToken: testCase.tokenGetter}

		cmd := conn.Command("SET", testCase.cmdToken, "true")
		if testCase.redisError {
			cmd = cmd.ExpectError(errors.New("Low level error!"))
		}

		result, err := cs.NewToken(ctx, in)

		assert.Equal(t, conn.Stats(cmd), testCase.resultCmdCount)
		assert.Equal(t, result.Token, testCase.resultToken)
		assert.Equal(t, err, testCase.resultErr)
	}
}

func TestUpdateCounterInvalidPath(t *testing.T) {
	conn := redigomock.NewConn()
	ctx := context.Background()
	cs := CounterService{RedisPool: getPool(conn), GetToken: getToken}

	req := manager_rpc.UpdateCounterRequest{Path: "", Token: validToken}
	_, err := cs.UpdateCounter(ctx, &req)

	assert.Equal(t, err.Error(), "rpc error: code = InvalidArgument desc = path cannot be empty")
}

func TestUpdateCounterEmptyToken(t *testing.T) {
	conn := redigomock.NewConn()
	ctx := context.Background()
	cs := CounterService{RedisPool: getPool(conn), GetToken: getToken}

	cmd := conn.Command("EXISTS", "")

	req := manager_rpc.UpdateCounterRequest{Path: "/kkk", Token: ""}
	result, err := cs.UpdateCounter(ctx, &req)

	assert.Equal(t, result.Success, false)
	assert.Equal(t, conn.Stats(cmd), 0)
	assert.Equal(t, err, nil)
}

func TestUpdateCounterInvalidToken(t *testing.T) {
	conn := redigomock.NewConn()
	ctx := context.Background()
	cs := CounterService{RedisPool: getPool(conn), GetToken: getToken}

	cmd := conn.Command("EXISTS", "i_am_invalid").Expect(int64(0))

	req := manager_rpc.UpdateCounterRequest{Path: "/kkk", Token: "i_am_invalid"}
	result, err := cs.UpdateCounter(ctx, &req)

	assert.Equal(t, result.Success, false)
	assert.Equal(t, conn.Stats(cmd), 1)
	assert.Equal(t, err, nil)
}
