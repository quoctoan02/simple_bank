package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"simple_bank/util"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", util.RandomString(6))
}

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

// func TestUpdateUser(t *testing.T) {
// 	user1 := createRandomUser(t)

// 	arg := AddUserBalanceParams{
// 		ID:     user1.ID,
// 		Amount: util.RandomMoney(),
// 	}
// 	//cannot initialize 2 variables with 1 valuescompilerWrongAssignCount

// 	user2, err := testQueries.AddUserBalance(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user1.Owner, user2.Owner)
// 	require.Equal(t, user1.Balance+arg.Amount, user2.Balance)
// 	require.Equal(t, user1.Currency, user2.Currency)
// 	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }

// func TestDeleteUser(t *testing.T) {
// 	user1 := createRandomUser(t)

// 	err := testQueries.DeleteUser(context.Background(), user1.ID)

// 	require.NoError(t, err)

// 	user2, err := testQueries.GetUser(context.Background(), user1.ID)

// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, user2)
// }

// func TestListUser(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomUser(t)
// 	}

// 	arg := ListUsersParams{
// 		Limit:  5,
// 		Offset: 0,
// 	}

// 	users, err := testQueries.ListUsers(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.Len(t, users, 5)

// 	for _, user := range users {
// 		require.NotEmpty(t, user)
// 	}
// }
