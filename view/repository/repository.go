package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/microsservices/view/model"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// TrackingRepository is a interface of repository
type TrackingRepository interface {
	Get(code string) (model.Tracking, error)
}
