package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
)

type AuthRepository struct {
	storage   *CacheStorage
	expiresAt time.Duration
}

func NewAuthRepository(storage *CacheStorage) *AuthRepository {
	return &AuthRepository{
		storage:   storage,
		expiresAt: storage.config.TknExpires,
	}
}

func (ar *AuthRepository) Set(ctx context.Context, token *model.Token) error {
	if err := ar.storage.redisClient.Set(token.ID, token, ar.expiresAt).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ar *AuthRepository) FindById(ctx context.Context, id string) (*model.Token, error) {
	res, err := ar.storage.redisClient.Get(id).Result()
	if err != nil {
		return nil, err
	}
	if res != "" {
		return nil, utils.ErrRecordNotFound
	}
	token := &model.Token{}
	if err := json.Unmarshal([]byte(res), token); err != nil {
		return nil, err
	}
	return token, nil
}

func (ar *AuthRepository) DeleteByUserId(ctx context.Context, id string) error {
	// ar.storage.redisClient.Del(id)
	log.Println(ar.storage.redisClient.Del(id))
	return nil
}
