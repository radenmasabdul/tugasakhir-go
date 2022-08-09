-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 09, 2022 at 02:14 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_crud`
--

-- --------------------------------------------------------

--
-- Table structure for table `task`
--

CREATE TABLE `task` (
  `id` int(10) UNSIGNED NOT NULL,
  `task` text NOT NULL,
  `assigne` varchar(30) NOT NULL,
  `deadline` date NOT NULL,
  `isdone` tinyint(1) DEFAULT 0,
  `updatedate` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `task`
--

INSERT INTO `task` (`id`, `task`, `assigne`, `deadline`, `isdone`, `updatedate`) VALUES
(2, 'Data untuk main.go baru saja di perbaharui. Sudah aktif, validasi OK!', 'Abdul Rahman Alhafizh', '2022-08-10', 1, '2022-08-08 04:22:51'),
(6, 'Data created task sudah di tambahkan validasi', 'Abdul Rahman Alhafizh', '2022-08-10', 1, '2022-08-08 04:59:41'),
(7, 'Data validasi pada Form Created Task sudah di perbaharui', 'Abdul Rahman Alhafizh', '2022-08-10', 1, '2022-08-08 04:09:12'),
(9, 'Data Main go, Add datatables editeds', 'Vicky Himawan', '2022-08-10', 1, '2022-08-08 05:00:19'),
(10, 'Data pada main.go sudah di tambahkan fungsi Is Done', 'Abdul Rahman Alhafizh', '2022-08-10', 1, '2022-08-08 04:59:52'),
(11, 'Data main.go di tambahkan button. Fungsi button OK!', 'Vicky Himawan', '2022-08-10', 1, '2022-08-08 05:00:20'),
(13, 'Cek all fungsi, OK!', 'Abdul Rahman Alhafizh', '2022-08-10', 1, '2022-08-09 12:04:19'),
(14, 'Tugas akhir DONE! Fungsi OKE! ALL DONE!', 'Abdul Rahman Alhafizh', '2022-08-10', 0, '2022-08-09 12:04:11');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `task`
--
ALTER TABLE `task`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `task`
--
ALTER TABLE `task`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
