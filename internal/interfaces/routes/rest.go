package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1732/vsearcher/internal/interfaces/handler/rest"
	"github.com/t1732/vsearcher/internal/registry"
	"gorm.io/gorm"
)

func Router(dbConn *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	repo := registry.NewRepository(dbConn)

	router.GET("/ping", rest.NewPing().Show)

	group := router.Group("/vtubers")
	vtuberHandler := rest.NewVtuber(repo)
	group.GET("", vtuberHandler.Index)
	group.GET("/:id", vtuberHandler.Show)

	return router
}
