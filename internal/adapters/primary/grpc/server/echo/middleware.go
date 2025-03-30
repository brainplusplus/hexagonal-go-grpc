package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"runtime/debug"
)

func Recoverer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if rvr := recover(); rvr != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"error": fmt.Sprintf("%v", rvr),
					"stack": string(debug.Stack()),
				})
			}
		}()
		return next(c)
	}
}
