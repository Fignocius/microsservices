package handler

import (
	"net/http"
	"time"

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
	var (
		tracking   model.Tracking
		parameters Parameters
	)
	err := c.Bind(&parameters)
	if err != nil {
		return echo.ErrBadRequest
	}
	tracking.Code = parameters.Code
	tracking.Status = parameters.Status
	tracking.UpdatedAt = time.Now().UTC()
	if parameters.Description != nil {
		tracking.Description = *parameters.Description
	}
	err = t.TrackingRepository.Update(tracking)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, "")
}

type Parameters struct {
	Code        string  `json:"code"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
}
