CREATE TABLE accounts(
  "id" BIGSERIAL PRIMARY KEY,
  "user_name" VARCHAR(10) UNIQUE NOT NULL,
  "first_name" VARCHAR(15) NOT NULL,
  "last_name" VARCHAR(15) NOT NULL,
  "password" VARCHAR(50) NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "created_on" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

-- CREATE TABLE users.accounts (id INT PRIMARY KEY, balance DECIMAL);
-- INSERT INTO users.accounts VALUES (1, "abebe", "abebe", "beso", "1234", "abebe@gmail.com", "March 05, 2018");


