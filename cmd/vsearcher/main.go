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
	"github.com/t1732/vsearcher/internal/config"
	lgg "github.com/t1732/vsearcher/internal/config/logger"
	"github.com/t1732/vsearcher/internal/interfaces/routes"
	"gorm.io/gorm/logger"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("running gin server " + port + " port")

	lgCnf := lgg.MysqlConfig{
		IsFile: gin.Mode() == gin.DebugMode,
	}
	if gin.Mode() == gin.DebugMode {
		lgCnf.LogLevel = logger.Info
	}
	dbConn, err := config.NewDB(lgCnf)
	if err != nil {
		panic(err)
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: routes.Router(dbConn),
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
