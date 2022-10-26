package repository

import (
	"context"
	"showcase/models"

	"github.com/pkg/errors"
)

const upsertProductQuery = `
	INSERT INTO product (id, name, price, trademark_id) 
    VALUES ($1, $2, $3, $4) ON CONFLICT (id) 
    DO UPDATE SET name=$2, price=$3, trademark_id=$4
`

const upsertCategoryQuery = `
	INSERT INTO category (id, name) 
    VALUES ($1, $2) ON CONFLICT (id) 
    DO UPDATE SET name=$2
`

func UpsertProduct(db PgxDB, product models.Product) error {
	if _, err := db.Exec(context.Background(), upsertProductQuery, product.ID, product.Name, product.Price,
		product.TrademarkID); err != nil {
		return errors.Wrapf(err, "can not upsert product %+v", product)
	}

	return nil
}

func ProductList(db PgxDB, query string) (models.ProductList, error) {
	var params []interface{}

	sql := "SELECT id, name, price, trademark_id FROM product"
	if query != "" {
		sql += " WHERE name ILIKE $1"

		params = append(params, "%"+query+"%")
	}

	rows, err := db.Query(context.Background(), sql, params...)
	if err != nil {
		return nil, errors.Wrap(err, "can not select from db")
	}

	res := make([]models.Product, 0)

	for rows.Next() {
		p := models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.TrademarkID); err != nil {
			return nil, errors.Wrap(err, "can not scan product")
		}

		res = append(res, p)
	}

	return res, nil
}

func ProductMap(db PgxDB, id int) (models.ProductList, error) {
	rows := db.QueryRow(context.Background(), "SELECT id, name, price, trademark_id FROM product where id = $1", id)
	res := make([]models.Product, 0)
	p := models.Product{}
	if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.TrademarkID); err != nil {
		return nil, errors.Wrap(err, "can not scan product")
	}

	res = append(res, p)

	return res, nil
}

func UpsertProducts(db PgxDB, products []models.Product) error {
	for _, product := range products {
		if err := UpsertProduct(db, product); err != nil {
			return errors.Wrapf(err, "can not upsert products %+v", product)
		}
	}

	return nil
}

func UpsertCategory(db PgxDB, category models.Category) error {
	if _, err := db.Exec(context.Background(), upsertCategoryQuery, category.ID, category.Name); err != nil {
		return errors.Wrapf(err, "can not upsert category %+v", category)
	}

	return nil
}
