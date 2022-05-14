package model

type Chat struct {
	ID         uint   `json:"_id,omitempty" gorm:"primaryKey"`
	ChatName   string `json:"chatName,omitempty"`
	MemberList []User `json:"memberList,omitempty" gorm:"many2many:user_chat"`
}
