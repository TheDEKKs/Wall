-- name: GetWallByUserID :one 
SELECT * FROM walls WHERE user_id = $1;

-- name: SetWall :exec
INSERT INTO walls 
    (user_id)
VALUES 
    ($1);