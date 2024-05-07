// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateFood(ctx context.Context, arg CreateFoodParams) (Food, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, username string) error
	GetFood(ctx context.Context, name string) (Food, error)
	GetProfile(ctx context.Context, id int64) (Profile, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListFoods(ctx context.Context, arg ListFoodsParams) ([]Food, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateFood(ctx context.Context, arg UpdateFoodParams) (Food, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
}

var _ Querier = (*Queries)(nil)
