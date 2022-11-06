package impl

import (
	"github.com/Despenrado/webMesk/internal/service"
	"github.com/Despenrado/webMesk/internal/storage"
	"go.uber.org/fx"
)

var Module fx.Option = fx.Provide(ProvideServiceImpl)

func ProvideServiceImpl(
	storage storage.Storage,
	cacheStorage storage.CacheStorage,
	userService *UserService,
	chatService *ChatService,
	messageService *MessageService,
	authService *AtuthService,
) service.Service {
	return NewService(
		storage,
		cacheStorage,
		userService,
		chatService,
		messageService,
		authService,
	)
}
