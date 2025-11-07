package service

import (
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type pwdProvider struct{}

func NewPwdProvider() *pwdProvider {
	return &pwdProvider{}
}

// Hash 生成密码哈希，自动生成盐
func (pwdProvider *pwdProvider) HashPwd(password string) (string, error) {
	// bcrypt 会自动生成随机盐
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Verify 验证密码
func (pwdProvider *pwdProvider) VerifyPwd(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			slog.Error("密钥验证出错", slog.String("Error", err.Error()))
			return false
		}
		return false
	}
	return true
}
