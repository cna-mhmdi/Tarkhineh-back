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
WHERE username = $1 LIMIT 1;


-- name: UpdateProfile :one
UPDATE profiles
set first_name = $3,
    last_name = $4,
    email = $5,
    phone_number = $6,
    birthday = $7,
    nickname = $8
WHERE id = $1 AND username = $2
RETURNING *;
