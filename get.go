package gecho

import (
	"github.com/labstack/echo/v4"
)

func Get[ReqT, ResT any](e *echo.Echo, path string, fn func(c echo.Context, req ReqT) (ResT, error), m ...echo.MiddlewareFunc) *echo.Route {
	return e.GET(path, func(c echo.Context) error {
		var rq ReqT
		if err := c.Bind(&rq); err != nil {
			return err
		}
		res, err := fn(c, rq)
		if err != nil {
			return err
		}
		return c.JSON(200, res)
	}, m...)
}