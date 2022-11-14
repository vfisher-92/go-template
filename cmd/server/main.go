package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/valyala/fasthttp"
	"go-template/config"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

var (
	cfg         *config.Config
	httpServer  fasthttp.Server
	redisClient *redis.Client
	db          *gorm.DB
	err         error
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	runGroup, gCtx := errgroup.WithContext(ctx)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	runGroup.Go(func() error {
		return httpServer.ListenAndServe(cfg.Address)
	})
	runGroup.Go(func() error {
		<-gCtx.Done()
		fmt.Println("Shutdown HTTP")
		return httpServer.Shutdown()
	})
	if err := runGroup.Wait(); err != nil {
		fmt.Printf("Exit reason: %s \n", err)
	}

	defer redisClient.Close()

}
