// 由./*_provider.go实现
package security

import (
	"adtmp/pkg/internal/domain/entities"
)

// - pwd_provider.go 实现
type PwdProvider interface {
	HashPwd(password string) (string, error)
	VerifyPwd(hashedPassword, password string) bool
}

// - token_provider.go 实现
type TokenProvider interface {
	GenerateToken(user *entities.User) (string, error)
	ValidateToken(tokenString string) (*entities.TokenClaims, error)
}
