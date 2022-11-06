package storage

import (
	"context"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
)

type CacheStorage interface {
	Auth() AuthRepository
}

type AuthRepository interface {
	Set(context.Context, *model.UserAuth, time.Duration) error
	FindById(context.Context, string) (*model.UserAuth, error)
	DeleteByUserId(context.Context, string) error
}
