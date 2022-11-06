package impl

import (
	"github.com/Despenrado/webMesk/internal/service"
	"github.com/Despenrado/webMesk/internal/storage"
	"go.uber.org/fx"
)

type Service struct {
	storage        storage.Storage
	cacheStorage   storage.CacheStorage
	userService    *UserService
	chatService    *ChatService
	messageService *MessageService
	atuthService   *AtuthService
}

func NewService(
	storage storage.Storage,
	cacheStorage storage.CacheStorage,
	userService *UserService,
	chatService *ChatService,
	messageService *MessageService,
	atuthService *AtuthService,
) *Service {
	s := &Service{
		storage:        storage,
		cacheStorage:   cacheStorage,
		userService:    userService,
		chatService:    chatService,
		messageService: messageService,
		atuthService:   atuthService,
	}
	return s
}

func (s *Service) User() service.UserService {
	if s.userService != nil {
		return s.userService
	}
	s.userService = NewUserService(s.storage)
	return s.userService
}

func (s *Service) Chat() service.ChatService {
	if s.chatService != nil {
		return s.chatService
	}
	s.chatService = NewChatService(s.storage)
	return s.chatService
}

func (s *Service) Message() service.MessageService {
	if s.messageService != nil {
		return s.messageService
	}
	s.messageService = NewMessageService(s.storage)
	return s.messageService
}

func (s *Service) Auth() service.AuthService {
	if s.atuthService != nil {
		return s.atuthService
	}
	s.User()
	a := &AtuthService{}
	fx.Extract(a)
	s.atuthService = a
	return s.atuthService
}
