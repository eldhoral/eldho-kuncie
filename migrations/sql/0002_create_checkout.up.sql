-- kuncie.checkout definition

CREATE TABLE `checkout` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `purchase_id` varchar(50) NOT NULL,
  `quantity` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `checkout_FK` (`product_id`),
  CONSTRAINT `checkout_FK` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
)