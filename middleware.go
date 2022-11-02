package cc4echo

import "github.com/labstack/echo/v4"

func Wrapper() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c = New(c)
			return next(c)
		}
	}
}
