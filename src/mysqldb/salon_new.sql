-- phpMyAdmin SQL Dump
-- version 4.6.6deb4
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Feb 02, 2020 at 03:56 PM
-- Server version: 10.1.41-MariaDB-0+deb9u1
-- PHP Version: 7.0.33-0+deb9u6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `salon_new`
--
CREATE DATABASE salon_new;

USE salon_new;
-- --------------------------------------------------------

--
-- Table structure for table `appointments`
--

CREATE TABLE `appointments` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `hairdresser_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `service_id` int(11) NOT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `service_status` varchar(45) DEFAULT NULL,
  `comments` text,
  `is_all_day` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `appointments_with_promotions`
--

CREATE TABLE `appointments_with_promotions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `hairdresser_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `promotion_id` int(11) NOT NULL,
  `appointment_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `bank_holidays`
--

CREATE TABLE `bank_holidays` (
  `id` int(11) NOT NULL,
  `summary` varchar(150) DEFAULT NULL,
  `date_start` date DEFAULT NULL,
  `date_end` date DEFAULT NULL,
  `country` varchar(25) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `bank_holidays_per_store`
--

CREATE TABLE `bank_holidays_per_store` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `bank_holidays_country` varchar(150) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `customer_name` varchar(100) NOT NULL,
  `customer_surname` varchar(100) NOT NULL,
  `gender_of_customer` varchar(20) NOT NULL,
  `date_of_birth` date DEFAULT NULL,
  `customer_address` varchar(300) NOT NULL,
  `customer_street_number` int(11) NOT NULL,
  `customer_city` varchar(100) NOT NULL,
  `customer_state` varchar(100) NOT NULL,
  `customer_zip_code` varchar(50) NOT NULL,
  `customer_country` varchar(70) NOT NULL,
  `home_phone_number` varchar(100) DEFAULT NULL,
  `mobile_phone_number` varchar(100) NOT NULL,
  `customer_email` varchar(70) NOT NULL,
  `customer_points` int(11) NOT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `pelatis_lianikis` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `user_id`, `customer_name`, `customer_surname`, `gender_of_customer`, `date_of_birth`, `customer_address`, `customer_street_number`, `customer_city`, `customer_state`, `customer_zip_code`, `customer_country`, `home_phone_number`, `mobile_phone_number`, `customer_email`, `customer_points`, `is_active`, `pelatis_lianikis`, `created_at`, `updated_at`) VALUES
(1, 1, 'ΜΑΡΙΑ', 'ΠΑΠΑΔΟΠΟΥΛΟΥ', 'female', '1978-06-23', 'Αυγουστίνου', 12, 'Άγιοι Θεόδωροι', '', '200 03', 'Ελλάδα', '', '6999111999', 'mariapapadopoulou@misel.gr', 0, 1, 0, '2020-01-14 10:22:00', '2020-01-14 10:22:00'),
(2, 1, 'Walk-In', 'Customer', 'other', '2020-01-23', '1231', 1231, 'Κεμπέκ', 'SP', '12323-400', 'Βραζιλία', '123', '123455667', 'info@info.infoasd', 0, 1, 1, '2020-01-23 19:21:42', '2020-01-23 19:21:42');

-- --------------------------------------------------------

--
-- Table structure for table `daily_values`
--

CREATE TABLE `daily_values` (
  `wdate` datetime DEFAULT NULL,
  `wval` decimal(10,0) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `expensestrans`
--

CREATE TABLE `expensestrans` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `expenses_list_id` int(11) NOT NULL,
  `expenses_price` decimal(5,2) DEFAULT NULL,
  `payment_type` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `expenses_list`
--

CREATE TABLE `expenses_list` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `expenses_name` varchar(100) NOT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `greek_name_days`
--

CREATE TABLE `greek_name_days` (
  `id` int(11) NOT NULL,
  `celebreation_date` datetime DEFAULT NULL,
  `name_day` varchar(150) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `hairdressers`
--

CREATE TABLE `hairdressers` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `hairdresser_name` varchar(200) NOT NULL,
  `hairdresser_mobile_phone` varchar(200) NOT NULL,
  `hairdresser_phone` varchar(100) NOT NULL,
  `display_order` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------------

--
-- Table structure for table `hairdresser_stores`
--

CREATE TABLE `hairdresser_stores` (
  `id` int(11) NOT NULL,
  `hairdresser_id` int(11) DEFAULT NULL,
  `store_id` int(11) DEFAULT NULL,
  `is_active_hairdresser` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `invoices`
--

CREATE TABLE `invoices` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `customer_id` int(11) DEFAULT NULL,
  `sub_total` float DEFAULT NULL,
  `discount` float DEFAULT NULL,
  `total_cost` float DEFAULT NULL,
  `status_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `invoices`
--

INSERT INTO `invoices` (`id`, `user_id`, `store_id`, `customer_id`, `sub_total`, `discount`, `total_cost`, `status_id`, `created_at`, `updated_at`) VALUES
(1, 1, 4, 1, 0, 0, 0, 1, '2020-01-14 15:28:07', '2020-01-14 15:28:07'),
(2, 1, 1, 0, 0, 0, 0, 1, '2020-01-15 04:30:38', '2020-01-15 04:30:38');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_code` varchar(150) NOT NULL,
  `user_id` int(11) NOT NULL,
  `supplier_id` int(11) NOT NULL,
  `supplier_name` varchar(200) NOT NULL,
  `product_name` varchar(200) NOT NULL,
  `product_ml` int(11) NOT NULL,
  `product_price` decimal(5,2) NOT NULL,
  `product_ml_per_price` decimal(5,0) NOT NULL,
  `store_id` int(11) DEFAULT NULL,
  `is_active_product` tinyint(1) DEFAULT NULL,
  `can_split_product` tinyint(1) DEFAULT NULL,
  `qty` float DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `product_code`, `user_id`, `supplier_id`, `supplier_name`, `product_name`, `product_ml`, `product_price`, `product_ml_per_price`, `store_id`, `is_active_product`, `can_split_product`, `qty`, `created_at`, `updated_at`) VALUES
(1, 'CV223', 1, 1, 'ASP', 'test', 23, '23.00', '100', 1, 0, 0, 2, '2020-01-22 02:49:08', '2020-01-22 02:49:08'),
(2, '123123', 1, 1, 'ASP', '123qwe', 2, '35.00', '1750', 3, 0, 0, 123123000, '2020-01-24 13:11:19', '2020-01-24 13:11:19'),
(3, '345345', 1, 2, '123123', 'threhd', 66, '60.00', '91', 4, 0, 0, 676767, '2020-01-27 12:05:26', '2020-01-27 12:05:26'),
(4, '234234', 1, 2, '123123', '23423', 234, '999.99', '1383', 2, 0, 0, 23423, '2020-01-29 02:32:17', '2020-01-29 02:32:17');



-- --------------------------------------------------------

--
-- Table structure for table `product_stores`
--

CREATE TABLE `product_stores` (
  `id` int(11) NOT NULL,
  `product_id` int(11) DEFAULT NULL,
  `store_id` int(11) DEFAULT NULL,
  `is_active_product` tinyint(1) DEFAULT NULL,
  `qty` float DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------------

--
-- Table structure for table `promotions`
--

CREATE TABLE `promotions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `promotion_title` varchar(150) NOT NULL,
  `days_duration` int(11) NOT NULL,
  `promotion_sale` decimal(5,2) NOT NULL,
  `promotion_description` text,
  `is_service` tinyint(1) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



-- --------------------------------------------------------

--
-- Table structure for table `promotions_commons`
--

CREATE TABLE `promotions_commons` (
  `id` int(11) NOT NULL,
  `promotion_id` int(11) NOT NULL,
  `promotion_service_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `required_products`
--

CREATE TABLE `required_products` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `appointment_service_id` int(11) NOT NULL,
  `product_used_ml` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `salestrans`
--

CREATE TABLE `salestrans` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `service_id` int(11) NOT NULL,
  `service_price` decimal(5,2) DEFAULT NULL,
  `service_qty` int(11) DEFAULT NULL,
  `service_discount` decimal(5,2) DEFAULT NULL,
  `payment_type` varchar(50) DEFAULT NULL,
  `service_line_total` decimal(5,2) DEFAULT NULL,
  `is_service` tinyint(1) DEFAULT NULL,
  `invoice_number` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `salestrans_common`
--

CREATE TABLE `salestrans_common` (
  `id` int(11) NOT NULL,
  `salestrans_id` int(11) NOT NULL,
  `hairdresser_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `appointment_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `sales_products_with_promotions`
--

CREATE TABLE `sales_products_with_promotions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `promotion_id` int(11) NOT NULL,
  `salestrans_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

CREATE TABLE `services` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `sub_category_id` int(11) NOT NULL,
  `service_name` varchar(100) NOT NULL,
  `service_duration` int(200) NOT NULL,
  `service_price` decimal(5,2) NOT NULL,
  `service_discount` decimal(5,2) DEFAULT NULL,
  `sub_category` varchar(100) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `switch_formula` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `services`
--

INSERT INTO `services` (`id`, `user_id`, `category_id`, `sub_category_id`, `service_name`, `service_duration`, `service_price`, `service_discount`, `sub_category`, `is_active`, `switch_formula`, `created_at`, `updated_at`) VALUES
(1, 1, 0, 0, 'ΚΟΥΡΕΜΑ', 30, '15.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:14:29', '2020-01-14 10:14:29'),
(2, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '20.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:02', '2020-01-14 10:16:02'),
(3, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '20.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:05', '2020-01-14 10:16:05'),
(4, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '20.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:07', '2020-01-14 10:16:07'),
(5, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '15.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:51', '2020-01-14 10:16:51'),
(6, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '15.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:53', '2020-01-14 10:16:53'),
(7, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '15.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:54', '2020-01-14 10:16:54'),
(8, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '15.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:55', '2020-01-14 10:16:55'),
(9, 1, 1, 1, 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ ΜΕ ΛΟΥΣΙΜΟ ΚΑΙ ΦΟΡΜΑΡΙΣΜΑ', 30, '15.00', '0.00', NULL, NULL, NULL, '2020-01-14 10:16:56', '2020-01-14 10:16:56'),
(10, 1, 1, 1, 'admin', 45, '50.00', '40.00', NULL, NULL, NULL, '2020-01-24 13:05:25', '2020-01-24 13:05:25'),
(11, 1, 1, 1, 'admin', 15, '500.00', '50.00', NULL, NULL, NULL, '2020-01-24 18:13:10', '2020-01-24 18:13:10'),
(12, 1, 1, 1, 'admin', 15, '500.00', '50.00', NULL, NULL, NULL, '2020-01-24 18:13:11', '2020-01-24 18:13:11'),
(13, 1, 1, 1, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:40:53', '2020-01-29 19:40:53'),
(14, 1, 1, 1, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:40:54', '2020-01-29 19:40:54'),
(15, 1, 1, 1, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:08', '2020-01-29 19:41:08'),
(16, 1, 1, 1, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:10', '2020-01-29 19:41:10'),
(17, 1, 1, 1, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:10', '2020-01-29 19:41:10'),
(18, 1, 1, 1, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:10', '2020-01-29 19:41:10'),
(19, 1, 0, 0, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:13', '2020-01-29 19:41:13'),
(20, 1, 0, 0, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:14', '2020-01-29 19:41:14'),
(21, 1, 0, 0, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:15', '2020-01-29 19:41:15'),
(22, 1, 0, 0, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:19', '2020-01-29 19:41:19'),
(23, 1, 0, 0, 'sdfsdf', 45, '999.99', '999.99', NULL, NULL, NULL, '2020-01-29 19:41:20', '2020-01-29 19:41:20');


-- --------------------------------------------------------

--
-- Table structure for table `services_stores`
--

CREATE TABLE `services_stores` (
  `id` int(11) NOT NULL,
  `service_id` int(11) DEFAULT NULL,
  `store_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `service_categories`
--

CREATE TABLE `service_categories` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `category_name` varchar(100) NOT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `service_categories`
--

INSERT INTO `service_categories` (`id`, `user_id`, `category_name`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 1, 'ΥΠΗΡΕΣΙΕΣ ΚΟΜΜΩΤΗΡΙΟΥ', 1, '2020-01-14 10:14:59', '2020-01-14 10:14:59');

-- --------------------------------------------------------

--
-- Table structure for table `service_sub_categories`
--

CREATE TABLE `service_sub_categories` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `category_id` int(11) DEFAULT NULL,
  `category_name` varchar(120) DEFAULT NULL,
  `sub_category_name` varchar(120) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `service_sub_categories`
--

INSERT INTO `service_sub_categories` (`id`, `user_id`, `category_id`, `category_name`, `sub_category_name`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'ΥΠΗΡΕΣΙΕΣ ΚΟΜΜΩΤΗΡΙΟΥ', 'ΓΥΝΑΙΚΕΙΟ ΚΟΥΡΕΜΑ', 1, '2020-01-14 10:15:29', '2020-01-14 10:15:29');

-- --------------------------------------------------------

--
-- Table structure for table `shops`
--

CREATE TABLE `shops` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `vat_number` varchar(30) NOT NULL,
  `company_name` varchar(200) NOT NULL,
  `company_address` varchar(300) NOT NULL,
  `company_street_number` int(20) NOT NULL,
  `company_city` varchar(100) NOT NULL,
  `company_state` varchar(200) DEFAULT NULL,
  `company_zip_code` varchar(80) NOT NULL,
  `tax_office` varchar(150) NOT NULL,
  `company_country` varchar(100) NOT NULL,
  `work_telephone` varchar(120) DEFAULT NULL,
  `mobile_telephone` varchar(120) DEFAULT NULL,
  `password` varchar(120) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `shops`
--

INSERT INTO `shops` (`id`, `user_id`, `vat_number`, `company_name`, `company_address`, `company_street_number`, `company_city`, `company_state`, `company_zip_code`, `tax_office`, `company_country`, `work_telephone`, `mobile_telephone`, `password`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 1, 'asd', 'qwe', '456456', 1231, 'Κεμπέκ', 'SP', '12323-400', '123234', 'Βραζιλία', '123', '123455667', '$2a$10$Jh9HM8cMlXUuTxF/v1FwqeMAE1WKmJ1744oJIDOlvtADoybh3HOGy', 1, '2020-01-14 07:34:51', '2020-01-29 20:16:40'),
(2, 1, 'qweqwe', 'qweqwe', '1231', 1231, 'Κεμπέκ', 'SP', '12323-400', '123234', 'Βραζιλία', '123', '123455667', '$2a$10$wf2dj4gjkt.svYNxlGv8WOhmj8vVfc0JdT/roJSIyUAvzYttZEOuG', 1, '2020-01-14 07:35:19', '2020-01-14 07:35:19'),
(3, 1, 'qweqweasd', 'qweqweasd', '1231', 1231, 'Κεμπέκ', 'SP', '12323-400', '123234', 'Βραζιλία', '123', '123455667', '$2a$10$76gWq4RXf3MQysqKrIpXaeL13R3tagYDxlZahbK4AbdSyClLX.wk.', 1, '2020-01-14 07:35:30', '2020-01-14 07:35:30'),
(4, 1, '999877897', 'ΑΙΓΑΛΕΩ', 'Ιερά Οδός', 236, 'Αιγάλεω', 'Αττική', '122 42', 'Αιγάλεω', 'Ελλάδα', '2113005030', '', '$2a$10$Owkce1yjH9TTTeCcLARGvuTfM8RYdRESrYdpWcR1UVPr0VjspDJBq', 1, '2020-01-14 10:04:11', '2020-01-14 10:04:11');

-- --------------------------------------------------------

--
-- Table structure for table `suppliers`
--

CREATE TABLE `suppliers` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `vat_number` varchar(30) NOT NULL,
  `supplier_name` varchar(200) NOT NULL,
  `supplier_address` varchar(300) NOT NULL,
  `supplier_street_number` int(20) NOT NULL,
  `supplier_city` varchar(100) NOT NULL,
  `supplier_state` varchar(200) DEFAULT NULL,
  `supplier_zip_code` varchar(80) NOT NULL,
  `field_of_business` varchar(150) NOT NULL,
  `supplier_country` varchar(100) NOT NULL,
  `work_telephone` varchar(120) DEFAULT NULL,
  `email` varchar(120) DEFAULT NULL,
  `website` varchar(250) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `suppliers`
--

INSERT INTO `suppliers` (`id`, `user_id`, `vat_number`, `supplier_name`, `supplier_address`, `supplier_street_number`, `supplier_city`, `supplier_state`, `supplier_zip_code`, `field_of_business`, `supplier_country`, `work_telephone`, `email`, `website`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 1, '1232123422', 'ASP', 'Ιερά Οδός', 236, 'Ελευσίνα', 'ΑΤΤΙΚΗ', '192 00', 'KSDKSBVK', 'Ελλάδα', '32324828929', 'KSKC@KSDK.COM', 'SKBCSB', 1, '2020-01-14 10:31:38', '2020-01-14 10:31:38'),
(2, 1, '123', '123123', '广东省深圳市南山区', 0, '深圳', '广东', '425100', 'asd', 'Ινδία', '123122', 'alonlong@126.com', 'qweasd', 1, '2020-01-27 12:04:04', '2020-01-27 12:04:04'),
(3, 1, '345345', 'dfaws', '', 0, '', '', '', 'cvhb', '', '231412', 'sdfng@facebook.com', 'sdfsdaf', 1, '2020-01-29 20:03:14', '2020-01-29 20:03:14'),
(4, 1, '214321', 'sdfsd', '', 324, 'Kolkatasdf', 'WB', '700091', 'cvhb', 'Ινδία', '32423', 'jg@facebook.com', 'sdfsdaf', 1, '2020-01-29 20:04:24', '2020-01-29 20:04:24');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `role` varchar(120) DEFAULT NULL,
  `forename` varchar(120) DEFAULT NULL,
  `surname` varchar(120) DEFAULT NULL,
  `work_telephone` varchar(120) DEFAULT NULL,
  `mobile_telephone` varchar(120) DEFAULT NULL,
  `company_name` varchar(200) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(120) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `logo_image` varchar(60) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `last_login` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `role`, `forename`, `surname`, `work_telephone`, `mobile_telephone`, `company_name`, `username`, `email`, `password`, `is_active`, `logo_image`, `created_at`, `updated_at`, `last_login`) VALUES
(1, 'administrator', 'admin', 'hi', '10000000000', '10000000000', 'admin', 'admin', 'admin@163.com', '$2a$10$cAsuLPm/gSddUGUxWWR05OuOd8cT/8uzY8TO2EJ/b9yiOIUqPiUCu', 1, 'stor_admin_admin.jpg', '2020-01-14 06:42:29', '2020-01-14 06:42:29', '2020-02-02 15:52:57'),
(2, 'storeadmin', 'Kleanthis', 'Kefalas', '2113005030', '6947300987', 'MISEL GROUP Hairdressing & Beauty ', 'miselgroupsalons', 'kleanthiskefalas@gmail.com', '$2a$10$hC6XFH3QvaqN3VYz4TvDhujoda4NWCp7zTzGylV2qIAx0KjakpCrK', 0, 'stor_admin_miselgroupsalons.jpg', '2020-01-14 11:37:45', '2020-01-14 11:37:45', '2020-01-14 11:38:05');

-- --------------------------------------------------------

--
-- Table structure for table `users_activation`
--

CREATE TABLE `users_activation` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  `email` varchar(50) NOT NULL,
  `activation_code` text,
  `creation_date` datetime DEFAULT NULL,
  `activation_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users_activation`
--

INSERT INTO `users_activation` (`id`, `user_id`, `username`, `email`, `activation_code`, `creation_date`, `activation_date`) VALUES
(1, 2, 'miselgroupsalons', 'kleanthiskefalas@gmail.com', '43_thzQcgfr2wbRpPZhgDcPffE4xvB6h', '2020-01-14 11:37:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `users_packages`
--

CREATE TABLE `users_packages` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `month_periods` int(11) NOT NULL,
  `number_of_stores` int(11) NOT NULL,
  `amount` decimal(5,2) NOT NULL,
  `credits` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `weekly_hours`
--

CREATE TABLE `weekly_hours` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `store_id` int(11) NOT NULL,
  `day_of_the_week` int(11) NOT NULL,
  `from_time` varchar(20) NOT NULL,
  `to_time` varchar(20) NOT NULL,
  `color` varchar(45) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `weekly_hours`
--

INSERT INTO `weekly_hours` (`id`, `user_id`, `store_id`, `day_of_the_week`, `from_time`, `to_time`, `color`, `created_at`, `updated_at`) VALUES
(2, 1, 4, 2, '10:00', '20:00', '', '2020-01-14 10:08:24', '2020-01-14 10:08:24'),
(3, 1, 4, 3, '10:00', '15:00', '', '2020-01-14 10:09:38', '2020-01-14 10:09:38'),
(4, 1, 4, 4, '10:00', '20:00', '', '2020-01-14 10:10:18', '2020-01-14 10:10:18'),
(5, 1, 4, 5, '10:00', '20:00', '', '2020-01-14 10:10:50', '2020-01-14 10:10:50'),
(6, 1, 4, 6, '10:00', '16:00', '', '2020-01-14 10:11:23', '2020-01-14 10:11:23');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `appointments`
--
ALTER TABLE `appointments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `hairdresser_id` (`hairdresser_id`),
  ADD KEY `customer_id` (`customer_id`),
  ADD KEY `store_id` (`store_id`);

--
-- Indexes for table `appointments_with_promotions`
--
ALTER TABLE `appointments_with_promotions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `appointment_id` (`appointment_id`),
  ADD KEY `promotion_id` (`promotion_id`),
  ADD KEY `hairdresser_id` (`hairdresser_id`),
  ADD KEY `customer_id` (`customer_id`),
  ADD KEY `store_id` (`store_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `bank_holidays`
--
ALTER TABLE `bank_holidays`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `date_start` (`summary`,`date_start`,`date_end`,`country`),
  ADD UNIQUE KEY `k_country` (`country`);

--
-- Indexes for table `bank_holidays_per_store`
--
ALTER TABLE `bank_holidays_per_store`
  ADD PRIMARY KEY (`id`),
  ADD KEY `bank_holidays_country` (`bank_holidays_country`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `customer_email` (`customer_email`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `expensestrans`
--
ALTER TABLE `expensestrans`
  ADD PRIMARY KEY (`id`),
  ADD KEY `expenses_list_id` (`expenses_list_id`),
  ADD KEY `store_id` (`store_id`);

--
-- Indexes for table `expenses_list`
--
ALTER TABLE `expenses_list`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `greek_name_days`
--
ALTER TABLE `greek_name_days`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `hairdressers`
--
ALTER TABLE `hairdressers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `hairdresser_stores`
--
ALTER TABLE `hairdresser_stores`
  ADD PRIMARY KEY (`d`),
  ADD KEY `hairdresser_id` (`hairdresser_id`),
  ADD KEY `store_id` (`store_id`);

--
-- Indexes for table `invoices`
--
ALTER TABLE `invoices`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `supplier_id` (`supplier_id`),
  ADD KEY `store_id` (`store_id`);

--
-- Indexes for table `product_stores`
--
ALTER TABLE `product_stores`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `store_id` (`store_id`);

--
-- Indexes for table `promotions`
--
ALTER TABLE `promotions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `store_id` (`store_id`);

--
-- Indexes for table `promotions_commons`
--
ALTER TABLE `promotions_commons`
  ADD PRIMARY KEY (`id`),
  ADD KEY `promotion_id` (`promotion_id`);

--
-- Indexes for table `required_products`
--
ALTER TABLE `required_products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `salestrans`
--
ALTER TABLE `salestrans`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `salestrans_common`
--
ALTER TABLE `salestrans_common`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sales_products_with_promotions`
--
ALTER TABLE `sales_products_with_promotions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `salestrans_id` (`salestrans_id`),
  ADD KEY `promotion_id` (`promotion_id`),
  ADD KEY `customer_id` (`customer_id`),
  ADD KEY `store_id` (`store_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `services`
--
ALTER TABLE `services`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `services_stores`
--
ALTER TABLE `services_stores`
  ADD PRIMARY KEY (`id`),
  ADD KEY `service_id` (`service_id`),
  ADD KEY `store_id` (`store_id`);

--
-- Indexes for table `service_categories`
--
ALTER TABLE `service_categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `service_sub_categories`
--
ALTER TABLE `service_sub_categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_id` (`category_id`);

--
-- Indexes for table `shops`
--
ALTER TABLE `shops`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `vat_number` (`vat_number`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `suppliers`
--
ALTER TABLE `suppliers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `vat_number` (`vat_number`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `users_activation`
--
ALTER TABLE `users_activation`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users_packages`
--
ALTER TABLE `users_packages`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `weekly_hours`
--
ALTER TABLE `weekly_hours`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `appointments`
--
ALTER TABLE `appointments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `appointments_with_promotions`
--
ALTER TABLE `appointments_with_promotions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `bank_holidays`
--
ALTER TABLE `bank_holidays`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `bank_holidays_per_store`
--
ALTER TABLE `bank_holidays_per_store`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `expensestrans`
--
ALTER TABLE `expensestrans`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `expenses_list`
--
ALTER TABLE `expenses_list`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `greek_name_days`
--
ALTER TABLE `greek_name_days`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `hairdressers`
--
ALTER TABLE `hairdressers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `hairdresser_stores`
--
ALTER TABLE `hairdresser_stores`
  MODIFY `d` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `invoices`
--
ALTER TABLE `invoices`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `product_stores`
--
ALTER TABLE `product_stores`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `promotions`
--
ALTER TABLE `promotions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `promotions_commons`
--
ALTER TABLE `promotions_commons`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `required_products`
--
ALTER TABLE `required_products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `salestrans`
--
ALTER TABLE `salestrans`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `salestrans_common`
--
ALTER TABLE `salestrans_common`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `sales_products_with_promotions`
--
ALTER TABLE `sales_products_with_promotions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `services`
--
ALTER TABLE `services`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;
--
-- AUTO_INCREMENT for table `services_stores`
--
ALTER TABLE `services_stores`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `service_categories`
--
ALTER TABLE `service_categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `service_sub_categories`
--
ALTER TABLE `service_sub_categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `shops`
--
ALTER TABLE `shops`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `suppliers`
--
ALTER TABLE `suppliers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `users_activation`
--
ALTER TABLE `users_activation`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `users_packages`
--
ALTER TABLE `users_packages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `weekly_hours`
--
ALTER TABLE `weekly_hours`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `appointments`
--
ALTER TABLE `appointments`
  ADD CONSTRAINT `appointments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_ibfk_2` FOREIGN KEY (`hairdresser_id`) REFERENCES `hairdressers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_ibfk_3` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_ibfk_4` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `appointments_with_promotions`
--
ALTER TABLE `appointments_with_promotions`
  ADD CONSTRAINT `appointments_with_promotions_ibfk_1` FOREIGN KEY (`appointment_id`) REFERENCES `appointments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_with_promotions_ibfk_2` FOREIGN KEY (`promotion_id`) REFERENCES `promotions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_with_promotions_ibfk_3` FOREIGN KEY (`hairdresser_id`) REFERENCES `hairdressers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_with_promotions_ibfk_4` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_with_promotions_ibfk_5` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `appointments_with_promotions_ibfk_6` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `bank_holidays_per_store`
--
ALTER TABLE `bank_holidays_per_store`
  ADD CONSTRAINT `bank_holidays_per_store_ibfk_1` FOREIGN KEY (`bank_holidays_country`) REFERENCES `bank_holidays` (`country`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `customers`
--
ALTER TABLE `customers`
  ADD CONSTRAINT `customers_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `expensestrans`
--
ALTER TABLE `expensestrans`
  ADD CONSTRAINT `expensestrans_ibfk_1` FOREIGN KEY (`expenses_list_id`) REFERENCES `expenses_list` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `expensestrans_ibfk_2` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `expenses_list`
--
ALTER TABLE `expenses_list`
  ADD CONSTRAINT `expenses_list_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `hairdressers`
--
ALTER TABLE `hairdressers`
  ADD CONSTRAINT `hairdressers_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `hairdresser_stores`
--
ALTER TABLE `hairdresser_stores`
  ADD CONSTRAINT `hairdresser_stores_ibfk_1` FOREIGN KEY (`hairdresser_id`) REFERENCES `hairdressers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `hairdresser_stores_ibfk_2` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `products_ibfk_3` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `product_stores`
--
ALTER TABLE `product_stores`
  ADD CONSTRAINT `product_stores_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `product_stores_ibfk_2` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `promotions`
--
ALTER TABLE `promotions`
  ADD CONSTRAINT `promotions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `promotions_ibfk_2` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `promotions_commons`
--
ALTER TABLE `promotions_commons`
  ADD CONSTRAINT `promotions_commons_ibfk_1` FOREIGN KEY (`promotion_id`) REFERENCES `promotions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `sales_products_with_promotions`
--
ALTER TABLE `sales_products_with_promotions`
  ADD CONSTRAINT `sales_products_with_promotions_ibfk_1` FOREIGN KEY (`salestrans_id`) REFERENCES `salestrans` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `sales_products_with_promotions_ibfk_2` FOREIGN KEY (`promotion_id`) REFERENCES `promotions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `sales_products_with_promotions_ibfk_3` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `sales_products_with_promotions_ibfk_4` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `sales_products_with_promotions_ibfk_5` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `services`
--
ALTER TABLE `services`
  ADD CONSTRAINT `services_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `services_stores`
--
ALTER TABLE `services_stores`
  ADD CONSTRAINT `services_stores_ibfk_1` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `services_stores_ibfk_2` FOREIGN KEY (`store_id`) REFERENCES `shops` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `service_categories`
--
ALTER TABLE `service_categories`
  ADD CONSTRAINT `service_categories_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `service_sub_categories`
--
ALTER TABLE `service_sub_categories`
  ADD CONSTRAINT `service_sub_categories_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `service_categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `shops`
--
ALTER TABLE `shops`
  ADD CONSTRAINT `shops_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `suppliers`
--
ALTER TABLE `suppliers`
  ADD CONSTRAINT `suppliers_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;



-----------------VIEWS----------------------

------------------------------------------------------
CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `root`@`localhost` 
    SQL SECURITY DEFINER
VIEW `hairdresser_view` AS
    SELECT 
        `hairdressers`.`user_id` AS `user_id`,
        `hairdressers`.`id` AS `hairdresser_id`,
        `hairdressers`.`hairdresser_name` AS `hairdresser_name`,
        `hairdressers`.`hairdresser_mobile_phone` AS `hairdresser_mobile_phone`,
        `hairdressers`.`hairdresser_phone` AS `hairdresser_phone`,
        `hairdressers`.`display_order` AS `display_order`,
        `hairdresser_stores`.`is_active_hairdresser` AS `is_active_hairdresser`,
        `shops`.`company_name` AS `company_name`,
        `shops`.`id` AS `company_id`,
        `weekly_hours`.`color` AS `color`,
        `hairdressers`.`created_at` AS `created_at`,
        `hairdressers`.`updated_at` AS `updated_at`
    FROM
        (((`hairdressers`
        JOIN `hairdresser_stores`)
        JOIN `shops`)
        JOIN `weekly_hours`)
    WHERE
        ((`hairdressers`.`id` = `hairdresser_stores`.`hairdresser_id`)
            AND (`weekly_hours`.`user_id` = `hairdressers`.`user_id`)
            AND (`shops`.`id` = `hairdresser_stores`.`store_id`))

---------------------------------------------------------------------------------

-- --------------------------------------------------------

CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `root`@`localhost` 
    SQL SECURITY DEFINER
VIEW `products_view` AS
    SELECT 
        `products`.`user_id` AS `user_id`,
        `products`.`id` AS `product_id`,
        `products`.`product_code` AS `product_code`,
        `products`.`product_name` AS `product_name`,
        `products`.`product_ml` AS `product_ml`,
        `products`.`product_price` AS `product_price`,
        `products`.`product_ml_per_price` AS `product_ml_per_price`,
        `products`.`qty` AS `qty`,
        `products`.`supplier_id` AS `supplier_id`,
        `products`.`supplier_name` AS `supplier_name`,
        `products`.`is_active_product` AS `is_active_product`,
        `products`.`can_split_product` AS `can_split_product`,
        `shops`.`id` AS `company_id`,
        `shops`.`company_name` AS `company_name`,
        `products`.`created_at` AS `created_at`,
        `products`.`updated_at` AS `updated_at`
        
    FROM
        (`products`
        JOIN `shops`)
        
    WHERE
        (`products`.`store_id` = `shops`.`id`)
         
------------------------------------------------------
CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `root`@`localhost` 
    SQL SECURITY DEFINER
VIEW `services_view` AS
    SELECT 
        `services`.`id` AS `service_id`,
        `services`.`user_id` AS `user_id`,
        `services`.`service_name` AS `service_name`,
        `services`.`service_duration` AS `service_duration`,
        `services`.`service_price` AS `service_price`,
        `services`.`service_discount` AS `service_discount`,
        `services_stores`.`store_id` AS `store_id`,
        `shops`.`company_name` AS `company_name`,
        `service_categories`.`category_name` AS `category_name`,
        `services`.`sub_category` AS `sub_category_name`,
        `services`.`is_active` AS `is_active`,
        `services`.`switch_formula` AS `switch_formula`,
        `services`.`created_at` AS `created_at`,
        `services`.`updated_at` AS `updated_at`
        
    FROM
        (((`services`
          JOIN `services_stores`)
          JOIN `shops`)
          JOIN `service_categories`)
            
    WHERE
        ((`services_stores`.`service_id` = `services`.`id`)
          AND (`services_stores`.`store_id` = `shops`.`id`)
          AND (`services`.`category_id` = `service_categories`.`id`))

-------------------------------------

CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `root`@`localhost` 
    SQL SECURITY DEFINER
VIEW `promotion_view` AS
    SELECT 
        `promotions`.`id` AS `promotions_id`,
        `promotions`.`promotion_title` AS `promotion_title`,
         GROUP_CONCAT(`services`.`service_name`) AS `service_name`,
         GROUP_CONCAT(`services`.`id`) AS `s_id`,
        `promotions`.`days_duration` AS `days_duration`,
        `promotions`.`promotion_sale` AS `promotion_sale`,
        `promotions`.`store_id` AS `store_id`,
        `promotions`.`user_id` AS `user_id`,
        `shops`.`company_name` AS `company_name`,
        `promotions`.`promotion_description` AS `promotion_description`,
        `promotions`.`is_service` AS `is_service`,
        `promotions`.`created_at` AS `created_at`,
        `promotions`.`updated_at` AS `updated_at`
        
    FROM
        ((`promotions`
          JOIN `services`)
          JOIN `shops`)
            
    WHERE
        ((`promotions`.`user_id` = `services`.`user_id`)
          AND (`promotions`.`store_id` = `shops`.`id`))

