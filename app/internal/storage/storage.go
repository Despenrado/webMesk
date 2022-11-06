package storage

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
)

type Storage interface {
	User() UserRepository
	Message() MessageRepository
	Chat() ChatRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error)
	FindById(ctx context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error
	GetUsersByFilter(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUserName(ctx context.Context, userName string) (*model.User, error)
}

type MessageRepository interface {
	Create(ctx context.Context, message *model.Message) (*model.Message, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error)
	FindById(ctx context.Context, id uint) (*model.Message, error)
	Update(ctx context.Context, message *model.Message) (*model.Message, error)
	Delete(ctx context.Context, message *model.Message) error
	FilterMessage(ctx context.Context, messageFilter *model.MessageFilter) ([]model.Message, error)
	MarkAsRead(ctx context.Context, id uint, user_id uint) error
}

type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.Chat, error)
	FindById(ctx context.Context, id uint) (*model.Chat, error)
	Update(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	Delete(ctx context.Context, id uint) error
	FindByUserId(ctx context.Context, id uint) ([]model.Chat, error)
	FilterChat(ctx context.Context, chatFilter *model.ChatFilter) ([]model.Chat, error)
}
