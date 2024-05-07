package db

import (
	"context"
	"github.com/cna-mhmdi/Tarkhineh-back/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createSomeFood(t *testing.T) Food {
	arg := CreateFoodParams{
		Name:        util.RandomString(6),
		Description: "this is decs for;" + util.RandomString(3),
		Price:       4,
		Rate:        3,
		Discount:    0,
		FoodTag:     "persian_food",
	}

	food, err := testQueries.CreateFood(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, food)

	return food
}

func TestCreateFood(t *testing.T) {
	createSomeFood(t)
}

func TestGetFood(t *testing.T) {
	food1 := createSomeFood(t)

	food2, err := testQueries.GetFood(context.Background(), food1.Name)

	require.NoError(t, err)
	require.NotEmpty(t, food2)

	require.Equal(t, food1.Name, food2.Name)
	require.Equal(t, food1.Description, food2.Description)
	require.Equal(t, food1.Price, food2.Price)
	require.Equal(t, food1.Rate, food2.Rate)
	require.Equal(t, food1.Discount, food2.Discount)
	require.Equal(t, food1.FoodTag, food2.FoodTag)
}

func TestListFoods(t *testing.T) {
	for i := 0; i < 10; i++ {
		createSomeFood(t)
	}

	arg := ListFoodsParams{
		Limit:  5,
		Offset: 5,
	}

	foods, err := testQueries.ListFoods(context.Background(), arg)
	require.NoError(t, err)

	for _, food := range foods {
		require.NotEmpty(t, food)
	}
}

func TestUpdateFood(t *testing.T) {
	arg := UpdateFoodParams{
		ID:          1,
		Name:        "1",
		Description: "2",
		Price:       3,
		Rate:        4,
		Discount:    5,
		FoodTag:     "6",
	}

	foods, err := testQueries.UpdateFood(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, foods)
}
