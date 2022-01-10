package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsearcher/internal/domain/model"
	"github.com/t1732/vsearcher/internal/registry"
	"github.com/t1732/vsearcher/internal/usecase"
	"github.com/t1732/vsearcher/pkg/utils/parser"
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
	vtuberRepo := v.repo.NewVtuberRepository()
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

	vtuberRepo := v.repo.NewVtuberRepository()
	var vtuber *model.Vtuber
	usecase := usecase.NewVtuber(vtuberRepo)
	vtuber, err = usecase.Show(id)
	if err != nil {
		panic(err)
	}

	c.SecureJSON(http.StatusOK, vtuber)
}
