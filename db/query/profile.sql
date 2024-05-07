-- name: CreateProfile :one
INSERT INTO profiles (
    username,
    first_name,
    last_name,
    email,
    phone_number,
    birthday,
    nickname
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)RETURNING *;


-- name: GetProfile :one
SELECT * FROM profiles
WHERE id = $1 LIMIT 1;


-- name: UpdateProfile :one
UPDATE profiles
set first_name = $2,
    last_name = $3,
    email = $4,
    phone_number = $5,
    birthday = $6,
    nickname = $7
WHERE id = $1
RETURNING *;
