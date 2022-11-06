package service

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
)

type Service interface {
	User() UserService
	Message() MessageService
	Chat() ChatService
	Auth() AuthService
}

type UserService interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error)
	FindById(ctx context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error
	FilterUser(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type MessageService interface {
	Create(ctx context.Context, message *model.Message) (*model.Message, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error)
	FindById(ctx context.Context, id uint) (*model.Message, error)
	Update(ctx context.Context, message *model.Message) (*model.Message, error)
	Delete(ctx context.Context, message *model.Message) error
	FilterMessage(ctx context.Context, messageFilter *model.MessageFilter) ([]model.Message, error)
	MarkAsRead(ctx context.Context, id uint, user_id uint) error
}

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	ReadAll(ctx context.Context, skip int, limit int) ([]model.Chat, error)
	FindById(ctx context.Context, id uint) (*model.Chat, error)
	Update(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	Delete(ctx context.Context, id uint) error
	FindByUserId(ctx context.Context, id uint) ([]model.Chat, error)
	FilterChat(ctx context.Context, chatFilter *model.ChatFilter) ([]model.Chat, error)
}

type AuthService interface {
	SignUp(context.Context, *model.User) (string, error)
	SignIn(context.Context, *model.User) (*model.User, string, error)
	ValidAuthorization(context.Context, string) (context.Context, error)
	Reauthorize(context.Context, string, string) error
	LogOut(context.Context, string) error
}
