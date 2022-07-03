package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Message struct {
	ID          uint              `json:"_id,omitempty" gorm:"primaryKey"`
	UserID      uint              `json:"user_id,omitempty" gorm:"index"`
	User        *User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChatID      uint              `json:"chat_id,omitempty" gorm:"index"`
	Chat        *Chat             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DateTime    time.Time         `json:"date_time,omitempty"`
	ReadBy      []uint            `json:"read_by,omitempty" gorm:"type:bigint[]"`
	MessageData map[string]string `json:"message_data,omitempty" gorm:"type:jsonb"`
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
