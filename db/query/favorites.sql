-- name: CreateFavorite :one
INSERT INTO favorites (
    username,
    food_id
) VALUES (
    $1, $2
)RETURNING *;

-- name: GetFavorites :many
SELECT * FROM favorites
WHERE username = $1;

-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE username = $1 AND food_id = $2;
