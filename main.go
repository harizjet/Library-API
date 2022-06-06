package main

import (
	"library/Basic-Golang-Api/routes"
)

func main() {
	router := routes.InitRoutes()
	router.Run("localhost:8080")
}
