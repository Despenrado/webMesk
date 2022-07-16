package storage

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
)

type CacheStorage interface {
	Auth() AuthRepository
}

type AuthRepository interface {
	Set(context.Context, *model.Token) error
	FindById(context.Context, string) (*model.Token, error)
	DeleteByUserId(context.Context, string) error
}
