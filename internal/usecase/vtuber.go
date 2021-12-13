package usecase

import (
	"github.com/t1732/vsercher/internal/domain/model"
	"github.com/t1732/vsercher/internal/domain/repository"
)

type Vtuber interface {
	Index() (*[]model.Vtuber, error)
	Show(id int64) (*model.Vtuber, error)
}

type vtuber struct {
	repo repository.Vtuber
}

func NewVtuber(repo repository.Vtuber) Vtuber {
	return &vtuber{repo}
}

func (v *vtuber) Index() (*[]model.Vtuber, error) {
	return v.repo.All()
}

func (v *vtuber) Show(id int64) (*model.Vtuber, error) {
	return v.repo.FindById(id)
}
