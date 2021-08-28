package main

import (
	"github.com/satriyoaji/echo-restful-api/routes"
)

func main()  {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}