package middleware

import "net/http"

type Mw interface {
	Auth(next http.HandlerFunc) http.HandlerFunc
}

type mw struct{}

func New() *mw {
	return &mw{}
}

func (mw *mw) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 认证逻辑
		next(w, r)
	}
}
