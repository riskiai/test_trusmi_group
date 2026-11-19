# Technical Test Programmer — TRUSMI GROUP

Repository ini berisi solusi lengkap **Technical Test Programmer TRUSMI GROUP** yang mencakup Query SQL & Normalisasi Database (MySQL), REST API Backend (Golang + GoFiber v3), Frontend Dashboard (React + Vite), serta konfigurasi Docker Development & Production.

## Struktur Project

```text
TRUSMI_GROUP/
├── docker/
│   ├── dev/
│   │   ├── backend.Dockerfile
│   │   ├── frontend.Dockerfile
│   │   └── docker-compose.dev.yml
│   └── prod/
│       ├── backend.Dockerfile
│       ├── frontend.Dockerfile
│       ├── docker-compose.prod.yml
│       └── nginx.conf
├── docs/
│   ├── soal_1_kpi_query.txt
│   ├── soal_2_ontime_late_query.txt
│   ├── soal_5_normalisasi.txt
│   ├── jawaban_soal_1_2_5.txt
│   ├── database_schema.sql
│   ├── DOKUMENTASI_SQL_SOAL_1_2_5.md
│   ├── API_DOCUMENTATION.md
│   └── DOCKER_DEPLOYMENT.md
├── trusmi_be/                      # Backend Golang + GoFiber v3
└── trusmi_fe/                      # Frontend React + Vite
```

## Endpoint REST API (`trusmi_be`)

- `GET /api/v1/health` — Pemeriksaan status service dan koneksi database MySQL.
- `GET /api/v1/dashboard-1` (atau `GET /api/v1/kpi/summary`) — Hasil perhitungan KPI Sales & Report per karyawan (Soal 1 & Soal 3).
- `GET /api/v1/dashboard-2` (atau `GET /api/v1/kpi/ontime-late`) — Hasil perhitungan jumlah & persentase tasklist On-Time dan Late (Soal 2 & Soal 4).
- `GET /api/v1/kpi/tasks` — Daftar detail 13 tasklist hasil relasi tabel normalisasi (Soal 5).

## Cara Menjalankan dengan Docker

- **Development (Lokal):**
  ```bash
  docker compose -f docker/dev/docker-compose.dev.yml up -d --build
  ```
- **Production (VPS):**
  Lihat panduan lengkap di [docs/DOCKER_DEPLOYMENT.md](docs/DOCKER_DEPLOYMENT.md).
