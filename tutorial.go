package main

import (
	"example/Basic-Golang-Api/routes"
)

func main() {
	router := routes.InitRoutes()
	router.Run("localhost:8080")
}
