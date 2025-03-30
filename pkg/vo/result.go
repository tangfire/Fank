package vo

import (
	bizErr "Fank/internal/error"
	"errors"
	"github.com/labstack/echo/v4"
	"time"
)

type Result struct {
	*bizErr.Err
	Data      interface{} `json:"data"`
	RequestId interface{} `json:"requestId"`
	TimeStamp interface{} `json:"timeStamp"`
}

// Success 成功返回
func Success(data interface{}, c echo.Context) Result {
	return Result{
		Err:       nil,
		Data:      data,
		RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
		TimeStamp: time.Now().Unix(),
	}
}

// Fail 失败返回
func Fail(data interface{}, err error, c echo.Context) Result {
	var newBizErr *bizErr.Err
	if ok := errors.As(err, &newBizErr); ok {
		return Result{
			Err:       newBizErr,
			Data:      data,
			RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
			TimeStamp: time.Now().Unix(),
		}
	}

	return Result{
		Err:       bizErr.New(bizErr.ServerError),
		Data:      data,
		RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
		TimeStamp: time.Now().Unix(),
	}
}
