package service

import (
	"context"
	"errors"

	"adtmp/internal/domain/entities/form"
	"adtmp/internal/domain/repositories"
)

var errPwd = errors.New("invalid password")

type authService struct {
	pwdProvider   PwdProvider
	tokenProvider TokenProvider
	authRepo      repositories.AuthRepository // 依赖接口
}

// 返回接口
func NewAuthService(pwdProvider PwdProvider, tokenProvider TokenProvider, authRepo repositories.AuthRepository) AuthService {
	return &authService{pwdProvider: pwdProvider, tokenProvider: tokenProvider, authRepo: authRepo}
}

func (authService *authService) AuthUser(ctx context.Context, user *form.UserLogin) (string, error) {
	repoUser, err := authService.authRepo.Auth(ctx, user)
	if err != nil {
		return "", err
	}
	// 验证密码
	if !authService.pwdProvider.VerifyPwd(repoUser.Password, user.Password) {
		return "", errPwd
	}
	// 生成token
	token, err := authService.tokenProvider.GenerateToken(&repoUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (authService *authService) GetUser(ctx context.Context, name string) {}
