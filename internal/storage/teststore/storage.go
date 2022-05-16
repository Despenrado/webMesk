package teststore

import (
	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
)

type Storage struct {
	userRepository    *UserRepository
	chatRepository    *ChatRepository
	messageRepository *MessageRepository
}

func NewStorage() *Storage {
	s := &Storage{}
	s.userRepository = NewUserRepository(make(map[uint]*model.User), s)
	s.chatRepository = NewChatRepository(make(map[uint]*model.Chat), s)
	s.messageRepository = NewMessageRepository(make(map[uint]*model.Message), s)
	return s
}

func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	return NewUserRepository(make(map[uint]*model.User), s)
}

func (s *Storage) Chat() storage.ChatRepository {
	if s.userRepository != nil {
		return s.chatRepository
	}
	return NewChatRepository(make(map[uint]*model.Chat), s)
}

func (s *Storage) Message() storage.MessageRepository {
	if s.userRepository != nil {
		return s.messageRepository
	}
	return NewMessageRepository(make(map[uint]*model.Message), s)
}
