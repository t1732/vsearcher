package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsercher/internal/domain/model"
	"github.com/t1732/vsercher/internal/infrastructure/dao"
	"gorm.io/gorm"
)

type VtuberHandler interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
}

type vtuber struct{}

func NewVtuber() VtuberHandler {
	return &vtuber{}
}

func (v *vtuber) Index(c *gin.Context) {
	var vtubers []model.Vtuber
	dao.DB().Find(&vtubers)
	c.SecureJSON(http.StatusOK, vtubers)
}

func (v *vtuber) Show(c *gin.Context) {
	var vtuber model.Vtuber
	result := dao.DB().First(&vtuber, c.Param("id"))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Error(result.Error)
		return
	}
	c.SecureJSON(http.StatusOK, vtuber)
}
