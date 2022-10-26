package handlers

import (
	"net/http"
	"showcase/repository"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

func Main(db *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		query := c.QueryParams().Get("q")

		productList, err := repository.ProductList(db, query)
		if err != nil {
			return errors.Wrap(err, "can not get product list")
		}

		data := map[string]interface{}{
			"ProductList": productList,
			"Query":       query,
		}

		return errors.Wrap(c.Render(http.StatusOK, "main.html", data), "can not render html")
	}
}
