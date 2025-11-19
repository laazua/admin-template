package entities

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	Path      string  `gorm:"size:100;not null"`
	Name      string  `gorm:"size:100;not null"`
	Component string  `gorm:"size:100;not null"`
	Title     string  `gorm:"size:100;not null"` // meta.title
	Icon      string  `gorm:"size:100"`          // meta.icon
	ParentID  *uint   `gorm:"index"`             // 层级结构
	Children  []Route `gorm:"foreignKey:ParentID"`
}

func (Route) TableName() string {
	return "sys_route"
}
