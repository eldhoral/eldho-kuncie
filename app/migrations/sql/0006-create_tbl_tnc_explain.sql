-- +migrate Up
CREATE TABLE `tbl_tnc_explain` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id_tnc` int NOT NULL,
  `created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `id_tnc_title` int NOT NULL,
  `id_tnc_subtitle` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tbl_tnc_explain_FK` (`id_tnc`),
  KEY `tbl_tnc_explain_FK_1` (`id_tnc_title`),
  KEY `tbl_tnc_explain_FK_2` (`id_tnc_subtitle`),
  CONSTRAINT `tbl_tnc_explain_FK` FOREIGN KEY (`id_tnc`) REFERENCES `tbl_tnc` (`id`),
  CONSTRAINT `tbl_tnc_explain_FK_1` FOREIGN KEY (`id_tnc_title`) REFERENCES `tbl_tnc_title` (`id`),
  CONSTRAINT `tbl_tnc_explain_FK_2` FOREIGN KEY (`id_tnc_subtitle`) REFERENCES `tbl_tnc_subtitle` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- +migrate Down
DROP TABLE tbl_tnc_explain;