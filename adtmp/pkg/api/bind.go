package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// BindJSON 解析 JSON 请求体
func BindJSON(r *http.Request, v any) error {
	// body, err := io.ReadAll(api.Request.Body)
	// if err != nil {
	// 	return err
	// }
	// defer api.Request.Body.Close()

	// return json.Unmarshal(body, v)
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return errors.New("Content-Type must be application/json")
	}
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return validateRequiredFields(v)
}

// 独立的验证方法，可按需调用
func validateRequiredFields(v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := typ.Field(i)
		jsonTag := fieldType.Tag.Get("json")

		if strings.Contains(jsonTag, "omitempty") {
			continue
		}

		if fieldValue.IsZero() {
			// 使用 JSON 标签名作为错误信息更友好
			fieldName := fieldType.Name
			if jsonName := strings.Split(jsonTag, ",")[0]; jsonName != "" {
				fieldName = jsonName
			}
			return fmt.Errorf("请求体参数[%s]不能为空", fieldName)
		}
	}
	return nil
}
