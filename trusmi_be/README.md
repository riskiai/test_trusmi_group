# TRUSMI GROUP — Marketing KPI Backend (`trusmi_be`)

REST API Backend menggunakan **Golang + GoFiber v3** dan **MySQL** (`database/sql` + `go-sql-driver/mysql`) untuk menyajikan data perhitungan KPI Marketing (Soal 3 / Dashboard 1 dan Soal 4 / Dashboard 2).

## Endpoint REST API (`/api/v1`)

| Method | Endpoint | Deskripsi |
|---|---|---|
| `GET` | `/api/v1/dashboard-1` | Hasil perhitungan KPI Sales & Report per karyawan (Soal 1 & Soal 3) |
| `GET` | `/api/v1/dashboard-2` | Hasil perhitungan jumlah & persentase tasklist Ontime vs Late beserta detail tasklist (Soal 2 & Soal 4) |

## Cara Menjalankan secara Lokal

```bash
go mod tidy
go test ./...
go run ./cmd/api
```
