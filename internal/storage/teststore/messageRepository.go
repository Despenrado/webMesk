package teststore

import (
	"context"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
)

type MessageRepository struct {
	messageDB map[uint]*model.Message
	storage   *Storage
}

func NewMessageRepository(db map[uint]*model.Message, storage *Storage) *MessageRepository {
	return &MessageRepository{
		messageDB: db,
		storage:   storage,
	}
}

func (mr *MessageRepository) Create(ctx context.Context, message *model.Message) (*model.Message, error) {
	return nil, nil
}

func (mr *MessageRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error) {
	return nil, nil
}

func (mr *MessageRepository) FindById(ctx context.Context, id uint) (*model.Message, error) {
	return nil, nil
}

func (mr *MessageRepository) Update(ctx context.Context, message *model.Message) (*model.Message, error) {
	return nil, nil
}

func (mr *MessageRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func (mr *MessageRepository) FindByUserId(ctx context.Context, id uint) ([]model.Message, error) {
	return nil, nil
}

func (mr *MessageRepository) FindByChatId(ctx context.Context, id uint) ([]model.Message, error) {
	return nil, nil
}

func (mr *MessageRepository) FindByChatIdAndAfterDateTime(ctx context.Context, dateTime time.Time) ([]model.Message, error) {
	return nil, nil
}
