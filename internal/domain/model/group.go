package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string `gorm:"site:30,uniqueIndex"`
	Memberships []Membership
}
