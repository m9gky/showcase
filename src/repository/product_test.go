package repository

import (
	"context"
	"showcase/models"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestProductList(t *testing.T) {
	mock, err := pgxmock.NewConn()

	t.Parallel()
	require.NoError(t, err)

	defer mock.Close(context.Background())

	prds := []models.Product{
		{ID: 1, Name: "post 1", Price: 5.0, TrademarkID: &[]int{1}[0]},
		{ID: 2, Name: "post 2", Price: 36.6, TrademarkID: nil},
	}

	rows := mock.NewRows([]string{"id", "name", "price", "trademark_id"}).
		AddRow(prds[0].ID, prds[0].Name, prds[0].Price, prds[0].TrademarkID).
		AddRow(prds[1].ID, prds[1].Name, prds[1].Price, prds[1].TrademarkID)

	mock.ExpectQuery("SELECT id, name, price, trademark_id FROM product").WillReturnRows(rows)
	res, err := ProductList(mock)

	require.NoError(t, err)
	require.Len(t, res, 2)
	require.Equal(t, prds[0], res[0])
	require.Equal(t, prds[1], res[1])

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
func TestUpsertProduct(t *testing.T) {
	mock, err := pgxmock.NewConn()

	t.Parallel()
	require.NoError(t, err)

	defer mock.Close(context.Background())

	p := models.Product{
		ID:    1,
		Name:  "test",
		Price: 0.4,
	}

	mock.ExpectExec("INSERT INTO product").
		WithArgs(p.ID, p.Name, p.Price, p.TrademarkID).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = UpsertProduct(mock, p)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
