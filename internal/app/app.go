package app

import (
	"go-sys/internal/app/printer"
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

	// json(echo) setup

	// create pid file

	// batch init

}
