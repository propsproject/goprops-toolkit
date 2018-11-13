package rediscache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/propsproject/goprops-toolkit/logging"
)

// connuri: localhost:6739
const connURIFormat = "%s:%v"

type PCache struct {
	logger *logging.Logger
}

func (pc *PCache) Redis(host string, port int, poolSize int) *RedisCache {
	return NewRedisCache(host, port, poolSize, pc.logger)
}

func PropsCache(logger *logging.Logger) *PCache {
	return &PCache{logger}
}

type RedisCache struct {
	client      *redis.Client
	logger      *logging.Logger
	initialized bool
}

func (rc *RedisCache) Client() *redis.Client {
	if !rc.initialized {
		rc.logger.Warn("attempting to use redis conn that hasn't been initialized")
	}
	return rc.client
}

func buildConnStr(port int, host string) string {
	return fmt.Sprintf(connURIFormat, host, port)
}

func NewRedisCache(host string, port int, poolSize int, logger *logging.Logger) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr:    buildConnStr(port, host),
		}),
		logger:      logger,
		initialized: true,
	}
}
