package psql

import (
	"context"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
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
	res := mr.storage.db.WithContext(ctx).Create(message)
	if res.RowsAffected != 1 {
		return message, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return message, nil
}

func (mr *MessageRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error) {
	messages := []model.Message{}
	res := mr.storage.db.WithContext(ctx).Limit(limit).Offset(skip).Find(&messages)
	return messages, res.Error
}

func (mr *MessageRepository) FindById(ctx context.Context, id uint) (*model.Message, error) {
	message := &model.Message{}
	res := mr.storage.db.WithContext(ctx).Preload("User").Preload("Chat").First(message, id)
	if res.RowsAffected != 1 {
		return nil, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return message, nil
}

func (mr *MessageRepository) Update(ctx context.Context, message *model.Message) (*model.Message, error) {
	res := mr.storage.db.WithContext(ctx).Save(message)
	if res.RowsAffected != 1 {
		return message, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return message, res.Error
}

func (mr *MessageRepository) Delete(ctx context.Context, id uint) error {
	res := mr.storage.db.WithContext(ctx).Delete(&model.Message{ID: id})
	if res.RowsAffected != 1 {
		return utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return res.Error
}

func (mr *MessageRepository) FindByUserId(ctx context.Context, id uint) ([]model.Message, error) {
	messages := []model.Message{}
	res := mr.storage.db.WithContext(ctx).Where(map[string]interface{}{"user_id": id}).Order("date_time esc").Find(&messages)
	return messages, res.Error
}

func (mr *MessageRepository) FindByChatId(ctx context.Context, id uint) ([]model.Message, error) {
	messages := []model.Message{}
	res := mr.storage.db.WithContext(ctx).Where(map[string]interface{}{"chat_id": id}).Order("date_time esc").Find(&messages)
	return messages, res.Error
}

func (mr *MessageRepository) FindByChatIdAndAfterDateTime(ctx context.Context, dateTime time.Time) ([]model.Message, error) {
	messages := []model.Message{}
	res := mr.storage.db.WithContext(ctx).Where("date_time >= ?", dateTime).Order("date_time esc").Find(&messages)
	return messages, res.Error
}
