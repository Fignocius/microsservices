package handler

import (
	"math/rand"
	"net/http"

	"github.com/fignocius/microsservices/view/repository"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// Tracking struct of handler
type Tracking struct {
	repository.TrackingRepository
}

// NewTrackingHandler is a function to instance a new TrackingHandler
func NewTrackingHandler(db *sqlx.DB) TrackingHandler {
	return &Tracking{
		repository.NewTracking(db),
	}
}

// View is a function to View the tracking status
func (t Tracking) View(c echo.Context) error {
	code := c.Param("code")
	tracking, err := t.TrackingRepository.Get(code)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, tracking)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
