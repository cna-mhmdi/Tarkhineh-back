-- name: CreateFood :one
INSERT INTO food (
    name,
    description,
    price,
    rate,
    discount,
    food_tag
) VALUES (
    $1, $2, $3, $4, $5, $6
)RETURNING *;

-- name: GetFood :one
SELECT * FROM food
WHERE name = $1 LIMIT 1;

-- name: GetFoodById :one
SELECT * FROM food
WHERE id = $1 LIMIT 1;

-- name: ListFoods :many
SELECT * FROM food
ORDER BY food_tag
LIMIT $1
OFFSET $2;

-- name: UpdateFood :one
UPDATE food
set name = $2,
    description = $3,
    price = $4,
    rate = $5,
    discount = $6,
    food_tag = $7
WHERE id = $1
RETURNING *;