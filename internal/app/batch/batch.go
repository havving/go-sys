package batch

import (
	"github.com/jasonlvhit/gocron"
	"go-sys/internal/app/printer"
)

func Start() {
	// test
	gocron.Every(5).Second().Do(printer.Printer())

	// start all the pending jobs
	<-gocron.Start()
}
