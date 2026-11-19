package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"trusmi_be/internal/models"
)

// KPIRepository mendefinisikan kontrak akses data MySQL untuk KPI Marketing.
type KPIRepository interface {
	Ping(ctx context.Context) error
	GetKPISummary(ctx context.Context, karyawanFilter string) ([]models.KPISummaryRow, error)
	GetOntimeLateSummary(ctx context.Context) (models.OntimeLateSummary, error)
	GetTasks(ctx context.Context, karyawanFilter, kpiFilter, statusFilter string) ([]models.TaskDetail, error)
}

type mysqlKPIRepository struct {
	db *sql.DB
}

// NewKPIRepository membuat instance repository berbasis MySQL database/sql.
func NewKPIRepository(db *sql.DB) KPIRepository {
	return &mysqlKPIRepository{db: db}
}

func (r *mysqlKPIRepository) Ping(ctx context.Context) error {
	if r.db == nil {
		return fmt.Errorf("database connection is nil")
	}
	return r.db.PingContext(ctx)
}

// GetKPISummary menjalankan query Soal 1 dari tabel hasil normalisasi (mst_karyawan, mst_kpi, trx_kpi_marketing).
func (r *mysqlKPIRepository) GetKPISummary(ctx context.Context, karyawanFilter string) ([]models.KPISummaryRow, error) {
	query := `
WITH rekap AS (
    SELECT
        k.id AS karyawan_id,
        k.nama_karyawan,
        kp.nama_kpi,
        kp.target_kpi,
        COUNT(t.aktual) AS actual,
        ROUND(
            COUNT(t.aktual) * 100.0 /
            NULLIF(kp.target_kpi, 0), 2
        ) AS pencapaian,
        ROUND(
            COUNT(t.aktual) * kp.bobot_kpi /
            NULLIF(kp.target_kpi, 0), 2
        ) AS bobot_actual,
        SUM(
            CASE WHEN t.aktual > t.deadline
                THEN 1 ELSE 0 END
        ) AS late_count,
        ROUND(
            SUM(
                CASE WHEN t.aktual > t.deadline
                    THEN 1 ELSE 0 END
            ) * kp.bobot_late, 2
        ) AS late
    FROM mst_karyawan k
    CROSS JOIN mst_kpi kp
    LEFT JOIN trx_kpi_marketing t
        ON t.karyawan_id = k.id
        AND t.kpi_id = kp.id
    GROUP BY
        k.id, k.nama_karyawan,
        kp.id, kp.nama_kpi,
        kp.target_kpi, kp.bobot_kpi, kp.bobot_late
)
SELECT
    karyawan_id,
    nama_karyawan AS nama,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Sales' THEN target_kpi END), 0) AS sales_target,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Sales' THEN actual END), 0) AS sales_actual,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Sales' THEN pencapaian END), 0) AS sales_pencapaian,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Sales' THEN bobot_actual END), 0) AS sales_bobot,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Sales' THEN late_count END), 0) AS sales_late_count,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Sales' THEN late END), 0) AS sales_late,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Sales' THEN ROUND(bobot_actual + late, 2) END), 0) AS sales_total_bobot,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Report' THEN target_kpi END), 0) AS report_target,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Report' THEN actual END), 0) AS report_actual,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Report' THEN pencapaian END), 0) AS report_pencapaian,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Report' THEN bobot_actual END), 0) AS report_bobot,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Report' THEN late_count END), 0) AS report_late_count,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Report' THEN late END), 0) AS report_late,
    COALESCE(MAX(CASE WHEN nama_kpi = 'Report' THEN ROUND(bobot_actual + late, 2) END), 0) AS report_total_bobot,
    COALESCE(ROUND(SUM(bobot_actual + late), 2), 0) AS total_kpi
FROM rekap
`
	var args []any
	if strings.TrimSpace(karyawanFilter) != "" {
		query += " WHERE LOWER(nama_karyawan) = LOWER(?) "
		args = append(args, strings.TrimSpace(karyawanFilter))
	}
	query += " GROUP BY karyawan_id, nama_karyawan ORDER BY karyawan_id;"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query kpi summary: %w", err)
	}
	defer rows.Close()

	var results []models.KPISummaryRow
	for rows.Next() {
		var row models.KPISummaryRow
		if err := rows.Scan(
			&row.KaryawanID,
			&row.Nama,
			&row.SalesTarget,
			&row.SalesActual,
			&row.SalesPencapaian,
			&row.SalesBobot,
			&row.SalesLateCount,
			&row.SalesLate,
			&row.SalesTotalBobot,
			&row.ReportTarget,
			&row.ReportActual,
			&row.ReportPencapaian,
			&row.ReportBobot,
			&row.ReportLateCount,
			&row.ReportLate,
			&row.ReportTotalBobot,
			&row.TotalKPI,
		); err != nil {
			return nil, fmt.Errorf("scan kpi summary row: %w", err)
		}
		results = append(results, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate kpi summary rows: %w", err)
	}

	return results, nil
}

// GetOntimeLateSummary menjalankan query Soal 2 beserta breakdown per karyawan dan per KPI.
func (r *mysqlKPIRepository) GetOntimeLateSummary(ctx context.Context) (models.OntimeLateSummary, error) {
	var summary models.OntimeLateSummary

	mainQuery := `
SELECT
    COUNT(*) AS total_task,
    COALESCE(SUM(CASE WHEN aktual <= deadline THEN 1 ELSE 0 END), 0) AS total_ontime,
    COALESCE(SUM(CASE WHEN aktual > deadline THEN 1 ELSE 0 END), 0) AS total_late,
    COALESCE(ROUND(
        SUM(CASE WHEN aktual <= deadline THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(*), 0), 2
    ), 0) AS persentase_ontime,
    COALESCE(ROUND(
        SUM(CASE WHEN aktual > deadline THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(*), 0), 2
    ), 0) AS persentase_late
FROM trx_kpi_marketing
WHERE aktual IS NOT NULL
  AND deadline IS NOT NULL;
`
	if err := r.db.QueryRowContext(ctx, mainQuery).Scan(
		&summary.TotalTask,
		&summary.TotalOntime,
		&summary.TotalLate,
		&summary.PersentaseOntime,
		&summary.PersentaseLate,
	); err != nil {
		return summary, fmt.Errorf("query ontime late summary: %w", err)
	}

	// Breakdown per karyawan
	empQuery := `
SELECT
    k.nama_karyawan AS nama,
    COUNT(t.id) AS total_task,
    COALESCE(SUM(CASE WHEN t.aktual <= t.deadline THEN 1 ELSE 0 END), 0) AS total_ontime,
    COALESCE(SUM(CASE WHEN t.aktual > t.deadline THEN 1 ELSE 0 END), 0) AS total_late,
    COALESCE(ROUND(
        SUM(CASE WHEN t.aktual <= t.deadline THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(t.id), 0), 2
    ), 0) AS persentase_ontime,
    COALESCE(ROUND(
        SUM(CASE WHEN t.aktual > t.deadline THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(t.id), 0), 2
    ), 0) AS persentase_late
FROM mst_karyawan k
LEFT JOIN trx_kpi_marketing t
    ON t.karyawan_id = k.id
    AND t.aktual IS NOT NULL
    AND t.deadline IS NOT NULL
GROUP BY k.id, k.nama_karyawan
ORDER BY k.id;
`
	empRows, err := r.db.QueryContext(ctx, empQuery)
	if err != nil {
		return summary, fmt.Errorf("query ontime late per karyawan: %w", err)
	}
	defer empRows.Close()

	for empRows.Next() {
		var item models.OntimeLateBreakdownKaryawan
		if err := empRows.Scan(
			&item.Nama,
			&item.TotalTask,
			&item.TotalOntime,
			&item.TotalLate,
			&item.PersentaseOntime,
			&item.PersentaseLate,
		); err != nil {
			return summary, fmt.Errorf("scan ontime late per karyawan: %w", err)
		}
		summary.PerKaryawan = append(summary.PerKaryawan, item)
	}

	// Breakdown per KPI
	kpiQuery := `
SELECT
    kp.nama_kpi AS kpi,
    COUNT(t.id) AS total_task,
    COALESCE(SUM(CASE WHEN t.aktual <= t.deadline THEN 1 ELSE 0 END), 0) AS total_ontime,
    COALESCE(SUM(CASE WHEN t.aktual > t.deadline THEN 1 ELSE 0 END), 0) AS total_late,
    COALESCE(ROUND(
        SUM(CASE WHEN t.aktual <= t.deadline THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(t.id), 0), 2
    ), 0) AS persentase_ontime,
    COALESCE(ROUND(
        SUM(CASE WHEN t.aktual > t.deadline THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(t.id), 0), 2
    ), 0) AS persentase_late
FROM mst_kpi kp
LEFT JOIN trx_kpi_marketing t
    ON t.kpi_id = kp.id
    AND t.aktual IS NOT NULL
    AND t.deadline IS NOT NULL
GROUP BY kp.id, kp.nama_kpi
ORDER BY kp.id;
`
	kpiRows, err := r.db.QueryContext(ctx, kpiQuery)
	if err != nil {
		return summary, fmt.Errorf("query ontime late per kpi: %w", err)
	}
	defer kpiRows.Close()

	for kpiRows.Next() {
		var item models.OntimeLateBreakdownKPI
		if err := kpiRows.Scan(
			&item.KPI,
			&item.TotalTask,
			&item.TotalOntime,
			&item.TotalLate,
			&item.PersentaseOntime,
			&item.PersentaseLate,
		); err != nil {
			return summary, fmt.Errorf("scan ontime late per kpi: %w", err)
		}
		summary.PerKPI = append(summary.PerKPI, item)
	}

	return summary, nil
}

// GetTasks menampilkan daftar detail task hasil relasi tabel normalisasi (mst_karyawan, mst_kpi, trx_kpi_marketing).
func (r *mysqlKPIRepository) GetTasks(ctx context.Context, karyawanFilter, kpiFilter, statusFilter string) ([]models.TaskDetail, error) {
	query := `
SELECT
    t.id,
    COALESCE(t.tasklist, '') AS tasklist,
    kp.nama_kpi AS kpi,
    k.nama_karyawan AS karyawan,
    COALESCE(DATE_FORMAT(t.deadline, '%Y-%m-%d'), '') AS deadline,
    COALESCE(DATE_FORMAT(t.aktual, '%Y-%m-%d'), '') AS aktual,
    CASE
        WHEN t.aktual IS NULL OR t.deadline IS NULL THEN 'Pending'
        WHEN t.aktual <= t.deadline THEN 'On-Time'
        ELSE 'Late'
    END AS status,
    COALESCE(DATEDIFF(t.aktual, t.deadline), 0) AS selisih_hari
FROM trx_kpi_marketing t
JOIN mst_karyawan k ON t.karyawan_id = k.id
JOIN mst_kpi kp ON t.kpi_id = kp.id
WHERE 1=1
`
	var args []any
	if strings.TrimSpace(karyawanFilter) != "" {
		query += " AND LOWER(k.nama_karyawan) = LOWER(?)"
		args = append(args, strings.TrimSpace(karyawanFilter))
	}
	if strings.TrimSpace(kpiFilter) != "" {
		query += " AND LOWER(kp.nama_kpi) = LOWER(?)"
		args = append(args, strings.TrimSpace(kpiFilter))
	}
	if strings.TrimSpace(statusFilter) != "" {
		st := strings.ToLower(strings.TrimSpace(statusFilter))
		if st == "on-time" || st == "ontime" {
			query += " AND t.aktual IS NOT NULL AND t.deadline IS NOT NULL AND t.aktual <= t.deadline"
		} else if st == "late" {
			query += " AND t.aktual IS NOT NULL AND t.deadline IS NOT NULL AND t.aktual > t.deadline"
		}
	}

	query += " ORDER BY CAST(SUBSTRING_INDEX(t.tasklist, ' ', -1) AS UNSIGNED) ASC, t.id ASC;"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.TaskDetail
	for rows.Next() {
		var item models.TaskDetail
		if err := rows.Scan(
			&item.ID,
			&item.Tasklist,
			&item.KPI,
			&item.Karyawan,
			&item.Deadline,
			&item.Aktual,
			&item.Status,
			&item.SelisihHari,
		); err != nil {
			return nil, fmt.Errorf("scan task detail: %w", err)
		}
		tasks = append(tasks, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate task details: %w", err)
	}

	return tasks, nil
}
