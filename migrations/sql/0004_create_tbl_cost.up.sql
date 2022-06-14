CREATE TABLE IF NOT EXISTS `tbl_cost` (
  `id` int NOT NULL AUTO_INCREMENT,
  `loan_option` varchar(25) NOT NULL,
  `interest` varchar(10) NOT NULL,
  `admin_fee` varchar(10) NOT NULL,
  `fine_per_day` varchar(10) NOT NULL,
  `description` varchar(100) NOT NULL,
  `id_loan_option` int NOT NULL,
  `is_visible` int NOT NULL,
  `created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

