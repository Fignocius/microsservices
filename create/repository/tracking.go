package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/microsservices/create/model"
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

// Create function to changes fields of the tracking
func (t Tracking) Create(tracking model.Tracking) (err error) {
	query := sq.Insert("trackings").
		Columns(
			"id",
			"code",
			"status",
			"description",
		).
		Values(
			tracking.ID,
			tracking.Code,
			tracking.Status,
			tracking.Description,
		)
	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := t.db.Preparex(statement)
	if err != nil {
		return
	}
	if _, err = stmt.Exec(args); err != nil {
		return
	}
	err = stmt.Close()
	return
}
