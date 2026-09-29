package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"renet/DB/repositories"
	"renet/config/db"
	"renet/config/env"
	"renet/controllers"
	sessionredis "renet/redis"
	"renet/router"
	"renet/services"
	"strconv"
	"strings"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Application struct {
	PORT string
}

func NewApplication() *Application { // point towards same memory address
	return &Application{
		PORT: env.GetString("PORT"),
	}
}

func (app *Application) Run() error {
	addr := app.PORT

	//add : if not exsists

	if !strings.HasPrefix(addr, ":") {
		addr = ":" + addr
	}
	database, err := db.InitDB()
	if err != nil {
		return err
	}

	repository := repositories.NewUserRespository(database)
	service := services.NewUserService(repository)

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisDB := 0
	if value := os.Getenv("REDIS_DB"); value != "" {
		redisDB, err = strconv.Atoi(value)
		if err != nil || redisDB < 0 {
			return fmt.Errorf("invalid REDIS_DB value %q", value)
		}
	}

	redisClient := goredis.NewClient(&goredis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		_ = redisClient.Close()
		return fmt.Errorf("failed to connect to Redis at %s: %w", redisAddr, err)
	}
	defer redisClient.Close()

	sessionStore := sessionredis.NewSessionStore(redisClient)
	sessionManager := sessionredis.NewSessionManager(sessionStore)
	controller := controllers.NewAuthController(service, sessionManager)
	appRouter := router.Router(controller)

	server := http.Server{
		Addr:         addr,
		Handler:      appRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running on PORT", addr)

	return server.ListenAndServe()
}
