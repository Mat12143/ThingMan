package main

import (
	"log"

	"github.com/Mat12143/thingMan/db"
	"github.com/Mat12143/thingMan/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Static("/static", "./static")

	e.GET("/", handlers.Index)

	e.POST("/add", handlers.AddItem)

	e.POST("/update/:id", handlers.UpdateItem)

	e.Logger.Fatal(e.Start(":8080"))
}
