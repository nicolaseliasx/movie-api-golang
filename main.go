package main

import (
	"github.com/nicolaseliasx/movie-apirest-golang/database"
	"github.com/nicolaseliasx/movie-apirest-golang/routes"
)

func main() {
	database.Connect()

	router := routes.SetupRouter()

	router.Run(":8080")
}
