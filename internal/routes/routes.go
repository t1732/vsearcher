package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsercher/internal/handler"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	router.GET("/ping", handler.NewPing().Show)

	group := router.Group("/vtubers")
	vh := handler.NewVtuber()
	group.GET("", vh.Index)

	return router
}
