package routes

import (
	"trusmi_be/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

// Register mendaftarkan seluruh route REST API di bawah prefix /api/v1.
func Register(app *fiber.App, kpiHandler *handlers.KPIHandler) {
	// Root endpoint info
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "TRUSMI GROUP - Marketing KPI API",
			"version": "1.0.0",
			"endpoints": fiber.Map{
				"health":      "GET /api/v1/health",
				"dashboard_1": "GET /api/v1/kpi/summary (atau GET /api/v1/dashboard-1)",
				"dashboard_2": "GET /api/v1/kpi/ontime-late (atau GET /api/v1/dashboard-2)",
				"tasks":       "GET /api/v1/kpi/tasks",
			},
		})
	})

	api := app.Group("/api/v1")

	api.Get("/health", kpiHandler.Health)

	// Alias langsung sesuai koleksi Postman (DASHBOARD 1 & DASHBOARD 2)
	api.Get("/dashboard-1", kpiHandler.GetKPISummary)
	api.Get("/dashboard-2", kpiHandler.GetOntimeLate)

	// Route utama sesuai spesifikasi /api/v1/kpi/...
	kpi := api.Group("/kpi")
	kpi.Get("/summary", kpiHandler.GetKPISummary)
	kpi.Get("/dashboard-1", kpiHandler.GetKPISummary)
	kpi.Get("/ontime-late", kpiHandler.GetOntimeLate)
	kpi.Get("/dashboard-2", kpiHandler.GetOntimeLate)
	kpi.Get("/tasks", kpiHandler.GetTasks)
}
