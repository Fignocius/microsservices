package update

import (
	"os"

	"github.com/fignocius/microsservices/update/handler"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type UpdateApp struct {
	handler.TrackingHandler
}

func Init() *UpdateApp {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return &UpdateApp{
		handler.NewTrackingHandler(db),
	}
}

func (u UpdateApp) Run(e echo.Echo) {

	e.POST("/update", u.Update)
}
