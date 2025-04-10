package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Mat12143/thingMan/db"
	"github.com/labstack/echo/v4"
)

func AddItem(c echo.Context) error {

	var newItem db.Item

	data, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No body found")
	}

	err = json.Unmarshal(data, &newItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON body")
	}

	if len(newItem.Name) == 0 || newItem.Quantity == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON body")
	}

	newItem.CreatedAt = time.Now().Unix()
	newItem.ModifiedAt = newItem.CreatedAt

	db.Istance.Create(&newItem)

	return c.JSON(http.StatusOK, newItem)
}
