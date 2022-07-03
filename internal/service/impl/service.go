package impl

import (
	"github.com/Despenrado/webMesk/internal/service"
	"github.com/Despenrado/webMesk/internal/storage"
)

type Service struct {
	storage        storage.Storage
	userService    *UserService
	chatService    *ChatService
	messageService *MessageService
}

func NewService(
	storage storage.Storage,
	userService *UserService,
	chatService *ChatService,
	messageService *MessageService,
) *Service {
	s := &Service{
		storage:        storage,
		userService:    userService,
		chatService:    chatService,
		messageService: messageService,
	}
	return s
}

func (s *Service) User() service.UserService {
	if s.userService != nil {
		return s.userService
	}
	return NewUserService(s.storage)
}

func (s *Service) Chat() service.ChatService {
	if s.userService != nil {
		return s.chatService
	}
	return NewChatService(s.storage)
}

func (s *Service) Message() service.MessageService {
	if s.userService != nil {
		return s.messageService
	}
	return NewMessageService(s.storage)
}
