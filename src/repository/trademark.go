package repository

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"showcase/models"
)

const upsertTrademarkQuery = "INSERT INTO trademark (id, name) VALUES ($1, $2) ON CONFLICT DO NOTHING"

func UpsertTrademark(db PgxDB, trademark models.Trademark) error {
	if _, err := db.Exec(context.Background(), upsertTrademarkQuery, trademark.ID, trademark.Name); err != nil {
		return errors.Wrapf(err, "can not upsert trademark %+v", trademark)
	}

	return nil
}

func UpsertNewTrademarks(db PgxDB, trademarks models.TrademarkMap, current models.TrademarkMap) error {
	for _, tm := range trademarks {
		if _, ok := current[tm.ID]; !ok {
			if err := UpsertTrademark(db, tm); err != nil {
				return err
			}

			current[tm.ID] = tm
		}
	}

	return nil
}

func Trademarks(db PgxDB) ([]models.Trademark, error) {
	rows, err := db.Query(context.Background(), "SELECT id, name FROM trademark")
	if err != nil {
		return nil, errors.Wrapf(err, "can not select trademarks")
	}
	defer rows.Close()

	res := []models.Trademark{}

	for rows.Next() {
		var tm models.Trademark
		if err := rows.Scan(&tm.ID, &tm.Name); err != nil {
			return nil, errors.Wrapf(err, "can not scan trademark")
		}

		res = append(res, tm)
	}

	return res, nil
}

func TrademarkMap(db PgxDB) (map[int]models.Trademark, error) {
	tm, err := Trademarks(db)
	if err != nil {
		return nil, err
	}

	res := map[int]models.Trademark{}
	for _, t := range tm {
		res[t.ID] = t
	}

	return res, nil
}
