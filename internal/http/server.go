package http

import (
	"fmt"
	"nba-task-main/internal/http/aggregateservice"
	"nba-task-main/internal/http/statservice"

	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
)

func NewAggregateServer(cfg Config, herdService *aggregateservice.Server) (*http.Server, error) {
	r, err := NewAggregateRouter(herdService)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gin.Engine: %w", err)
	}

	return &http.Server{
		Addr:    cfg.Host,
		Handler: r,
	}, nil
}

func NewAggregateRouter(herdService *aggregateservice.Server) (*echo.Echo, error) {
	swagger, err := aggregateservice.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("failed to read swagger: %w", err)
	}
	swagger.Servers = nil

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.OapiRequestValidator(swagger))

	aggregateservice.RegisterHandlers(e, aggregateservice.NewStrictHandler(herdService, nil))

	return e, nil
}

func NewStatServer(cfg Config, herdService *statservice.Server) (*http.Server, error) {
	r, err := NewStatRouter(herdService)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gin.Engine: %w", err)
	}

	return &http.Server{
		Addr:    cfg.Host,
		Handler: r,
	}, nil
}

func NewStatRouter(herdService *statservice.Server) (*echo.Echo, error) {
	swagger, err := statservice.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("failed to read swagger: %w", err)
	}
	swagger.Servers = nil

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.OapiRequestValidator(swagger))

	statservice.RegisterHandlers(e, statservice.NewStrictHandler(herdService, nil))

	return e, nil
}
