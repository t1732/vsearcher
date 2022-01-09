package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsercher/internal/interfaces/handler/rest"
	registory "github.com/t1732/vsercher/internal/registry"
	"gorm.io/gorm"
)

func Router(dbConn *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	repo := registory.NewRepository(dbConn)

	router.GET("/ping", rest.NewPing().Show)

	group := router.Group("/vtubers")
	vtuberHandler := rest.NewVtuber(repo)
	group.GET("", vtuberHandler.Index)
	group.GET("/:id", vtuberHandler.Show)

	return router
}
