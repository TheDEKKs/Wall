-- +goose Up
-- +goose StatementBegin
CREATE TABLE wall(
    Id_Wall int PRIMARY KEY, 
    Id_Creator int UNIQUE NOT NULL,

    Mat BOOLEAN 
);

DROP TABLE wall;

