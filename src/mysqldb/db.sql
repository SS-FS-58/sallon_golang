CREATE DATABASE salon_new;

USE salon_new;

CREATE TABLE users(
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  role VARCHAR(120),
  forename VARCHAR(120),
  surname VARCHAR(120),
  work_telephone VARCHAR(120) NULL,
  mobile_telephone VARCHAR(120) NULL,
  company_name VARCHAR(200) NOT NULL,
  username VARCHAR (50),
  email VARCHAR(50) NOT NULL,
  password VARCHAR (120),
  is_active BOOL,
  logo_image VARCHAR(60),
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  last_login DATETIME,
  UNIQUE KEY username (username)
) CHARSET = utf8;

CREATE TABLE shops (
  id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  vat_number varchar(30) NOT NULL,
  company_name varchar(200) NOT NULL,
  company_address varchar(300) NOT NULL,
  company_street_number int(20) NOT NULL,
  company_city varchar(100) NOT NULL,
  company_state varchar(200) NULL,
  company_zip_code varchar(80) NOT NULL,
  tax_office VARCHAR(150) NOT NULL,
  company_country varchar(100) NOT NULL,
  work_telephone varchar(120) NULL,
  mobile_telephone varchar(120) NULL,
  password VARCHAR(120),
  is_active bool,
  created_at DATETIME,
  updated_at DATETIME,
  UNIQUE KEY vat_number (vat_number),
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE services (
  id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  category_id INT NOT NULL,
  sub_category_id INT NOT NULL,
  service_name varchar(100) NOT NULL,
  service_duration int(200) NOT NULL,
  service_price DECIMAL(5, 2) NOT NULL,
  service_discount DECIMAL(5, 2) NULL,
  sub_category VARCHAR(100),
  is_active bool,
  switch_formula bool,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE suppliers (
  id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  vat_number varchar(30) NOT NULL,
  supplier_name varchar(200) NOT NULL,
  supplier_address varchar(300) NOT NULL,
  supplier_street_number int(20) NOT NULL,
  supplier_city varchar(100) NOT NULL,
  supplier_state varchar(200) NULL,
  supplier_zip_code varchar(80) NOT NULL,
  field_of_business VARCHAR(150) NOT NULL,
  supplier_country varchar(100) NOT NULL,
  work_telephone varchar(120) NULL,
  email varchar(120) NULL,
  website varchar(250) NULL,
  is_active bool,
  created_at DATETIME,
  updated_at DATETIME,
  UNIQUE KEY vat_number (vat_number),
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE products (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  product_code VARCHAR(150) NOT NULL,
  user_id INT NOT NULL,
  supplier_id INT NOT NULL,
  supplier_name VARCHAR(200) NOT NULL,
  product_name VARCHAR(200) NOT NULL,
  product_ml INT NOT NULL,
  product_price DECIMAL(5, 2) NOT NULL,
  product_ml_per_price DECIMAL(5.3) NOT NULL,
  store_id INT,
  is_active_product BOOL,
  can_split_product BOOL,
  qty float,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (supplier_id) REFERENCES suppliers(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE hairdressers (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  hairdresser_name VARCHAR(200) NOT NULL,
  hairdresser_mobile_phone VARCHAR(200) NOT NULL,
  hairdresser_phone VARCHAR(100) NOT NULL,
  display_order INT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE customers (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  customer_name VARCHAR(100) NOT NULL,
  customer_surname VARCHAR(100) NOT NULL,
  gender_of_customer VARCHAR(20) NOT NULL,
  date_of_birth DATE,
  customer_address VARCHAR(300) NOT NULL,
  customer_street_number INT NOT NULL,
  customer_city VARCHAR(100) NOT NULL,
  customer_state VARCHAR(100) NOT NULL,
  customer_zip_code VARCHAR(50) NOT NULL,
  customer_country VARCHAR(70) NOT NULL,
  home_phone_number VARCHAR(100) NULL,
  mobile_phone_number VARCHAR(100) NOT NULL,
  customer_email VARCHAR(70) NOT NULL,
  customer_points int NOT NULL,
  is_active BOOL,
  pelatis_lianikis BOOL,
  created_at DATETIME,
  updated_at DATETIME,
  UNIQUE KEY customer_email (customer_email),
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE service_categories (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  category_name VARCHAR(100) NOT NULL,
  is_active BOOL,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE service_sub_categories(
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  category_id INT,
  category_name VARCHAR(120),
  sub_category_name VARCHAR(120),
  is_active BOOL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (category_id) REFERENCES service_categories(id) ON UPDATE CASCADE ON DELETE CASCADE
) DEFAULT CHARSET = utf8;

CREATE TABLE hairdresser_stores (
  d INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  hairdresser_id INT,
  store_id INT,
  is_active_hairdresser bool,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (hairdresser_id) REFERENCES hairdressers(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE product_stores (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  product_id INT,
  store_id INT,
  is_active_product bool,
  qty float,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE services_stores (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  service_id INT,
  store_id INT,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (service_id) REFERENCES services(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE appointments (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  hairdresser_id INT NOT NULL,
  customer_id INT NOT NULL,
  store_id INT NOT NULL,
  service_id INT NOT NUll,
  start_time DATETIME,
  end_time DATETIME,
  service_status VARCHAR(45),
  comments text,
  is_all_day BOOL,
  created_at DATETIME,
  updated_at DATETIME,
  KEY k_service_id(service_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (hairdresser_id) REFERENCES hairdressers(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE required_products (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  product_id INT NOT NULL,
  appointment_service_id INT NOT NULL,
  product_used_ml INT,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (appointment_service_id) REFERENCES appointments(service_id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE salestrans (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  service_id INT NOT NULL,
  service_price DECIMAL(5, 2),
  service_qty INT,
  service_discount DECIMAL(5, 2),
  payment_type VARCHAR(50),
  service_line_total DECIMAL(5, 2),
  is_service BOOL,
  invoice_number INT(11) NULL,
  created_at DATETIME,
  updated_at DATETIME
) CHARSET = utf8;

CREATE TABLE salestrans_common (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  salestrans_id INT NOT NULL,
  hairdresser_id INT NOT NUll,
  store_id INT NOT NUll,
  customer_id INT NOT NULL,
  appointment_id INT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME
) CHARSET = utf8;

CREATE TABLE invoices (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  store_id INT NOT NUll,
  customer_id int,
  sub_total float,
  discount float,
  total_cost float,
  status_id int,
  created_at DATETIME,
  updated_at DATETIME
) CHARSET = utf8;

CREATE TABLE expenses_list (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  store_id INT NOT NUll,
  expenses_name VARCHAR(100) NOT NULL,
  is_active BOOL,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE expensestrans (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  store_id INT NOT NUll,
  expenses_list_id INT NOT NULL,
  expenses_price DECIMAL(5, 2),
  payment_type BOOL,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (expenses_list_id) REFERENCES expenses_list(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE promotions (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  store_id INT NOT NULL,
  promotion_title VARCHAR(150) NOT NULL,
  days_duration INT NOT NULL,
  promotion_sale DECIMAL(5, 2) NOT NULL,
  promotion_description TEXT NULL,
  is_service BOOL NOT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE promotions_commons (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  promotion_id INT NOT NULL,
  promotion_service_id INT NOT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (promotion_id) REFERENCES promotions(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE bank_holidays (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  summary varchar(150) DEFAULT NULL,
  date_start date DEFAULT NULL,
  date_end date DEFAULT NUll,
  country varchar(25) DEFAULT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  KEY k_country(country),
  UNIQUE KEY date_start (summary, date_start, date_end, country)
) CHARSET = utf8;

CREATE TABLE bank_holidays_per_store (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  store_id INT NOT NULL,
  bank_holidays_country varchar(150) DEFAULT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (bank_holidays_country) REFERENCES bank_holidays(country) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE weekly_hours (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  store_id INT NOT NULL,
  day_of_the_week INT NOT NULL,
  from_time VARCHAR(20) NOT NULL,
  to_time VARCHAR(20) NOT NULL,
  color VARCHAR(45) NOT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL
) CHARSET = utf8;

-- ############ 24
CREATE TABLE appointments_with_promotions (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  store_id INT NOT NULL,
  hairdresser_id INT NOT NULL,
  customer_id INT NOT NULL,
  promotion_id INT NOT NUll,
  appointment_id INT NOT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (promotion_id) REFERENCES promotions(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (hairdresser_id) REFERENCES hairdressers(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE sales_products_with_promotions (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  store_id INT NOT NULL,
  customer_id INT NOT NULL,
  promotion_id INT NOT NUll,
  salestrans_id INT NOT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL,
  FOREIGN KEY (salestrans_id) REFERENCES salestrans(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (promotion_id) REFERENCES promotions(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES shops(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
) CHARSET = utf8;

CREATE TABLE greek_name_days (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  celebreation_date DATETIME NULL,
  name_day varchar(150) DEFAULT NULL,
  created_at DATETIME NULL,
  updated_at DATETIME NULL
) CHARSET = utf8;

CREATE TABLE users_activation (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id int NOT NUll,
  username VARCHAR (50),
  email VARCHAR(50) NOT NULL,
  activation_code text,
  creation_date DATETIME NULL,
  activation_date DATETIME NULL
) CHARSET = utf8;

CREATE TABLE users_packages (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  month_periods int NOT NUll,
  number_of_stores Int NOT NULL,
  amount DECIMAL(5, 2) NOT NULL,
  credits int NOT NUll,
  created_at DATETIME NULL,
  updated_at DATETIME NULL
) CHARSET = utf8;

create table daily_values (wdate datetime, wval decimal);

-- ########### 30
-- DROP PROCEDURE IF EXISTS filldates;
-- DELIMITER | CREATE PROCEDURE filldates(dateStart DATE, dateEnd DATE) BEGIN WHILE dateStart <= dateEnd DO
-- INSERT INTO
--   daily_values (wdate, wval)
-- VALUES
--   (dateStart, 0);
-- SET
--   dateStart = date_add(dateStart, INTERVAL 1 DAY);
-- END WHILE;
-- END;
-- | DELIMITER;
-- CALL filldates('2018-01-01', '2025-12-31');