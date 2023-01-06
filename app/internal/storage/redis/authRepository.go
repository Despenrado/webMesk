package redis

import (
	"context"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
)

type AuthRepository struct {
	storage *CacheStorage
}

func NewAuthRepository(storage *CacheStorage) *AuthRepository {
	return &AuthRepository{
		storage: storage,
	}
}

func (ar *AuthRepository) Set(ctx context.Context, token *model.UserAuth, expiresAt time.Duration) error {
	if err := ar.storage.redisClient.Set(token.ID, token, expiresAt).Err(); err != nil {
		return err
	}
	return nil
}

func (ar *AuthRepository) FindById(ctx context.Context, id string) (*model.UserAuth, error) {
	res, err := ar.storage.redisClient.Get(id).Result()
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, utils.ErrRecordNotFound
	}
	token := &model.UserAuth{}
	if err := token.UnmarshalBinary([]byte(res)); err != nil {
		return nil, err
	}
	return token, nil
}

func (ar *AuthRepository) DeleteByUserId(ctx context.Context, id string) error {
	// ar.storage.redisClient.Del(id)
	res, err := ar.storage.redisClient.Del(id).Result()
	if err != nil {
		return err
	}
	if res != 1 {
		return utils.ErrRecordNotFound
	}
	return nil
}
