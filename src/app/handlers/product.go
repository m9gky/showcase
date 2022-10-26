package handlers

import (
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"showcase/repository"
	"strconv"
)

func Product(db *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		getID := c.Param("id")
		id, err := strconv.Atoi(getID)
		if err != nil {
			return errors.Wrap(err, "can not get id")
		}
		product, err := repository.ProductMap(db, id)
		if err != nil {
			return errors.Wrap(c.Render(http.StatusBadRequest, "product.html", err), "can not render html")
		}
		data := map[string]interface{}{
			"Product": product,
		}
		return errors.Wrap(c.Render(http.StatusOK, "product.html", data), "can not render html")
	}
}
