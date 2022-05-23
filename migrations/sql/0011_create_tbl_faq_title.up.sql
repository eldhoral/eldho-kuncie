CREATE TABLE IF NOT EXISTS `tbl_faq_title` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_faq` int NOT NULL,
  `title` varchar(100) NOT NULL,
  `created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `description` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id_order` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tbl_faq_title_FK` (`id_faq`),
  CONSTRAINT `tbl_faq_title_FK` FOREIGN KEY (`id_faq`) REFERENCES `tbl_faq` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

