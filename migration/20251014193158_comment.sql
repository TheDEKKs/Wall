-- +goose Up
CREATE TABLE comments (
    Id_Comment SERIAL PRIMARY KEY,
    Id_Wall INT NOT NULL,
    Id_Commentator INT NOT NULL,

    Text_Comment VARCHAR(128) NOT NULL
);


-- +goose Down
DROP TABLE comments;

