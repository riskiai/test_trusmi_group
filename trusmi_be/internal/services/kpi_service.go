package services

import (
	"context"
	"math"
	"time"

	"trusmi_be/internal/models"
	"trusmi_be/internal/repositories"
)

// KPIService mengelola business logic KPI Marketing.
type KPIService struct {
	repo repositories.KPIRepository
}

// NewKPIService membuat instance KPIService baru.
func NewKPIService(repo repositories.KPIRepository) *KPIService {
	return &KPIService{repo: repo}
}

// CheckHealth memeriksa koneksi database dengan timeout 3 detik.
func (s *KPIService) CheckHealth(ctx context.Context) error {
	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return s.repo.Ping(pingCtx)
}

// GetKPISummary mengambil rekap KPI per karyawan dan menghitung metadata ringkasan.
func (s *KPIService) GetKPISummary(ctx context.Context, karyawanFilter string) (models.KPISummaryResponse, error) {
	qCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := s.repo.GetKPISummary(qCtx, karyawanFilter)
	if err != nil {
		return models.KPISummaryResponse{}, err
	}
	if rows == nil {
		rows = []models.KPISummaryRow{}
	}

	meta := CalculateSummaryMeta(rows)

	return models.KPISummaryResponse{
		Status: "success",
		Data:   rows,
		Meta:   meta,
	}, nil
}

// GetOntimeLateSummary mengambil ringkasan ketepatan waktu (Soal 2).
func (s *KPIService) GetOntimeLateSummary(ctx context.Context) (models.OntimeLateResponse, error) {
	qCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	data, err := s.repo.GetOntimeLateSummary(qCtx)
	if err != nil {
		return models.OntimeLateResponse{}, err
	}
	if data.PerKaryawan == nil {
		data.PerKaryawan = []models.OntimeLateBreakdownKaryawan{}
	}
	if data.PerKPI == nil {
		data.PerKPI = []models.OntimeLateBreakdownKPI{}
	}

	return models.OntimeLateResponse{
		Status: "success",
		Data:   data,
	}, nil
}

// GetTasks mengambil daftar detail tasklist beserta status On-Time / Late.
func (s *KPIService) GetTasks(ctx context.Context, karyawanFilter, kpiFilter, statusFilter string) (models.TaskListResponse, error) {
	qCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tasks, err := s.repo.GetTasks(qCtx, karyawanFilter, kpiFilter, statusFilter)
	if err != nil {
		return models.TaskListResponse{}, err
	}
	if tasks == nil {
		tasks = []models.TaskDetail{}
	}

	return models.TaskListResponse{
		Status: "success",
		Data:   tasks,
		Total:  len(tasks),
	}, nil
}

// CalculateSummaryMeta menghitung rata-rata KPI, skor tertinggi, dan karyawan terbaik.
func CalculateSummaryMeta(rows []models.KPISummaryRow) models.KPISummaryMeta {
	if len(rows) == 0 {
		return models.KPISummaryMeta{}
	}

	var total float64
	highest := rows[0].TotalKPI
	lowest := rows[0].TotalKPI
	bestName := rows[0].Nama

	for _, r := range rows {
		total += r.TotalKPI
		if r.TotalKPI > highest {
			highest = r.TotalKPI
			bestName = r.Nama
		}
		if r.TotalKPI < lowest {
			lowest = r.TotalKPI
		}
	}

	avg := Round2(total / float64(len(rows)))

	return models.KPISummaryMeta{
		TotalKaryawan:   len(rows),
		RataRataKPI:     avg,
		KPITertinggi:    Round2(highest),
		KaryawanTerbaik: bestName,
		KPITerendah:     Round2(lowest),
	}
}

// CalculateKPIMetric adalah fungsi murni untuk memverifikasi rumus business rule KPI pada unit test.
func CalculateKPIMetric(target, actual, lateCount int, bobotKPI, bobotLate float64) (pencapaian, bobotActual, penaltiLate, totalBobot float64) {
	if target <= 0 {
		return 0, 0, 0, 0
	}
	pencapaian = Round2(float64(actual) * 100.0 / float64(target))
	bobotActual = Round2(float64(actual) * bobotKPI / float64(target))
	penaltiLate = Round2(float64(lateCount) * bobotLate)
	totalBobot = Round2(bobotActual + penaltiLate)
	return
}

// ClassifyTaskStatus menentukan apakah task On-Time, Late, atau Pending berdasarkan tanggal YYYY-MM-DD.
func ClassifyTaskStatus(deadline, aktual string) string {
	if deadline == "" || aktual == "" {
		return "Pending"
	}
	if aktual <= deadline {
		return "On-Time"
	}
	return "Late"
}

// Round2 membulatkan angka float64 ke 2 angka di belakang koma.
func Round2(val float64) float64 {
	return math.Round(val*100) / 100
}
