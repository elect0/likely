package main

import (
	"log"
	"os"
	"time"

	"github.com/elect0/likely/internal/handler"
	"github.com/elect0/likely/internal/repository"
	"github.com/elect0/likely/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	dbCfg := repository.DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := repository.NewPostgresDB(dbCfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	defer db.Close()

	// redis
	// redisAddr := os.Getenv("REDIS_ADDR")
	// redisClient, err := repository.NewRedisClient(redisAddr)
	// if err != nil {
	// 	log.Fatalf("Could not connect to Redis: %v", err)
	// }

	// defer redisClient.Close()

	userRepo := repository.NewPostgresRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	tokenTTL := time.Hour * 24
	userService := service.NewUserService(userRepo, jwtSecret, tokenTTL)

	httpHandler := handler.NewHTTPHandler(userService)

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	httpHandler.RegisterRoutes(e)
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
