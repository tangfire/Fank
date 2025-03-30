package account

import (
	bizErr "Fank/internal/error"
	"Fank/pkg/serve/controller/account/dto"
	"Fank/pkg/vo"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterAcc(c echo.Context) error {
	req := new(dto.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

}
