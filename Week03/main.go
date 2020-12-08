package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type UserHandler struct{}

func (UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	server := http.Server{
		Addr:    ":8080",
		Handler: UserHandler{},
	}

	// signal notify
	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

		select {
		case <-quit:
			log.Println("Shutting down server...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown:", err)
				return err
			}
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}

		return nil
	})

	// listen serve
	g.Go(func() error {
		errChan := make(chan error)
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Printf("Err: server.ListenAndServe(), %v", err)
				errChan <- err
			}
		}()

		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("Err: wait, %v", err)
	}

}
