package response

import (
	"errors"
	"myapp/pkg/e"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Code      int       `json:"code"`
	Msg       string    `json:"msg"`
	Data      T         `json:"data,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type PageResponse[T any] struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
	List     []T   `json:"list"`
}

func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Response[T]{
		Code:      e.CodeSuccess.Code,
		Msg:       e.CodeSuccess.Msg,
		Data:      data,
		Timestamp: time.Now(),
	})
}

func PageSuccess[T any](c *gin.Context, list []T, total int64, page, pageSize int) {
	Success(c, PageResponse[T]{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

func HandleError(c *gin.Context, err error) {
	var appError *e.AppError
	if errors.As(err, &appError) {
		c.JSON(http.StatusOK, Response[any]{
			Code: appError.Code,
			Msg:  appError.Msg,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, Response[any]{
		Code: e.CodeInternalError.Code,
		Msg:  err.Error(),
	})
}

func Wrapper[T any](handler func(c *gin.Context) (T, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handler(c)
		if err != nil {
			HandleError(c, err)
			return
		}
		Success(c, data)
	}
}
