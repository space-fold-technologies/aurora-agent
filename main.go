package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/space-fold-technologies/aurora-agent/app"
)

func main() {
	mode := flag.String("mode", "run", " for the app run mode")
	flag.Parse()
	if *mode == "run" || *mode == "RUN" {
		application := app.Application{}
		application.Start()
		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		application.Stop()
	}
}
