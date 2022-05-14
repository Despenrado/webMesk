package model

import (
	"time"
)

type User struct {
	ID         uint      `json:"_id,omitempty" gorm:"primaryKey"`
	UserName   string    `json:"UserName,omitempty"`
	Email      string    `json:"Email,omitempty"`
	Password   string    `json:"Password,omitempty"`
	SessionId  string    `json:"SessionId,omitempty"`
	LastOnline time.Time `json:"LastOnline,omitempty"`
}
