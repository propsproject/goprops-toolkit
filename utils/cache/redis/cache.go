package rediscache


import (
	"github.com/mediocregopher/radix.v2/redis"
	"fmt"
	propslogger "github.com/propsproject/goprops-toolkit/logger"
)

type RedisCache struct {
	instance   *redis.Client
	logger *propslogger.Wrapper
	initialized bool
}

func (rc *RedisCache) Client() *redis.Client {
	if !rc.initialized {
		rc.logger.Warn("attempting to use redis conn that hasn't been initialized")
	}
	return rc.instance
}

func (rc *RedisCache) CMD() *command {
	return &command{instance: rc.instance}
}

func NewRedisCache(host string, port int, logger *propslogger.Wrapper) (*RedisCache, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		return nil, fmt.Errorf(errOnOpenConnection, err)
	}

	if resp := conn.Cmd(PING); resp.Err != nil {
		conn.Close()
		return nil, fmt.Errorf(errOnOpenConnection, err)
	}

	return &RedisCache {instance: conn,logger: logger,initialized: true}, nil
}
