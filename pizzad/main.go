package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	util "pizza-delivery/apputil"
	"pizza-delivery/bootstrapper"
	"pizza-delivery/router"
	"syscall"
)

func main() {
	bootstrapper.StartUp()
	router := router.InitRoutes()
	server := &http.Server{
		Addr:     bootstrapper.AppConfig.Server,
		Handler:  router,
		ErrorLog: log.New(util.LogWriter, "", 0),
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Running the HTTP server
	go func() {
		server.ListenAndServe()
	}()

	interruptSignal := <-interrupt
	switch interruptSignal {
	case os.Kill:
		util.Error("Got SIGKILL...")
	case os.Interrupt:
		util.Error("Got SIGINT...")
	case syscall.SIGTERM:
		util.Error("Got SIGTERM...")
	}

	util.Info("The service is shutting down...")
	server.Shutdown(context.Background())
	util.Info("Shut down is done")
}
