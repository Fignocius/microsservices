package repository

import (
	"github.com/fignocius/microsservices/create/model"
)

// TrackingRepository is a interface of repository
type TrackingRepository interface {
	Create(model.Tracking) error
}
