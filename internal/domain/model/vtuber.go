package model

import "gorm.io/gorm"

type Vtuber struct {
	gorm.Model
	Name   string  `gorm:"size:30,uniqueIndex"`
	Groups []Group `gorm:"many2many:memberships;"`
}
