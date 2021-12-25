CREATE TABLE IF NOT EXISTS gorm_users(
    id SERIAL PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    password BYTEA UNIQUE NOT NULL,
    phonenumber BYTEA NOT NULL,
    color BYTEA NOT NULL
)