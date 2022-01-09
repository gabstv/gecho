package gecho

import (
	"context"

	"github.com/labstack/echo/v4"
)

func Middleware[ReqT any](fn func(c echo.Context, req ReqT) (ReqT, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if ibound := c.Get("_bound_"); ibound != nil {
				b := ibound.(ReqT)
				b2, err := fn(c, b)
				if err != nil {
					return err
				}
				c.Set("_bound_", b2)
				return next(c)
			}
			var b ReqT
			if err := c.Bind(&b); err != nil {
				return err
			}
			b2, err := fn(c, b)
			if err != nil {
				return err
			}
			c.Set("_bound_", b2)
			return next(c)
		}
	}
}

func Get[ReqT, ResT any](e *echo.Echo, path string, fn func(c echo.Context, req ReqT) (ResT, error), m ...echo.MiddlewareFunc) *echo.Route {
	return e.GET(path, func(c echo.Context) error {
		var rq ReqT
		if ibound := c.Get("_bound_"); ibound != nil {
			rq = ibound.(ReqT)
		} else {
			if err := c.Bind(&rq); err != nil {
				return err
			}
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
		if ibound := c.Get("_bound_"); ibound != nil {
			rq = ibound.(ReqT)
		} else {
			if err := c.Bind(&rq); err != nil {
				return err
			}
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
		if ibound := c.Get("_bound_"); ibound != nil {
			rq = ibound.(ReqT)
		} else {
			if err := c.Bind(&rq); err != nil {
				return err
			}
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
		if ibound := c.Get("_bound_"); ibound != nil {
			rq = ibound.(ReqT)
		} else {
			if err := c.Bind(&rq); err != nil {
				return err
			}
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
		if ibound := c.Get("_bound_"); ibound != nil {
			rq = ibound.(ReqT)
		} else {
			if err := c.Bind(&rq); err != nil {
				return err
			}
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

// CtxGet converts to T. Returns the zero value if not found
func CtxGet[T any](c echo.Context, key string) T {
	v, _ := c.Get(key).(T)
	return v
}

func CtxSet[T any](c echo.Context, key string, v T) {
	c.Set(key, v)
}
