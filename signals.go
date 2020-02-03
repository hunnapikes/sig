package sig

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var defaultSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
}

func New(sigs ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	Notify(cancel, sigs...)
	return ctx, cancel
}

func Notify(cancelFunc context.CancelFunc, sigs ...os.Signal) {
	go func() {
		var signals = make(chan os.Signal)
		if len(sigs) == 0 {
			sigs = defaultSignals
		}
		signal.Notify(signals, sigs...)
		sig := <-signals
		cancelFunc()
		log.Println(sig)
	}()
}
