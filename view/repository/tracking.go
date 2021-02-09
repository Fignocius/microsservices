package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/microsservices/view/model"
	"github.com/jmoiron/sqlx"
)

// Tracking struct repository
type Tracking struct {
	db *sqlx.DB
}

// NewTracking function to instance new tracking
func NewTracking(db *sqlx.DB) TrackingRepository {
	return &Tracking{
		db: db,
	}
}

// Get function to get fields of the tracking
func (t Tracking) Get(code string) (tracking model.Tracking, err error) {
	query := psql.Select("*").
		From("trackings").
		Where(sq.Eq{"code": code})

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := t.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Get(&tracking, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}
