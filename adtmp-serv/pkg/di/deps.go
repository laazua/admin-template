// 依赖注入

package di

import (
	"net/http"

	"adtmp/internal/handler"
	"adtmp/internal/repository/store"
	"adtmp/internal/service"

	"gorm.io/gorm"
)

type controller interface {
	Register(mux *http.ServeMux)
}

func Denps(db *gorm.DB) []controller {
	var handlers []controller

	pwdProvider := service.NewPwdProvider()
	tokenProvider := service.NewJWTProvider()

	// 用户认证
	authRepo := store.NewAuthRepository(db)
	authService := service.NewAuthService(pwdProvider, tokenProvider, authRepo)
	authHandler := handler.NewAuthHandler(authService)
	handlers = append(handlers, authHandler)

	// 用户接口
	// userRepo := store.NewUserRepository(db)
	uRepo := store.NewUserRepo(db)
	userService := service.NewUserService(pwdProvider, uRepo)
	userHandler := handler.NewUserHandler(userService)
	handlers = append(handlers, userHandler)

	return handlers
}
