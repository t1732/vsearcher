package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("running gin server " + port + " port")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
