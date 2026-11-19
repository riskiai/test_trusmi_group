package models

// KPISummaryRow merepresentasikan satu baris rekapitulasi KPI per karyawan (Dashboard 1 - Soal 1 & 3).
type KPISummaryRow struct {
	KaryawanID       int     `json:"karyawan_id"`
	Nama             string  `json:"nama"`
	SalesTarget      int     `json:"sales_target"`
	SalesActual      int     `json:"sales_actual"`
	SalesPencapaian  float64 `json:"sales_pencapaian"`
	SalesBobot       float64 `json:"sales_bobot"`
	SalesLateCount   int     `json:"sales_late_count"`
	SalesLate        float64 `json:"sales_late"`
	SalesTotalBobot  float64 `json:"sales_total_bobot"`
	ReportTarget     int     `json:"report_target"`
	ReportActual     int     `json:"report_actual"`
	ReportPencapaian float64 `json:"report_pencapaian"`
	ReportBobot      float64 `json:"report_bobot"`
	ReportLateCount  int     `json:"report_late_count"`
	ReportLate       float64 `json:"report_late"`
	ReportTotalBobot float64 `json:"report_total_bobot"`
	TotalKPI         float64 `json:"total_kpi"`
}

// KPISummaryMeta memuat ringkasan agregat untuk Dashboard 1.
type KPISummaryMeta struct {
	TotalKaryawan   int     `json:"total_karyawan"`
	RataRataKPI     float64 `json:"rata_rata_kpi"`
	KPITertinggi    float64 `json:"kpi_tertinggi"`
	KaryawanTerbaik string  `json:"karyawan_terbaik"`
	KPITerendah     float64 `json:"kpi_terendah"`
}

// KPISummaryResponse adalah payload JSON untuk GET /api/v1/dashboard-1.
type KPISummaryResponse struct {
	Status string          `json:"status"`
	Data   []KPISummaryRow `json:"data"`
	Meta   KPISummaryMeta  `json:"meta"`
}

// OntimeLateBreakdownKaryawan merepresentasikan ketepatan waktu per karyawan.
type OntimeLateBreakdownKaryawan struct {
	Nama             string  `json:"nama"`
	TotalTask        int     `json:"total_task"`
	TotalOntime      int     `json:"total_ontime"`
	TotalLate        int     `json:"total_late"`
	PersentaseOntime float64 `json:"persentase_ontime"`
	PersentaseLate   float64 `json:"persentase_late"`
}

// OntimeLateBreakdownKPI merepresentasikan ketepatan waktu per kategori KPI (Sales / Report).
type OntimeLateBreakdownKPI struct {
	KPI              string  `json:"kpi"`
	TotalTask        int     `json:"total_task"`
	TotalOntime      int     `json:"total_ontime"`
	TotalLate        int     `json:"total_late"`
	PersentaseOntime float64 `json:"persentase_ontime"`
	PersentaseLate   float64 `json:"persentase_late"`
}

// TaskDetail merepresentasikan satu baris tasklist.
type TaskDetail struct {
	ID          int    `json:"id"`
	Tasklist    string `json:"tasklist"`
	KPI         string `json:"kpi"`
	Karyawan    string `json:"karyawan"`
	Deadline    string `json:"deadline"`
	Aktual      string `json:"aktual"`
	Status      string `json:"status"`
	SelisihHari int    `json:"selisih_hari"`
}

// OntimeLateSummary merepresentasikan hasil perhitungan Dashboard 2 (Soal 2 & 4).
type OntimeLateSummary struct {
	TotalTask        int                           `json:"total_task"`
	TotalOntime      int                           `json:"total_ontime"`
	TotalLate        int                           `json:"total_late"`
	PersentaseOntime float64                       `json:"persentase_ontime"`
	PersentaseLate   float64                       `json:"persentase_late"`
	PerKaryawan      []OntimeLateBreakdownKaryawan `json:"per_karyawan"`
	PerKPI           []OntimeLateBreakdownKPI      `json:"per_kpi"`
	Tasks            []TaskDetail                  `json:"tasks"`
}

// OntimeLateResponse adalah payload JSON untuk GET /api/v1/dashboard-2.
type OntimeLateResponse struct {
	Status string            `json:"status"`
	Data   OntimeLateSummary `json:"data"`
}
