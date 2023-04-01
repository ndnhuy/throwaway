DROP TABLE IF EXISTS delivery_order;
CREATE TABLE delivery_order (
  id INT AUTO_INCREMENT NOT NULL,
  sale_order_id VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY (sale_order_id)
);

DROP TABLE IF EXISTS delivery_item;
CREATE TABLE delivery_item (
  id INT AUTO_INCREMENT NOT NULL,
  delivery_order_id INT,
  product_id VARCHAR(128) NOT NULL,
  quantity INT,
  PRIMARY KEY (`id`),
  UNIQUE KEY (delivery_order_id, product_id)
);

DROP TABLE IF EXISTS trip;
CREATE TABLE trip (
  id INT AUTO_INCREMENT NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS delivery_trip;
CREATE TABLE delivery_trip (
  delivery_order_id INT,
  trip_id INT,
  UNIQUE KEY (delivery_order_id, trip_id)
);