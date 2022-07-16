package psql

import (
	"context"
	"log"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
	"gorm.io/gorm"
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
	res := mr.storage.db.WithContext(ctx).First(message, id)
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

func (mr *MessageRepository) FilterMessage(ctx context.Context, messageFilter *model.MessageFilter) ([]model.Message, error) {
	query := mr.storage.db.WithContext(ctx)
	if !messageFilter.DateTime.IsZero() && messageFilter.DateTimeComparation != "" {
		filter := "date_time " + messageFilter.DateTimeComparation + " ?"
		query = query.Where(filter, messageFilter.DateTime)
	}
	if messageFilter.UserID != 0 {
		query = query.Joins("JOIN chats ON messages.chat_id = chats.id").Joins("JOIN user_chat ON chats.id = user_chat.chat_id").Where("messages.user_id = ? OR user_chat.user_id = ?", messageFilter.UserID, messageFilter.UserID)
	}
	if messageFilter.ChatID != 0 {
		query = query.Where("chat_id = ?", messageFilter.ChatID)
	}
	query = query.Offset(int(messageFilter.Skip))
	if messageFilter.Limit != 0 {
		query = query.Limit(int(messageFilter.Limit))
	}
	log.Println(query)
	messages := []model.Message{}
	res := query.Find(&messages)
	return messages, res.Error
}

func (mr *MessageRepository) MarkAsRead(ctx context.Context, id uint, user_id uint) error {
	log.Println(user_id)
	// res := mr.storage.db.Model(&model.Message{}).Where("id = ?", id).Update("read_by", gorm.Expr("array_append(read_by, ?)", user_id))
	// res := mr.storage.db.Model(&model.Message{}).Where("id = ?", id).Update("read_by", gorm.Expr("(SELECT array_agg(distinct e) FROM unnest(read_by || ?) AS e)", pq.Int64Array([]int64{int64(user_id)})))
	res := mr.storage.db.Model(&model.Message{}).Where("id = ?", id).Update("read_by", gorm.Expr("(SELECT array_agg(distinct e) FROM unnest(read_by || (SELECT user_chat.user_id FROM user_chat, messages WHERE user_chat.chat_id = messages.chat_id AND messages.id = ? AND user_chat.user_id = ?)) AS e)", id, user_id))
	return res.Error
}
