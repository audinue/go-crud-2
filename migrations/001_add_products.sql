-- Write your migrate up statements here

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255)
);

---- create above / drop below ----

DROP TABLE products;
