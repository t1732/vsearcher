package registry

import (
	"github.com/t1732/vsercher/internal/domain/repository"
	"github.com/t1732/vsercher/internal/infrastructure/dao"
	"gorm.io/gorm"
)

type Repository interface {
	NewVtuberRepository() repository.Vtuber
}

type repositoryImpl struct {
	dbConn *gorm.DB
}

func NewRepository(dbConn *gorm.DB) Repository {
	return &repositoryImpl{
		dbConn: dbConn,
	}
}

func (r *repositoryImpl) NewVtuberRepository() repository.Vtuber {
	return dao.NewVtuber(r.dbConn)
}
