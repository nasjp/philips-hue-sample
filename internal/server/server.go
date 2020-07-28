package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nasjp/philips-hue-sample/internal/http"
)

const port = 8080

func Run() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run(ctx context.Context) error {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	s := http.NewServer(port)
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-termCh:
		if err := s.Stop(ctx); err != nil {
			return err
		}

		return nil
	case <-errCh:
		return nil
	}
}
