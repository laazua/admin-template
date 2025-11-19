package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"adtmp/internal/domain/entities"
	"adtmp/pkg/api"
	"adtmp/pkg/security"
)

type Mw interface {
	Auth(next http.HandlerFunc) http.HandlerFunc
	Limit(next http.HandlerFunc) http.HandlerFunc
}

// middleware 中间件实现
type middleware struct {
	// 可以在这里添加中间件依赖，如 JWT 服务、日志服务等
}

// New 创建中间件实例
// 返回具体实现而不是单个大接口，便于扩展和在 di 层按需依赖更小的接口或具体类型。
// *middleware 仍然实现当前的 Mw 接口，di 包可以继续使用或改为依赖更细粒度的接口。
func New() *middleware {
	return &middleware{}
}

func (m *middleware) Limit(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// 这里实现限流逻辑，例如基于 IP 地址的简单限流
		// 这只是一个示例，实际应用中应使用更复杂的限流算法和存储
		// ip := r.RemoteAddr
		// if !rateLimiter.Allow(ip) {
		// 	api.Failure(w, http.StatusTooManyRequests, "请求过于频繁，请稍后再试")
		// 	return
		// }
		next(w, r)
	}
}

// =============== 认证中间件实现 =================
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
