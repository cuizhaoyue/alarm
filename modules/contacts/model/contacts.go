package model

import (
	"time"
)

// Contacts 联系人组
type Contacts struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	CreatedAtTs int64     `gorm:"-"`
	UpdatedAtTs int64     `gorm:"-"`
}

func (*Contacts) TableName() string {
	return "contacts"
}

// ContactUserRelation 联系人和用户的关联结构
type ContactUserRelation struct {
	Id        int       `gorm:"primaryKey"`
	ContactId int       // 联系人组id
	UserId    string    // 用户id
	Username  string    // 用户名称
	Email     string    // 用户邮箱
	Telephone string    // 用户电话
	CreatedAt time.Time // 加入用户组的时间
}

func (*ContactUserRelation) TableName() string {
	return "contact_user_relation"
}

// ContactList 联系人组列表
type ContactList struct {
	Total int64
	Items []*Contacts
}
