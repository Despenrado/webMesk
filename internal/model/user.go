package model

import (
	"errors"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint      `json:"_id,omitempty" gorm:"primaryKey"`
	UserName   string    `json:"user_name,omitempty" gorm:"index,unique"`
	Email      string    `json:"email,omitempty" gorm:"index,unique"`
	Password   string    `json:"password,omitempty"`
	SessionId  string    `json:"sessionId,omitempty"`
	LastOnline time.Time `json:"last_online,omitempty"`
}

type UserFilter struct {
	UserName  string `json:"user_name,omitempty"`
	Email     string `json:"email,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	Skip      uint   `json:"skip,omitempty"`
	Limit     uint   `json:"limit,omitempty"`
}

func (u *User) Validate() error {
	fmt.Println("______________________user Validation___________________")
	return validation.ValidateStruct(
		u,
		validation.Field(&u.UserName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) <= 0 {
		return errors.New("password can not be empty")
	}
	enc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(enc)
	return nil
}

func (u *User) VerifyPassword(p string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) == nil
}

func (u *User) Sanitize() {
	u.Password = ""
}
