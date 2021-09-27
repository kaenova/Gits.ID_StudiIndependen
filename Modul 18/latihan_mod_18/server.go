package main

import (
	"latihan_mod_18/db"
	"latihan_mod_18/routes"
)

func main() {
	db.CreateConnection()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
