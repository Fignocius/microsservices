package handler

import (
	"github.com/fignocius/microsservices/update/model"
	"github.com/fignocius/microsservices/update/repository"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

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

// Update is a function to Update the tracking status
func (t Tracking) Update(c echo.Context) error {
	var tracking model.Tracking
	err := c.Bind(tracking)
	if err != nil {
		return echo.ErrBadRequest
	}
	err = t.TrackingRepository.Update(tracking)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return echo.ErrBadRequest
}
