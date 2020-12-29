CREATE TABLE offers (
   offer_id int PRIMARY KEY,
   name varchar NOT NULL,
   price int NOT NULL,
   quantity int NOT NULL,
   available boolean NOT NULL,
   seller_id int NOT NULL
);