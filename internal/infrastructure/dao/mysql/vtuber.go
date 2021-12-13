package mysql

import (
	"github.com/t1732/vsercher/internal/domain/model"
	"gorm.io/gorm"
)

type Vtuber interface {
	All() (*[]model.Vtuber, error)
	FindById(id int64) (*model.Vtuber, error)
}

type vtuberImpl struct {
	conn *gorm.DB
}

func NewVtuber() Vtuber {
	return &vtuberImpl{conn: Connection}
}

func (v *vtuberImpl) All() (*[]model.Vtuber, error) {
	var vtubers []model.Vtuber
	result := v.conn.Find(&vtubers)

	if result.Error != nil {
		return nil, result.Error
	}
	return &vtubers, nil
}

func (v *vtuberImpl) FindById(id int64) (*model.Vtuber, error) {
	var vtuber model.Vtuber
	result := v.conn.First(&vtuber, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vtuber, nil
}
