package main

import (
	"mythical-mysfit-api/src/config"
	"mythical-mysfit-api/src/router"
)

func main() {

	config.Init()
	router.Init()

}
