package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/echo-restful-api/controllers"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world!")
	})

	e.GET("/employees", controllers.FetchAllEmployees)

	return e
}
