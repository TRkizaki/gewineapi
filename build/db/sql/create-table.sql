DROP TABLE IF EXISTS winelist;
CREATE TABLE winelist (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  brand    VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO winelist
  (title, brand, price)
VALUES
  ('Teleda Orgo', 'Saperavi', 73.00),
  ('Kapistoni Shavkapito', 'kartli', 81.00),
  ('Gotsa Tsitska', 'Tsitska', 70.00),
  ('Teliani Valley', 'Rkatsiteli', 19.00);