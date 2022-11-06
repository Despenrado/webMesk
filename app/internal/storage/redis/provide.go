package redis

import (
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/utils"
	"go.uber.org/fx"
)

var Module fx.Option = fx.Provide(ProvideRedisCacheStorage)

func ProvideRedisCacheStorage(
	config *utils.Config,
) (storage.CacheStorage, error) {
	rClient := NewRedisConnection(*config.RedisConfig)
	return NewCacheStorage(rClient, config.RedisConfig)
}
