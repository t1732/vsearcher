package model

import "time"

type Membership struct {
	ID        uint
	CreatedAt time.Time
	VtuberID  uint
	Vtuber    Vtuber
	GroupID   uint
	Group     Group
}
