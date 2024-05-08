package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createUserAddresses(t *testing.T) Address {
	user := createUserAccount(t)
	arg := CreateAddressParams{
		Username:    user.Username,
		AddressLine: "this is an address for user 1",
		AddressTag:  "wokplace",
		PhoneNumber: "09100460435",
	}

	address, err := testQueries.CreateAddress(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, address.Username)
	require.NotEmpty(t, address.AddressLine)
	require.NotEmpty(t, address.AddressTag)
	require.NotEmpty(t, address.PhoneNumber)

	return address
}

func TestCreateAddress(t *testing.T) {
	createUserAddresses(t)
}

func TestGetAddresses(t *testing.T) {
	Address1 := createUserAddresses(t)
	Address2, err := testQueries.GetAddresses(context.Background(), Address1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Address2)
}

func TestDeleteAddress(t *testing.T) {
	address1 := createUserAddresses(t)
	err := testQueries.DeleteAddress(context.Background(), address1.ID)
	require.NoError(t, err)

	address2, err := testQueries.GetAddresses(context.Background(), address1.ID)
	require.NoError(t, err)
	require.Empty(t, address2)
}

func TestUpdateAddress(t *testing.T) {
	Address1 := createUserAddresses(t)

	arg := UpdateAddressParams{
		ID:          Address1.ID,
		AddressLine: "2",
		AddressTag:  "3",
		PhoneNumber: "4",
	}

	address2, err := testQueries.UpdateAddress(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, address2)
	require.NotEqual(t, Address1.AddressLine, address2.AddressLine)
	require.NotEqual(t, Address1.AddressTag, address2.AddressTag)
	require.NotEqual(t, Address1.PhoneNumber, address2.PhoneNumber)

}
