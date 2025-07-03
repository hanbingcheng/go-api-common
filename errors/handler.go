package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleError Gin用のエラーハンドラー
func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *AppError:
		c.JSON(e.StatusCode, gin.H{
			"error": gin.H{
				"code":    e.Code,
				"message": e.Message,
				"details": e.Details,
			},
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"code":    ErrInternalServer.Code,
				"message": ErrInternalServer.Message,
			},
		})
	}
}

// ErrorResponse エラーレスポンスの構造体
type ErrorResponse struct {
	Error struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}

// NewErrorResponse エラーレスポンスを作成
func NewErrorResponse(err *AppError) *ErrorResponse {
	res := &ErrorResponse{}
	res.Error.Code = err.Code
	res.Error.Message = err.Message
	res.Error.Details = err.Details
	return res
}
