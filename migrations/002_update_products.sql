-- Write your migrate up statements here

ALTER TABLE products SET UNLOGGED;

---- create above / drop below ----

ALTER TABLE products SET LOGGED;
