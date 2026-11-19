package routes

import (
	"trusmi_be/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

// Register mendaftarkan route REST API untuk Dashboard 1 dan Dashboard 2.
func Register(app *fiber.App, kpiHandler *handlers.KPIHandler) {
	api := app.Group("/api/v1")

	api.Get("/dashboard-1", kpiHandler.GetDashboard1)
	api.Get("/dashboard-2", kpiHandler.GetDashboard2)
}
