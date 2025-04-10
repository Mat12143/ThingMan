package handlers

import (
	"context"

	"github.com/Mat12143/thingMan/components"
	"github.com/Mat12143/thingMan/db"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var items []db.Item

	db.Istance.Find(&items)

	index := components.Index(items)
	index.Render(context.Background(), c.Response().Writer)

	return nil
}
