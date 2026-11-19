package handlers

import (
	"log"

	"trusmi_be/internal/services"

	"github.com/gofiber/fiber/v3"
)

// KPIHandler menangani HTTP request untuk endpoint health dan KPI Marketing.
type KPIHandler struct {
	service *services.KPIService
}

// NewKPIHandler membuat instance KPIHandler baru.
func NewKPIHandler(service *services.KPIService) *KPIHandler {
	return &KPIHandler{service: service}
}

// Health menangani GET /api/v1/health
func (h *KPIHandler) Health(c fiber.Ctx) error {
	if err := h.service.CheckHealth(c.Context()); err != nil {
		log.Printf("[ERROR] database healthcheck failed: %v", err)
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":   "degraded",
			"service":  "trusmi-be",
			"database": "unreachable",
			"message":  "Database service is currently unavailable",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "ok",
		"service":  "trusmi-be",
		"database": "connected",
	})
}

// GetKPISummary menangani GET /api/v1/kpi/summary (Soal 1 & Soal 3)
func (h *KPIHandler) GetKPISummary(c fiber.Ctx) error {
	karyawan := c.Query("karyawan")

	res, err := h.service.GetKPISummary(c.Context(), karyawan)
	if err != nil {
		log.Printf("[ERROR] failed to fetch KPI summary: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data ringkasan KPI",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// GetOntimeLate menangani GET /api/v1/kpi/ontime-late (Soal 2 & Soal 4)
func (h *KPIHandler) GetOntimeLate(c fiber.Ctx) error {
	res, err := h.service.GetOntimeLateSummary(c.Context())
	if err != nil {
		log.Printf("[ERROR] failed to fetch Ontime/Late summary: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data ketepatan waktu tasklist",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// GetTasks menangani GET /api/v1/kpi/tasks (Soal 5 & Tabel Task Detail)
func (h *KPIHandler) GetTasks(c fiber.Ctx) error {
	karyawan := c.Query("karyawan")
	kpi := c.Query("kpi")
	status := c.Query("status")

	res, err := h.service.GetTasks(c.Context(), karyawan, kpi, status)
	if err != nil {
		log.Printf("[ERROR] failed to fetch task details: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil daftar tasklist KPI",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
