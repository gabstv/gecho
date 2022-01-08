package gecho

import (
	"context"

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

func Post[ReqT, ResT any](e *echo.Echo, path string, fn func(c echo.Context, req ReqT) (ResT, error), m ...echo.MiddlewareFunc) *echo.Route {
	return e.POST(path, func(c echo.Context) error {
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

func Put[ReqT, ResT any](e *echo.Echo, path string, fn func(c echo.Context, req ReqT) (ResT, error), m ...echo.MiddlewareFunc) *echo.Route {
	return e.PUT(path, func(c echo.Context) error {
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

func Patch[ReqT, ResT any](e *echo.Echo, path string, fn func(c echo.Context, req ReqT) (ResT, error), m ...echo.MiddlewareFunc) *echo.Route {
	return e.PATCH(path, func(c echo.Context) error {
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

func Delete[ReqT, ResT any](e *echo.Echo, path string, fn func(c echo.Context, req ReqT) (ResT, error), m ...echo.MiddlewareFunc) *echo.Route {
	return e.DELETE(path, func(c echo.Context) error {
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

func Wrap[ReqT, ResT any](fn func(c context.Context, req ReqT) (ResT, error)) func(c echo.Context, req ReqT) (ResT, error) {
	return func(c echo.Context, req ReqT) (ResT, error) {
		return fn(c.Request().Context(), req)
	}
}
