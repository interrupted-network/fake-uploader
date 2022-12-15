package app

import (
	"os"
	"os/signal"
	"syscall"
)

func (a *app) Start() {
	a.coordinator.UseCase.Start()

	done := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		a.logger.DebugF("exiting...")
		close(done)
	}()
	<-done
}
