package redis

import (
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/utils"
	"github.com/go-redis/redis"
)

type CacheStorage struct {
	redisClient    *redis.Client
	authRepository *AuthRepository
	config         *utils.RedisConfig
}

func NewRedisConnection(config utils.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})
}

func NewCacheStorage(redisClient *redis.Client, config *utils.RedisConfig) *CacheStorage {
	s := &CacheStorage{
		redisClient: redisClient,
		config:      config,
	}
	s.authRepository = NewAuthRepository(s)
	return s
}

func (cs *CacheStorage) Auth() storage.AuthRepository {
	if cs.authRepository != nil {
		return cs.authRepository
	}
	return NewAuthRepository(cs)
}
