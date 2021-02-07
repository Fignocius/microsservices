package repository

import (
	"github.com/fignocius/microsservices/update/model"
)

// TrackingRepository is a interface of repository
type TrackingRepository interface {
	Update(model.Tracking) error
}
