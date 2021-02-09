package main

import (
	"fmt"
	"os"

	"github.com/fignocius/microsservices/view/handler"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq" //Sqlx dependency
)

func main() {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	defer db.Close()

	trackingHandler := handler.NewTrackingHandler(db)

	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.GET("/view/:code", trackingHandler.View)
	e.Logger.Fatal(e.Start(":8080"))
}
