-- +migrate Up
CREATE TABLE IF NOT EXISTS `tbl_loan_explain` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `interest` int NOT NULL DEFAULT '0',
  `admin_fee` decimal(10,0) NOT NULL DEFAULT '0',
  `fine` float NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- +migrate Down
DROP TABLE tbl_loan_explain;