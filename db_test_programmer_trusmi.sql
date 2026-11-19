-- phpMyAdmin SQL Dump
-- version 5.2.3
-- https://www.phpmyadmin.net/
--
-- Host: mysql-global:3306
-- Generation Time: Nov 19, 2026 at 10:19 AM
-- Server version: 8.4.13
-- PHP Version: 8.3.36

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_test_programmer_trusmi`
--

-- --------------------------------------------------------

--
-- Table structure for table `mst_karyawan`
--

CREATE TABLE `mst_karyawan` (
  `id` int NOT NULL,
  `nama_karyawan` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `mst_karyawan`
--

INSERT INTO `mst_karyawan` (`id`, `nama_karyawan`) VALUES
(1, 'Budi'),
(2, 'Adi'),
(3, 'Rara'),
(4, 'Doni');

-- --------------------------------------------------------

--
-- Table structure for table `mst_kpi`
--

CREATE TABLE `mst_kpi` (
  `id` int NOT NULL,
  `nama_kpi` varchar(100) NOT NULL,
  `target_kpi` int NOT NULL,
  `bobot_kpi` decimal(5,2) NOT NULL,
  `bobot_late` decimal(5,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `mst_kpi`
--

INSERT INTO `mst_kpi` (`id`, `nama_kpi`, `target_kpi`, `bobot_kpi`, `bobot_late`) VALUES
(1, 'Sales', 2, 50.00, -7.00),
(2, 'Report', 2, 50.00, -5.00);

-- --------------------------------------------------------

--
-- Table structure for table `table_kpi_marketing`
--

CREATE TABLE `table_kpi_marketing` (
  `id` int NOT NULL,
  `tasklist` varchar(100) DEFAULT NULL,
  `kpi` varchar(100) DEFAULT NULL,
  `karyawan` varchar(100) DEFAULT NULL,
  `deadline` date DEFAULT NULL,
  `aktual` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `table_kpi_marketing`
--

INSERT INTO `table_kpi_marketing` (`id`, `tasklist`, `kpi`, `karyawan`, `deadline`, `aktual`) VALUES
(1, 'Tasklist 1', 'Sales', 'Budi', '2022-01-10', '2022-01-09'),
(2, 'Tasklist 2', 'Sales', 'Budi', '2022-01-10', '2022-01-08'),
(3, 'Tasklist 3', 'Report', 'Budi', '2022-01-10', '2022-01-07'),
(4, 'Tasklist 4', 'Report', 'Budi', '2022-01-10', '2022-01-12'),
(5, 'Tasklist 5', 'Sales', 'Adi', '2022-01-10', '2022-01-09'),
(6, 'Tasklist 6', 'Sales', 'Adi', '2022-01-10', '2022-01-12'),
(7, 'Tasklist 7', 'Report', 'Adi', '2022-01-10', '2022-01-07'),
(8, 'Tasklist 8', 'Report', 'Adi', '2022-01-10', '2022-01-07'),
(9, 'Tasklist 9', 'Sales', 'Rara', '2022-01-10', '2022-01-12'),
(10, 'Tasklist 10', 'Sales', 'Rara', '2022-01-10', '2022-01-09'),
(11, 'Tasklist 11', 'Report', 'Rara', '2022-01-10', '2022-01-12'),
(12, 'Tasklist 12', 'Report', 'Doni', '2022-01-10', '2022-01-09'),
(13, 'Tasklist 13', 'Sales', 'Doni', '2022-01-10', '2022-01-12');

-- --------------------------------------------------------

--
-- Table structure for table `trx_kpi_marketing`
--

CREATE TABLE `trx_kpi_marketing` (
  `id` int NOT NULL,
  `tasklist` varchar(100) DEFAULT NULL,
  `karyawan_id` int NOT NULL,
  `kpi_id` int NOT NULL,
  `deadline` date DEFAULT NULL,
  `aktual` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `trx_kpi_marketing`
--

INSERT INTO `trx_kpi_marketing` (`id`, `tasklist`, `karyawan_id`, `kpi_id`, `deadline`, `aktual`) VALUES
(1, 'Tasklist 4', 1, 2, '2022-01-10', '2022-01-12'),
(2, 'Tasklist 3', 1, 2, '2022-01-10', '2022-01-07'),
(3, 'Tasklist 2', 1, 1, '2022-01-10', '2022-01-08'),
(4, 'Tasklist 1', 1, 1, '2022-01-10', '2022-01-09'),
(5, 'Tasklist 8', 2, 2, '2022-01-10', '2022-01-07'),
(6, 'Tasklist 7', 2, 2, '2022-01-10', '2022-01-07'),
(7, 'Tasklist 6', 2, 1, '2022-01-10', '2022-01-12'),
(8, 'Tasklist 5', 2, 1, '2022-01-10', '2022-01-09'),
(9, 'Tasklist 11', 3, 2, '2022-01-10', '2022-01-12'),
(10, 'Tasklist 10', 3, 1, '2022-01-10', '2022-01-09'),
(11, 'Tasklist 9', 3, 1, '2022-01-10', '2022-01-12'),
(12, 'Tasklist 13', 4, 1, '2022-01-10', '2022-01-12'),
(13, 'Tasklist 12', 4, 2, '2022-01-10', '2022-01-09');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `mst_karyawan`
--
ALTER TABLE `mst_karyawan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mst_kpi`
--
ALTER TABLE `mst_kpi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `table_kpi_marketing`
--
ALTER TABLE `table_kpi_marketing`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `trx_kpi_marketing`
--
ALTER TABLE `trx_kpi_marketing`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_trx_karyawan` (`karyawan_id`),
  ADD KEY `fk_trx_kpi` (`kpi_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `mst_karyawan`
--
ALTER TABLE `mst_karyawan`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `mst_kpi`
--
ALTER TABLE `mst_kpi`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `trx_kpi_marketing`
--
ALTER TABLE `trx_kpi_marketing`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `trx_kpi_marketing`
--
ALTER TABLE `trx_kpi_marketing`
  ADD CONSTRAINT `fk_trx_karyawan` FOREIGN KEY (`karyawan_id`) REFERENCES `mst_karyawan` (`id`),
  ADD CONSTRAINT `fk_trx_kpi` FOREIGN KEY (`kpi_id`) REFERENCES `mst_kpi` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
