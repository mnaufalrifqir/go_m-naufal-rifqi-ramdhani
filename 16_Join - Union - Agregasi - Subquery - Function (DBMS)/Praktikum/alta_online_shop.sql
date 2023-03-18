SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

CREATE DATABASE IF NOT EXISTS `alta_online_shop` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `alta_online_shop`;

CREATE TABLE IF NOT EXISTS `operators` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `payment_methods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_type_id` int(11) NOT NULL,
  `operator_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `product_type_id` (`product_type_id`),
  KEY `operator_id` (`operator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `product_descriptions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `product_types` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `total_qty` int(11) NOT NULL,
  `total_price` decimal(25,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `payment_method_id` (`payment_method_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `transaction_details` (
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` decimal(25,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  KEY `transaction_id` (`transaction_id`),
  KEY `product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `dob` date NOT NULL,
  `status_user` smallint(6) NOT NULL,
  `gender` char(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`);

ALTER TABLE `product_descriptions`
  ADD CONSTRAINT `product_descriptions_ibfk_2` FOREIGN KEY (`id`) REFERENCES `products` (`id`);

ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`);

ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- INSERT
INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES 
  (NULL, 'Operator Satu', current_timestamp(), current_timestamp()), 
  (NULL, 'Operator Dua', current_timestamp(), current_timestamp()), 
  (NULL, 'Operator Tiga', current_timestamp(), current_timestamp()), 
  (NULL, 'Operator Empat', current_timestamp(), current_timestamp()), 
  (NULL, 'Operator Lima', current_timestamp(), current_timestamp());

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES 
  (NULL, 'Pakaian', current_timestamp(), current_timestamp()), 
  (NULL, 'Makanan', current_timestamp(), current_timestamp()), 
  (NULL, 'Minuman', current_timestamp(), current_timestamp());

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES 
  (NULL, '1', '3', 'p4k414n1', 'Baju Seragam', '1', current_timestamp(), current_timestamp()), 
  (NULL, '1', '3', 'p4k414n2', 'Baju Kaos', '1', current_timestamp(), current_timestamp());

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES 
  (NULL, '2', '1', 'm4k4n4n1', 'Nasi Goreng', '1', current_timestamp(), current_timestamp()), 
  (NULL, '2', '1', 'm4k4n4n2', 'Ayam Goreng', '1', current_timestamp(), current_timestamp()), 
  (NULL, '2', '1', 'm4k4n4n3', 'Ayam Bakar', '1', current_timestamp(), current_timestamp());

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES 
  (NULL, '3', '4', 'm1num4n1', 'Air Putih', '1', current_timestamp(), current_timestamp()), 
  (NULL, '3', '4', 'm1num4n2', 'Es Jeruk', '1', current_timestamp(), current_timestamp()), 
  (NULL, '3', '4', 'm1num4n3', 'Es Kelapa', '1', current_timestamp(), current_timestamp());

INSERT INTO `product_descriptions` (`id`, `description`, `created_at`, `updated_at`) VALUES 
  ('1', 'Baju seragam SMA', current_timestamp(), current_timestamp()), 
  ('2', 'Baju kaos bola emyu', current_timestamp(), current_timestamp()), 
  ('3', 'Nasi goreng spesial pedes dikit', current_timestamp(), current_timestamp()), 
  ('4', 'Ayam goreng Ibu Lemoe', current_timestamp(), current_timestamp()), 
  ('5', 'Ayam bakar kampung', current_timestamp(), current_timestamp()), 
  ('6', 'Air putih ada manis-manisnya', current_timestamp(), current_timestamp()), 
  ('7', 'Es jeruk purut', current_timestamp(), current_timestamp()), 
  ('8', 'Es kelapa muda', current_timestamp(), current_timestamp());

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES 
  (NULL, 'OPO', '1', current_timestamp(), current_timestamp()), 
  (NULL, 'GoPei', '1', current_timestamp(), current_timestamp()), 
  (NULL, 'Mobel Bengking', '1', current_timestamp(), current_timestamp());

INSERT INTO `users` (`id`, `name`, `address`, `dob`, `status_user`, `gender`, `created_at`, `updated_at`) VALUES 
  (NULL, 'Muhammad', 'Sukabirus', '2001-11-13', '1', 'M', current_timestamp(), current_timestamp()), 
  (NULL, 'Naufal', 'Sukabumi', '2001-11-14', '2', 'M', current_timestamp(), current_timestamp()), 
  (NULL, 'Rifqi', 'Bandung', '2001-11-15', '3', 'F', current_timestamp(), current_timestamp()), 
  (NULL, 'Ramdhani', 'Sukapura', '2001-11-16', '4', 'M', current_timestamp(), current_timestamp()), 
  (NULL, 'Awaliyah', 'Sukaraja', '2001-11-17', '5', 'F', current_timestamp(), current_timestamp());

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES 
  (NULL, '1', '1', '1', '3', '90000', current_timestamp(), current_timestamp()), 
  (NULL, '1', '2', '2', '3', '60000', current_timestamp(), current_timestamp()), 
  (NULL, '1', '3', '3', '3', '30000', current_timestamp(), current_timestamp()), 
  (NULL, '2', '1', '1', '3', '45000', current_timestamp(), current_timestamp()), 
  (NULL, '2', '2', '2', '3', '30000', current_timestamp(), current_timestamp()), 
  (NULL, '2', '3', '3', '3', '15000', current_timestamp(), current_timestamp()), 
  (NULL, '3', '1', '1', '3', '15000', current_timestamp(), current_timestamp()), 
  (NULL, '3', '2', '2', '3', '12000', current_timestamp(), current_timestamp()), 
  (NULL, '3', '3', '3', '3', '9000', current_timestamp(), current_timestamp()), 
  (NULL, '4', '1', '1', '3', '120000', current_timestamp(), current_timestamp()), 
  (NULL, '4', '2', '2', '3', '90000', current_timestamp(), current_timestamp()), 
  (NULL, '4', '3', '3', '3', '60000', current_timestamp(), current_timestamp()), 
  (NULL, '5', '1', '1', '3', '210000', current_timestamp(), current_timestamp()), 
  (NULL, '5', '2', '2', '3', '180000', current_timestamp(), current_timestamp()), 
  (NULL, '5', '3', '3', '3', '150000', current_timestamp(), current_timestamp());

INSERT INTO `transaction_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES 
  ('1', '1', '1', '1', '30000', current_timestamp(), current_timestamp()), 
  ('1', '4', '2', '2', '30000', current_timestamp(), current_timestamp()), 
  ('1', '7', '3', '3', '30000', current_timestamp(), current_timestamp()), 
  ('2', '7', '1', '4', '20000', current_timestamp(), current_timestamp()), 
  ('2', '2', '2', '1', '20000', current_timestamp(), current_timestamp()), 
  ('2', '4', '3', '2', '20000', current_timestamp(), current_timestamp()), 
  ('3', '5', '1', '5', '10000', current_timestamp(), current_timestamp()), 
  ('3', '8', '2', '1', '10000', current_timestamp(), current_timestamp()), 
  ('3', '7', '3', '2', '10000', current_timestamp(), current_timestamp()), 
  ('4', '6', '1', '3', '15000', current_timestamp(), current_timestamp()), 
  ('4', '6', '2', '3', '15000', current_timestamp(), current_timestamp()), 
  ('4', '4', '3', '5', '15000', current_timestamp(), current_timestamp()), 
  ('5', '6', '1', '1', '10000', current_timestamp(), current_timestamp()), 
  ('5', '8', '2', '5', '10000', current_timestamp(), current_timestamp()), 
  ('5', '7', '3', '2', '10000', current_timestamp(), current_timestamp()), 
  ('6', '7', '1', '1', '5000', current_timestamp(), current_timestamp()), 
  ('6', '1', '2', '1', '5000', current_timestamp(), current_timestamp()), 
  ('6', '4', '3', '1', '5000', current_timestamp(), current_timestamp()), 
  ('7', '1', '1', '1', '5000', current_timestamp(), current_timestamp()), 
  ('7', '7', '2', '1', '5000', current_timestamp(), current_timestamp()), 
  ('7', '4', '3', '1', '5000', current_timestamp(), current_timestamp()), 
  ('8', '5', '1', '1', '4000', current_timestamp(), current_timestamp()), 
  ('8', '7', '2', '1', '4000', current_timestamp(), current_timestamp()), 
  ('8', '4', '3', '1', '4000', current_timestamp(), current_timestamp()), 
  ('9', '4', '1', '1', '3000', current_timestamp(), current_timestamp()), 
  ('9', '7', '2', '1', '3000', current_timestamp(), current_timestamp()), 
  ('9', '5', '3', '1', '3000', current_timestamp(), current_timestamp()), 
  ('10', '7', '1', '8', '40000', current_timestamp(), current_timestamp()), 
  ('10', '6', '2', '2', '40000', current_timestamp(), current_timestamp()), 
  ('10', '5', '3', '4', '40000', current_timestamp(), current_timestamp()), 
  ('11', '7', '1', '6', '30000', current_timestamp(), current_timestamp()), 
  ('11', '4', '2', '3', '30000', current_timestamp(), current_timestamp()), 
  ('11', '3', '3', '3', '30000', current_timestamp(), current_timestamp()), 
  ('12', '7', '1', '4', '20000', current_timestamp(), current_timestamp()), 
  ('12', '4', '2', '2', '20000', current_timestamp(), current_timestamp()), 
  ('12', '5', '3', '4', '20000', current_timestamp(), current_timestamp()), 
  ('13', '7', '1', '14', '70000', current_timestamp(), current_timestamp()), 
  ('13', '3', '2', '10', '70000', current_timestamp(), current_timestamp()), 
  ('13', '5', '3', '7', '70000', current_timestamp(), current_timestamp()), 
  ('14', '1', '1', '10', '60000', current_timestamp(), current_timestamp()), 
  ('14', '8', '2', '6', '60000', current_timestamp(), current_timestamp()), 
  ('14', '7', '3', '12', '60000', current_timestamp(), current_timestamp()), 
  ('15', '6', '1', '25', '50000', current_timestamp(), current_timestamp()), 
  ('15', '8', '2', '5', '50000', current_timestamp(), current_timestamp()), 
  ('15', '7', '3', '10', '50000', current_timestamp(), current_timestamp());
/*
-- SELECT
SELECT name FROM users WHERE users.gender = 'M';

SELECT * FROM products WHERE id = '3';

SELECT * FROM users 
WHERE created_at 
BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() 
AND name LIKE '%a%';

SELECT COUNT(*) AS users_female FROM users WHERE gender = 'F';

SELECT * FROM users ORDER BY name ASC;

SELECT * FROM products LIMIT 5;

-- Update
UPDATE products SET name = 'product dummy' WHERE products.id = 1;

UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

-- DELETE
DELETE FROM products WHERE id = 1;

DELETE FROM products WHERE product_type_id = 1;

-- Join, Union, Sub query, Function
SELECT * FROM transactions WHERE user_id IN (1, 2);

SELECT SUM(total_price) AS total_harga FROM transactions WHERE user_id = 1;

SELECT COUNT(td.transaction_id) AS total_transaksi
FROM transaction_details td
INNER JOIN products p ON p.id = td.product_id
WHERE p.product_type_id = 2;

SELECT p.*, pt.name AS product_types_name FROM product_types pt INNER JOIN products p ON p.product_type_id = pt.id;

SELECT t.*, p.name AS product_name, u.name AS user_name 
FROM transactions t 
INNER JOIN transaction_details td ON t.id = td.transaction_id 
INNER JOIN products p ON td.product_id = p.id 
INNER JOIN users u ON t.user_id = u.id;

DELIMITER $$
CREATE TRIGGER delete_transaction_details
AFTER DELETE ON transactions FOR EACH ROW
BEGIN
DECLARE t_id INT;
SET t_id = OLD.id;
DELETE FROM transaction_details WHERE transaction_id = t_id;
END$$

DELIMITER $$
CREATE TRIGGER min_qty
AFTER DELETE ON transaction_details FOR EACH ROW
BEGIN
UPDATE transactions SET total_qty = total_qty - OLD.qty
WHERE id = OLD.transaction_id;
END$$

SELECT * FROM products WHERE id NOT IN ( SELECT product_id FROM transaction_details );
*/ 
