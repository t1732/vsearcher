package repository

import "github.com/t1732/vsercher/internal/domain/model"

type Vtuber interface {
	All() (*[]model.Vtuber, error)
	FindById(int64) (*model.Vtuber, error)
}
