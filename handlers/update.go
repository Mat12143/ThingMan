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

	var item db.Item
	dbItem := db.Istance.Where("id = ?", id).First(&item)

	if dbItem.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id provided")
	}

	data, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON provided")
	}

	err = json.Unmarshal(data, &item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON provided")
	}

	if item.Quantity <= 0 {
		db.Istance.Delete(&db.Item{}, item.Id)
		return nil
	}

	item.ModifiedAt = time.Now().Unix()
	db.Istance.Save(&item)

	return c.JSON(http.StatusOK, item)
}
