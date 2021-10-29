package app

import (
	"go-sys/internal/app/printer"
	"go-sys/internal/pkg/database"
	"log"
)

var Log *log.Logger

func Start() {
	log.Println("System Collector Start")
	defer log.Println("System Collector Stopped")

	// log setup

	// args setup
	printer.Setup()

	// database setup
	database.Setup()

	// json(echo) setup

	// create pid file

	// batch init

}
