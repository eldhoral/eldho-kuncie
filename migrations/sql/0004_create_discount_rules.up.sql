-- kuncie.discount_rules definition

CREATE TABLE `discount_rules` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `rules` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `discount_rules_FK` (`product_id`),
  KEY `discount_rules_FK_1` (`rules`),
  CONSTRAINT `discount_rules_FK` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
  CONSTRAINT `discount_rules_FK_1` FOREIGN KEY (`rules`) REFERENCES `criterias_rules` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;