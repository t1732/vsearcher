package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string   `gorm:"site:30,uniqueIndex"`
	Vtubers []Vtuber `gorm:"many2many:memberships;"`
}
