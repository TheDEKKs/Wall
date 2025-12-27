-- name: NewComment :exec
INSERT INTO comments 
    (user_id, wall_id, text)
VALUES 
    ($1, $2, $3);