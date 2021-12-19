package middleware

import (
	"errors"
	"net/http"

	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/web"
	"github.com/labstack/echo/v4"
)

func ServeHTTP(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("X-API-KEY") == "RAHASIA" {
			return next(c)
		} else {
			helper.PanicIfError(errors.New("gagal"))
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data:   nil,
			}
			return c.JSON(http.StatusUnauthorized, webResponse)
		}
	}
}
