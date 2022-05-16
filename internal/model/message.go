package model

import (
	"time"
)

type Message struct {
	ID          uint              `json:"_id,omitempty" gorm:"primaryKey"`
	UserID      uint              `json:"user_id,omitempty" gorm:"index"`
	User        User              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChatID      uint              `json:"chatId,omitempty" gorm:"index"`
	Chat        Chat              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DateTime    time.Time         `json:"date_time,omitempty"`
	ReadBy      []uint            `json:"read_by,omitempty" gorm:"type:bigint[]"`
	MessageData map[string]string `json:"message_data,omitempty" gorm:"type:jsonb"`
}
