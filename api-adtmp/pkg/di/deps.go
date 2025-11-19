// 依赖注入

package di

import (
	"net/http"

	"adtmp/internal/handler"
	"adtmp/internal/repository/store"
	"adtmp/internal/service"
	"adtmp/pkg/middleware"
	"adtmp/pkg/security"

	"gorm.io/gorm"
)

type controller interface {
	Register(mux *http.ServeMux)
}

func Denps(db *gorm.DB) []controller {
	var handlers []controller
	// 中间件
	mwProvider := middleware.New()
	mwApplier := NewMiddlewareApplier(mwProvider) // 创建中间件应用器
	pwdProvider := security.NewPwdProvider()
	tokenProvider := security.NewJWTProvider()

	// 用户认证
	authRepo := store.NewAuthRepository(db)
	authService := service.NewAuthService(pwdProvider, tokenProvider, authRepo)
	authHandler := handler.NewAuthHandler(authService)
	handlers = append(handlers, &authHandlerWithMw{authHandler, mwApplier})

	// 用户接口
	// userRepo := store.NewUserRepository(db)
	uRepo := store.NewUserRepo(db)
	userService := service.NewUserService(pwdProvider, uRepo)
	userHandler := handler.NewUserHandler(userService)
	handlers = append(handlers, &userHandlerWithMw{userHandler, mwApplier})

	// 角色接口
	rRepo := store.NewRoleRepo(db)
	roleService := service.NewRoleService(rRepo)
	roleHandler := handler.NewRoleHandler(roleService)
	handlers = append(handlers, &roleHandlerWithMw{roleHandler, mwApplier})

	return handlers
}
