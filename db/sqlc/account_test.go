package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{Owner: "eni", Balance: 150, Currency: "USD"}

	ctx := context.Background()
	account, err := testQueries.CreateAccount(ctx, arg)

	assert.Nil(t, err)
	assert.NotEmpty(t, account)

	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.Equal(t, arg.Currency, account.Currency)

	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

}

func TestGetAccount(t *testing.T) {

	account, err := testQueries.GetAccount(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, int32(1), account.ID)
}

func TestUpdateAccount(t *testing.T) {
	arg := UpdateAccountParams{ID: int32(1), Balance: 200}

	ctx := context.Background()
	account, err := testQueries.UpdateAccount(ctx, arg)

	assert.Nil(t, err)
	assert.NotEmpty(t, account)

	assert.Equal(t, arg.Balance, account.Balance)

	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

}

func TestDeleteAccount(t *testing.T) {

	ctx := context.Background()
	err := testQueries.DeleteAccount(ctx, int32(1))

	assert.Nil(t, err)

}

func TestListAccount(t *testing.T) {

	ctx := context.Background()
	arg := ListAccountParams{Limit: int32(1), Offset: int32(1)}
	accounts, err := testQueries.ListAccount(ctx, arg)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(accounts))

	for _, v := range accounts {
		assert.NotEmpty(t, v.Owner)
		assert.NotEmpty(t, v.Balance)
		assert.NotEmpty(t, v.Currency)

	}
}
