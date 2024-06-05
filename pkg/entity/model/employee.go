package model

import (
	"github.com/google/uuid"
)

type Detail struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Position string    `json:"position"`
	Salary   float64   `json:"salary"`
}
