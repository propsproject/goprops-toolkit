package rediscache


import (
	"github.com/mediocregopher/radix.v2/redis"
	"fmt"
	propslogger "github.com/propsproject/goprops-toolkit/logger"
	"github.com/mediocregopher/radix.v2/pool"
)

// connuri: localhost:6739
const connURIFormat  = "%s:%v"

type PCache struct {
	logger *propslogger.Wrapper
}

func (pc *PCache) Redis(host string, port int, poolSize int) (*RedisCache, error) {
	return NewRedisCache(host, port, poolSize, pc.logger)
}

func PropsCache(logger *propslogger.Wrapper) *PCache {
	return &PCache{logger}
}

type RedisCache struct {
	pool   *pool.Pool
	logger *propslogger.Wrapper
	initialized bool
}

func (rc *RedisCache) Client() (*redis.Client, error) {
	if !rc.initialized {
		rc.logger.Warn("attempting to use redis conn that hasn't been initialized")
	}
	return rc.pool.Get()
}

func buildConnStr(port int, host string) string {
	return fmt.Sprintf(connURIFormat, host, port)
}

func NewRedisCache(host string, port int, poolSize int, logger *propslogger.Wrapper) (*RedisCache, error) {
	p, err := pool.New("tcp", buildConnStr(port, host), poolSize)
	if err != nil {
		return nil, fmt.Errorf(errOnOpenConnection, err)
	}

	return &RedisCache {pool: p, logger: logger,initialized: true}, nil
}
