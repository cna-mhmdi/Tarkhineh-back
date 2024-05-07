package db

import (
	"context"
	"fmt"
	"github.com/cna-mhmdi/Tarkhineh-back/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createUserProfile(t *testing.T) Profile {
	user := createUserAccount(t)
	arg := CreateProfileParams{
		Username:  user.Username,
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
		Email:     util.RandomEmail(),
		Birthday:  util.RandomBirthday(),
		Nickname:  util.RandomString(6),
	}

	profile, err := testQueries.CreateProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile)
	require.Equal(t, arg.Username, profile.Username)
	require.Equal(t, arg.FirstName, profile.FirstName)
	require.Equal(t, arg.LastName, profile.LastName)
	require.Equal(t, arg.Email, profile.Email)
	require.Equal(t, arg.Birthday, profile.Birthday)
	require.Equal(t, arg.Nickname, profile.Nickname)

	return profile
}

func TestCreateProfile(t *testing.T) {
	user1 := createUserProfile(t)
	require.NotEmpty(t, user1)

	arg := CreateProfileParams{
		Username: user1.Username,
	}

	user, err := testQueries.CreateProfile(context.Background(), arg)
	require.NoError(t, err)
	fmt.Println(user)
	require.NotEmpty(t, user.Username)
	require.Empty(t, user.FirstName)
	require.Empty(t, user.LastName)
	require.Empty(t, user.Email)
	require.Empty(t, user.Birthday)
	require.Empty(t, user.Nickname)
}

func TestGetProfile(t *testing.T) {
	profile1 := createUserProfile(t)
	profile2, err := testQueries.GetProfile(context.Background(), profile1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.Username, profile2.Username)
	require.Equal(t, profile1.FirstName, profile2.FirstName)
	require.Equal(t, profile1.LastName, profile2.LastName)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, profile1.Birthday, profile2.Birthday)
	require.Equal(t, profile1.Nickname, profile2.Nickname)
}

func TestUpdateProfile(t *testing.T) {
	user1 := createUserProfile(t)

	arg := UpdateProfileParams{
		ID:        user1.ID,
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
		Email:     util.RandomEmail(),
		Birthday:  util.RandomBirthday(),
		Nickname:  util.RandomString(6),
	}

	user2, err := testQueries.UpdateProfile(context.Background(), arg)
	require.NoError(t, err)
	fmt.Println(user1)
	fmt.Println(user2)
	require.Equal(t, user1.Username, user2.Username)
	require.NotEqual(t, user1.FirstName, user2.FirstName)
	require.NotEqual(t, user1.LastName, user2.LastName)
	require.NotEqual(t, user1.Email, user2.Email)
	require.Equal(t, user1.Birthday, user2.Birthday)
	require.NotEqual(t, user1.Nickname, user2.Nickname)
}
