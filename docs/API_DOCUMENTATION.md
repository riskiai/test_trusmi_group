# Dokumentasi REST API — TRUSMI GROUP (`trusmi_be`)

Base URL Lokal: `http://localhost:8090`  
Base URL Production: `https://trusmi.riskiahmadiilham.my.id`

Backend menyediakan **2 endpoint `GET`** untuk menyajikan data grafik dan tabel **Dashboard 1** (Soal 1 & 3) serta **Dashboard 2** (Soal 2 & 4).

---

## 1. Dashboard 1 — Perhitungan KPI Marketing per Karyawan (Soal 1 & Soal 3)

- **Method:** `GET`
- **Endpoint:** `/api/v1/dashboard-1`
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

## 2. Dashboard 2 — Tasklist Ontime & Late Beserta Persentasenya (Soal 2 & Soal 4)

- **Method:** `GET`
- **Endpoint:** `/api/v1/dashboard-2`
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
        { "nama": "Budi", "total_task": 4, "total_ontime": 3, "total_late": 1, "persentase_ontime": 75, "persentase_late": 25 },
        { "nama": "Adi", "total_task": 4, "total_ontime": 3, "total_late": 1, "persentase_ontime": 75, "persentase_late": 25 },
        { "nama": "Rara", "total_task": 3, "total_ontime": 1, "total_late": 2, "persentase_ontime": 33.33, "persentase_late": 66.67 },
        { "nama": "Doni", "total_task": 2, "total_ontime": 1, "total_late": 1, "persentase_ontime": 50, "persentase_late": 50 }
      ],
      "per_kpi": [
        { "kpi": "Sales", "total_task": 7, "total_ontime": 4, "total_late": 3, "persentase_ontime": 57.14, "persentase_late": 42.86 },
        { "kpi": "Report", "total_task": 6, "total_ontime": 4, "total_late": 2, "persentase_ontime": 66.67, "persentase_late": 33.33 }
      ],
      "tasks": [
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
  }
  ```
