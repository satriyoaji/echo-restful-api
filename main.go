package echo_restful_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main()  {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}