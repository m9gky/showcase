package repository

import (
	"context"
	"showcase/models"

	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestUpsertTrademark(t *testing.T) {
	mock, err := pgxmock.NewConn()

	t.Parallel()
	require.NoError(t, err)

	defer mock.Close(context.Background())

	tm := models.Trademark{
		ID:   1,
		Name: "Test",
	}

	mock.ExpectExec("INSERT INTO trademark").
		WithArgs(tm.ID, tm.Name).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = UpsertTrademark(mock, tm)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestTrademarks(t *testing.T) {
	mock, err := pgxmock.NewConn()

	t.Parallel()
	require.NoError(t, err)

	tm := models.Trademark{
		ID:   1,
		Name: "test",
	}

	mock.ExpectQuery("SELECT id, name FROM trademark").
		WillReturnRows(pgxmock.NewRows([]string{"id", "name"}).AddRow(tm.ID, tm.Name))

	res, err := Trademarks(mock)
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, tm, res[0])
}

func TestTrademarkMap(t *testing.T) {
	mock, err := pgxmock.NewConn()

	t.Parallel()
	require.NoError(t, err)

	tm := models.Trademark{
		ID:   1,
		Name: "test",
	}

	mock.ExpectQuery("SELECT id, name FROM trademark").
		WillReturnRows(pgxmock.NewRows([]string{"id", "name"}).AddRow(tm.ID, tm.Name))

	res, err := TrademarkMap(mock)
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, tm, res[tm.ID])
}
