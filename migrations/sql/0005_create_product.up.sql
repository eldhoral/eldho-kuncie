INSERT INTO kuncie.product (id, sku, name, price, quantity) VALUES(1, '120P90', 'Google Home', 49.99, 100);
INSERT INTO kuncie.product (id, sku, name, price, quantity) VALUES(2, '43N23P', 'Macbook Pro', 5399.99, 100);
INSERT INTO kuncie.product (id, sku, name, price, quantity) VALUES(3, 'A304SD', 'Alexa Speaker', 109.50, 100);
INSERT INTO kuncie.product (id, sku, name, price, quantity) VALUES(4, '234234', 'Raspberry Pi B', 30.00, 100);

INSERT INTO kuncie.criterias_rules (id, criteria, reward) VALUES(1, 1, 4);
INSERT INTO kuncie.criterias_rules (id, criteria, reward) VALUES(2, 3, 2);
INSERT INTO kuncie.criterias_rules (id, criteria, reward) VALUES(3, 3, 10);

INSERT INTO kuncie.discount_rules (id, product_id, rules) VALUES(1, 2, 1);
INSERT INTO kuncie.discount_rules (id, product_id, rules) VALUES(2, 1, 2);
INSERT INTO kuncie.discount_rules (id, product_id, rules) VALUES(3, 3, 3);


