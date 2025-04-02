-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 20, 2025 at 10:15 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.1.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `intern_admin_v3`
--

-- --------------------------------------------------------

--
-- Table structure for table `chi_tiet_hoa_don_nhap_kho`
--

CREATE TABLE `chi_tiet_hoa_don_nhap_kho` (
  `id` int(11) NOT NULL,
  `hoa_don_id` int(11) DEFAULT NULL,
  `san_pham_id` int(11) DEFAULT NULL,
  `ctsp_id` int(11) DEFAULT NULL,
  `sku` varchar(255) DEFAULT NULL,
  `so_luong` varchar(255) DEFAULT NULL,
  `don_vi_tinh` varchar(255) DEFAULT NULL,
  `ke` varchar(255) DEFAULT NULL,
  `gia_nhap` varchar(255) DEFAULT NULL,
  `gia_ban` varchar(255) DEFAULT NULL,
  `chiet_khau` varchar(255) DEFAULT NULL,
  `thanh_tien` double DEFAULT NULL,
  `han_su_dung` timestamp NULL DEFAULT NULL,
  `la_qua_tang` tinyint(4) DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- Table structure for table `chi_tiet_hoa_don_xuat_kho`
--

CREATE TABLE `chi_tiet_hoa_don_xuat_kho` (
  `id` int(11) NOT NULL,
  `hoa_don_id` int(11) DEFAULT NULL,
  `san_pham_id` int(11) DEFAULT NULL,
  `ctsp_id` int(11) DEFAULT NULL,
  `sku` varchar(255) DEFAULT NULL,
  `don_vi_tinh` varchar(255) DEFAULT NULL,
  `so_luong_ban` varchar(255) DEFAULT NULL,
  `gia_ban` varchar(255) DEFAULT NULL,
  `chiet_khau` varchar(255) DEFAULT NULL,
  `thanh_tien` double DEFAULT NULL,
  `gia_nhap` varchar(255) DEFAULT NULL,
  `loi_nhuan` varchar(255) DEFAULT NULL,
  `la_qua_tang` tinyint(4) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- Table structure for table `chi_tiet_san_pham`
--

CREATE TABLE `chi_tiet_san_pham` (
  `id` int(11) NOT NULL,
  `san_pham_id` int(11) DEFAULT NULL,
  `ten_phan_loai` varchar(255) DEFAULT NULL,
  `hinh_anh` varchar(255) DEFAULT NULL,
  `gia_nhap` float DEFAULT NULL,
  `gia_ban` varchar(255) DEFAULT NULL,
  `so_luong` varchar(255) DEFAULT NULL,
  `trang_thai` tinyint(4) DEFAULT NULL,
  `khong_phan_loai` tinyint(4) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `chi_tiet_san_pham`
--

INSERT INTO `chi_tiet_san_pham` (`id`, `san_pham_id`, `ten_phan_loai`, `hinh_anh`, `gia_nhap`, `gia_ban`, `so_luong`, `trang_thai`, `khong_phan_loai`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 'bla bla', 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAA+8AAAFjCAYAAABMlzqvAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAE4VSURBVHgB7d1trF1Xnef5/z7nJkV1QuKgJkQk4O0BDQ9BlZsXFZQCKcd2mjBSQW6qCfBiKr43SY9mmmJiB00JesT43qGHQt0itlUwPYgk9zozIxVQg6+be', 0, '0', '0', 1, 0, '2025-03-14 03:58:33', '2025-03-14 03:58:33', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `chuc_nang`
--

CREATE TABLE `chuc_nang` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `show_in_menu` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `chuc_nang`
--

INSERT INTO `chuc_nang` (`id`, `ten`, `code`, `type`, `show_in_menu`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Xem thời gian bảo hành', 'view-thoi-gian-bao-hanh', 'thoi-gian-bao-hanh', 'Thời gian bảo hành', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(2, 'Thêm thời gian bảo hành', 'create-thoi-gian-bao-hanh', 'thoi-gian-bao-hanh', 'Thời gian bảo hành', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(3, 'Cập nhật thời gian bảo hành', 'update-thoi-gian-bao-hanh', 'thoi-gian-bao-hanh', 'Thời gian bảo hành', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(4, 'Xóa thời gian bảo hành', 'delete-thoi-gian-bao-hanh', 'thoi-gian-bao-hanh', 'Thời gian bảo hành', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(5, 'Xem loại giảm giá', 'view-loai-giam-gia', 'loai-giam-gia', 'Loại giảm giá', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(6, 'Thêm loại giảm giá', 'create-loai-giam-gia', 'loai-giam-gia', 'Loại giảm giá', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(7, 'Cập nhật loại giảm giá', 'update-loai-giam-gia', 'loai-giam-gia', 'Loại giảm giá', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(8, 'Xóa loại giảm giá', 'delete-loai-giam-gia', 'loai-giam-gia', 'Loại giảm giá', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(9, 'Xem đơn vị tính', 'view-don-vi-tinh', 'don-vi-tinh', 'Đơn vị tính', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(10, 'Thêm đơn vị tính', 'create-don-vi-tinh', 'don-vi-tinh', 'Đơn vị tính', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(11, 'Cập nhật đơn vị tính', 'update-don-vi-tinh', 'don-vi-tinh', 'Đơn vị tính', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(12, 'Xóa đơn vị tính', 'delete-don-vi-tinh', 'don-vi-tinh', 'Đơn vị tính', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(13, 'Xem loại sản phẩm', 'view-loai-san-pham', 'loai-san-pham', 'Loại sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(14, 'Thêm loại sản phẩm', 'create-loai-san-pham', 'loai-san-pham', 'Loại sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(15, 'Cập nhật loại sản phẩm', 'update-loai-san-pham', 'loai-san-pham', 'Loại sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(16, 'Xóa loại sản phẩm', 'delete-loai-san-pham', 'loai-san-pham', 'Loại sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(17, 'Xem sản phẩm', 'view-san-pham', 'san-pham', 'Sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(18, 'Thêm sản phẩm', 'create-san-pham', 'san-pham', 'Sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(19, 'Cập nhật sản phẩm', 'update-san-pham', 'san-pham', 'Sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(20, 'Xóa sản phẩm', 'delete-san-pham', 'san-pham', 'Sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(21, 'Xem nhân viên', 'view-nhan-vien', 'nhan-vien', 'Nhân viên', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(22, 'Thêm nhân viên', 'create-nhan-vien', 'nhan-vien', 'Nhân viên', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(23, 'Cập nhật nhân viên', 'update-nhan-vien', 'nhan-vien', 'Nhân viên', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(24, 'Xóa nhân viên', 'delete-nhan-vien', 'nhan-vien', 'Nhân viên', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(25, 'Xem chức vụ', 'view-chuc-vu', 'chuc-vu', 'Chức vụ', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(26, 'Thêm chức vụ', 'create-chuc-vu', 'chuc-vu', 'Chức vụ', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(27, 'Cập nhật chức vụ', 'update-chuc-vu', 'chuc-vu', 'Chức vụ', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(28, 'Xóa chức vụ', 'delete-chuc-vu', 'chuc-vu', 'Chức vụ', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(29, 'Xem kho', 'view-kho', 'kho', 'Kho', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(30, 'Thêm kho', 'create-kho', 'kho', 'Kho', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(31, 'Cập nhật kho', 'update-kho', 'kho', 'Kho', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(32, 'Xóa kho', 'delete-kho', 'kho', 'Kho', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(33, 'Xem quyền', 'view-quyen', 'quyen', 'Quyền', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(34, 'Chỉnh sửa quyền', 'modify-quyen', 'quyen', 'Quyền', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(35, 'Xem chi tiết sản phẩm', 'view-chi-tiet-san-pham', 'chi-tiet-san-pham', 'Chi tiết sản phẩm', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(36, 'Xem tồn kho', 'view-ton-kho', 'ton-kho', 'Tồn kho', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(37, 'Xem công nợ khách hàng', 'view-cong-no-khach-hang', 'cong-no-khach-hang', 'Công nợ khách hàng', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(38, 'Xem công nợ nhà phân phối', 'view-cong-no-nha-phan-phoi', 'cong-no-nha-phan-phoi', 'Công nợ nhà phân phối', '2025-03-20 08:59:34', '2025-03-20 08:59:34', NULL),
(39, 'Xem khách hàng', 'view-khach-hang', 'khach-hang', 'Khách hàng', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(40, 'Thêm khách hàng', 'create-khach-hang', 'khach-hang', 'Khách hàng', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(41, 'Cập nhật khách hàng', 'update-khach-hang', 'khach-hang', 'Khách hàng', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(42, 'Xóa khách hàng', 'delete-khach-hang', 'khach-hang', 'Khách hàng', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(43, 'Xem nhà phân phối', 'view-nha-phan-phoi', 'nha-phan-phoi', 'Nhà phân phối', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(44, 'Thêm nhà phân phối', 'create-nha-phan-phoi', 'nha-phan-phoi', 'Nhà phân phối', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(45, 'Cập nhật nhà phân phối', 'update-nha-phan-phoi', 'nha-phan-phoi', 'Nhà phân phối', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(46, 'Xóa nhà phân phối', 'delete-nha-phan-phoi', 'nha-phan-phoi', 'Nhà phân phối', '2025-03-20 09:03:03', '2025-03-20 09:03:03', NULL),
(47, 'Xem hóa đơn nhập kho', 'view-hoa-don-nhap-kho', 'hoa-don-nhap-kho', 'Hóa đơn nhập kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(48, 'Thêm hóa đơn nhập kho', 'create-hoa-don-nhap-kho', 'hoa-don-nhap-kho', 'Hóa đơn nhập kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(49, 'Cập nhật hóa đơn nhập kho', 'update-hoa-don-nhap-kho', 'hoa-don-nhap-kho', 'Hóa đơn nhập kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(50, 'Khóa hóa đơn nhập kho', 'lock-hoa-don-nhap-kho', 'hoa-don-nhap-kho', 'Hóa đơn nhập kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(51, 'Trả nợ hóa đơn nhập kho', 'tra-no-hoa-don-nhap-kho', 'hoa-don-nhap-kho', 'Hóa đơn nhập kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(52, 'Trả hàng hóa đơn nhập kho', 'tra-hang-hoa-don-nhap-kho', 'hoa-don-nhap-kho', 'Hóa đơn nhập kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(53, 'Xem hóa đơn xuất kho', 'view-hoa-don-xuat-kho', 'hoa-don-xuat-kho', 'Hóa đơn xuất kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(54, 'Thêm hóa đơn xuất kho', 'create-hoa-don-xuat-kho', 'hoa-don-xuat-kho', 'Hóa đơn xuất kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(55, 'Cập nhật hóa đơn xuất kho', 'update-hoa-don-xuat-kho', 'hoa-don-xuat-kho', 'Hóa đơn xuất kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(56, 'Khóa hóa đơn xuất kho', 'lock-hoa-don-xuat-kho', 'hoa-don-xuat-kho', 'Hóa đơn xuất kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(57, 'Trả nợ hóa đơn xuất kho', 'tra-no-hoa-don-xuat-kho', 'hoa-don-xuat-kho', 'Hóa đơn xuất kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL),
(58, 'Trả hàng hóa đơn xuất kho', 'tra-hang-hoa-don-xuat-kho', 'hoa-don-xuat-kho', 'Hóa đơn xuất kho', '2025-03-20 09:05:18', '2025-03-20 09:05:18', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `chuc_vu`
--

CREATE TABLE `chuc_vu` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `chuc_vu`
--

INSERT INTO `chuc_vu` (`id`, `ten`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Super Admin', '2025-02-03 03:22:40', '2025-02-03 03:22:40', NULL),
(2, 'Quản trị viên', '2025-02-03 03:22:40', '2025-02-03 03:22:40', NULL),
(8, 'Chức vụ test', '2025-02-12 09:16:15', '2025-02-12 09:16:15', '2025-02-12 09:16:38'),
(9, '1234', '2025-02-17 02:45:04', '2025-02-17 02:46:08', '2025-02-17 02:54:46'),
(10, 'test', '2025-02-25 10:31:11', '2025-02-25 10:31:11', NULL),
(11, 'aaa', '2025-02-28 02:00:38', '2025-02-28 02:00:38', '2025-02-28 02:01:17'),
(12, 'kien', '2025-02-28 02:51:12', '2025-02-28 02:51:12', NULL),
(13, 'Sale', '2025-02-28 06:57:46', '2025-02-28 06:57:46', NULL),
(14, 'Giao hàng', '2025-02-28 06:57:53', '2025-02-28 06:57:53', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `don_vi_tinh`
--

CREATE TABLE `don_vi_tinh` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `don_vi_tinh`
--

INSERT INTO `don_vi_tinh` (`id`, `ten`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'cai', '2025-03-14 03:57:49', '2025-03-14 03:57:49', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `hoa_don_nhap_kho`
--

CREATE TABLE `hoa_don_nhap_kho` (
  `id` int(11) NOT NULL,
  `so_hoa_don` int(11) DEFAULT NULL,
  `ma_hoa_don` varchar(255) DEFAULT NULL,
  `nha_phan_phoi_id` int(11) DEFAULT NULL,
  `kho_id` int(11) DEFAULT NULL,
  `ngay_nhap` timestamp NULL DEFAULT NULL,
  `tong_tien` double DEFAULT NULL,
  `tra_truoc` double DEFAULT NULL,
  `con_lai` double DEFAULT NULL,
  `ghi_chu` varchar(255) DEFAULT NULL,
  `khoa_don` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- Table structure for table `hoa_don_xuat_kho`
--

CREATE TABLE `hoa_don_xuat_kho` (
  `id` int(11) NOT NULL,
  `so_hoa_don` int(11) DEFAULT NULL,
  `ma_hoa_don` varchar(255) DEFAULT NULL,
  `khach_hang_id` int(11) DEFAULT NULL,
  `nhan_vien_sale_id` int(11) DEFAULT NULL,
  `nhan_vien_giao_hang_id` int(11) DEFAULT NULL,
  `ngay_xuat` timestamp NULL DEFAULT NULL,
  `tong_tien` double DEFAULT NULL,
  `vat` float DEFAULT NULL,
  `thanh_tien` double DEFAULT NULL,
  `tra_truoc` double DEFAULT NULL,
  `con_lai` double DEFAULT NULL,
  `tong_gia_nhap` double DEFAULT NULL,
  `loi_nhuan` double DEFAULT NULL,
  `ghi_chu` varchar(255) DEFAULT NULL,
  `da_giao_hang` tinyint(4) DEFAULT NULL,
  `loai_chiet_khau` tinyint(4) DEFAULT NULL,
  `gia_tri_chiet_khau` varchar(255) DEFAULT NULL,
  `khoa_don` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- Table structure for table `khach_hang`
--

CREATE TABLE `khach_hang` (
  `id` int(11) NOT NULL,
  `ho_ten` varchar(255) DEFAULT NULL,
  `dia_chi` varchar(255) DEFAULT NULL,
  `dien_thoai` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `khach_hang`
--

INSERT INTO `khach_hang` (`id`, `ho_ten`, `dia_chi`, `dien_thoai`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'bin', '123', '0869610949', '2025-03-14 11:57:32', '2025-03-14 11:57:32', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `kho`
--

CREATE TABLE `kho` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `dia_chi` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `kho`
--

INSERT INTO `kho` (`id`, `ten`, `dia_chi`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Binn', '123', '2025-03-14 05:02:26', '2025-03-14 05:02:26', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `loai_giam_gia`
--

CREATE TABLE `loai_giam_gia` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `gia_tri` float DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- Table structure for table `loai_san_pham`
--

CREATE TABLE `loai_san_pham` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `hinh_anh` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `loai_san_pham`
--

INSERT INTO `loai_san_pham` (`id`, `ten`, `hinh_anh`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'mày ngu', 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAA+8AAAFjCAYAAABMlzqvAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAE4VSURBVHgB7d1trF1Xnef5/z7nJkV1QuKgJkQk4O0BDQ9BlZsXFZQCKcd2mjBSQW6qCfBiKr43SY9mmmJiB00JesT43qGHQt0itlUwPYgk9zozIxVQg6+be', '2025-03-14 03:56:53', '2025-03-14 03:56:53', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `nhan_vien`
--

CREATE TABLE `nhan_vien` (
  `id` int(11) NOT NULL,
  `ten_dang_nhap` varchar(255) DEFAULT NULL,
  `mat_khau` varchar(255) DEFAULT NULL,
  `ho_ten` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `dien_thoai` varchar(255) DEFAULT NULL,
  `dia_chi` varchar(255) DEFAULT NULL,
  `avatar` mediumtext DEFAULT NULL,
  `chuc_vu_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `nhan_vien`
--

INSERT INTO `nhan_vien` (`id`, `ten_dang_nhap`, `mat_khau`, `ho_ten`, `email`, `dien_thoai`, `dia_chi`, `avatar`, `chuc_vu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(9, 'moderator', 'None', 'nhân viên sale', 'vanquocdalay@gmail.com', '03887800522', 'HCM2', 'avatar_1.png', 13, '2025-02-03 03:23:34', '2025-02-03 03:23:34', NULL),
(10, 'nhanviengiaohang', 'None', 'giao hàng 1', 'giaohang@gmail.com', '0388780052', '359 phạm văn chiêu', 'avatar_10.png', 14, NULL, NULL, NULL),
(12, 'suki', '$2a$10$Ldv3z7NRZ4GiCY5qpmPeJObXnaO.aPfON9fEJVAHQ4Cr/Ue7.RFji', 'Zinh', 'suki@gmail.com', '123456', 'None', 'None', 1, '2025-03-14 03:25:56', '2025-03-14 03:25:56', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `nha_phan_phoi`
--

CREATE TABLE `nha_phan_phoi` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `dia_chi` varchar(255) DEFAULT NULL,
  `dien_thoai` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `nha_phan_phoi`
--

INSERT INTO `nha_phan_phoi` (`id`, `ten`, `dia_chi`, `dien_thoai`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Bin Phan Phoi', '', '', '', '2025-03-14 03:59:49', '2025-03-14 04:04:03', '2025-03-14 04:04:03'),
(2, 'Tao Phan Phoi ne', 'blabla', '0869610949', 'blalba@gmail.com', '2025-03-14 04:04:29', '2025-03-14 04:29:58', '2025-03-14 04:29:58'),
(3, 'Test ne', '123', '0869610949', 'vinh@gmal.com', '2025-03-14 04:20:21', '2025-03-14 04:30:13', '2025-03-14 04:30:13'),
(4, 'test', '123', '0869610949', 'vinh@gmail.com', '2025-03-14 04:30:29', '2025-03-14 04:38:29', '2025-03-14 04:38:29'),
(5, 'test', '123', '0869610949', '123@gmail.com', '2025-03-14 04:38:45', '2025-03-14 04:38:48', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `quyen`
--

CREATE TABLE `quyen` (
  `id` int(11) NOT NULL,
  `chuc_vu_id` int(11) DEFAULT NULL,
  `chuc_nang_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `quyen`
--

INSERT INTO `quyen` (`id`, `chuc_vu_id`, `chuc_nang_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 1, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(2, 1, 2, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(3, 1, 3, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(4, 1, 4, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(5, 1, 5, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(6, 1, 6, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(7, 1, 7, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(8, 1, 8, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(9, 1, 9, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(10, 1, 10, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(11, 1, 11, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(12, 1, 12, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(13, 1, 13, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(14, 1, 14, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(15, 1, 15, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(16, 1, 16, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(17, 1, 17, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(18, 1, 18, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(19, 1, 19, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(20, 1, 20, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(21, 1, 21, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(22, 1, 22, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(23, 1, 23, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(24, 1, 24, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(25, 1, 25, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(26, 1, 26, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(27, 1, 27, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(28, 1, 28, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(29, 1, 29, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(30, 1, 30, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(31, 1, 31, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(32, 1, 32, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(33, 1, 33, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(34, 1, 34, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(35, 1, 35, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(36, 1, 36, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(37, 1, 37, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(38, 1, 38, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(39, 1, 39, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(40, 1, 40, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(41, 1, 41, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(42, 1, 42, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(43, 1, 43, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(44, 1, 44, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(45, 1, 45, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(46, 1, 46, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(47, 1, 47, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(48, 1, 48, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(49, 1, 49, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(50, 1, 50, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(51, 1, 51, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(52, 1, 52, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(53, 1, 53, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(54, 1, 54, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(55, 1, 55, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(56, 1, 56, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(57, 1, 57, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL),
(58, 1, 58, '2025-03-20 09:13:40', '2025-03-20 09:13:40', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `san_pham`
--

CREATE TABLE `san_pham` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `upc` varchar(255) DEFAULT NULL,
  `loai_san_pham_id` int(11) DEFAULT NULL,
  `hinh_anh` varchar(255) DEFAULT NULL,
  `don_vi_tinh_id` int(11) DEFAULT NULL,
  `vat` float DEFAULT NULL,
  `mo_ta` longblob DEFAULT NULL,
  `trang_thai` tinyint(4) DEFAULT NULL,
  `loai_giam_gia_id` int(11) DEFAULT NULL,
  `thoi_gian_bao_hanh_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `san_pham`
--

INSERT INTO `san_pham` (`id`, `ten`, `upc`, `loai_san_pham_id`, `hinh_anh`, `don_vi_tinh_id`, `vat`, `mo_ta`, `trang_thai`, `loai_giam_gia_id`, `thoi_gian_bao_hanh_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'binn', 'binn', 1, 'data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/4gIoSUNDX1BST0ZJTEUAAQEAAAIYanhsIARAAABtbnRyUkdCIFhZWiAH4wAMAAEAAAAAAABhY3NwQVBQTAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAA9tYAAQAAAADTLWp4bCACufkBQHM6b/D/A/Tw9worAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAtkZXNj', 1, 0, '', 1, 0, 0, '2025-03-14 03:58:33', '2025-03-14 04:38:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `san_pham_nha_phan_phoi`
--

CREATE TABLE `san_pham_nha_phan_phoi` (
  `id` int(11) NOT NULL,
  `nha_phan_phoi_id` int(11) DEFAULT NULL,
  `san_pham_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Dumping data for table `san_pham_nha_phan_phoi`
--

INSERT INTO `san_pham_nha_phan_phoi` (`id`, `nha_phan_phoi_id`, `san_pham_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 3, 1, NULL, NULL, NULL),
(6, 4, 1, NULL, NULL, NULL),
(7, 4, 1, '2025-03-14 04:30:33', '2025-03-14 04:30:33', NULL),
(8, 5, 1, NULL, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `thoi_gian_bao_hanh`
--

CREATE TABLE `thoi_gian_bao_hanh` (
  `id` int(11) NOT NULL,
  `ten` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- Table structure for table `ton_kho`
--

CREATE TABLE `ton_kho` (
  `id` int(11) NOT NULL,
  `san_pham_id` int(11) DEFAULT NULL,
  `ctsp_id` int(11) DEFAULT NULL,
  `sku` varchar(255) DEFAULT NULL,
  `so_luong_ton` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `chi_tiet_hoa_don_nhap_kho`
--
ALTER TABLE `chi_tiet_hoa_don_nhap_kho`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`hoa_don_id`,`ctsp_id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `chi_tiet_hoa_don_xuat_kho`
--
ALTER TABLE `chi_tiet_hoa_don_xuat_kho`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`hoa_don_id`,`san_pham_id`,`ctsp_id`,`sku`,`deleted_at`) USING BTREE;

--
-- Indexes for table `chi_tiet_san_pham`
--
ALTER TABLE `chi_tiet_san_pham`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`san_pham_id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `chuc_nang`
--
ALTER TABLE `chuc_nang`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`,`code`) USING BTREE;

--
-- Indexes for table `chuc_vu`
--
ALTER TABLE `chuc_vu`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `don_vi_tinh`
--
ALTER TABLE `don_vi_tinh`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `hoa_don_nhap_kho`
--
ALTER TABLE `hoa_don_nhap_kho`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`nha_phan_phoi_id`,`ngay_nhap`,`deleted_at`) USING BTREE;

--
-- Indexes for table `hoa_don_xuat_kho`
--
ALTER TABLE `hoa_don_xuat_kho`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `khach_hang`
--
ALTER TABLE `khach_hang`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `kho`
--
ALTER TABLE `kho`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `loai_giam_gia`
--
ALTER TABLE `loai_giam_gia`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `loai_san_pham`
--
ALTER TABLE `loai_san_pham`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `nhan_vien`
--
ALTER TABLE `nhan_vien`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`ten_dang_nhap`,`chuc_vu_id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `nha_phan_phoi`
--
ALTER TABLE `nha_phan_phoi`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `quyen`
--
ALTER TABLE `quyen`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`chuc_vu_id`,`chuc_nang_id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `san_pham`
--
ALTER TABLE `san_pham`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`upc`,`deleted_at`) USING BTREE;

--
-- Indexes for table `san_pham_nha_phan_phoi`
--
ALTER TABLE `san_pham_nha_phan_phoi`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- Indexes for table `thoi_gian_bao_hanh`
--
ALTER TABLE `thoi_gian_bao_hanh`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`deleted_at`) USING BTREE;

--
-- Indexes for table `ton_kho`
--
ALTER TABLE `ton_kho`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx` (`id`,`san_pham_id`,`ctsp_id`,`sku`,`deleted_at`) USING BTREE;

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `chi_tiet_hoa_don_nhap_kho`
--
ALTER TABLE `chi_tiet_hoa_don_nhap_kho`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `chi_tiet_hoa_don_xuat_kho`
--
ALTER TABLE `chi_tiet_hoa_don_xuat_kho`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `chi_tiet_san_pham`
--
ALTER TABLE `chi_tiet_san_pham`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `chuc_nang`
--
ALTER TABLE `chuc_nang`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=59;

--
-- AUTO_INCREMENT for table `chuc_vu`
--
ALTER TABLE `chuc_vu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `don_vi_tinh`
--
ALTER TABLE `don_vi_tinh`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `hoa_don_nhap_kho`
--
ALTER TABLE `hoa_don_nhap_kho`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `hoa_don_xuat_kho`
--
ALTER TABLE `hoa_don_xuat_kho`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `khach_hang`
--
ALTER TABLE `khach_hang`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `kho`
--
ALTER TABLE `kho`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `loai_giam_gia`
--
ALTER TABLE `loai_giam_gia`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `loai_san_pham`
--
ALTER TABLE `loai_san_pham`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `nhan_vien`
--
ALTER TABLE `nhan_vien`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `nha_phan_phoi`
--
ALTER TABLE `nha_phan_phoi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `quyen`
--
ALTER TABLE `quyen`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=59;

--
-- AUTO_INCREMENT for table `san_pham`
--
ALTER TABLE `san_pham`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `san_pham_nha_phan_phoi`
--
ALTER TABLE `san_pham_nha_phan_phoi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `thoi_gian_bao_hanh`
--
ALTER TABLE `thoi_gian_bao_hanh`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ton_kho`
--
ALTER TABLE `ton_kho`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
