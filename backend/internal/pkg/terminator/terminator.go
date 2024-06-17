package terminator

import (
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
)

type Func func() error

type Terminator struct {
	clients []Func
	ch      chan os.Signal
}

func NewTerminator(clients []Func) Terminator {
	return Terminator{clients: clients, ch: make(chan os.Signal, 1)}
}

func (t Terminator) Listen() {
	go func() {
		signal.Notify(t.ch, os.Interrupt, syscall.SIGTERM)
		<-t.ch

		s := spinner.New(spinner.CharSets[2], 100*time.Millisecond)
		s.Suffix = " shutting down clients"
		s.Start()

		wg := sync.WaitGroup{}

		for i := range t.clients {
			wg.Add(1)
			go func(f func() error) {
				defer wg.Done()
				if err := f(); err != nil {
					slog.Error("error while terminating client", err)
				}
			}(t.clients[i])
		}
		wg.Wait()

		s.Stop()
		os.Exit(0)
	}()
}
