// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error)
	CreateFavorite(ctx context.Context, arg CreateFavoriteParams) (Favorite, error)
	CreateFood(ctx context.Context, arg CreateFoodParams) (Food, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAddress(ctx context.Context, id int64) error
	DeleteFavorite(ctx context.Context, arg DeleteFavoriteParams) error
	DeleteUser(ctx context.Context, username string) error
	GetAddresses(ctx context.Context, id int64) ([]Address, error)
	GetFavorites(ctx context.Context, username string) ([]Favorite, error)
	GetFood(ctx context.Context, name string) (Food, error)
	GetFoodById(ctx context.Context, id int64) (Food, error)
	GetProfile(ctx context.Context, username string) (Profile, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListFoods(ctx context.Context, arg ListFoodsParams) ([]Food, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error)
	UpdateFood(ctx context.Context, arg UpdateFoodParams) (Food, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
}

var _ Querier = (*Queries)(nil)
