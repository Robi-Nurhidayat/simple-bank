package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEntry(t *testing.T) {

	arg := CreateEntryParams{AccountID: int64(2), Amount: 120}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	assert.Nil(t, err)
	assert.Equal(t, int64(120), entry.Amount)

}

func TestGetEntry(t *testing.T) {

	entry, err := testQueries.GetEntry(context.Background(), int32(1))

	assert.Nil(t, err)
	assert.Equal(t, int64(2), entry.AccountID)

}

func TestUpdateEntry(t *testing.T) {
	arg := UpdateEntryParams{ID: int32(1), Amount: int64(300)}

	entry, err := testQueries.UpdateEntry(context.Background(), arg)

	assert.Nil(t, err)
	assert.Equal(t, int64(300), entry.Amount)
	assert.Equal(t, int32(1), entry.ID)
	assert.NotZero(t, entry.AccountID)
}

func TestDeleteEntry(t *testing.T) {

	err := testQueries.DeleteEntry(context.Background(), int32(3))
	assert.Nil(t, err)

}

func TestListEntry(t *testing.T) {

	arg := ListEntryParams{Limit: 3, Offset: 0}
	entry, err := testQueries.ListEntry(context.Background(), arg)

	assert.Nil(t, err)

	assert.Equal(t, 3, len(entry))

	for _, v := range entry {
		assert.NotEmpty(t, v.ID)
		assert.NotEmpty(t, v.Amount)
		assert.NotEmpty(t, v.AccountID)

		fmt.Println(v.Amount)
	}
}
