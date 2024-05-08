-- name: CreateFavorite :one
INSERT INTO favorites (
    username,
    food_id
) VALUES (
    $1, $2
)RETURNING *;

-- name: GetFavorites :many
SELECT * FROM favorites
WHERE id = $1;

-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE id = $1;
