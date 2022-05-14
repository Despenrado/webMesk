package storage

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
)

type Storage interface {
	User() UserRepository
	// Message() MessageRepository
	// Chat() ChatRepository
	// Auth() AuthRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetAll(ctx context.Context) (*[]model.User, error)
	GetById(ctx context.Context, id uint64) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint64) (*model.User, error)
}

// type MessageRepository interface {
// 	Create(ctx context.Context, message *model.Message) (*model.Message, error)
// 	GetAll(ctx context.Context) (*[]model.Message, error)
// 	GetById(ctx context.Context, id uint64) (*model.Message, error)
// 	Update(ctx context.Context, message *model.Message) (*model.Message, error)
// 	Delete(ctx context.Context, id uint64) (*model.Message, error)
// }

// type ChatRepository interface {
// 	Create(ctx context.Context, chat *model.Chat) (*model.Chat, error)
// 	GetAll(ctx context.Context) (*[]model.Chat, error)
// 	GetById(ctx context.Context, id uint64) (*model.Chat, error)
// 	Update(ctx context.Context, chat *model.Chat) (*model.Chat, error)
// 	Delete(ctx context.Context, id uint64) (*model.Chat, error)
// }

// type AuthRepository interface {
// }
