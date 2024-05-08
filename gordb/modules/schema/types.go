package schema

import "time"

type Table struct {
	Name      string    `json:"name"  binding:"required"`
	Columns   []Column  `json:"columns" binding:"required"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Column struct {
	Name      string    `json:"name" binding:"required"`
	Type      string    `json:"type" binding:"required"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
