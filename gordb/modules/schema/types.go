package schema

import "time"

type Table struct {
	Name    string   `json:"name"  binding:"required"`
	Columns []Column `json:"columns" binding:"required"`
}

type Column struct {
	Name      string    `json:"name" binding:"required"`
	Type      string    `json:"type" binding:"required"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Meta struct {
	TableName string
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TableRepoResponse struct {
	Name    string `json:"name" binding:"required"`
	Message string `json:"message" binding:"required"`
	Error   string `json:"error"`
	Path    string `json:"path"`
}
