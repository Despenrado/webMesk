package model

import (
	"time"

	"github.com/Despenrado/webMesk/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	pq "github.com/lib/pq"
)

type Message struct {
	ID          uint          `json:"_id,omitempty" gorm:"primaryKey"`
	UserID      uint          `json:"user_id,omitempty" gorm:"index"`
	User        *User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChatID      uint          `json:"chat_id,omitempty" gorm:"index"`
	Chat        *Chat         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DateTime    time.Time     `json:"date_time,omitempty"`
	ReadBy      pq.Int64Array `json:"read_by,omitempty" gorm:"type:bigint[]"`
	MessageData utils.JSONB   `json:"message_data,omitempty" gorm:"type:jsonb"`
}

type MessageFilter struct {
	UserID              uint      `json:"user_id,omitempty"`
	ChatID              uint      `json:"chat_id,omitempty"`
	DateTime            time.Time `json:"date_time,omitempty"`
	DateTimeComparation string    `json:"date_time_comparation,omitempty"`
	Skip                uint      `json:"skip,omitempty"`
	Limit               uint      `json:"limit,omitempty"`
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
