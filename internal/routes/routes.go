package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsercher/internal/handler"
	registory "github.com/t1732/vsercher/internal/registry"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	repo := registory.NewRepository()

	router.GET("/ping", handler.NewPing().Show)

	group := router.Group("/vtubers")
	vtuberHandler := handler.NewVtuber(repo)
	group.GET("", vtuberHandler.Index)
	group.GET("/:id", vtuberHandler.Show)

	return router
}
