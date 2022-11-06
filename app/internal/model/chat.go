package model

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Chat struct {
	ID         uint    `json:"id,omitempty" gorm:"primaryKey"`
	ChatName   string  `json:"chat_name,omitempty"`
	MemberList []*User `json:"member_list,omitempty" gorm:"many2many:user_chat"`
}

type ChatFilter struct {
	UserID   uint   `schema:"user_id,omitempty"`
	ChatName string `schema:"chat_name,omitempty"`
	Skip     uint   `schema:"skip,omitempty"`
	Limit    uint   `schema:"limit,omitempty"`
}

func (c *Chat) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.ChatName, validation.Required),
		// validation.Field(&c.MemberList, validation.Length(2, 200)),
	)
}

func (c *Chat) BeforeCreate() error {
	return nil
}

func (c *Chat) Sanitize() {
	if c.MemberList != nil {
		for _, v := range c.MemberList {
			v.Sanitize()
		}
	}
}

func (c *Chat) CheckPermissions(userId uint) bool {
	if c.MemberList == nil {
		return false
	}
	log.Println(userId)
	for _, v := range c.MemberList {
		log.Println(v)
		if v.ID == userId {
			return true
		}
	}
	return false
}
