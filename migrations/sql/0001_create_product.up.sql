-- kuncie.product definition

CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sku` varchar(10) NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` decimal(13,2) NOT NULL,
  `quantity` int NOT NULL,
  PRIMARY KEY (`id`)
)