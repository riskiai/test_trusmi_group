# Dokumentasi REST API — TRUSMI GROUP Marketing KPI (`trusmi_be`)

Base URL Lokal: `http://localhost:8090`  
Base URL Production: `https://api-trusmi.riskiahmadilham.my.id`

Seluruh endpoint berada di bawah prefix `/api/v1` dan bersifat **read-only** (`GET`).

---

## 1. Health Check
Memeriksa status layanan backend dan konektivitas pool database MySQL (`db_test_programmer_trusmi`).

- **Endpoint:** `GET /api/v1/health`
- **Response `200 OK`:**
  ```json
  {
    "database": "connected",
    "service": "trusmi-be",
    "status": "ok"
  }
  ```

---

## 2. Rekapitulasi KPI Marketing per Karyawan (Soal 1 & Soal 3)
Mengembalikan hasil perhitungan KPI Marketing (`Sales` dan `Report`) untuk setiap karyawan beserta metadata ringkasan.

- **Endpoint:** `GET /api/v1/kpi/summary`
- **Query Parameter Opsional:**
  - `karyawan` (string): Filter berdasarkan nama karyawan (contoh: `?karyawan=Budi`)
- **Response `200 OK`:**
  ```json
  {
    "status": "success",
    "data": [
      {
        "karyawan_id": 1,
        "nama": "Budi",
        "sales_target": 2,
        "sales_actual": 2,
        "sales_pencapaian": 100,
        "sales_bobot": 50,
        "sales_late_count": 0,
        "sales_late": 0,
        "sales_total_bobot": 50,
        "report_target": 2,
        "report_actual": 2,
        "report_pencapaian": 100,
        "report_bobot": 50,
        "report_late_count": 1,
        "report_late": -5,
        "report_total_bobot": 45,
        "total_kpi": 95
      },
      {
        "karyawan_id": 2,
        "nama": "Adi",
        "sales_target": 2,
        "sales_actual": 2,
        "sales_pencapaian": 100,
        "sales_bobot": 50,
        "sales_late_count": 1,
        "sales_late": -7,
        "sales_total_bobot": 43,
        "report_target": 2,
        "report_actual": 2,
        "report_pencapaian": 100,
        "report_bobot": 50,
        "report_late_count": 0,
        "report_late": 0,
        "report_total_bobot": 50,
        "total_kpi": 93
      },
      {
        "karyawan_id": 3,
        "nama": "Rara",
        "sales_target": 2,
        "sales_actual": 2,
        "sales_pencapaian": 100,
        "sales_bobot": 50,
        "sales_late_count": 1,
        "sales_late": -7,
        "sales_total_bobot": 43,
        "report_target": 2,
        "report_actual": 1,
        "report_pencapaian": 50,
        "report_bobot": 25,
        "report_late_count": 1,
        "report_late": -5,
        "report_total_bobot": 20,
        "total_kpi": 63
      },
      {
        "karyawan_id": 4,
        "nama": "Doni",
        "sales_target": 2,
        "sales_actual": 1,
        "sales_pencapaian": 50,
        "sales_bobot": 25,
        "sales_late_count": 1,
        "sales_late": -7,
        "sales_total_bobot": 18,
        "report_target": 2,
        "report_actual": 1,
        "report_pencapaian": 50,
        "report_bobot": 25,
        "report_late_count": 0,
        "report_late": 0,
        "report_total_bobot": 25,
        "total_kpi": 43
      }
    ],
    "meta": {
      "total_karyawan": 4,
      "rata_rata_kpi": 73.5,
      "kpi_tertinggi": 95,
      "karyawan_terbaik": "Budi",
      "kpi_terendah": 43
    }
  }
  ```

---

## 3. Ringkasan Tasklist On-Time & Late (Soal 2 & Soal 4)
Mengembalikan jumlah dan persentase tasklist yang selesai tepat waktu (`aktual <= deadline`) maupun terlambat (`aktual > deadline`), dilengkapi breakdown per karyawan dan per kategori KPI.

- **Endpoint:** `GET /api/v1/kpi/ontime-late`
- **Response `200 OK`:**
  ```json
  {
    "status": "success",
    "data": {
      "total_task": 13,
      "total_ontime": 8,
      "total_late": 5,
      "persentase_ontime": 61.54,
      "persentase_late": 38.46,
      "per_karyawan": [
        {
          "nama": "Budi",
          "total_task": 4,
          "total_ontime": 3,
          "total_late": 1,
          "persentase_ontime": 75,
          "persentase_late": 25
        }
      ],
      "per_kpi": [
        {
          "kpi": "Sales",
          "total_task": 7,
          "total_ontime": 4,
          "total_late": 3,
          "persentase_ontime": 57.14,
          "persentase_late": 42.86
        }
      ]
    }
  }
  ```

---

## 4. Daftar Detail Tasklist Hasil Relasi Normalisasi (Soal 5)
Mengembalikan daftar seluruh tasklist dari hasil `JOIN` tabel `trx_kpi_marketing`, `mst_karyawan`, dan `mst_kpi`.

- **Endpoint:** `GET /api/v1/kpi/tasks`
- **Query Parameter Opsional:**
  - `karyawan` (string): Filter nama karyawan (`Budi`, `Adi`, `Rara`, `Doni`)
  - `kpi` (string): Filter jenis KPI (`Sales`, `Report`)
  - `status` (string): Filter status ketepatan waktu (`On-Time`, `Late`)
- **Response `200 OK`:**
  ```json
  {
    "status": "success",
    "total": 13,
    "data": [
      {
        "id": 4,
        "tasklist": "Tasklist 1",
        "kpi": "Sales",
        "karyawan": "Budi",
        "deadline": "2022-01-10",
        "aktual": "2022-01-09",
        "status": "On-Time",
        "selisih_hari": -1
      }
    ]
  }
  ```
