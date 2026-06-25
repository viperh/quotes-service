package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"QuotesService/internal/api/controllers"
	"QuotesService/internal/api/middlewares"
	"QuotesService/internal/api/routes"
	"QuotesService/internal/config"
	"QuotesService/internal/provider"
	"QuotesService/internal/seed"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	Controller *controllers.Controller
	Config     *config.Config
	Provider   *provider.Postgres
	Engine     *gin.Engine
}

func New() *App {
	cfg := config.New()
	prv := provider.New(cfg)

	if err := prv.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	if err := seed.Run(prv); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	cnt := controllers.New(prv)

	engine := gin.Default()
	engine.Use(middlewares.CORSMiddleware())
	routes.DefineRoutes(engine, cnt)

	return &App{
		Controller: cnt,
		Config:     cfg,
		Provider:   prv,
		Engine:     engine,
	}
}

func (a *App) Run() {
	srv := &http.Server{
		Addr:    ":" + a.Config.Port,
		Handler: a.Engine,
	}

	go func() {
		log.Printf("Server starting on port %s", a.Config.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
