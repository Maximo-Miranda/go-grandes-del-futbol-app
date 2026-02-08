package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/goravel/framework/facades"

	"grandesdelfutbol/bootstrap"
)

func main() {
	bootstrap.Boot()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route Run error: %v", err)
			fmt.Fprintf(os.Stderr, "Error: no se pudo iniciar el servidor: %v\n", err)
			os.Exit(1)
		}
	}()

	<-quit
	if err := facades.Route().Shutdown(); err != nil {
		facades.Log().Errorf("Route Shutdown error: %v", err)
	}
}
