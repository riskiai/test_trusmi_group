-- ============================================================================
-- DATABASE SCHEMA & NORMALIZATION SCRIPT
-- Database: db_test_programmer_trusmi
-- ============================================================================

-- 1. Tabel Original (table_kpi_marketing)
CREATE TABLE IF NOT EXISTS `table_kpi_marketing` (
  `id` int NOT NULL,
  `tasklist` varchar(100) DEFAULT NULL,
  `kpi` varchar(100) DEFAULT NULL,
  `karyawan` varchar(100) DEFAULT NULL,
  `deadline` date DEFAULT NULL,
  `aktual` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- 2. Tabel Master Karyawan (mst_karyawan)
CREATE TABLE IF NOT EXISTS `mst_karyawan` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama_karyawan` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- 3. Tabel Master KPI (mst_kpi)
CREATE TABLE IF NOT EXISTS `mst_kpi` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama_kpi` varchar(100) NOT NULL,
  `target_kpi` int NOT NULL,
  `bobot_kpi` decimal(5,2) NOT NULL,
  `bobot_late` decimal(5,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- 4. Tabel Transaksi KPI Marketing (trx_kpi_marketing)
CREATE TABLE IF NOT EXISTS `trx_kpi_marketing` (
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
