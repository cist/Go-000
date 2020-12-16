package main

import (
	v1 "Week04/api/user/v1"
	"context"
	"fmt"
	xerrors "github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init service api
	usrsrv := InitUserService()
	// register grpc service
	s := grpc.NewServer()
	v1.RegisterUserServer(s, usrsrv)

	// context

	g, ctx := errgroup.WithContext(context.Background())

	// start grpc server
	g.Go(func() error {
		listen, err := net.Listen("tcp", ":5001")
		if err != nil {
			return err
		}
		go func() {
			<-ctx.Done()
			s.GracefulStop()
			log.Printf("grpc server stop")
		}()
		return s.Serve(listen)
	})
	// catch signals
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			return xerrors.New(fmt.Sprintf("signal caught: %s, ready to quit...", sig.String()))
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("error: %+v", err)
	}
}
