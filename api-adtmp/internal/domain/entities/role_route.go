package entities

type RoleRoute struct {
	RoleID  uint `gorm:"primaryKey"`
	RouteID uint `gorm:"primaryKey"`
}

func (RoleRoute) TableName() string {
	return "sys_role_route"
}
