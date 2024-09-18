package app

import (
	"context"

	"auth-athlevo/api"
	"auth-athlevo/api/handlers"
	"auth-athlevo/config"
	kafka "auth-athlevo/internal/kafka/consumer"
	prd "auth-athlevo/internal/kafka/producer"
	"auth-athlevo/internal/storage/postgres"
	"auth-athlevo/service"

	l "github.com/azizbek-qodirov/logger"
	"github.com/go-redis/redis/v8"
)

func Run(cfg *config.Config) {
	logger, err := l.NewLogger(&l.LogFileConfigs{
		Directory: "internal/logs/info/",
		Filename:  "app.log",
		Stdout:    false,
		Include:   l.DateTime | l.Loglevel | l.ShortFileName,
	})
	if err != nil {
		panic(err)
	}

	logger.INFO.Println("Application started")
	// Postgres Connection
	db, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		logger.ERROR.Printf("can't connect to db: %v", err)
	}
	logger.INFO.Println("dd", db.Db)
	defer db.Db.Close()
	logger.INFO.Println("Connected to Postgres")

	// Redis Connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		logger.ERROR.Panicf("Failed to connect to Redis: %v", err)
	}
	logger.INFO.Println("Connected to Redis")

	authService := service.NewAuthService(db)
	userService := service.NewUserService(db)

	// Kafka
	brokers := []string{"kafka:9092"}
	cm := kafka.NewKafkaConsumerManager()
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		logger.ERROR.Println("Failed to create Kafka producer:", err)
		return
	}

	Reader(brokers, cm, authService, userService, logger)

	// HTTP Server
	h := handlers.NewHandler(authService, userService, rdb, &pr, logger)

	router := api.Engine(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.AUTH_PORT); err != nil {
		logger.ERROR.Panicf("can't start server: %v", err)
	}

	logger.INFO.Printf("REST server started on port %s", cfg.AUTH_PORT)
}
