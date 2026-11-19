package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"trusmi_be/internal/config"
	"trusmi_be/internal/database"
	"trusmi_be/internal/handlers"
	"trusmi_be/internal/middleware"
	"trusmi_be/internal/repositories"
	"trusmi_be/internal/routes"
	"trusmi_be/internal/services"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.Load()

	db, err := database.NewMySQL(cfg)
	if err != nil {
		log.Fatalf("[FATAL] Database connection failed: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("[WARN] Error closing database connection: %v", err)
		}
	}()

	app := fiber.New(fiber.Config{
		AppName: "TRUSMI GROUP - Marketing KPI API",
	})

	middleware.Setup(app, cfg)

	kpiRepo := repositories.NewKPIRepository(db)
	kpiService := services.NewKPIService(kpiRepo)
	kpiHandler := handlers.NewKPIHandler(kpiService)

	routes.Register(app, kpiHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	listenAddr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	log.Printf("[INFO] Starting TRUSMI KPI API server on %s (env=%s)", listenAddr, cfg.AppEnv)

	if err := app.Listen(listenAddr, fiber.ListenConfig{
		GracefulContext: ctx,
		ShutdownTimeout: 10 * time.Second,
	}); err != nil {
		log.Fatalf("[FATAL] Server stopped with error: %v", err)
	}
}
