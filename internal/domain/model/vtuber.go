package model

import "time"

type Vtuber struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
