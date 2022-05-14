package psql

import (
	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db                *gorm.DB
	userRepository    *UserRepository
	chatRepository    *ChatRepository
	messageRepository *MessageRepository
}

func NewConnection(dsn string, autoMigration bool) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if autoMigration {
		db.AutoMigrate(&model.User{}, &model.Chat{}, &model.Message{})
	}
	return db, err
}

func NewStorage(db *gorm.DB) *Storage {
	s := &Storage{
		db: db,
	}
	s.userRepository = NewUserRepository(s)
	s.chatRepository = NewChatRepository(s)
	s.messageRepository = NewMessageRepository(s)
	return s
}

func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	return NewUserRepository(s)
}

// func (s *Storage) Chat() storage.ChatRepository {
// 	if s.userRepository != nil {
// 		return s.chatRepository
// 	}
// 	return NewChatRepository(s)
// }

// func (s *Storage) Message() storage.MessageRepository {
// 	if s.userRepository != nil {
// 		return s.messageRepository
// 	}
// 	return NewMessageRepository(s)
// }
