package model

import (
	"time"
)

type Message struct {
	ID          uint              `json:"_id,omitempty" gorm:"primaryKey"`
	UserID      uint              `json:"userId,omitempty"`
	User        User              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ChatID      uint              `json:"chatId,omitempty"`
	Chat        Chat              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DateTime    time.Time         `json:"dateTime,omitempty"`
	ReadBy      []uint            `json:"readBy,omitempty" gorm:"type:bigint[]"`
	MessageData map[string]string `json:"messageData,omitempty" gorm:"type:jsonb"`
}
