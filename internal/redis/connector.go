package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/ntec-io/Nostradamus/internal/logger"
)

type Client struct {
	rdb *redis.Client
	ctx context.Context
}

type wrongResultError struct {
	msg string
}

func (e wrongResultError) Error() string {
	return e.msg
}

func NewClient(pw string) (client Client, err error) {
	logger.Log().Debug("Creating redis client")
	client.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: pw,
		DB:       0,
	})
	client.ctx = context.Background()
	err = client.testConnection()
	return
}

func (c Client) testConnection() (err error) {
	logger.Log().Debug("redis connection test: started setting test value")
	err = c.rdb.Set(c.ctx, "test", "test123", 0).Err()
	if err != nil {
		logger.Log().Error(err)
		return
	}
	logger.Log().Debug("redis connection test: started reading test value")
	val, err := c.rdb.Get(c.ctx, "test").Result()
	if err != nil {
		logger.Log().Error(err)
		return
	}
	if val != "test123" {
		logger.Log().Error(err)
		err = wrongResultError{"test connection failed: results are not the same"}
		return
	}
	logger.Log().Debug("redis connection test: started deleting test value")
	err = c.rdb.Del(c.ctx, "test").Err()
	if err != nil {
		logger.Log().Error(err)
		return
	}
	logger.Log().Info("All redis connection tests successfull")
	return
}
