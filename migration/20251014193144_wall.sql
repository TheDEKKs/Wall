-- +goose Up
CREATE TABLE walls(
    Id_Wall SERIAL PRIMARY KEY, 
    Id_Creator INT UNIQUE NOT NULL,

    Mat BOOLEAN 
);

-- +goose Down

DROP TABLE walls;

