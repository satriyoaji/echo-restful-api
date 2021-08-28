package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/echo-restful-api/models"
	"net/http"
)

func FetchAllEmployees(e echo.Context) error {
	result, err := models.FetchAllEmployees()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, result)
}
