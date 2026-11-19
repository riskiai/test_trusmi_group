package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"trusmi_be/internal/handlers"
	"trusmi_be/internal/models"
	"trusmi_be/internal/routes"
	"trusmi_be/internal/services"

	"github.com/gofiber/fiber/v3"
)

type mockKPIRepo struct{}

func (m *mockKPIRepo) Ping(_ context.Context) error {
	return nil
}

func (m *mockKPIRepo) GetKPISummary(_ context.Context, _ string) ([]models.KPISummaryRow, error) {
	return []models.KPISummaryRow{
		{
			KaryawanID:       1,
			Nama:             "Budi",
			SalesTarget:      2,
			SalesActual:      2,
			SalesPencapaian:  100,
			SalesBobot:       50,
			SalesLateCount:   0,
			SalesLate:        0,
			SalesTotalBobot:  50,
			ReportTarget:     2,
			ReportActual:     2,
			ReportPencapaian: 100,
			ReportBobot:      50,
			ReportLateCount:  1,
			ReportLate:       -5,
			ReportTotalBobot: 45,
			TotalKPI:         95,
		},
		{
			KaryawanID:       2,
			Nama:             "Adi",
			SalesTarget:      2,
			SalesActual:      2,
			SalesPencapaian:  100,
			SalesBobot:       50,
			SalesLateCount:   1,
			SalesLate:        -7,
			SalesTotalBobot:  43,
			ReportTarget:     2,
			ReportActual:     2,
			ReportPencapaian: 100,
			ReportBobot:      50,
			ReportLateCount:  0,
			ReportLate:       0,
			ReportTotalBobot: 50,
			TotalKPI:         93,
		},
		{
			KaryawanID:       3,
			Nama:             "Rara",
			SalesTarget:      2,
			SalesActual:      2,
			SalesPencapaian:  100,
			SalesBobot:       50,
			SalesLateCount:   1,
			SalesLate:        -7,
			SalesTotalBobot:  43,
			ReportTarget:     2,
			ReportActual:     1,
			ReportPencapaian: 50,
			ReportBobot:      25,
			ReportLateCount:  1,
			ReportLate:       -5,
			ReportTotalBobot: 20,
			TotalKPI:         63,
		},
		{
			KaryawanID:       4,
			Nama:             "Doni",
			SalesTarget:      2,
			SalesActual:      1,
			SalesPencapaian:  50,
			SalesBobot:       25,
			SalesLateCount:   1,
			SalesLate:        -7,
			SalesTotalBobot:  18,
			ReportTarget:     2,
			ReportActual:     1,
			ReportPencapaian: 50,
			ReportBobot:      25,
			ReportLateCount:  0,
			ReportLate:       0,
			ReportTotalBobot: 25,
			TotalKPI:         43,
		},
	}, nil
}

func (m *mockKPIRepo) GetOntimeLateSummary(_ context.Context) (models.OntimeLateSummary, error) {
	return models.OntimeLateSummary{
		TotalTask:        13,
		TotalOntime:      8,
		TotalLate:        5,
		PersentaseOntime: 61.54,
		PersentaseLate:   38.46,
	}, nil
}

func (m *mockKPIRepo) GetTasks(_ context.Context, _, _, _ string) ([]models.TaskDetail, error) {
	return []models.TaskDetail{
		{
			ID:          4,
			Tasklist:    "Tasklist 1",
			KPI:         "Sales",
			Karyawan:    "Budi",
			Deadline:    "2022-01-10",
			Aktual:      "2022-01-09",
			Status:      "On-Time",
			SelisihHari: -1,
		},
	}, nil
}

func setupTestApp() *fiber.App {
	app := fiber.New()
	svc := services.NewKPIService(&mockKPIRepo{})
	h := handlers.NewKPIHandler(svc)
	routes.Register(app, h)
	return app
}

func TestEndpoints(t *testing.T) {
	app := setupTestApp()

	t.Run("GET /api/v1/dashboard-1", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/dashboard-1", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status = %d; want %d", resp.StatusCode, http.StatusOK)
		}

		var body models.KPISummaryResponse
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatalf("decode json failed: %v", err)
		}
		if len(body.Data) != 4 {
			t.Fatalf("len(Data) = %d; want 4", len(body.Data))
		}
		if body.Data[0].TotalKPI != 95 || body.Data[3].TotalKPI != 43 {
			t.Errorf("unexpected KPI scores: %+v", body.Data)
		}
	})

	t.Run("GET /api/v1/dashboard-2", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/dashboard-2", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("app.Test failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status = %d; want %d", resp.StatusCode, http.StatusOK)
		}

		var body models.OntimeLateResponse
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatalf("decode json failed: %v", err)
		}
		if body.Data.TotalTask != 13 || body.Data.TotalOntime != 8 || body.Data.TotalLate != 5 {
			t.Errorf("unexpected ontime/late data: %+v", body.Data)
		}
	})
}
