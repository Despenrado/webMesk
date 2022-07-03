package storage

import (
	"context"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
)

type Storage interface {
	User() UserRepository
	Message() MessageRepository
	Chat() ChatRepository
	// Auth() AuthRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error)
	FindById(ctx context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error

	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUserName(ctx context.Context, userName string) (*model.User, error)
}

type MessageRepository interface {
	Create(ctx context.Context, message *model.Message) (*model.Message, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error)
	FindById(ctx context.Context, id uint) (*model.Message, error)
	Update(ctx context.Context, message *model.Message) (*model.Message, error)
	Delete(ctx context.Context, id uint) error
	FindByUserId(ctx context.Context, id uint) ([]model.Message, error)
	FindByChatId(ctx context.Context, id uint) ([]model.Message, error)
	FindByChatIdAndAfterDateTime(ctx context.Context, dateTime time.Time) ([]model.Message, error)
}

type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.Chat, error)
	FindById(ctx context.Context, id uint) (*model.Chat, error)
	Update(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	Delete(ctx context.Context, id uint) error
	FindByUserId(ctx context.Context, id uint) ([]model.Chat, error)
}

// type AuthRepository interface {
// }
