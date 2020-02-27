package questionsvc

import (
	"github.com/miraclew/tao-demo/question"
	"e.coding.net/miraclew/tao/pkg/ac"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type handler struct {
	Service question.Service
}

func (h *handler) RegisterRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	e.POST("/v1/question/create", h.Create, m...)
	e.POST("/v1/question/delete", h.Delete, m...)
	e.POST("/v1/question/update", h.Update, m...)
	e.POST("/v1/question/get", h.Get, m...)
	e.POST("/v1/question/query", h.Query, m...)
}

func (h *handler) Create(c echo.Context) error {
	ctx := ac.FromEcho(c)
	req := new(question.CreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}

	res, err := h.Service.Create(ctx, req)
	if err != nil {
		return errors.Wrap(err, "handler: get error")
	}

	return c.JSON(200, res)
}

func (h *handler) Delete(c echo.Context) error {
	ctx := ac.FromEcho(c)
	req := new(question.DeleteRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}

	res, err := h.Service.Delete(ctx, req)
	if err != nil {
		return errors.Wrap(err, "handler: get error")
	}

	return c.JSON(200, res)
}

func (h *handler) Update(c echo.Context) error {
	ctx := ac.FromEcho(c)
	req := new(question.UpdateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}

	res, err := h.Service.Update(ctx, req)
	if err != nil {
		return errors.Wrap(err, "handler: get error")
	}

	return c.JSON(200, res)
}

func (h *handler) Get(c echo.Context) error {
	ctx := ac.FromEcho(c)
	req := new(question.GetRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}

	res, err := h.Service.Get(ctx, req)
	if err != nil {
		return errors.Wrap(err, "handler: get error")
	}

	return c.JSON(200, res)
}

func (h *handler) Query(c echo.Context) error {
	ctx := ac.FromEcho(c)
	req := new(question.QueryRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}

	res, err := h.Service.Query(ctx, req)
	if err != nil {
		return errors.Wrap(err, "handler: get error")
	}

	return c.JSON(200, res)
}