package psql

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
)

type MessageRepository struct {
	storage *Storage
}

func NewMessageRepository(storage *Storage) *MessageRepository {
	return &MessageRepository{
		storage: storage,
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
