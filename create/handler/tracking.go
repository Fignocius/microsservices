package handler

import (
	"fmt"
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

	trackingIn := Create{}
	trackingOut := model.Tracking{
		ID:   uuid.NewV4(),
		Code: randStringBytes(14),
	}

	err := c.Bind(&trackingIn)
	if err != nil {
		fmt.Println(trackingIn)
		return echo.ErrBadRequest
	}
	trackingOut.Status = trackingIn.Status
	trackingOut.Description = trackingIn.Description
	err = t.TrackingRepository.Create(trackingOut)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, trackingOut)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type Create struct {
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}
