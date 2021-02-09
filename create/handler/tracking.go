package handler

import (
	"math/rand"
	"net/http"

	"github.com/fignocius/microsservices/create/model"
	"github.com/fignocius/microsservices/create/repository"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
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

// Create is a function to Create the tracking status
func (t Tracking) Create(c echo.Context) error {

	tracking := model.Tracking{
		ID:          uuid.NewV4(),
		Code:        randStringBytes(14),
		Status:      "Coletado",
		Description: "Entregue no posto de coleta",
	}

	err := t.TrackingRepository.Create(tracking)
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
