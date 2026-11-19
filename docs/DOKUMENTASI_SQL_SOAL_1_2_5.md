# Dokumentasi Pengujian Query SQL — Technical Test Programmer TRUSMI GROUP

Dokumen ini memuat dokumentasi query SQL beserta hasil pengujian di MySQL (`db_test_programmer_trusmi`) untuk **Soal Nomor 1, 2, dan 5**.

---

## Soal 1 — Query Perhitungan KPI Marketing per Karyawan

### Query SQL (MySQL)

```sql
WITH rekap AS (
    SELECT
        k.id AS karyawan_id,
        k.nama_karyawan,
        kp.nama_kpi,
        kp.target_kpi,
        COUNT(t.aktual) AS actual,

        ROUND(
            COUNT(t.aktual) * 100.0 /
            NULLIF(kp.target_kpi, 0), 2
        ) AS pencapaian,

        ROUND(
            COUNT(t.aktual) * kp.bobot_kpi /
            NULLIF(kp.target_kpi, 0), 2
        ) AS bobot_actual,

        SUM(
            CASE WHEN t.aktual > t.deadline
                THEN 1 ELSE 0 END
        ) * kp.bobot_late AS late

    FROM mst_karyawan k
    CROSS JOIN mst_kpi kp
    LEFT JOIN trx_kpi_marketing t
        ON t.karyawan_id = k.id
        AND t.kpi_id = kp.id

    GROUP BY
        k.id, k.nama_karyawan,
        kp.id, kp.nama_kpi,
        kp.target_kpi, kp.bobot_kpi, kp.bobot_late
)

SELECT
    nama_karyawan AS Nama,

    MAX(CASE WHEN nama_kpi = 'Sales'
        THEN target_kpi END) AS sales_target,
    MAX(CASE WHEN nama_kpi = 'Sales'
        THEN actual END) AS sales_actual,
    MAX(CASE WHEN nama_kpi = 'Sales'
        THEN pencapaian END) AS sales_pencapaian,
    MAX(CASE WHEN nama_kpi = 'Sales'
        THEN bobot_actual END) AS bobot_sales,
    MAX(CASE WHEN nama_kpi = 'Sales'
        THEN late END) AS late_sales,
    MAX(CASE WHEN nama_kpi = 'Sales'
        THEN bobot_actual + late END) AS total_bobot_sales,

    MAX(CASE WHEN nama_kpi = 'Report'
        THEN target_kpi END) AS report_target,
    MAX(CASE WHEN nama_kpi = 'Report'
        THEN actual END) AS report_actual,
    MAX(CASE WHEN nama_kpi = 'Report'
        THEN pencapaian END) AS report_pencapaian,
    MAX(CASE WHEN nama_kpi = 'Report'
        THEN bobot_actual END) AS bobot_report,
    MAX(CASE WHEN nama_kpi = 'Report'
        THEN late END) AS late_report,
    MAX(CASE WHEN nama_kpi = 'Report'
        THEN bobot_actual + late END) AS total_bobot_report,

    ROUND(SUM(bobot_actual + late), 2) AS KPI

FROM rekap
GROUP BY karyawan_id, nama_karyawan
ORDER BY karyawan_id;
```

### Hasil Output MySQL

```text
+------+--------------+--------------+------------------+-------------+------------+-------------------+---------------+---------------+-------------------+--------------+-------------+--------------------+-------+
| Nama | sales_target | sales_actual | sales_pencapaian | bobot_sales | late_sales | total_bobot_sales | report_target | report_actual | report_pencapaian | bobot_report | late_report | total_bobot_report | KPI   |
+------+--------------+--------------+------------------+-------------+------------+-------------------+---------------+---------------+-------------------+--------------+-------------+--------------------+-------+
| Budi |            2 |            2 |           100.00 |       50.00 |       0.00 |             50.00 |             2 |             2 |            100.00 |        50.00 |       -5.00 |              45.00 | 95.00 |
| Adi  |            2 |            2 |           100.00 |       50.00 |      -7.00 |             43.00 |             2 |             2 |            100.00 |        50.00 |        0.00 |              50.00 | 93.00 |
| Rara |            2 |            2 |           100.00 |       50.00 |      -7.00 |             43.00 |             2 |             1 |             50.00 |        25.00 |       -5.00 |              20.00 | 63.00 |
| Doni |            2 |            1 |            50.00 |       25.00 |      -7.00 |             18.00 |             2 |             1 |             50.00 |        25.00 |        0.00 |              25.00 | 43.00 |
+------+--------------+--------------+------------------+-------------+------------+-------------------+---------------+---------------+-------------------+--------------+-------------+--------------------+-------+
4 rows in set (0.01 sec)
```

---

## Soal 2 — Query Perhitungan Jumlah Tasklist Ontime & Late Beserta Persentasenya

### Query SQL (MySQL)

```sql
SELECT
    COUNT(*) AS total_task,

    SUM(
        CASE WHEN aktual <= deadline
            THEN 1 ELSE 0 END
    ) AS total_ontime,

    SUM(
        CASE WHEN aktual > deadline
            THEN 1 ELSE 0 END
    ) AS total_late,

    ROUND(
        SUM(CASE WHEN aktual <= deadline
            THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(*), 0), 2
    ) AS persentase_ontime,

    ROUND(
        SUM(CASE WHEN aktual > deadline
            THEN 1 ELSE 0 END) * 100.0 /
        NULLIF(COUNT(*), 0), 2
    ) AS persentase_late

FROM trx_kpi_marketing
WHERE aktual IS NOT NULL
  AND deadline IS NOT NULL;
```

### Hasil Output MySQL

```text
+------------+--------------+------------+-------------------+-----------------+
| total_task | total_ontime | total_late | persentase_ontime | persentase_late |
+------------+--------------+------------+-------------------+-----------------+
|         13 |            8 |          5 |             61.54 |           38.46 |
+------------+--------------+------------+-------------------+-----------------+
1 row in set (0.00 sec)
```

---

## Soal 5 — Normalisasi Struktur `table_kpi_marketing`

### 1. Daftar Tabel di Database

```sql
SHOW TABLES;
```

```text
+-------------------------------------+
| Tables_in_db_test_programmer_trusmi |
+-------------------------------------+
| mst_karyawan                        |
| mst_kpi                             |
| table_kpi_marketing                 |
| trx_kpi_marketing                   |
+-------------------------------------+
4 rows in set (0.00 sec)
```

### 2. Deskripsi Struktur Tabel (`DESCRIBE`)

#### `table_kpi_marketing` (Tabel Original Sebelum Normalisasi)
```text
+----------+--------------+------+-----+---------+-------+
| Field    | Type         | Null | Key | Default | Extra |
+----------+--------------+------+-----+---------+-------+
| id       | int          | NO   | PRI | NULL    |       |
| tasklist | varchar(100) | YES  |     | NULL    |       |
| kpi      | varchar(100) | YES  |     | NULL    |       |
| karyawan | varchar(100) | YES  |     | NULL    |       |
| deadline | date         | YES  |     | NULL    |       |
| aktual   | date         | YES  |     | NULL    |       |
+----------+--------------+------+-----+---------+-------+
```

#### `mst_karyawan` (Tabel Master Karyawan)
```text
+---------------+--------------+------+-----+---------+----------------+
| Field         | Type         | Null | Key | Default | Extra          |
+---------------+--------------+------+-----+---------+----------------+
| id            | int          | NO   | PRI | NULL    | auto_increment |
| nama_karyawan | varchar(100) | NO   |     | NULL    |                |
+---------------+--------------+------+-----+---------+----------------+
```

#### `mst_kpi` (Tabel Master KPI)
```text
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int          | NO   | PRI | NULL    | auto_increment |
| nama_kpi   | varchar(100) | NO   |     | NULL    |                |
| target_kpi | int          | NO   |     | NULL    |                |
| bobot_kpi  | decimal(5,2) | NO   |     | NULL    |                |
| bobot_late | decimal(5,2) | NO   |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
```

#### `trx_kpi_marketing` (Tabel Transaksi KPI Marketing)
```text
+-------------+--------------+------+-----+---------+----------------+
| Field       | Type         | Null | Key | Default | Extra          |
+-------------+--------------+------+-----+---------+----------------+
| id          | int          | NO   | PRI | NULL    | auto_increment |
| tasklist    | varchar(100) | YES  |     | NULL    |                |
| karyawan_id | int          | NO   | MUL | NULL    |                |
| kpi_id      | int          | NO   | MUL | NULL    |                |
| deadline    | date         | YES  |     | NULL    |                |
| aktual      | date         | YES  |     | NULL    |                |
+-------------+--------------+------+-----+---------+----------------+
```

### 3. Struktur DDL Tabel Hasil Normalisasi (`SHOW CREATE TABLE`)

```sql
CREATE TABLE `mst_karyawan` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama_karyawan` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

CREATE TABLE `mst_kpi` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama_kpi` varchar(100) NOT NULL,
  `target_kpi` int NOT NULL,
  `bobot_kpi` decimal(5,2) NOT NULL,
  `bobot_late` decimal(5,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

CREATE TABLE `trx_kpi_marketing` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tasklist` varchar(100) DEFAULT NULL,
  `karyawan_id` int NOT NULL,
  `kpi_id` int NOT NULL,
  `deadline` date DEFAULT NULL,
  `aktual` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_trx_karyawan` (`karyawan_id`),
  KEY `fk_trx_kpi` (`kpi_id`),
  CONSTRAINT `fk_trx_karyawan` FOREIGN KEY (`karyawan_id`) REFERENCES `mst_karyawan` (`id`),
  CONSTRAINT `fk_trx_kpi` FOREIGN KEY (`kpi_id`) REFERENCES `mst_kpi` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
```

### 4. Query & Hasil Relasi Tabel Normalisasi

```sql
SELECT
    t.id,
    t.tasklist,
    kp.nama_kpi AS kpi,
    k.nama_karyawan AS karyawan,
    t.deadline,
    t.aktual
FROM trx_kpi_marketing t
JOIN mst_karyawan k
    ON t.karyawan_id = k.id
JOIN mst_kpi kp
    ON t.kpi_id = kp.id
ORDER BY t.tasklist;
```

```text
+----+-------------+--------+----------+------------+------------+
| id | tasklist    | kpi    | karyawan | deadline   | aktual     |
+----+-------------+--------+----------+------------+------------+
|  4 | Tasklist 1  | Sales  | Budi     | 2022-01-10 | 2022-01-09 |
| 10 | Tasklist 10 | Sales  | Rara     | 2022-01-10 | 2022-01-09 |
|  9 | Tasklist 11 | Report | Rara     | 2022-01-10 | 2022-01-12 |
| 13 | Tasklist 12 | Report | Doni     | 2022-01-10 | 2022-01-09 |
| 12 | Tasklist 13 | Sales  | Doni     | 2022-01-10 | 2022-01-12 |
|  3 | Tasklist 2  | Sales  | Budi     | 2022-01-10 | 2022-01-08 |
|  2 | Tasklist 3  | Report | Budi     | 2022-01-10 | 2022-01-07 |
|  1 | Tasklist 4  | Report | Budi     | 2022-01-10 | 2022-01-12 |
|  8 | Tasklist 5  | Sales  | Adi      | 2022-01-10 | 2022-01-09 |
|  7 | Tasklist 6  | Sales  | Adi      | 2022-01-10 | 2022-01-12 |
|  6 | Tasklist 7  | Report | Adi      | 2022-01-10 | 2022-01-07 |
|  5 | Tasklist 8  | Report | Adi      | 2022-01-10 | 2022-01-07 |
| 11 | Tasklist 9  | Sales  | Rara     | 2022-01-10 | 2022-01-12 |
+----+-------------+--------+----------+------------+------------+
13 rows in set (0.00 sec)
```
