// 依赖注入

package di

import (
	"net/http"

	"adtmp/internal/handler"
	"adtmp/internal/repository/store"
	"adtmp/internal/service"
	"adtmp/pkg/middleware"

	"gorm.io/gorm"
)

type controller interface {
	Register(mux *http.ServeMux)
}

func Denps(db *gorm.DB) []controller {
	var handlers []controller
	// 中间件
	mwProvider := middleware.New()

	pwdProvider := service.NewPwdProvider()
	tokenProvider := service.NewJWTProvider()

	// 用户认证
	authRepo := store.NewAuthRepository(db)
	authService := service.NewAuthService(pwdProvider, tokenProvider, authRepo)
	authHandler := handler.NewAuthHandler(mwProvider, authService)
	handlers = append(handlers, authHandler)

	// 用户接口
	// userRepo := store.NewUserRepository(db)
	uRepo := store.NewUserRepo(db)
	userService := service.NewUserService(pwdProvider, uRepo)
	userHandler := handler.NewUserHandler(mwProvider, userService)
	handlers = append(handlers, userHandler)

	// 角色接口
	rRepo := store.NewRoleRepo(db)
	roleService := service.NewRoleService(rRepo)
	roleHandler := handler.NewRoleHandler(mwProvider, roleService)
	handlers = append(handlers, roleHandler)

	return handlers
}
