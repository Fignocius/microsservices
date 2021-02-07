package handler

import (
	"github.com/labstack/echo/v4"
)

// TrackingHandler interface
type TrackingHandler interface {
	Create(c echo.Context) error
}
