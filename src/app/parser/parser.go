package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"showcase/models"

	"github.com/pkg/errors"
)

const itemsURL = "https://www.sima-land.ru/api/v3/item/?category_id=%v"

func decodeCategoryResponse(body io.Reader) (models.ProductList, models.TrademarkMap, error) {
	var items struct {
		Items []struct {
			models.Product
			Trademark *models.Trademark
		}
	}

	err := json.NewDecoder(body).Decode(&items)

	p := models.ProductList{}
	t := models.TrademarkMap{}

	for _, i := range items.Items {
		if i.Trademark != nil {
			t[i.Trademark.ID] = *i.Trademark
			i.Product.TrademarkID = &i.Trademark.ID
		}

		p = append(p, i.Product)
	}

	return p, t, errors.Wrap(err, "can not decode response")
}

func ParseCategory(categoryID int) (models.ProductList, models.TrademarkMap, error) {
	url := fmt.Sprintf(itemsURL, categoryID)

	resp, err := http.Get(url) //nolint: gosec
	if err != nil {
		return nil, nil, errors.Wrapf(err, "read url %v", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, errors.Wrapf(err, "invalid status %v of reading url %v", resp.StatusCode, url)
	}

	return decodeCategoryResponse(resp.Body)
}
