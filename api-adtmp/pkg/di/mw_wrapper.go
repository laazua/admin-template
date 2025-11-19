package di

import (
	"net/http"

	"adtmp/pkg/middleware"
)

type middlewareApplier struct {
	mw middleware.Mw
}

func NewMiddlewareApplier(mw middleware.Mw) *middlewareApplier {
	return &middlewareApplier{mw: mw}
}

// Apply 应用中间件到处理器
func (ma *middlewareApplier) Apply(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// WithAuth 便捷方法：应用认证中间件
func (ma *middlewareApplier) WithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return ma.Apply(handler, ma.mw.Auth)
}

// WithMws 应用多个中间件
func (ma *middlewareApplier) WithMws(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	return ma.Apply(handler, middlewares...)
}
