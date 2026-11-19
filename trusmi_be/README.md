# TRUSMI GROUP — Marketing KPI Backend (`trusmi_be`)

REST API Backend menggunakan **Golang + GoFiber v3** dan **MySQL** (`database/sql` + `go-sql-driver/mysql`) untuk menyajikan perhitungan KPI Marketing (Soal 1–5 Technical Test Programmer TRUSMI GROUP).

## Struktur Folder

```text
trusmi_be/
├── cmd/
│   └── api/
│       └── main.go                   # Entrypoint aplikasi Fiber v3
├── internal/
│   ├── config/
│   │   └── config.go                 # Pembacaan ENV & .env
│   ├── database/
│   │   └── mysql.go                  # Koneksi pool MySQL (database/sql)
│   ├── handlers/
│   │   ├── kpi_handler.go            # HTTP handlers
│   │   └── kpi_handler_test.go       # Unit test HTTP endpoint Fiber v3
│   ├── middleware/
│   │   └── middleware.go             # Recover, Logger, CORS allowlist
│   ├── models/
│   │   └── kpi.go                    # Structs & JSON DTO
│   ├── repositories/
│   │   └── kpi_repository.go         # Query SQL eksplisit ke MySQL
│   ├── routes/
│   │   └── routes.go                 # Registrasi endpoint /api/v1
│   └── services/
│       ├── kpi_service.go            # Business logic KPI
│       └── kpi_service_test.go       # Table-driven unit test rumus KPI
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Endpoint REST API (`/api/v1`)

| Method | Endpoint | Deskripsi |
|---|---|---|
| `GET` | `/api/v1/health` | Pemeriksaan status service dan koneksi database MySQL |
| `GET` | `/api/v1/kpi/summary` | Rekapitulasi perhitungan KPI Sales & Report per karyawan (Soal 1 & Soal 3) |
| `GET` | `/api/v1/kpi/ontime-late` | Perhitungan jumlah & persentase task On-Time vs Late beserta breakdown per karyawan & per KPI (Soal 2 & Soal 4) |
| `GET` | `/api/v1/kpi/tasks` | Daftar detail 13 tasklist hasil relasi tabel normalisasi beserta status On-Time / Late (Soal 5) |

## Cara Menjalankan secara Lokal

1. Salin file `.env.example` menjadi `.env` dan sesuaikan kredensial MySQL lokal Anda:
   ```bash
   cp .env.example .env
   ```
2. Unduh dependensi, jalankan unit test, dan jalankan server:
   ```bash
   go mod tidy
   go test ./...
   go vet ./...
   go run ./cmd/api
   ```
3. Build binary:
   ```bash
   go build -o bin/trusmi-api ./cmd/api
   ```
