CREATE TABLE IF NOT EXISTS users (
   id serial PRIMARY KEY,
   firstName VARCHAR NOT NULL,
   lastName VARCHAR NOT NULL,
   password VARCHAR NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);