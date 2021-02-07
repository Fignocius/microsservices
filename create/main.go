package main

import (
	"os"

	"github.com/fignocius/microsservices/create/handler"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq" //Sqlx dependency
)

var db *sqlx.DB

func init() {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func main() {
	trackingHandler := handler.NewTrackingHandler(db)
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.POST("/create", trackingHandler.Create)
	e.Logger.Fatal(e.Start(":8080"))
}
