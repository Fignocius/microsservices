package handler

import (
	"github.com/labstack/echo/v4"
)

// TrackingHandler interface
type TrackingHandler interface {
	Update(c echo.Context) error
}
