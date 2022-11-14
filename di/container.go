package di

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go-template/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"math/rand"
	"time"
)

type DependencyContainer struct {
	Db          *gorm.DB
	RedisClient *redis.Client
	Config      config.Config
}

var instance DependencyContainer

func InitDC() {
	config.InitConfig()
	cfg := config.GetInstance()
	instance.Config = *cfg

	rand.Seed(time.Now().UnixNano())

	dbLogger := gormlogger.Default.LogMode(gormlogger.Error)
	if cfg.DevMode {
		dbLogger = gormlogger.Default.LogMode(gormlogger.Info)
	}
	dsn := "host=" + cfg.DB.Host + " user=" + cfg.DB.User + " dbname=" + cfg.DB.Dbname + " sslmode=disable password=" + cfg.DB.Password + " port=" + cfg.DB.Port
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                                   dbLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {

	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(cfg.DB.MaxOpenIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	instance.Db = db

	// Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	pingResult, err := redisClient.Ping(context.Background()).Result()
	if err != nil {

	}
	if pingResult != "PONG" {

	}
	instance.RedisClient = redisClient

	instance.Config = *cfg
}

func GetDependencyContainer() DependencyContainer {
	return instance
}

func (dc *DependencyContainer) Destroy() {
	dc.RedisClient.Close()
}
