package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsercher/internal/domain/model"
	"github.com/t1732/vsercher/internal/registry"
	"github.com/t1732/vsercher/internal/usecase"
	"github.com/t1732/vsercher/pkg/utils/parser"
)

type VtuberHandler interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
}

type vtuber struct {
	repo registry.Repository
}

func NewVtuber(repo registry.Repository) VtuberHandler {
	return &vtuber{repo}
}

func (v *vtuber) Index(c *gin.Context) {
	vtuberRepo := v.repo.NewVtuber()
	usecase := usecase.NewVtuber(vtuberRepo)
	vtubers, err := usecase.Index()
	if err != nil {
		panic(err)
	}

	c.SecureJSON(http.StatusOK, vtubers)
}

func (v *vtuber) Show(c *gin.Context) {
	id, err := parser.ToInt64(c.Param("id"))
	if err != nil {
		panic(err)
	}

	vtuberRepo := v.repo.NewVtuber()
	var vtuber *model.Vtuber
	usecase := usecase.NewVtuber(vtuberRepo)
	vtuber, err = usecase.Show(id)
	if err != nil {
		panic(err)
	}

	c.SecureJSON(http.StatusOK, vtuber)
}
