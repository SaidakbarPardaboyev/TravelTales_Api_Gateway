package main

import (
	"travel/api"
	"travel/config"
)

func main() {

	r := api.NewRouter()
	r.Run(":" + config.Load().API_GATEWAY_PORT)

}
