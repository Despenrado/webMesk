package impl

import (
	"context"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
)

type MessageService struct {
	storage storage.Storage
}

func NewMessageService(storage storage.Storage) *MessageService {
	return &MessageService{
		storage: storage,
	}
}

func (ms *MessageService) Create(ctx context.Context, message *model.Message) (*model.Message, error) {
	if err := message.BeforeCreate(); err != nil {
		return message, err
	}
	message, err := ms.storage.Message().Create(ctx, message)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (ms *MessageService) ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error) {
	return ms.storage.Message().ReadAll(ctx, skip, limit)
}

func (ms *MessageService) FindById(ctx context.Context, id uint) (*model.Message, error) {
	message, err := ms.storage.Message().FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (ms *MessageService) Update(ctx context.Context, message *model.Message) (*model.Message, error) {
	message, err := ms.storage.Message().Update(ctx, message)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (ms *MessageService) Delete(ctx context.Context, id uint) error {
	return ms.storage.Message().Delete(ctx, id)
}

func (ms *MessageService) FindByUserId(ctx context.Context, id uint) ([]model.Message, error) {
	messages, err := ms.storage.Message().FindByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (ms *MessageService) FindByChatId(ctx context.Context, id uint) ([]model.Message, error) {
	messages, err := ms.storage.Message().FindByChatId(ctx, id)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (ms *MessageService) FindByChatIdAndAfterDateTime(ctx context.Context, dateTime time.Time) ([]model.Message, error) {
	messages, err := ms.storage.Message().FindByChatIdAndAfterDateTime(ctx, dateTime)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
