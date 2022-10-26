package parser

import (
	"bytes"
	"showcase/models"
	"testing"

	"github.com/stretchr/testify/require"
)

const categoryAPIResponse = `
{
  "items": [
    {
      "id": 1,
      "name": "Товар 1",
      "price": 0.7,
	  "trademark_id": 1,	
	  "trademark": {
		  "id": 1,
          "name": "ТМ 1"
	  }	
    },
    {
      "id": 2,
      "name": "Товар 2",
      "price": 10,
	  "trademark_id": null,
	  "trademark": null
    }
  ]
}
`

func TestParseResponse(t *testing.T) {
	buf := bytes.NewBuffer([]byte(categoryAPIResponse))
	products, trademarks, err := decodeCategoryResponse(buf)

	t.Parallel()
	require.NoError(t, err)
	require.Len(t, products, 2)
	require.Equal(t, models.Product{ID: 1, Name: "Товар 1", Price: 0.7, TrademarkID: &[]int{1}[0]}, products[0])
	require.Equal(t, models.Product{ID: 2, Name: "Товар 2", Price: 10, TrademarkID: nil}, products[1])
	require.Len(t, trademarks, 1)
	require.Equal(t, models.Trademark{ID: 1, Name: "ТМ 1"}, trademarks[1])
}
