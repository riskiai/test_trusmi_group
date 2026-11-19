package services

import (
	"testing"

	"trusmi_be/internal/models"
)

func TestCalculateKPIMetric(t *testing.T) {
	tests := []struct {
		name            string
		salesActual     int
		salesLateCount  int
		reportActual    int
		reportLateCount int
		wantSalesBobot  float64
		wantSalesTotal  float64
		wantReportBobot float64
		wantReportTotal float64
		wantTotalKPI    float64
	}{
		{
			name:            "Budi - Sales 2 (0 late), Report 2 (1 late)",
			salesActual:     2,
			salesLateCount:  0,
			reportActual:    2,
			reportLateCount: 1,
			wantSalesBobot:  50.00,
			wantSalesTotal:  50.00,
			wantReportBobot: 50.00,
			wantReportTotal: 45.00,
			wantTotalKPI:    95.00,
		},
		{
			name:            "Adi - Sales 2 (1 late), Report 2 (0 late)",
			salesActual:     2,
			salesLateCount:  1,
			reportActual:    2,
			reportLateCount: 0,
			wantSalesBobot:  50.00,
			wantSalesTotal:  43.00,
			wantReportBobot: 50.00,
			wantReportTotal: 50.00,
			wantTotalKPI:    93.00,
		},
		{
			name:            "Rara - Sales 2 (1 late), Report 1 (1 late)",
			salesActual:     2,
			salesLateCount:  1,
			reportActual:    1,
			reportLateCount: 1,
			wantSalesBobot:  50.00,
			wantSalesTotal:  43.00,
			wantReportBobot: 25.00,
			wantReportTotal: 20.00,
			wantTotalKPI:    63.00,
		},
		{
			name:            "Doni - Sales 1 (1 late), Report 1 (0 late)",
			salesActual:     1,
			salesLateCount:  1,
			reportActual:    1,
			reportLateCount: 0,
			wantSalesBobot:  25.00,
			wantSalesTotal:  18.00,
			wantReportBobot: 25.00,
			wantReportTotal: 25.00,
			wantTotalKPI:    43.00,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, salesBobot, _, salesTotal := CalculateKPIMetric(2, tc.salesActual, tc.salesLateCount, 50.0, -7.0)
			_, reportBobot, _, reportTotal := CalculateKPIMetric(2, tc.reportActual, tc.reportLateCount, 50.0, -5.0)
			totalKPI := Round2(salesTotal + reportTotal)

			if salesBobot != tc.wantSalesBobot {
				t.Errorf("salesBobot = %.2f; want %.2f", salesBobot, tc.wantSalesBobot)
			}
			if salesTotal != tc.wantSalesTotal {
				t.Errorf("salesTotal = %.2f; want %.2f", salesTotal, tc.wantSalesTotal)
			}
			if reportBobot != tc.wantReportBobot {
				t.Errorf("reportBobot = %.2f; want %.2f", reportBobot, tc.wantReportBobot)
			}
			if reportTotal != tc.wantReportTotal {
				t.Errorf("reportTotal = %.2f; want %.2f", reportTotal, tc.wantReportTotal)
			}
			if totalKPI != tc.wantTotalKPI {
				t.Errorf("totalKPI = %.2f; want %.2f", totalKPI, tc.wantTotalKPI)
			}
		})
	}
}

func TestClassifyTaskStatus(t *testing.T) {
	dataset := []struct {
		tasklist string
		deadline string
		aktual   string
		want     string
	}{
		{"Tasklist 1", "2022-01-10", "2022-01-09", "On-Time"},
		{"Tasklist 2", "2022-01-10", "2022-01-08", "On-Time"},
		{"Tasklist 3", "2022-01-10", "2022-01-07", "On-Time"},
		{"Tasklist 4", "2022-01-10", "2022-01-12", "Late"},
		{"Tasklist 5", "2022-01-10", "2022-01-09", "On-Time"},
		{"Tasklist 6", "2022-01-10", "2022-01-12", "Late"},
		{"Tasklist 7", "2022-01-10", "2022-01-07", "On-Time"},
		{"Tasklist 8", "2022-01-10", "2022-01-07", "On-Time"},
		{"Tasklist 9", "2022-01-10", "2022-01-12", "Late"},
		{"Tasklist 10", "2022-01-10", "2022-01-09", "On-Time"},
		{"Tasklist 11", "2022-01-10", "2022-01-12", "Late"},
		{"Tasklist 12", "2022-01-10", "2022-01-09", "On-Time"},
		{"Tasklist 13", "2022-01-10", "2022-01-12", "Late"},
	}

	var ontimeCount, lateCount int
	for _, tc := range dataset {
		got := ClassifyTaskStatus(tc.deadline, tc.aktual)
		if got != tc.want {
			t.Errorf("%s: ClassifyTaskStatus(%s, %s) = %s; want %s", tc.tasklist, tc.deadline, tc.aktual, got, tc.want)
		}
		if got == "On-Time" {
			ontimeCount++
		} else if got == "Late" {
			lateCount++
		}
	}

	if ontimeCount != 8 || lateCount != 5 {
		t.Errorf("ontimeCount=%d, lateCount=%d; want 8 and 5", ontimeCount, lateCount)
	}

	ontimePct := Round2(float64(ontimeCount) * 100.0 / float64(len(dataset)))
	latePct := Round2(float64(lateCount) * 100.0 / float64(len(dataset)))
	if ontimePct != 61.54 || latePct != 38.46 {
		t.Errorf("percentages = %.2f%% / %.2f%%; want 61.54%% / 38.46%%", ontimePct, latePct)
	}
}

func TestCalculateSummaryMeta(t *testing.T) {
	rows := []models.KPISummaryRow{
		{Nama: "Budi", TotalKPI: 95},
		{Nama: "Adi", TotalKPI: 93},
		{Nama: "Rara", TotalKPI: 63},
		{Nama: "Doni", TotalKPI: 43},
	}

	meta := CalculateSummaryMeta(rows)
	if meta.TotalKaryawan != 4 {
		t.Errorf("TotalKaryawan = %d; want 4", meta.TotalKaryawan)
	}
	if meta.RataRataKPI != 73.5 {
		t.Errorf("RataRataKPI = %.2f; want 73.50", meta.RataRataKPI)
	}
	if meta.KPITertinggi != 95 || meta.KaryawanTerbaik != "Budi" {
		t.Errorf("KPITertinggi = %.2f (%s); want 95.00 (Budi)", meta.KPITertinggi, meta.KaryawanTerbaik)
	}
}
