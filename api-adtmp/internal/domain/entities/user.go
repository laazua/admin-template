package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `grom:"size:50;not null"`
	Email    string `gorm:"size:100;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`
	Phone    string `gorm:"size:25;not null"`
	Roles    []Role `gorm:"many2many:sys_user_role;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Avatar   string `gorm:"size:255;not null"`
}

func (User) TableName() string {
	return "sys_user"
}
