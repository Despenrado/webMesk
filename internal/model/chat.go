package model

type Chat struct {
	ID         uint   `json:"_id,omitempty" gorm:"primaryKey"`
	ChatName   string `json:"chat_name,omitempty" gorm:"index"`
	MemberList []User `json:"member_list,omitempty" gorm:"many2many:user_chat"`
}
