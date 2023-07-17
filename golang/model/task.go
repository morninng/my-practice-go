package model

import (
	"github.com/volatiletech/null/v8"
)

type TaskResponse struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}
