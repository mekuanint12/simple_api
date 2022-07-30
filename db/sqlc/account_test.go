package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingAccount(t *testing.T) {
	arg := CreateAccountParams{
		UserName:  "testuser",
		FirstName: "meku",
		LastName:  "last",
		Password:  "1234",
		Email:     "test@test.com",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserName, account.UserName)
	require.Equal(t, arg.FirstName, account.FirstName)
	require.Equal(t, arg.LastName, account.LastName)
	require.Equal(t, arg.Email, account.Email)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedOn)
}
