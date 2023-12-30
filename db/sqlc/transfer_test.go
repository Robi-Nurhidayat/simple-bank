package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransfer(t *testing.T) {
	arg := CreateTransferParams{FromAccountID: 2, ToAccountID: 4, Amount: 30}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	assert.Nil(t, err)

	assert.Equal(t, int64(30), transfer.Amount)
}

func TestGetTransfer(t *testing.T) {
	transfer, err := testQueries.GetTransfer(context.Background(), int32(1))

	assert.Nil(t, err)

	assert.Equal(t, int32(1), transfer.ID)
	assert.Equal(t, int64(30), transfer.Amount)
}

func TestUpdateTransfer(t *testing.T) {

	arg := UpdateTransferParams{
		ID:     int32(1),
		Amount: 500,
	}

	transfer, err := testQueries.UpdateTransfer(context.Background(), arg)

	assert.Nil(t, err)

	assert.Equal(t, int64(500), transfer.Amount)

}

func TestDeleteTransfer(t *testing.T) {
	err := testQueries.DeleteTransfer(context.Background(), int32(1))

	assert.Nil(t, err)
}

func TestListTransfer(t *testing.T) {
	arg := ListTransferParams{Limit: 10, Offset: 0}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)

	assert.Nil(t, err)

	assert.Equal(t, 2, len(transfers))
}
