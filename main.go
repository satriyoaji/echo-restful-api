package main

import (
	"github.com/satriyoaji/echo-restful-api/database"
	"github.com/satriyoaji/echo-restful-api/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
