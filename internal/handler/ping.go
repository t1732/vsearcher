package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler interface {
	Show(c *gin.Context)
}

type ping struct {}

func NewPing() PingHandler {
	return &ping{}
}

func (p *ping) Show(c *gin.Context) {
	c.SecureJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
