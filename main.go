package main

import (
	"context"
	"errors"
	"github.com/g3co/twtest/pkg/chainviewer"
	"github.com/g3co/twtest/pkg/rest"
	"github.com/g3co/twtest/pkg/service"
	"github.com/g3co/twtest/pkg/storage/memory"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

const httpPort = ":8080"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	viewer := chainviewer.NewViewer()
	storage := memory.NewStorage()

	svc := service.NewService(storage, viewer)

	eg, errCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		log.Println("Starting service...")
		return svc.Run(errCtx)
	})

	restSrv := &http.Server{
		Addr:    httpPort,
		Handler: rest.NewHTTPROuter(svc),
	}

	eg.Go(func() error {
		log.Printf("Starting HTTP service on port %s", httpPort)
		if srvErr := restSrv.ListenAndServe(); srvErr != nil && !errors.Is(srvErr, http.ErrServerClosed) {
			return srvErr
		}
		return nil
	})

	eg.Go(func() error {
		<-errCtx.Done()
		log.Println("Halting HTTP service...")
		return restSrv.Shutdown(context.Background())
	})

	if err := eg.Wait(); errors.Is(err, context.Canceled) || err != nil {
		log.Printf("error occurred %s", err)
	}
}
