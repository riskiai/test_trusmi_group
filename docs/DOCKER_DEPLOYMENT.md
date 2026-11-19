# Dokumentasi Docker Development & Production — TRUSMI GROUP

Dokumen ini menjelaskan arsitektur, konfigurasi, pengujian lokal, dan rencana deployment **Docker Development** serta **Docker Production** pada proyek **TRUSMI GROUP**.

---

## 1. Struktur Docker Development dan Production

```text
TRUSMI_GROUP/
├── docker/
│   ├── dev/
│   │   ├── backend.Dockerfile          # Dockerfile Backend Development (Go 1.25 Alpine)
│   │   ├── frontend.Dockerfile         # Dockerfile Frontend Development (Node 22 Alpine + Vite HMR)
│   │   └── docker-compose.dev.yml      # Orchestrator Development Lokal
│   └── prod/
│       ├── backend.Dockerfile          # Multi-stage Dockerfile Backend Production (Go -> Alpine 3.22 non-root)
│       ├── frontend.Dockerfile         # Multi-stage Dockerfile Frontend Production (Node build -> Nginx Alpine)
│       ├── docker-compose.prod.yml     # Orchestrator Production VPS (Terhubung ke MySQL Global)
│       └── nginx.conf                  # Konfigurasi Nginx internal container Frontend (SPA Fallback)
├── docs/
│   └── DOCKER_DEPLOYMENT.md
├── trusmi_be/
│   ├── .env                            # Environment Development Lokal (Git-ignored)
│   ├── .env.example                    # Template Environment Development
│   ├── .env.prod                       # Environment Production VPS (Git-ignored)
│   └── .env.prod.example               # Template Environment Production (Aman untuk Repository)
├── trusmi_fe/
│   ├── .env                            # Environment Frontend Development (Git-ignored)
│   ├── .env.example                    # Template Environment Frontend Development
│   ├── .env.prod                       # Environment Frontend Production (Git-ignored)
│   └── .env.prod.example               # Template Environment Frontend Production
├── .dockerignore
└── .gitignore
```

---

## 2. Cara Mengatur Environment

Konfigurasi **Development** dan **Production** dipisahkan secara tegas agar tidak ada kredensial development (`root` tanpa password) atau URL `localhost` yang terbawa ke production.

### 2.1 Backend Environment

- **Development (`trusmi_be/.env` — contoh di `trusmi_be/.env.example`)**:
  ```dotenv
  APP_ENV=development
  APP_HOST=0.0.0.0
  APP_PORT=8090

  DB_HOST=127.0.0.1
  DB_PORT=3306
  DB_NAME=db_test_programmer_trusmi
  DB_USER=root
  DB_PASSWORD=

  CORS_ALLOWED_ORIGINS=http://localhost:5173,http://127.0.0.1:5173,http://localhost:5174,http://127.0.0.1:5174,http://localhost:18081,https://trusmi.riskiahmadiilham.my.id
  ```

- **Production (`trusmi_be/.env.prod` — contoh di `trusmi_be/.env.prod.example`)**:
  ```dotenv
  APP_ENV=production
  APP_HOST=0.0.0.0
  APP_PORT=8090

  DB_HOST=mysql-global
  DB_PORT=3306
  DB_NAME=db_test_programmer_trusmi
  DB_USER=trusmi_user
  DB_PASSWORD=<DIISI_MANUAL>

  CORS_ALLOWED_ORIGINS=https://trusmi.riskiahmadiilham.my.id
  ```

### 2.2 Frontend Environment

- **Development (`trusmi_fe/.env` — contoh di `trusmi_fe/.env.example`)**:
  ```dotenv
  VITE_API_BASE_URL=http://localhost:8090
  ```
- **Production (`trusmi_fe/.env.prod` — contoh di `trusmi_fe/.env.prod.example`)**:
  ```dotenv
  VITE_API_BASE_URL=/api/v1
  ```
  Pada mode production, Frontend menggunakan **relative API URL** (`/api/v1`) pada domain yang sama (`https://trusmi.riskiahmadiilham.my.id`). Fungsi `buildApiUrl()` di `trusmi_fe/src/services/api.js` otomatis menormalisasi prefix sehingga tidak terjadi duplikasi `/api/v1/api/v1`.

---

## 3. Cara Menjalankan Docker Development (Lokal)

Pada mode Development:
- Container `trusmi-dev-backend` berjalan pada port `8090` dan mengakses MySQL lokal Windows melalui `host.docker.internal:3306`.
- Container `trusmi-dev-frontend` berjalan pada port `5173` dengan bind mount `src/`, `public/`, `index.html`, dan `vite.config.js` sehingga **Vite Hot Module Replacement (HMR)** tetap berfungsi.
- Tidak membuat container MySQL tambahan di lokal.

### Perintah dari Root Project (`TRUSMI_GROUP/`)

```bash
# 1. Validasi konfigurasi Docker Compose Development
docker compose -f docker/dev/docker-compose.dev.yml config --quiet

# 2. Build dan jalankan Docker Development
docker compose -f docker/dev/docker-compose.dev.yml up -d --build

# 3. Periksa status container
docker compose -f docker/dev/docker-compose.dev.yml ps

# 4. Periksa log Backend & Frontend
docker compose -f docker/dev/docker-compose.dev.yml logs --tail=50 backend
docker compose -f docker/dev/docker-compose.dev.yml logs --tail=50 frontend

# 5. Menghentikan Docker Development
docker compose -f docker/dev/docker-compose.dev.yml down
```

*(Catatan: Jika port `5173` di host Windows sedang dipakai oleh proses `npm run dev` lokal, Anda dapat menjalankan container frontend di port alternatif, misalnya `5174`, dengan variabel `FE_DEV_PORT=5174`).*

---

## 4. Koneksi Backend ke MySQL Global & External Docker Network

Pada VPS Ubuntu (`Ubuntu 24.04`), database MySQL sudah disediakan oleh container **MySQL Global** yang berjalan di network Docker eksternal `mysql_shared`.

Oleh karena itu, pada `docker/prod/docker-compose.prod.yml`:
1. **Tidak ada service MySQL baru** dan tidak ada volume MySQL baru.
2. Service `backend` (`trusmi-prod-backend`) terhubung ke **dua network**:
   - `trusmi_prod_net` (`driver: bridge`): Jaringan internal antara `trusmi-prod-frontend` dan `trusmi-prod-backend`.
   - `mysql_shared` (`external: true`): Jaringan eksternal tempat container `mysql-global` berada.
3. Dengan tergabung ke `mysql_shared`, container `trusmi-prod-backend` dapat me-resolve hostname `mysql-global` pada port internal `3306` menggunakan akun non-root `trusmi_user` dan database `db_test_programmer_trusmi`.

---

## 5. Penjelasan Port Mapping

| Environment | Container | Binding Host | Port Internal Container | Keterangan |
|---|---|---|---|---|
| **Development** | `trusmi-dev-frontend` | `0.0.0.0:5173` (atau `FE_DEV_PORT`) | `5173` | Vite Dev Server dengan HMR |
| **Development** | `trusmi-dev-backend` | `0.0.0.0:8090` | `8090` | Golang Fiber REST API |
| **Production** | `trusmi-prod-frontend` | `127.0.0.1:18081` | `80` | Nginx static SPA container (hanya diakses Nginx Host VPS) |
| **Production** | `trusmi-prod-backend` | `127.0.0.1:18082` | `8090` | Golang Fiber binary non-root (hanya diakses Nginx Host VPS) |

> **Keamanan Port Production:** Binding `127.0.0.1:18081` dan `127.0.0.1:18082` memastikan container tidak terekspos langsung ke internet publik (bypass UFW) dan tidak bentrok dengan aplikasi `jastip-core` (`8080`, `8081`, `15432`).

---

## 6. Cara Validasi dan Build Docker Production

```bash
# 1. Validasi sintaks dan konfigurasi Docker Compose Production
docker compose -f docker/prod/docker-compose.prod.yml config --quiet

# 2. Build image production (Backend & Frontend) tanpa menjalankan container
docker compose -f docker/prod/docker-compose.prod.yml build
```

---

## 7. Contoh Konfigurasi Nginx Host VPS (Satu Subdomain)

Pada VPS Ubuntu, **satu domain** `https://trusmi.riskiahmadiilham.my.id` digunakan untuk melayani Frontend (`/`) dan Backend (`/api/`):
- Request ke `/` diteruskan ke `http://127.0.0.1:18081` (`trusmi-prod-frontend`).
- Request ke `/api/` diteruskan ke `http://127.0.0.1:18082` (`trusmi-prod-backend`).

Contoh file `/etc/nginx/sites-available/trusmi.riskiahmadiilham.my.id`:

```nginx
server {
    listen 80;
    server_name trusmi.riskiahmadiilham.my.id;

    # Backend Golang Fiber REST API (/api/v1/...)
    location /api/ {
        proxy_pass http://127.0.0.1:18082;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Frontend React Vite (Nginx Container)
    location / {
        proxy_pass http://127.0.0.1:18081;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

---

## 8. Rencana Deployment VPS Secara Manual

1. **Persiapan Folder di VPS (`/opt/apps/trusmi-group`)**:
   ```bash
   sudo mkdir -p /opt/apps/trusmi-group
   sudo chown -R $USER:$USER /opt/apps/trusmi-group
   cd /opt/apps/trusmi-group
   ```
2. **Clone / Salin Repository & Folder Docker**:
   - Clone repository `trusmi_be` dan `trusmi_fe` ke dalam `/opt/apps/trusmi-group/`.
   - Salin folder `docker/` dan `.dockerignore` ke `/opt/apps/trusmi-group/`.
3. **Pastikan External Network `mysql_shared` Tersedia**:
   ```bash
   docker network ls | grep mysql_shared
   ```
4. **Siapkan File `trusmi_be/.env.prod`**:
   ```bash
   cp trusmi_be/.env.prod.example trusmi_be/.env.prod
   nano trusmi_be/.env.prod
   # Isi DB_PASSWORD dengan password user MySQL trusmi_user yang sebenarnya
   ```
5. **Build dan Jalankan Container Production**:
   ```bash
   docker compose -f docker/prod/docker-compose.prod.yml up -d --build
   ```
6. **Aktifkan Virtual Host Nginx & SSL Certbot**:
   ```bash
   sudo ln -s /etc/nginx/sites-available/trusmi.riskiahmadiilham.my.id /etc/nginx/sites-enabled/
   sudo nginx -t
   sudo systemctl reload nginx
   sudo certbot --nginx -d trusmi.riskiahmadiilham.my.id
   ```

---

## 9. Cara Verifikasi Setelah Deployment

```bash
# 1. Pastikan container trusmi-prod-backend dan trusmi-prod-frontend berstatus Up
docker compose -f docker/prod/docker-compose.prod.yml ps

# 2. Uji endpoint health backend melalui localhost VPS
curl -i http://127.0.0.1:18082/api/v1/health

# 3. Uji endpoint KPI summary melalui localhost VPS
curl -i http://127.0.0.1:18082/api/v1/kpi/summary

# 4. Uji frontend container melalui localhost VPS
curl -I http://127.0.0.1:18081/

# 5. Uji melalui domain HTTPS publik
curl -i https://trusmi.riskiahmadiilham.my.id/api/v1/health
```

---

## 10. Troubleshooting Umum

1. **`network mysql_shared declared as external, but could not be found`**:
   - Di mesin lokal (atau jika network belum ada), buat network dengan `docker network create mysql_shared` atau periksa nama network MySQL Global di VPS dengan `docker network ls`.
2. **Backend gagal konek ke MySQL lokal pada Docker Development (`connect: connection refused`)**:
   - Pastikan service MySQL Windows aktif di port `3306` dan mengizinkan koneksi TCP dari `host.docker.internal` (atau `127.0.0.1`), serta `DB_NAME=db_test_programmer_trusmi` sudah tersedia.
3. **Port `5173` sudah terpakai saat menjalankan Docker Development**:
   - Hentikan `npm run dev` yang berjalan di terminal host, atau jalankan Docker Development dengan port alternatif:
     - PowerShell: `$env:FE_DEV_PORT="5174"; docker compose -f docker/dev/docker-compose.dev.yml up -d`
     - Bash: `FE_DEV_PORT=5174 docker compose -f docker/dev/docker-compose.dev.yml up -d`
4. **Perubahan `VITE_API_BASE_URL` tidak berdampak di Production**:
   - Vite menyisipkan variabel `VITE_*` pada saat proses `npm run build` di image builder. Jalankan ulang `docker compose -f docker/prod/docker-compose.prod.yml up -d --build frontend` setelah mengubah `VITE_API_BASE_URL`.
