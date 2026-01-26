-- name: NewTelegramRecord :one
INSERT INTO telegram
    (telegram_id, first_name, last_name, username) 
VALUES 
    ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByTelegramID :one 
SELECT * FROM telegram WHERE telegram_id = $1;