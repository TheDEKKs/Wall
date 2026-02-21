-- name: RegistrationUser :one
INSERT INTO users
    (user_name, password_hash, telegram_record_id)
VALUES  
    ($1, $2, $3)
RETURNING *;

-- name: LoginUser :one 
SELECT * FROM users WHERE user_name = $1;

-- name: GetUserByUserID :one 
SELECT * FROM users WHERE id = $1;

-- name: GetUserIDByUserName :one
SELECT id FROM users WHERE user_name = $1;

-- name: SetNewPasswordByUserID :exec 
UPDATE users SET password_hash = $2 WHERE id = $1;

-- name: SetNewUserNameByUserID :exec
UPDATE users SET user_name = $2 WHERE id = $1;


-- name: SetNewTelegramByUserID :exec
UPDATE users SET telegram_record_id = $2 WHERE id = $1;