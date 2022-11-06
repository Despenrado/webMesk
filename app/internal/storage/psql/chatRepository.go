package psql

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
	"gorm.io/gorm/clause"
)

type ChatRepository struct {
	storage *Storage
}

func NewChatRepository(storage *Storage) *ChatRepository {
	return &ChatRepository{
		storage: storage,
	}
}

func (cr *ChatRepository) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	res := cr.storage.DB.WithContext(ctx).Omit("MemberList.*").Create(chat)
	if res.RowsAffected != 1 {
		return chat, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return chat, nil
}

func (cr *ChatRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.Chat, error) {
	chats := []model.Chat{}
	res := cr.storage.DB.WithContext(ctx).Limit(limit).Offset(skip).Find(&chats)
	return chats, res.Error
}

func (cr *ChatRepository) FindById(ctx context.Context, id uint) (*model.Chat, error) {
	chat := &model.Chat{}
	res := cr.storage.DB.WithContext(ctx).Preload("MemberList").First(chat, id)
	if res.RowsAffected != 1 {
		return nil, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return chat, nil
}

func (cr *ChatRepository) Update(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	tx := cr.storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}
	res := tx.WithContext(ctx).Omit("MemberList.*").Save(chat)
	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		tx.Rollback()
		return nil, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	err := tx.WithContext(ctx).Model(&chat).Association("MemberList").Replace(chat.MemberList)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return chat, nil
}

func (cr *ChatRepository) Delete(ctx context.Context, id uint) error {
	res := cr.storage.DB.WithContext(ctx).Select(clause.Associations).Delete(&model.Chat{ID: id})
	if res.RowsAffected < 1 {
		return utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return nil
}

func (cr *ChatRepository) FindByUserId(ctx context.Context, id uint) ([]model.Chat, error) {
	// chat := &model.Chat{}
	// res := cr.storage.db.WithContext(ctx).Where(map[string]interface{}{"user_name": userName}).First(chat)
	// if res.RowsAffected != 1 {
	// 	return nil, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	// }
	chats := []model.Chat{}
	res := cr.storage.DB.WithContext(ctx).Preload("MemberList", map[string]interface{}{"id": id}).Find(&chats)
	return chats, res.Error
}

func (cr *ChatRepository) FilterChat(ctx context.Context, chatFilter *model.ChatFilter) ([]model.Chat, error) {
	query := cr.storage.DB.WithContext(ctx)
	if chatFilter.ChatName != "" {
		query = query.Where("chat_name LIKE ?", "%"+chatFilter.ChatName+"%")
	}
	if chatFilter.UserID != 0 {
		query = query.Joins("JOIN user_chat ON chats.id = user_chat.chat_id").Where("user_chat.user_id = ?", chatFilter.UserID)
	}
	query = query.Offset(int(chatFilter.Skip))
	if chatFilter.Limit != 0 {
		query = query.Limit(int(chatFilter.Limit))
	}
	chats := []model.Chat{}
	res := query.Debug().Find(&chats)
	return chats, res.Error
}
