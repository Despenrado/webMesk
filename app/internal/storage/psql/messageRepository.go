package psql

import (
	"context"

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
	res := mr.storage.DB.WithContext(ctx).Create(message)
	if res.RowsAffected != 1 {
		return message, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return message, nil
}

func (mr *MessageRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error) {
	messages := []model.Message{}
	res := mr.storage.DB.WithContext(ctx).Limit(limit).Offset(skip).Find(&messages)
	return messages, res.Error
}

func (mr *MessageRepository) FindById(ctx context.Context, id uint) (*model.Message, error) {
	message := &model.Message{}
	res := mr.storage.DB.WithContext(ctx).Preload("Chat").Preload("Chat.MemberList").First(message, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return message, nil
}

func (mr *MessageRepository) Update(ctx context.Context, message *model.Message) (*model.Message, error) {
	res := mr.storage.DB.WithContext(ctx).Where("id = ? AND user_id = ?", message.ID, message.UserID).Omit("date_time", "read_by", "user_id", "chat_id").Updates(message)
	if res.RowsAffected != 1 {
		return message, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return message, res.Error
}

func (mr *MessageRepository) Delete(ctx context.Context, message *model.Message) error {
	res := mr.storage.DB.WithContext(ctx).Delete(&model.Message{ID: message.ID, UserID: message.UserID})
	if res.RowsAffected != 1 {
		return utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return res.Error
}

func (mr *MessageRepository) FilterMessage(ctx context.Context, messageFilter *model.MessageFilter) ([]model.Message, error) {
	query := mr.storage.DB.WithContext(ctx)
	if !messageFilter.DateTime.IsZero() && messageFilter.DateTimeComparationType != "" {
		filter := "date_time " + messageFilter.DateTimeComparationType + " ?"
		query = query.Where(filter, messageFilter.DateTime)
	}
	if messageFilter.UserID != 0 {
		if messageFilter.OwnerOnly {
			query = query.Where("messages.user_id = ?", messageFilter.UserID)
		} else {
			query = query.Joins("JOIN chats ON chats.id = messages.chat_id").Joins("JOIN user_chat ON chats.id = user_chat.chat_id").Where("user_chat.user_id = ?", messageFilter.UserID)
		}
	}
	if messageFilter.UnreadOnly && messageFilter.UserID > 0 {
		query = query.Where("NOT (? = ANY(read_by))", messageFilter.UserID)
	}
	if messageFilter.ChatID != 0 {
		query = query.Where("chat_id = ?", messageFilter.ChatID)
	}
	query = query.Offset(int(messageFilter.Skip))
	if messageFilter.Limit != 0 {
		query = query.Limit(int(messageFilter.Limit))
	}
	messages := []model.Message{}
	res := query.Find(&messages)
	return messages, res.Error
}

func (mr *MessageRepository) MarkAsRead(ctx context.Context, id uint, user_id uint) error {
	res := mr.storage.DB.Model(&model.Message{}).Where("id = ?", id).Update("read_by", gorm.Expr("(SELECT array_agg(distinct e) FROM unnest(read_by || (SELECT user_chat.user_id FROM user_chat, messages WHERE user_chat.chat_id = messages.chat_id AND messages.id = ? AND user_chat.user_id = ?)) AS e)", id, user_id))
	return res.Error
}
