-- +goose Up
CREATE TABLE users (
    Id_User SERIAL PRIMARY KEY,
    Name_User VARCHAR(16) UNIQUE NOT NULL,
    Id_Telegram int UNIQUE,
    Id_Wall INT UNIQUE,
    Password TEXT NOT NULL
);
-- +goose Down
DROP TABLE users;
