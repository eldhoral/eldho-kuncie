-- +migrate Up
CREATE TABLE `tbl_tnc_explain` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `id_tnc` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- +migrate Down
DROP TABLE tbl_tnc_explain;