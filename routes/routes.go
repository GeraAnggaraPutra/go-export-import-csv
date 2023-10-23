package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"exportimportcsv/controller"
	"exportimportcsv/db"
	"exportimportcsv/repository"
	"exportimportcsv/service"
)

func Init() error {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	// Routes
	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "Application is Running",
		})
	})

	e.Static("/file/csv", "public/csv")

	e.POST("import", controller.ImportCSV)
	e.GET("export", controller.ExportCSV)

	return e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
