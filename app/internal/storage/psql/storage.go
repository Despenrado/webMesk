package psql

import (
	"log"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var counter = 0

type Storage struct {
	DB                *gorm.DB
	userRepository    *UserRepository
	chatRepository    *ChatRepository
	messageRepository *MessageRepository
}

func NewConnection(config *utils.PostgreSQLConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + config.Master.Host + " port=" + config.Master.Port + " dbname=" + config.Master.DBName +
			" user=" + config.Master.User + " password=" + config.Master.Password +
			" sslmode=" + config.Master.SSLMode + " TimeZone=" + config.Master.TimeZone,
		PreferSimpleProtocol: true,
	}), &gorm.Config{SkipDefaultTransaction: true})
	dialector := []gorm.Dialector{
		postgres.New(postgres.Config{
			DSN: "host=" + config.Master.Host + " port=" + config.Master.Port + " dbname=" + config.Master.DBName +
				" user=" + config.Master.User + " password=" + config.Master.Password +
				" sslmode=" + config.Master.SSLMode + " TimeZone=" + config.Master.TimeZone,
			PreferSimpleProtocol: true,
		})}
	log.Println("------------------------------------------------------------------------------------------")
	log.Println(config)
	log.Println(config.Slave)
	for _, v := range config.Slave {
		log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		dialector = append(dialector, postgres.New(postgres.Config{
			DSN: "host=" + v.Host + " port=" + v.Port + " dbname=" + config.Master.DBName +
				" user=" + config.Master.User + " password=" + config.Master.Password +
				" sslmode=" + config.Master.SSLMode + " TimeZone=" + config.Master.TimeZone,
			PreferSimpleProtocol: true,
		}),
		)
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: dialector,
		Policy:   RoundRobinPolicy{counter: &counter},
	}).
		SetConnMaxIdleTime(100).
		SetMaxOpenConns(200))

	db.AutoMigrate(&model.User{}, &model.Chat{}, &model.Message{})
	return db, err
}

type RoundRobinPolicy struct {
	counter *int
}

func (p RoundRobinPolicy) Resolve(connPools []gorm.ConnPool) gorm.ConnPool {
	if *p.counter >= len(connPools)-1 {
		*p.counter = 0
	} else {
		*p.counter++
	}
	return connPools[*p.counter]
}

func NewStorage(db *gorm.DB) *Storage {
	s := &Storage{
		DB: db,
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

func (s *Storage) Chat() storage.ChatRepository {
	if s.userRepository != nil {
		return s.chatRepository
	}
	return NewChatRepository(s)
}

func (s *Storage) Message() storage.MessageRepository {
	if s.userRepository != nil {
		return s.messageRepository
	}
	return NewMessageRepository(s)
}
