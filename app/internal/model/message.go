package model

import (
	"time"

	"github.com/Despenrado/webMesk/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	pq "github.com/lib/pq"
)

type Message struct {
	ID          uint          `json:"id,omitempty" gorm:"primaryKey"`
	UserID      uint          `json:"user_id,omitempty" gorm:"index"`
	User        *User         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChatID      uint          `json:"chat_id,omitempty" gorm:"index"`
	Chat        *Chat         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DateTime    time.Time     `json:"date_time,omitempty"`
	ReadBy      pq.Int64Array `json:"read_by,omitempty" gorm:"type:bigint[]"`
	MessageData utils.JSONB   `json:"message_data,omitempty" gorm:"type:jsonb"`
}

type MessageFilter struct {
	UserID                  uint      `schema:"user_id,omitempty"`
	ChatID                  uint      `schema:"chat_id,omitempty"`
	DateTime                time.Time `schema:"date_time,omitempty"`
	DateTimeComparationType string    `schema:"date_time_comparation_type,omitempty"`
	UnreadOnly              bool      `schema:"unread_only,omitempty"`
	OwnerOnly               bool      `schema:"owner_only,omitempty"`
	Skip                    uint      `schema:"skip,omitempty"`
	Limit                   uint      `schema:"limit,omitempty"`
}

func (m *MessageFilter) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.UserID, validation.Required),
	)
}

func (m *Message) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.UserID, validation.Required),
		validation.Field(&m.ChatID, validation.Required),
	)
}

func (m *Message) BeforeCreate() error {
	m.DateTime = time.Now()
	m.ReadBy = make(pq.Int64Array, 0)
	return nil
}

func (m *Message) Sanitize() {
	if m.User != nil {
		m.User.Sanitize()
	}
	if m.Chat != nil {
		m.Chat.Sanitize()
	}
}

func (m *Message) CheckPermissions(userId uint) bool {
	if m.UserID == userId {
		return true
	}
	if m.Chat == nil {
		return false
	}
	return m.Chat.CheckPermissions(userId)
}
