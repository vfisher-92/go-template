package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go-template/di"
	"go-template/internal/handlers"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	dc di.DependencyContainer
)

func init() {
	di.InitDC()
	dc = di.GetDependencyContainer()
}

func main() {

	r := mux.NewRouter()

	handlers.NewAppHandler(r)
	http.Handle("/", r)

	httpServer := fasthttp.Server{
		Handler:      fasthttpadaptor.NewFastHTTPHandler(http.DefaultServeMux),
		IdleTimeout:  time.Duration(dc.Config.HttpIdleTimeout) * time.Second,
		WriteTimeout: time.Duration(dc.Config.WorkersTimeout) * time.Second,
	}

	ctx, cancel := context.WithCancel(context.Background())
	runGroup, gCtx := errgroup.WithContext(ctx)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	runGroup.Go(func() error {
		return httpServer.ListenAndServe(dc.Config.Address)
	})
	runGroup.Go(func() error {
		<-gCtx.Done()
		fmt.Println("Shutdown HTTP")
		return httpServer.Shutdown()
	})
	if err := runGroup.Wait(); err != nil {
		fmt.Printf("Exit reason: %s \n", err)
	}

	defer dc.Destroy()
}
