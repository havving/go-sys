package app

import (
	"go-sys/internal/app/batch"
	"go-sys/internal/app/printer"
	"go-sys/internal/pkg/database"
	"log"
)

func Start() {
	log.Println("System Collector Start")
	defer log.Println("System Collector Stopped")

	// args setup
	printer.Setup()

	// database setup
	database.Setup()

	// json(echo) setup
	//e := echo.New()
	//e.Start(":8080")

	// batch init (system collector)
	batch.Start()

}
