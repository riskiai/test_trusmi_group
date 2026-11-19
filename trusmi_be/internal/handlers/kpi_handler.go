package handlers

import (
	"log"

	"trusmi_be/internal/services"

	"github.com/gofiber/fiber/v3"
)

// KPIHandler menangani HTTP request untuk Dashboard 1 dan Dashboard 2.
type KPIHandler struct {
	service *services.KPIService
}

// NewKPIHandler membuat instance KPIHandler baru.
func NewKPIHandler(service *services.KPIService) *KPIHandler {
	return &KPIHandler{service: service}
}

// GetDashboard1 menangani GET /api/v1/dashboard-1 (Soal 1 & Soal 3)
func (h *KPIHandler) GetDashboard1(c fiber.Ctx) error {
	karyawan := c.Query("karyawan")

	res, err := h.service.GetKPISummary(c.Context(), karyawan)
	if err != nil {
		log.Printf("[ERROR] failed to fetch Dashboard 1 KPI summary: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data Dashboard 1",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// GetDashboard2 menangani GET /api/v1/dashboard-2 (Soal 2 & Soal 4)
func (h *KPIHandler) GetDashboard2(c fiber.Ctx) error {
	res, err := h.service.GetOntimeLateSummary(c.Context())
	if err != nil {
		log.Printf("[ERROR] failed to fetch Dashboard 2 Ontime/Late summary: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data Dashboard 2",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
