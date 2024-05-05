// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    password_hash
) VALUES (
    $1, $2
)RETURNING username, password_hash, created_at
`

type CreateUserParams struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.PasswordHash)
	var i User
	err := row.Scan(&i.Username, &i.PasswordHash, &i.CreatedAt)
	return i, err
}
