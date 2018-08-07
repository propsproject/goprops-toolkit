package rediscache


import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	propslogger "github.com/propsproject/goprops-toolkit/logger"
	"time"
)

// connuri: localhost:6739
const connURIFormat  = "%s:%v"

type PCache struct {
	logger *propslogger.Wrapper
}

func (pc *PCache) Redis(host string, port int, poolSize int) *RedisCache {
	return NewRedisCache(host, port, poolSize, pc.logger)
}

func PropsCache(logger *propslogger.Wrapper) *PCache {
	return &PCache{logger}
}

type RedisCache struct {
	pool   *redis.Pool
	logger *propslogger.Wrapper
	initialized bool
}

func (rc *RedisCache) Client() redis.Conn {
	if !rc.initialized {
		rc.logger.Warn("attempting to use redis conn that hasn't been initialized")
	}
	return rc.pool.Get()
}

func newPool(host string, port int, poolSize int) *redis.Pool {
	return &redis.Pool{
		MaxIdle: poolSize,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			return redis.Dial("tcp", buildConnStr(port, host))
		},
	}
}

func buildConnStr(port int, host string) string {
	return fmt.Sprintf(connURIFormat, host, port)
}

func NewRedisCache(host string, port int, poolSize int, logger *propslogger.Wrapper) *RedisCache {
	pool := newPool(host, port, poolSize)
	return &RedisCache {pool: pool, logger: logger,initialized: true}
}
