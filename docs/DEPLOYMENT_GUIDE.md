# Panduan Docker Development (`docker/dev`) & Production (`docker/prod`) — TRUSMI GROUP

Struktur folder `docker/` mengikuti pola `jastip-core` (`docker/dev` dan `docker/prod`), membaca kredensial langsung dari `trusmi_be/.env` dan `trusmi_fe/.env`, serta menggunakan port yang **tidak bentrok** dengan `jastip-core` di server (`srv2005276`).

---

## 1. Perbandingan Port Server (`jastip-core` vs `trusmi-core`)

| Layanan | `jastip-core` (Sudah Ada di Server) | `trusmi-core` (Baru — Tidak Bentrok) |
|---|---|---|
| **Frontend Prod** | `127.0.0.1:8081->80/tcp` | `127.0.0.1:18081->80/tcp` |
| **Backend Prod** | `8080/tcp` | `127.0.0.1:18082->8090/tcp` (`APP_PORT=8090`) |
| **Database Prod** | `127.0.0.1:15432->5432/tcp` (Postgres) | `127.0.0.1:13306->3306/tcp` (MySQL) |
| **Backend Dev (Lokal)** | `8080` | `8090:8090` (`http://localhost:8090`) |
| **Frontend Dev (Lokal)** | `5173` | `5173:5173` (`http://localhost:5173`) |

---

## 2. Struktur Folder `docker/`

```text
TRUSMI_GROUP/
├── docker/
│   ├── dev/
│   │   ├── backend.Dockerfile
│   │   ├── docker-compose.dev.yml
│   │   └── frontend.Dockerfile
│   └── prod/
│       ├── backend.Dockerfile
│       ├── docker-compose.prod.yml
│       ├── frontend.Dockerfile
│       └── nginx.conf
├── trusmi_be/
│   └── .env
└── trusmi_fe/
    └── .env
```

---

## 3. Mode Development Lokal (`docker/dev`)

```bash
cd docker/dev
docker compose -f docker-compose.dev.yml up -d --build
```

---

## 4. Mode Production VPS (`docker/prod`)

```bash
cd docker/prod
docker compose -f docker-compose.prod.yml up -d --build
```
