package db

import (
	"context"
	"database/sql"
	"github.com/cna-mhmdi/Tarkhineh-back/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createUserAccount(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomPassword())
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomUsername(),
		PasswordHash: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCrateUser(t *testing.T) {
	createUserAccount(t)
}

func TestGetUser(t *testing.T) {
	user1 := createUserAccount(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createUserAccount(t)
	err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createUserAccount(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}
