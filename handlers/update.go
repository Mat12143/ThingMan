package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Mat12143/thingMan/db"
	"github.com/labstack/echo/v4"
)

func UpdateItem(c echo.Context) error {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id provided")
	}

	var selectedItem db.Item
	var updatedItem db.Item

	dbItem := db.Istance.Where("id = ?", id).First(&selectedItem)

	if dbItem.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id provided")
	}

	data, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON provided")
	}

	err = json.Unmarshal(data, &updatedItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON provided")
	}

	if updatedItem.Quantity <= 0 {
		db.Istance.Delete(&db.Item{}, updatedItem.Id)
		return nil
	}

	if updatedItem.Location != selectedItem.Location {
		updatedItem.ModifiedAt = time.Now().Unix()
	}

	updatedItem.Id = selectedItem.Id
	updatedItem.CreatedAt = selectedItem.CreatedAt

	db.Istance.Save(&updatedItem)

	return c.JSON(http.StatusOK, updatedItem)
}
