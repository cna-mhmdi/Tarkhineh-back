package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createUserFavorite(t *testing.T) Favorite {
	user := createUserAccount(t)
	food := createSomeFood(t)

	arg := CreateFavoriteParams{
		Username: user.Username,
		FoodID:   food.ID,
	}

	favorite, err := testQueries.CreateFavorite(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, favorite.Username)
	require.NotEmpty(t, favorite.FoodID)
	require.NotEmpty(t, favorite.AddedAt)

	return favorite
}

func TestCreateFavorite(t *testing.T) {
	createUserFavorite(t)
}

func TestGetUserFavorite(t *testing.T) {
	favorite := createUserFavorite(t)

	favorites, err := testQueries.GetFavorites(context.Background(), favorite.ID)
	require.NoError(t, err)
	require.NotEmpty(t, favorites)
}

func TestDeleteFavorite(t *testing.T) {
	favorite := createUserFavorite(t)

	err := testQueries.DeleteFavorite(context.Background(), favorite.ID)
	require.NoError(t, err)

	favorite1, err := testQueries.GetFavorites(context.Background(), favorite.ID)
	require.NoError(t, err)
	require.Empty(t, favorite1)
}
