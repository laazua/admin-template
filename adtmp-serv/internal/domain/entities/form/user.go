package form

import "adtmp/internal/domain/entities"

// 登录
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 新增用户
type UserCreate struct {
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Phone    string          `json:"phone,omitempty"`
	Avatar   string          `json:"avatar,omitempty"`
	Roles    []entities.Role `json:"roles"`
}

// 更新用户
type UserUpdate struct {
	UserCreate
}

// -------------------------------------------------------------

// 定义用户表单接口
type UserForm interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAvatar() string
	GetRoles() []entities.Role
}

// 为表单类型实现方法
func (f *UserCreate) GetName() string           { return f.Name }
func (f *UserCreate) GetEmail() string          { return f.Email }
func (f *UserCreate) GetPassword() string       { return f.Password }
func (f *UserCreate) GetAvatar() string         { return f.Avatar }
func (f *UserCreate) GetRoles() []entities.Role { return f.Roles }

func (f *UserUpdate) GetName() string           { return f.Name }
func (f *UserUpdate) GetEmail() string          { return f.Email }
func (f *UserUpdate) GetPassword() string       { return f.Password }
func (f *UserUpdate) GetAvatar() string         { return f.Avatar }
func (f *UserUpdate) GetRoles() []entities.Role { return f.Roles }

// 转换函数
func ToUserDb[T UserForm](t T) *entities.User {
	if any(t) == nil {
		return nil
	}
	form := any(t).(UserForm)
	return &entities.User{
		Name:     form.GetName(),
		Email:    form.GetEmail(),
		Password: form.GetPassword(),
		Avatar:   form.GetAvatar(),
		Roles:    form.GetRoles(),
	}
}
