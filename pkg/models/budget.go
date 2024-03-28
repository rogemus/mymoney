package models

import (
	"time"
)

type Budget struct {
	ID          int
	Uuid        string
	Created     time.Time
	Description string
	Title       string
}
