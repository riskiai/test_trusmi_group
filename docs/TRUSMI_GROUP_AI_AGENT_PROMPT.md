# TRUSMI GROUP — Panduan Pengerjaan Dashboard & Docker Deployment

**Dokumen pengganti prompt sebelumnya**  
**Fokus saat ini:** menyelesaikan **tampilan/dashboard React Vite**, mengintegrasikannya dengan backend **Go + Fiber** dan database **MySQL** yang sudah disiapkan, serta menyiapkan **satu konfigurasi Docker Compose terpusat** untuk deployment satu VPS Ubuntu.

> **UNTUK AI CODING AGENT:** Baca dokumen ini sebagai instruksi kerja. Lakukan audit project yang sudah ada terlebih dahulu, lalu kerjakan bagian yang belum selesai. Jangan membuat ulang database, SQL, atau backend yang telah dikerjakan. Jangan menjalankan deployment production atau mengubah layanan VPS tanpa persetujuan pemilik project.

## 1. Konteks Technical Test

Referensi soal: <https://docs.google.com/spreadsheets/d/1qt1oQiZmHN58f4ArzpWWzqSEcMfCUaatbTy8CIoaIzw/edit?gid=0#gid=0>

Tes terdiri dari lima kebutuhan:

1. SQL perhitungan KPI Marketing per karyawan.
2. SQL jumlah dan persentase task **on-time** dan **late**.
3. Dashboard grafik untuk hasil SQL soal 1.
4. Dashboard grafik untuk hasil SQL soal 2.
5. Dokumentasi normalisasi `table_kpi_marketing`.

**Status dari informasi pemilik project:**
- [x] Tabel original dan tabel normalisasi sudah dibuat di MySQL.
- [x] Data 13 task sudah tersedia.
- [x] Query SQL soal 1 dan 2 sudah disusun dalam pembahasan sebelumnya; **agent tetap harus mengecek apakah sudah disimpan dan benar saat dieksekusi**.
- [x] Struktur master dan transaksi untuk soal 5 sudah tersedia; **bukti relasi/normalisasi tetap perlu diverifikasi**.
- [ ] **Prioritas utama:** buat/selesaikan tampilan dashboard React untuk soal 3 dan 4.
- [ ] Hubungkan dashboard ke API Go Fiber yang ada; buat endpoint minimal hanya bila belum tersedia.
- [ ] Siapkan Docker Compose terpusat agar FE, BE, dan MySQL dapat berjalan pada satu VPS.
- [ ] Testing, dokumentasi pengumpulan, dan screenshot dashboard yang benar-benar berjalan.

**Bukan prioritas:** authentication, user management, CRUD kompleks, Redis, Kafka, CI/CD, fitur SaaS, atau penggantian framework.

## 2. Stack dan Target Deployment

| Lapisan | Pilihan | Catatan |
|---|---|---|
| Frontend | React + Vite (JavaScript) | Dashboard, grafik dan tabel; boleh Recharts |
| Backend | Golang + GoFiber | Ikuti versi Fiber yang sudah digunakan; utamakan v3 bila kompatibel |
| Database | MySQL | `db_test_programmer_trusmi` |
| Container | Docker + **satu** Docker Compose | Satu orchestration stack berisi FE, BE, MySQL |
| Reverse proxy publik | Nginx **di host VPS** | Pertahankan Nginx yang sudah ada jika digunakan aplikasi lain |
| Hosting | Satu VPS Ubuntu | Tidak menggunakan Vercel |
| DNS Frontend | `trusmi.riskiahmadilham.my.id` | A record ke IP publik VPS |
| DNS Backend | `api-trusmi.riskiahmadilham.my.id` | A record ke IP publik VPS |
| Git | Dua repo terpisah | Repo `trusmi_fe` dan `trusmi_be` |

**Penting:** Go tidak harus terpasang dengan versi yang sama di host VPS jika backend **dibuild di Docker**. Yang wajib kompatibel adalah versi Go pada *builder image* dengan `go.mod` dan versi Fiber. Saat dokumentasi ini disusun, dokumentasi resmi Fiber v3 mencantumkan syarat Go **1.26+**; agent wajib membaca ulang `go.mod` dan dokumentasi resmi sebelum memilih image builder. Jangan menaikkan versi Go pada VPS host tanpa alasan dan tanpa persetujuan.

Referensi: <https://docs.gofiber.io/>

## 3. Struktur Folder Final — SATU Folder `docker/` di Root

**Wajib ikuti susunan berikut.** Folder `docker` **bukan** di dalam `trusmi_fe`, `trusmi_be`, atau `docs`.

```text
TRUSMI_GROUP/
├── docs/
│   ├── TRUSMI_GROUP_AI_AGENT_PROMPT.md
│   ├── soal_1_kpi_query.txt           # jika sudah ada, gunakan ulang
│   ├── soal_2_ontime_late_query.txt   # jika sudah ada, gunakan ulang
│   ├── soal_5_normalisasi.txt         # jika sudah ada, gunakan ulang
│   ├── API_DOCUMENTATION.md
│   ├── DEPLOYMENT_GUIDE.md
│   └── screenshots/
│       ├── soal_3_dashboard_kpi.png
│       └── soal_4_dashboard_ontime_late.png
│
├── trusmi_be/                         # Git repo Backend (jangan reset)
│   ├── cmd/api/main.go                 # contoh, sesuaikan kondisi existing
│   ├── internal/
│   ├── go.mod
│   ├── go.sum
│   └── .env.example
│
├── trusmi_fe/                         # Git repo Frontend (sudah Vite)
│   ├── public/
│   ├── src/
│   ├── package.json
│   ├── package-lock.json
│   └── vite.config.js
│
├── docker/                             # SATU folder Docker untuk seluruh stack
│   ├── docker-compose.yml             # HANYA SATU file Compose utama
│   ├── backend.Dockerfile
│   ├── frontend.Dockerfile
│   ├── frontend.nginx.conf
│   ├── .env.prod.example              # template aman, tanpa password asli
│   └── .env.prod                      # rahasia, tidak boleh commit
│
└── .dockerignore                      # shared build context
```

**Interpretasi “Docker satu saja”:** hanya **satu folder `docker/` + satu `docker-compose.yml`** untuk mengelola tiga service. **Tetap membutuhkan Dockerfile FE dan BE masing-masing** karena bahasa/build prosesnya berbeda. Tidak perlu Docker Compose terpisah di folder FE atau BE.

**Catatan dua repository:** karena `docker/` berada di luar dua repo, folder ini tidak ikut ter-push saat `git push` dari FE/BE. Sertakan prosedur `scp`/`rsync` untuk menyalin folder Docker dan docs ke VPS, atau simpan folder infrastruktur pada repositori tersendiri **hanya jika pemilik project menyetujui**. Jangan otomatis membuat repo ketiga.

## 4. Tahap Awal — Audit Tanpa Menimpa File

Sebelum coding:

1. Baca struktur existing `trusmi_fe/src`, `package.json`, dan konfigurasi routing/styling.
2. Baca `trusmi_be/go.mod`, route, handler, service, repository, port, dan konfigurasi MySQL.
3. Periksa API yang **benar-benar ada** dan tes responsnya. Jangan menganggap endpoint tersedia jika belum dilihat.
4. Verifikasi data MySQL original dan normalisasi, jumlah task serta hasil SQL; **jangan melakukan DROP/TRUNCATE/RESET**.
5. Periksa apakah ada Dockerfile/Compose existing sebelum membuat file baru, lalu konsolidasikan menjadi satu folder `docker/` tanpa menghapus konfigurasi yang diperlukan.
6. Di VPS, identifikasi port serta Nginx/container yang sudah berjalan. **Jangan bind 80/443 dari container** jika sudah digunakan aplikasi lain.
7. Buat perubahan paling kecil yang menyelesaikan dashboard dan deployment. Jangan refactor backend secara luas.

Ringkas hasil audit dengan kategori: `SUDAH ADA`, `BELUM ADA`, `PERLU DISESUAIKAN` sebelum implementasi lanjutan.

## 5. Data & Aturan KPI — Jangan Berubah

**Database:** `db_test_programmer_trusmi`

Tabel yang sudah dimiliki pemilik project:

- `table_kpi_marketing(id, tasklist, kpi, karyawan, deadline, aktual)` — data asli soal.
- `mst_karyawan(id, nama_karyawan)`.
- `mst_kpi(id, nama_kpi, target_kpi, bobot_kpi, bobot_late)`.
- `trx_kpi_marketing(id, tasklist, karyawan_id, kpi_id, deadline, aktual)`.

Isi `mst_kpi` yang sudah dikonfirmasi:

| KPI | Target | Bobot | Penalti per task late |
|---|---:|---:|---:|
| Sales | 2 | 50.00 | -7.00 |
| Report | 2 | 50.00 | -5.00 |

Rumus:

- On-time jika `aktual <= deadline`.
- Late jika `aktual > deadline`.
- Actual KPI = banyak task selesai dengan `aktual` terisi.
- Pencapaian = `actual / target * 100`.
- Bobot aktual = `actual / target * bobot_kpi`.
- Penalti late = `jumlah late * bobot_late` (nilai `bobot_late` di DB sudah negatif).
- Total bobot KPI = `bobot aktual + penalti late`.
- Total KPI per karyawan = total bobot Sales + Report.

**Asumsi penilaian:** penalti dikenakan per task yang terlambat. Jika target terlampaui, hasil dihitung sesuai rumus di atas kecuali ada aturan pembatasan yang dikonfirmasi. Nilai yang berpotensi negatif tidak boleh diam-diam diubah menjadi nol tanpa aturan soal.

**Expected data test (bukan hardcode hasil):**

| Nama | Sales | Report | Total KPI |
|---|---:|---:|---:|
| Budi | 50% | 45% | **95%** |
| Adi | 43% | 50% | **93%** |
| Rara | 43% | 20% | **63%** |
| Doni | 18% | 25% | **43%** |

Task keseluruhan: **13**, on-time: **8 (61,54%)**, late: **5 (38,46%)**.

**Ketentuan soal 1:** tabel hasil harus menampilkan target, actual, pencapaian, bobot, late, total bobot untuk masing-masing Sales/Report, kemudian total KPI. Tampilkan persen dengan benar di UI. **Jangan hardcode hasil SQL di frontend**.

## 6. Prioritas Utama — Dashboard Frontend React Vite

Gunakan struktur existing; contoh modularisasi (jangan paksa mengganti proyek yang sudah ada):

```text
trusmi_fe/src/
├── components/
│   ├── layout/
│   ├── common/
│   ├── charts/
│   └── tables/
├── pages/
│   └── Dashboard.jsx
├── services/
│   └── api.js
├── hooks/
│   └── useDashboardData.js
├── utils/
│   └── formatters.js
├── App.jsx
└── main.jsx
```

### Fitur wajib untuk Soal 3

- Header dashboard: **TRUSMI GROUP — Marketing KPI Dashboard**.
- Summary cards: total karyawan, rata-rata skor KPI, total task, skor KPI tertinggi (hanya jika data tersedia).
- Bar chart: **Sales vs Report** per karyawan (total bobot setelah late).
- Bar chart: **Total KPI** per karyawan.
- Tabel **KPI Marketing** dengan kolom kelompok Sales, Report, dan KPI akhir; gunakan scroll horizontal di layar kecil agar tidak menyembunyikan detail.
- Tooltip grafik, label persen, dan legenda yang konsisten.

### Fitur wajib untuk Soal 4

- Summary cards: total task, on-time, late, persentase on-time dan late.
- Pie/donut chart: **On-time vs Late**.
- Bar chart opsional: **On-time vs Late per karyawan**, jika endpoint tersedia.
- Tabel ringkas task status; hitung status menggunakan perbandingan `aktual` dan `deadline` bila respons API belum menyertakan status.

### Prinsip desain

- Bersih, modern, profesional; desktop-first tetapi responsive untuk mobile.
- Fokus readability dan data: judul, angka, satuan persentase, axis, empty/loading/error state.
- React functional components dan hooks; gunakan **Recharts** bila cocok dengan dependensi existing.
- Pengambilan data via HTTP API, **bukan SQL dari browser**, **bukan data dummy/hardcoded**.
- Gunakan `import.meta.env.VITE_API_BASE_URL` untuk base URL API.
- Tangani kegagalan fetch dengan pesan yang jelas dan tidak menampilkan detail internal server.
- Hindari fitur tambahan (login, CRUD, design system kompleks) yang memperlambat tes.
- Bila memungkinkan, sediakan tombol `Refresh` yang benar-benar memuat data ulang.

## 7. Backend Go Fiber — Hanya Integrasi Minimal

**Backend sudah dipilih: Go + Fiber, bukan PHP Native.** Pemeriksaan awal menentukan apa yang telah tersedia. Jika endpoint sudah ada, **gunakan dan pertahankan**. Jika belum ada, tambahkan endpoint read-only berikut tanpa membuat ulang seluruh project:

```http
GET /api/v1/health
GET /api/v1/kpi/summary
GET /api/v1/kpi/ontime-late
GET /api/v1/kpi/tasks
```

Respons summary idealnya berisi satu baris per karyawan dengan field:

```json
{
  "data": [
    {
      "nama": "Budi",
      "sales_target": 2,
      "sales_actual": 2,
      "sales_pencapaian": 100,
      "sales_bobot": 50,
      "sales_late": 0,
      "sales_total_bobot": 50,
      "report_target": 2,
      "report_actual": 2,
      "report_pencapaian": 100,
      "report_bobot": 50,
      "report_late": -5,
      "report_total_bobot": 45,
      "total_kpi": 95
    }
  ]
}
```

Ini **contoh kontrak API**, bukan klaim bahwa API sekarang sudah menghasilkan bentuk ini. Jika schema API existing berbeda, buat mapping frontend atau sesuaikan kontrak secara minimal dan dokumentasikan.

Persyaratan backend:

- Menggunakan `database/sql` dan MySQL driver; query terparameterisasi.
- Perhitungan berdasarkan data DB dan nilai target/bobot di `mst_kpi`.
- Gunakan `.env` untuk lokal dan environment Docker untuk production.
- Set host MySQL saat container berjalan menjadi `mysql`, port `3306`.
- Terapkan CORS origin yang eksplisit: lokal Vite dan domain FE production.
- Hindari memperlihatkan error SQL/credential ke client.
- Validasi `/health` dan endpoint KPI sebelum integrasi FE.
- Gunakan **Fiber v3** hanya jika kompatibel; lihat `go.mod` dan dokumentasi resmi. API v3 memiliki signature handler `func(c fiber.Ctx) error`.

## 8. Docker — Implementasikan Hanya SATU Compose Terpusat

Agent harus membuat file-file berikut di **`TRUSMI_GROUP/docker/`**, bukan di dalam source FE/BE. **Contoh konfigurasi ini adalah baseline yang harus disesuaikan dengan source code aktual** (entrypoint Go, environment names, port backend, dan router).

### 8.1 `docker/docker-compose.yml`

```yaml
name: trusmi-group

services:
  mysql:
    image: mysql:8.4
    restart: unless-stopped
    environment:
      MYSQL_DATABASE: ${MYSQL_DATABASE:?MYSQL_DATABASE required}
      MYSQL_USER: ${MYSQL_USER:?MYSQL_USER required}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD:?MYSQL_PASSWORD required}
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD:?MYSQL_ROOT_PASSWORD required}
    volumes:
      - mysql_data:/var/lib/mysql
    networks:
      - trusmi_internal
    healthcheck:
      test: ["CMD-SHELL", "MYSQL_PWD=\"$$MYSQL_ROOT_PASSWORD\" mysqladmin ping -h 127.0.0.1 -u root --silent"]
      interval: 10s
      timeout: 5s
      retries: 12
      start_period: 30s

  backend:
    build:
      context: ..
      dockerfile: docker/backend.Dockerfile
    restart: unless-stopped
    environment:
      APP_PORT: "8080"
      DB_HOST: mysql
      DB_PORT: "3306"
      DB_NAME: ${MYSQL_DATABASE:?MYSQL_DATABASE required}
      DB_USER: ${MYSQL_USER:?MYSQL_USER required}
      DB_PASSWORD: ${MYSQL_PASSWORD:?MYSQL_PASSWORD required}
      CORS_ALLOWED_ORIGINS: ${CORS_ALLOWED_ORIGINS:?CORS_ALLOWED_ORIGINS required}
    ports:
      - "127.0.0.1:${TRUSMI_API_HOST_PORT:-18082}:8080"
    depends_on:
      mysql:
        condition: service_healthy
    networks:
      - trusmi_internal

  frontend:
    build:
      context: ..
      dockerfile: docker/frontend.Dockerfile
      args:
        VITE_API_BASE_URL: ${VITE_API_BASE_URL:?VITE_API_BASE_URL required}
    restart: unless-stopped
    ports:
      - "127.0.0.1:${TRUSMI_FE_HOST_PORT:-18081}:80"
    depends_on:
      - backend
    networks:
      - trusmi_internal

networks:
  trusmi_internal:
    driver: bridge

volumes:
  mysql_data:
```

**Mengapa `context: ..`?** Agar Dockerfile yang disimpan di folder `docker/` dapat `COPY trusmi_fe/...` dan `COPY trusmi_be/...` dari konteks project root. Perhatikan path relatif terhadap Compose file. Jangan menulis `COPY ../...` di Dockerfile jika sumber berada di luar konteks build.

**Security:** hanya port FE/BE yang di-bind ke `127.0.0.1`; MySQL tidak memiliki mapping `ports:`. Port 80/443 publik tetap ditangani oleh Nginx host. Jangan otomatis mengganti container lain yang sudah memakai port publik.

### 8.2 `docker/backend.Dockerfile` — template

```dockerfile
# Cocokkan versi image builder dengan go.mod & Fiber.
# Fiber v3 pada dokumentasi terkini memerlukan Go 1.26+.
FROM golang:1.26-alpine AS builder
WORKDIR /src

COPY trusmi_be/go.mod trusmi_be/go.sum ./
RUN go mod download
COPY trusmi_be/ ./

# Periksa path entrypoint yang benar di project existing.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /out/trusmi-api ./cmd/api

FROM alpine:3.22
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app && adduser -S app -G app
COPY --from=builder /out/trusmi-api /usr/local/bin/trusmi-api
USER app
EXPOSE 8080
CMD ["/usr/local/bin/trusmi-api"]
```

> Jika entrypoint bukan `./cmd/api`, Go memiliki dependency CGO, atau `go.sum` belum dibuat, agent wajib menyesuaikan Dockerfile terhadap kondisi project. Jangan mengubah modul/versi Fiber secara serampangan.

### 8.3 `docker/frontend.Dockerfile` — template

```dockerfile
FROM node:22-alpine AS builder
WORKDIR /app

COPY trusmi_fe/package*.json ./
RUN npm ci
COPY trusmi_fe/ ./

ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=${VITE_API_BASE_URL}
RUN npm run build

FROM nginx:alpine
COPY docker/frontend.nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

Vite menyisipkan nilai `VITE_API_BASE_URL` **saat build**, bukan saat container frontend baru dijalankan. Setelah mengubah URL API, lakukan rebuild FE. `VITE_` adalah variabel publik, jangan pernah menaruh secret di dalamnya.

### 8.4 `docker/frontend.nginx.conf` — static SPA di container

```nginx
server {
    listen 80;
    server_name _;
    root /usr/share/nginx/html;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /assets/ {
        try_files $uri =404;
        expires 7d;
    }
}
```

Nginx **dalam container FE** hanya menyajikan hasil `dist/`; Nginx **host VPS** menangani domain dan TLS. Keduanya punya fungsi berbeda.

### 8.5 `docker/.env.prod.example` — tanpa secret asli

```dotenv
MYSQL_DATABASE=db_test_programmer_trusmi
MYSQL_USER=trusmi_app
MYSQL_PASSWORD=GANTI_DENGAN_PASSWORD_KUAT
MYSQL_ROOT_PASSWORD=GANTI_DENGAN_PASSWORD_ROOT_KUAT

TRUSMI_FE_HOST_PORT=18081
TRUSMI_API_HOST_PORT=18082

VITE_API_BASE_URL=https://api-trusmi.riskiahmadilham.my.id
CORS_ALLOWED_ORIGINS=https://trusmi.riskiahmadilham.my.id
```

Agent harus menambahkan **`docker/.env.prod` ke `.gitignore`** dan tidak memasukkan password nyata ke output/log/repository. Template ini hanya contoh; gunakan kata sandi kuat sendiri di production. Pastikan nama environment variables disesuaikan jika backend existing memakai penamaan berbeda.

### 8.6 `.dockerignore` di root

```gitignore
**/.git
**/node_modules
**/dist
**/coverage
**/.env
**/.env.*
!**/.env.example
!**/.env.prod.example
**/*.log
**/.DS_Store
```

Catatan: pastikan konfigurasi `.dockerignore` tidak mengecualikan file yang betul-betul diperlukan dalam build. `docs/` tidak perlu dikirim ke build jika tidak digunakan.

## 9. Lokal dan Production — Command yang Disarankan

Dari root `TRUSMI_GROUP/`, setelah `.env.prod` dibuat dengan benar:

```bash
# Validasi konfigurasi tanpa menjalankan container
# Hati-hati: output config berpotensi memuat ENV; jangan bagikan log-nya mentah.
docker compose --env-file docker/.env.prod -f docker/docker-compose.yml config --quiet

# Build dan jalankan seluruh service
docker compose --env-file docker/.env.prod -f docker/docker-compose.yml up -d --build

# Cek status
docker compose --env-file docker/.env.prod -f docker/docker-compose.yml ps

# Periksa log tanpa membuka file secret
docker compose --env-file docker/.env.prod -f docker/docker-compose.yml logs --tail=80 backend

# Nonaktifkan service tanpa menghapus volume database
docker compose --env-file docker/.env.prod -f docker/docker-compose.yml down
```

**Jangan jalankan `docker compose down -v` untuk database yang sudah berisi data.** `down -v` dapat menghapus volume MySQL. Untuk pembaruan FE/BE gunakan build/recreate sesuai kebutuhan tanpa menghapus volume.

**Perhatian MySQL:** volume baru masih kosong walaupun data di MySQL lokal sudah lengkap. Database harus di-*backup* dari lokal dan di-*import* ke container melalui prosedur terverifikasi. Jangan menganggap Compose otomatis memigrasikan database lokal ke VPS. Jangan import ulang dengan metode destruktif tanpa backup dan persetujuan.

## 10. DNS, Nginx Host, dan HTTPS

Di DNS Rumahweb (jika DNS domain dikelola di sana), buat A record:

| Type | Host | Target |
|---|---|---|
| A | `trusmi` | **IP publik VPS yang benar** |
| A | `api-trusmi` | **IP publik VPS yang benar** |

**Jangan hardcode atau mengarang IP VPS.** Periksa IP yang saat ini berlaku, propagasi DNS, dan konflik record sebelumnya.

Konfigurasi **Nginx host VPS** (contoh lokasi `/etc/nginx/sites-available/`; jangan overwrite virtual host existing):

```nginx
server {
    listen 80;
    server_name trusmi.riskiahmadilham.my.id;

    location / {
        proxy_pass http://127.0.0.1:18081;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}

server {
    listen 80;
    server_name api-trusmi.riskiahmadilham.my.id;

    location / {
        proxy_pass http://127.0.0.1:18082;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Setelah memeriksa vhost existing, DNS, dan respons HTTP:

1. Uji `sudo nginx -t` sebelum reload.
2. Pastikan Nginx yang aktif memang berada pada host, bukan container reverse proxy lain.
3. Aktifkan sertifikat HTTPS dengan Certbot/Let's Encrypt memakai cara yang sesuai instalasi Nginx server.
4. Uji URL final FE dan API lewat HTTPS.
5. Pastikan CORS hanya mengizinkan origin FE production (ditambah lokal untuk development).
6. Jangan membuka port 3306, 18081, atau 18082 ke publik via firewall; keduanya yang terakhir sudah bind localhost.

**Do not deploy automatically:** minta izin pemilik project sebelum membuat/mengubah DNS, menulis vhost Nginx, menjalankan Certbot, membuka firewall, mengubah container existing, atau mengimpor database production.

## 11. Pengujian & Kriteria Selesai

### Database dan API

- [ ] Data empat karyawan dan 13 task bisa dibaca dari database yang digunakan API.
- [ ] Query menghasilkan KPI: Budi 95, Adi 93, Rara 63, Doni 43.
- [ ] On-time = 8/13 (61,54%) dan late = 5/13 (38,46%).
- [ ] API `/health` merespons sukses.
- [ ] Endpoint KPI mengembalikan JSON yang dapat dipakai frontend.
- [ ] Konfigurasi CORS lokal/production benar.

### Dashboard

- [ ] Grafik Sales dan Report per karyawan tampil dari API.
- [ ] Total KPI per karyawan tampil dari API.
- [ ] Grafik jumlah/persentase on-time vs late tampil dari API.
- [ ] Tabel output mengikuti kolom dan label pada soal 1.
- [ ] UI responsive, ada loading/empty/error state.
- [ ] Tidak ada nilai KPI yang di-hardcode di frontend.
- [ ] `npm run build` sukses.

### Docker & deployment

- [ ] Hanya ada satu Compose utama di root `docker/`.
- [ ] FE, BE, MySQL berjalan dan saling terhubung.
- [ ] Docker hanya bind FE/BE pada localhost host VPS.
- [ ] MySQL memakai volume persisten, tidak expose port publik.
- [ ] FE dapat diakses lewat subdomain, API lewat subdomain lain, keduanya HTTPS.
- [ ] Aplikasi server lain tidak terganggu.

### Pengumpulan technical test

- [ ] `docs/soal_1_kpi_query.txt` berisi query SQL yang benar.
- [ ] `docs/soal_2_ontime_late_query.txt` berisi query SQL yang benar.
- [ ] `docs/soal_5_normalisasi.txt` menjelaskan sebelum/sesudah normalisasi dan relasi tabel.
- [ ] Screenshot **asli** Soal 3 dan Soal 4 disimpan di `docs/screenshots/`.
- [ ] Source FE/BE dan cara menjalankannya terdokumentasi.
- [ ] Prompt AI yang digunakan dilampirkan sesuai ketentuan tes.

## 12. Instruksi Eksekusi untuk Agent — Prioritas dan Guardrail

**Urutan kerja:**

1. **Audit** FE, BE, docs, MySQL, dan port server (read-only terlebih dahulu).
2. **Bangun dashboard React** Soal 3 dan 4. Ini fokus utama; reuse komponen yang ada.
3. **Integrasikan API** Go Fiber; buat atau koreksi endpoint minimal hanya jika diperlukan.
4. **Verifikasi angka KPI** dan UI terhadap dataset tes.
5. **Buat folder `docker/` terpusat** dan sesuaikan Compose, Dockerfile, Nginx, `.dockerignore`.
6. **Jalankan uji lokal** bila environment tersedia.
7. **Siapkan** dokumentasi deployment, screenshot, file jawaban `.txt`, dan catatan prompt AI.
8. **Deployment VPS production hanya setelah ada izin eksplisit.**

**Aturan keselamatan perubahan:**

- Jangan hapus/reset database atau file proyek.
- Jangan mengganti Go Fiber menjadi PHP/Laravel atau FE menjadi Next.js.
- Jangan membuat Docker Compose lain di dalam FE atau BE.
- Jangan membuat repository baru atau `git push` tanpa izin.
- Jangan mengubah aplikasi lain pada VPS atau mengambil alih port 80/443.
- Jangan mencetak password, `.env`, access token, atau private key.
- Jangan mengklaim tes atau screenshot berhasil jika belum benar-benar dijalankan/diambil.
- Jangan memaksakan migrasi ulang tabel asli; utamakan implementasi UI sekarang.

**Output akhir dari agent:** sampaikan ringkas daftar file baru/diubah, hasil tes yang benar-benar berhasil, hasil dashboard, cara menjalankan satu Docker Compose, dan hal yang masih butuh keputusan/izin pemilik proyek.

---

## Rujukan teknis

- Soal test: <https://docs.google.com/spreadsheets/d/1qt1oQiZmHN58f4ArzpWWzqSEcMfCUaatbTy8CIoaIzw/edit?gid=0#gid=0>
- Fiber: <https://docs.gofiber.io/>
- Docker Compose build: <https://docs.docker.com/reference/compose-file/build/>
- Docker Compose CLI: <https://docs.docker.com/reference/cli/docker/compose/>
- Vite environment variables: <https://vite.dev/guide/env-and-mode>
- Certbot: <https://certbot.eff.org/instructions>
