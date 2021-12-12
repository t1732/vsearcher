package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/t1732/vsercher/internal/infrastructure/dao"
	"github.com/t1732/vsercher/internal/routes"
)

func main() {
	db, err := dao.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dao.Migrate()
	dao.Seed()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("running gin server " + port + " port")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: routes.Router(),
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
