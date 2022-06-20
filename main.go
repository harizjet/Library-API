package main

import (
	"library/Basic-Golang-Api/config"
	"library/Basic-Golang-Api/consts"
	"library/Basic-Golang-Api/routes"
)

func main() {
	cfg := config.LoadConfig(consts.ENV_CONFIG_PATH)

	router := routes.InitRoutes(cfg)
	router.Run(cfg.ServerConfig.Host + ":" + cfg.ServerConfig.Port)
}
