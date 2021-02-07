package model

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Tracking struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (m Tracking) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return string("")
	}
	return string(b)
}
