package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Chat struct {
	ID         uint    `json:"_id,omitempty" gorm:"primaryKey"`
	ChatName   string  `json:"chat_name,omitempty" gorm:"index"`
	MemberList []*User `json:"member_list,omitempty" gorm:"many2many:user_chat"`
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
