package repositories_test

import (
	"context"
	"testing"
	"time"

	"trusmi_be/internal/config"
	"trusmi_be/internal/database"
	"trusmi_be/internal/repositories"
)

func TestMySQLIntegrationReadOnly(t *testing.T) {
	cfg := config.Load()
	db, err := database.NewMySQL(cfg)
	if err != nil {
		t.Skipf("Skipping live MySQL integration test (DB not reachable): %v", err)
	}
	defer db.Close()

	repo := repositories.NewKPIRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Verify Soal 1 Query
	summary, err := repo.GetKPISummary(ctx, "")
	if err != nil {
		t.Fatalf("GetKPISummary failed: %v", err)
	}
	if len(summary) != 4 {
		t.Fatalf("len(summary) = %d; want 4", len(summary))
	}

	expectedKPI := map[string]float64{
		"Budi": 95.0,
		"Adi":  93.0,
		"Rara": 63.0,
		"Doni": 43.0,
	}
	for _, row := range summary {
		if want, ok := expectedKPI[row.Nama]; ok {
			if row.TotalKPI != want {
				t.Errorf("%s TotalKPI = %.2f; want %.2f", row.Nama, row.TotalKPI, want)
			}
		}
	}

	// 2. Verify Soal 2 Query
	ontimeLate, err := repo.GetOntimeLateSummary(ctx)
	if err != nil {
		t.Fatalf("GetOntimeLateSummary failed: %v", err)
	}
	if ontimeLate.TotalTask != 13 || ontimeLate.TotalOntime != 8 || ontimeLate.TotalLate != 5 {
		t.Errorf("OntimeLate = %+v; want Total=13, Ontime=8, Late=5", ontimeLate)
	}

	// 3. Verify Soal 5 Tasks Query
	tasks, err := repo.GetTasks(ctx, "", "", "")
	if err != nil {
		t.Fatalf("GetTasks failed: %v", err)
	}
	if len(tasks) != 13 {
		t.Errorf("len(tasks) = %d; want 13", len(tasks))
	}
}
