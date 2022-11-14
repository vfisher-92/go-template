package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go-template/config"
	"go-template/internal/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	config.InitConfig()
	cfg = config.GetInstance()

	rand.Seed(time.Now().UnixNano())

	dbLogger := gormlogger.Default.LogMode(gormlogger.Error)
	if cfg.DevMode {
		dbLogger = gormlogger.Default.LogMode(gormlogger.Info)
	}
	dsn := "host=" + cfg.DB.Host + " user=" + cfg.DB.User + " dbname=" + cfg.DB.Dbname + " sslmode=disable password=" + cfg.DB.Password + " port=" + cfg.DB.Port
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                                   dbLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		// log
	}

	// Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	pingResult, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		// log error
	}
	if pingResult != "PONG" {
		// log
	}

	r := mux.NewRouter()
	//TODO init hadlers
	handlers.NewAppHandler(r)
	http.Handle("/", r)

	httpServer = fasthttp.Server{
		Handler:      fasthttpadaptor.NewFastHTTPHandler(http.DefaultServeMux),
		IdleTimeout:  time.Duration(cfg.HttpIdleTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WorkersTimeout) * time.Second,
	}

}
