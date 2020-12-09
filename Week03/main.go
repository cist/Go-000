package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server is a HTTP server.
type Server struct {
	srv *http.Server
}

// NewServer : to create a http_server
func NewServer() *Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(
		func(resp http.ResponseWriter, request *http.Request) {
			// mock
			fmt.Fprintln(resp, "Hello,world!")
			time.Sleep(100 * time.Second)
		},
	))
	srv := &http.Server{
		Addr:    "127.0.0.1:5000",
		Handler: mux,
	}
	return &Server{srv: srv}
}

// Start the http_server
func (s *Server) Start(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := s.Shutdown(stopCtx); err != nil {
			log.Printf("Server forced to shutdown: %v", err)
		}
		log.Printf("Shutdown App")
	}()
	log.Printf("[HTTP] Listening on: %s\n", s.srv.Addr)
	return s.srv.ListenAndServe()
}

// Shutdown to close the http_server
func (s *Server) Shutdown(ctx context.Context) error {
	log.Printf("start to shoutdown http_server")
	return s.srv.Shutdown(ctx)
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	srv := NewServer()
	// start http_server
	g.Go(func() error {
		return srv.Start(ctx)
	})

	// listen OS_signal
	g.Go(func() error {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			log.Printf("waiting for quit signal")
			select {
			case <-ctx.Done():
				log.Printf("signal ctx done")
				return ctx.Err()
			case <-signalChan:
				return errors.New("receive quit signal")
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("Server Error:%v\n", err)
	}
}
