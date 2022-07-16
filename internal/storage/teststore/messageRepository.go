package teststore

import (
	"context"

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

func (mr *MessageRepository) FilterMessage(ctx context.Context, essageFilter *model.MessageFilter) ([]model.Message, error) {
	return nil, nil
}

func (mr *MessageRepository) MarkAsRead(ctx context.Context, id uint, user_id uint) error {
	return nil
}
