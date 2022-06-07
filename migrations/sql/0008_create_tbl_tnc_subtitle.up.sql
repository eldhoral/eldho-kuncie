CREATE TABLE IF NOT EXISTS `tbl_tnc_subtitle` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_tnc_title` int NOT NULL,
  `subtitle` varchar(100) NOT NULL,
  `created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `tbl_tnc_subtitle_FK` (`id_tnc_title`),
  CONSTRAINT `tbl_tnc_subtitle_FK` FOREIGN KEY (`id_tnc_title`) REFERENCES `tbl_tnc_title` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
