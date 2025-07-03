package errors

import (
	"fmt"
	"net/http"
)

// AppError アプリケーション共通のエラー構造体
type AppError struct {
	Code       int         `json:"code"`              // アプリケーション固有のエラーコード
	StatusCode int         `json:"status_code"`       // HTTPステータスコード
	Message    string      `json:"message"`           // ユーザー向けメッセージ
	Details    interface{} `json:"details,omitempty"` // 詳細情報（任意）
	InnerError error       `json:"-"`                 // 元になったエラー（デバッグ用）
}

// Error errorインターフェースの実装
func (e *AppError) Error() string {
	if e.InnerError != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.InnerError)
	}
	return e.Message
}

// Wrap 既存のエラーをラップする
func (e *AppError) Wrap(err error) *AppError {
	e.InnerError = err
	return e
}

// 事前定義済みエラー
var (
	ErrInternalServer = &AppError{
		Code:       1000,
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
	}

	ErrNotFound = &AppError{
		Code:       1001,
		StatusCode: http.StatusNotFound,
		Message:    "Requested resource not found",
	}

	ErrUnauthorized = &AppError{
		Code:       1002,
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized access",
	}

	ErrInvalidInput = &AppError{
		Code:       1003,
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid input data",
	}
)

// New 新しいカスタムエラーを作成
func New(code, statusCode int, message string) *AppError {
	return &AppError{
		Code:       code,
		StatusCode: statusCode,
		Message:    message,
	}
}
