package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsercher/internal/domain/model"
	"github.com/t1732/vsercher/internal/infrastructure/dao"
)

type VtuberHandler interface {
	Index(c *gin.Context)
}

type vtuber struct {}

func NewVtuber() VtuberHandler {
	return &vtuber{}
}

func (v *vtuber) Index(c *gin.Context) {
	var vtubers []model.Vtuber
	dao.DB().Find(&vtubers)
	c.SecureJSON(http.StatusOK, vtubers)
}
