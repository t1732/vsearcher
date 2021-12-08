package main

import (
	"github.com/gin-gonic/gin"
	"github.com/t1732/vsercher/internal/handler"
	"github.com/t1732/vsercher/internal/infrastructure/dao"
)

func main() {
	db, err := dao.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dao.Migrate()
	dao.Seed()

	router := gin.Default()
	router.GET("/ping", handler.NewPing().Show)
	router.GET("/vtubers", handler.NewVtuber().Index)
	router.Run()
}
