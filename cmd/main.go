package main

import (
	"diplom/route-service/internal/app"
	"diplom/route-service/pkg/config"
	"log"
)

func main() {
	conf, err := config.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Start(conf)
	if err != nil {
		log.Fatal(err)
	}
}
