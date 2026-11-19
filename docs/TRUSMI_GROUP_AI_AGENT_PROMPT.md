# Lampiran Prompt AI — Pengerjaan Technical Test Programmer TRUSMI GROUP

Dokumen ini berisi daftar prompt yang saya gunakan saat meminta bantuan AI Assistant untuk mempercepat setup boilerplate backend, frontend dashboard grafik, serta konfigurasi Docker setelah saya menyelesaikan perancangan tabel normalisasi dan pengujian query SQL di MySQL lokal.

---

## Prompt 1 — Setup Awal Project Frontend (`trusmi_fe`)
> Untuk folder `trusmi_fe` saya ingin anda melakukan setup menggunakan React + Vite, coba buatkan dulu struktur awalnya.

---

## Prompt 2 — Merapikan Dokumentasi Query SQL yang Sudah Saya Uji di MySQL
> Saya sekarang sedang melakukan dokumentasi untuk query MySQL yang sudah saya uji di database `db_test_programmer_trusmi` (Soal 1 perhitungan KPI Sales & Report, Soal 2 perhitungan jumlah dan persentase tasklist ontime & late, serta Soal 5 proses normalisasi tabel `table_kpi_marketing` menjadi `mst_karyawan`, `mst_kpi`, dan `trx_kpi_marketing`). Tolong buatkan file `.txt` dan `.md`-nya di dalam folder `docs` berdasarkan hasil query MySQL yang sudah saya jalankan.

---

## Prompt 3 — Pembuatan REST API Backend (`trusmi_be`)
> Tolong buatkan backend di folder `trusmi_be` menggunakan Golang (GoFiber) yang terhubung ke database MySQL lokal saya (`db_test_programmer_trusmi`, user `root`). Buat 2 route `GET` khusus untuk soal dashboard saja:
> - `GET /api/v1/dashboard-1` (mengambil data hasil query Soal 1 untuk grafik Soal 3)
> - `GET /api/v1/dashboard-2` (mengambil data hasil query Soal 2 untuk grafik Soal 4)

---

## Prompt 4 — Pembuatan Tampilan Dashboard Grafik Soal 3 dan Soal 4 (`trusmi_fe`)
> Tolong terapkan tampilan dashboard grafik di `trusmi_fe` menggunakan React + Vite dan Recharts yang mengambil data dari endpoint `GET /api/v1/dashboard-1` dan `GET /api/v1/dashboard-2`. Bagi menjadi 2 tab sederhana:
> - **Dashboard 1** (Soal 3): menampilkan summary, grafik perbandingan bobot Sales vs Report, grafik Total KPI per karyawan, dan tabel perhitungan KPI sesuai format tabel soal.
> - **Dashboard 2** (Soal 4): menampilkan summary, grafik persentase/jumlah tasklist Ontime vs Late, grafik Ontime vs Late per karyawan, dan tabel data tasklist.
> Buat tampilannya bersih, simpel, dan natural tanpa teks atau elemen dekorasi yang berlebihan.

---

## Prompt 5 — Konfigurasi Docker Development dan Production (`docker/`)
> Buatkan konfigurasi Dockerfile dan Docker Compose di folder `docker/dev` (untuk menjalankan backend port `8090` dan frontend Vite di lokal yang terhubung ke MySQL lokal) serta `docker/prod` (multi-stage build untuk deployment di VPS Ubuntu menggunakan Nginx dan koneksi ke MySQL container).
