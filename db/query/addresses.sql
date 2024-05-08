-- name: CreateAddress :one
INSERT INTO addresses(
    username,
    address_line,
    address_tag,
    phone_number
) VALUES (
    $1, $2, $3, $4
)RETURNING *;

-- name: GetAddresses :many
SELECT * FROM addresses
WHERE id = $1;

-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = $1;

-- name: UpdateAddress :one
UPDATE addresses
set address_line = $2,
    address_tag = $3,
    phone_number = $4
WHERE id = $1
RETURNING *;