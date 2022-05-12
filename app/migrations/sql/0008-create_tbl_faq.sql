-- +migrate Up
CREATE TABLE IF NOT EXISTS `tbl_faq` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- +migrate Down
DROP TABLE tbl_faq;