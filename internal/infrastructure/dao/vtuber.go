package dao

import (
	"github.com/t1732/vsearcher/internal/domain/model"
	"gorm.io/gorm"
)

type Vtuber interface {
	All() (*[]model.Vtuber, error)
	FindById(id int64) (*model.Vtuber, error)
}

type vtuberImpl struct {
	dbConn *gorm.DB
}

func NewVtuber(dbConn *gorm.DB) Vtuber {
	return &vtuberImpl{dbConn: dbConn}
}

func (v *vtuberImpl) All() (*[]model.Vtuber, error) {
	var vtubers []model.Vtuber

	if err := v.dbConn.Preload("Groups", func(db *gorm.DB) *gorm.DB {
	 	return db.Select("ID", "Name")
	}).Select("ID", "Name").Order("id desc").Find(&vtubers).Error; err != nil {
		return nil, err
	}

	return &vtubers, nil
}

func (v *vtuberImpl) FindById(id int64) (*model.Vtuber, error) {
	var vtuber model.Vtuber
	result := v.dbConn.First(&vtuber, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vtuber, nil
}
