package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransferTX(t *testing.T) {
	store := NewStore(testDB)
	account1, erraC := testQueries.GetAccount(context.Background(), 6)
	assert.Nil(t, erraC)
	account2, erraC := testQueries.GetAccount(context.Background(), 7)
	assert.Nil(t, erraC)

	fmt.Println("before >> : ", account1.Balance, account2.Balance)
	//  run n concurrent transfer transaction
	n := 2
	amount := int64(50)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {

		go func() {

			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: int64(account1.ID),
				ToAccountID:   int64(account2.ID),
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	//	check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		assert.Nil(t, err)

		rs := <-results
		assert.NotEmpty(t, rs)

		//	check transfer

		transfer := rs.Transfer
		assert.NotEmpty(t, transfer)
		assert.Equal(t, int64(3), transfer.FromAccountID)
		assert.Equal(t, int64(2), transfer.ToAccountID)
		assert.Equal(t, amount, transfer.Amount)
		assert.NotZero(t, transfer.ID)
		assert.NotZero(t, transfer.CreatedAt)

		//	check accounts
		fromAccount := rs.FromAccount
		assert.NotEmpty(t, fromAccount)
		assert.Equal(t, int32(3), fromAccount.ID)

		toAccount := rs.ToAccount
		assert.NotEmpty(t, toAccount)
		assert.Equal(t, int32(2), toAccount.ID)

		//	check account balance

		fmt.Println(">> tx : ", fromAccount.Balance, toAccount.Balance)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		assert.Equal(t, diff1, diff2)
		assert.True(t, diff1 > 0)
		assert.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		assert.True(t, k >= 1 && k <= n)
		assert.NotContains(t, existed, k)
		existed[k] = true
	}

	// check di final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	assert.NoError(t, err)

	fmt.Println("after >> : ", updatedAccount1.Balance, updatedAccount2.Balance)
	assert.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	assert.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}
