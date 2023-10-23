package controller

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	"exportimportcsv/service"

)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) ImportCSV(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	data, err := c.service.ImportCSV(file)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully import data from CSV",
		"data":    data,
	})
}

func (c *Controller) ExportCSV(ctx echo.Context) error {
	data, err := c.service.ExportCSV()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	exportedFileName := "exported_data_user.csv"
	exportedFilePath := "./public/csv/" + exportedFileName

	file, err := os.Create(exportedFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"No", "Email", "Role", "Created Date"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, user := range data {
		row := []string{
			strconv.Itoa(user.NO),
			user.Email,
			user.RoleName,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	baseURL := os.Getenv("BASE_URL")
	fullURL := baseURL + "/file/csv/" + exportedFileName

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully export csv data",
		"data": fullURL,
	})
}
