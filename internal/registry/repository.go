package registry

import (
	"github.com/t1732/vsercher/internal/domain/repository"
	"github.com/t1732/vsercher/internal/infrastructure/dao/mysql"
)

type Repository interface {
	NewVtuber() repository.Vtuber
}

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) NewVtuber() repository.Vtuber {
	return mysql.NewVtuber()
}
