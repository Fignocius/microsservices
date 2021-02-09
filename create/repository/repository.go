package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/microsservices/create/model"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// TrackingRepository is a interface of repository
type TrackingRepository interface {
	Create(model.Tracking) error
}
