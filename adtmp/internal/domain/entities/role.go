package entities

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string  `gorm:"size:50;uniqueIndex;not null"`
	Description string  `gorm:"size:512;not null"`
	Routes      []Route `gorm:"many2many:sys_role_route;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Role) TableName() string {
	return "sys_role"
}
