package main

import (
	"anoneko-adopt-api/src/config"
	"anoneko-adopt-api/src/router"
)

func main() {

	config.Init()
	router.Init()

}
