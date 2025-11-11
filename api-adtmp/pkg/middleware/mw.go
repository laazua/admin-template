package middleware

import (
	"adtmp/internal/domain/entities"
	"adtmp/pkg/api"
	"adtmp/pkg/security"
	"context"
	"log/slog"
	"net/http"
	"strings"
)

type Mw interface {
	Auth(next http.HandlerFunc) http.HandlerFunc
}

// middleware 中间件实现
type middleware struct {
	// 可以在这里添加中间件依赖，如 JWT 服务、日志服务等
}

// New 创建中间件实例
func New() Mw {
	return &middleware{}
}

type contextKey string

const ContextUserKey = contextKey("token")

// Auth 认证中间件实现
func (m *middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		const BearerSchema = "Bearer "
		if !strings.HasPrefix(tokenStr, BearerSchema) {
			slog.Error("token格式错误", slog.Any("Token", tokenStr))
			api.Failure(w, http.StatusUnauthorized, "token格式错误")
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(tokenStr, BearerSchema))

		tokenProvider := security.NewJWTProvider()
		claims, err := tokenProvider.ValidateToken(tokenString)
		if err != nil {
			slog.Error("token认证失败", slog.String("Error", err.Error()))
			api.Failure(w, http.StatusUnauthorized, "token认证失败")
			return
		}
		// 将用户放入 context
		tokenclaims := &entities.TokenClaims{
			Name:  claims.Name,
			Email: claims.Email,
		}
		ctx := context.WithValue(r.Context(), ContextUserKey, tokenclaims)
		next(w, r.WithContext(ctx))
	}
}
