package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type M map[string]any

// respFormat 统一响应格式
type respFormat struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

// JSON 发送 JSON 响应
func jSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 成功响应
func Success(w http.ResponseWriter, data any, message string) {
	jSON(w, http.StatusOK, respFormat{
		Code:      http.StatusOK,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

// 失败响应
func Failure(w http.ResponseWriter, code int, message string) {
	jSON(w, code, respFormat{
		Code:      code,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}
