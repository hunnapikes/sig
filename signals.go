package sig

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var defaultSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
}

type Log func(s os.Signal)

func New(log Log, sigs ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	Notify(cancel, log, sigs...)
	return ctx, cancel
}

func Notify(cancelFunc context.CancelFunc, log Log, sigs ...os.Signal) {
	go func() {
		if len(sigs) == 0 {
			sigs = defaultSignals
		}

		signals := make(chan os.Signal)
		signal.Notify(signals, sigs...)

		sig := <-signals
		log(sig)
		cancelFunc()
	}()
}
