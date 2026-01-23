-- name: RegistrationUser :one
INSERT INTO users
    (user_name, password_hash, telegram_id)
VALUES  
    ($1, $2, $3)
RETURNING *;

-- name: LoginUser :one 
SELECT * FROM users WHERE user_name = $1;

-- name: GetUserIDByUserName :one
SELECT id FROM users WHERE user_name = $1;