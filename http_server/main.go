package main

import (
	"admin-v1/app"
	"admin-v1/app/helpers"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redisContext = context.Background()

// @title CRM
// @version 1.0
// @host 192.168.0.120:8000
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	helpers.Ctx = redisContext

	if os.Getenv("DOCKER_ENV") != "true" {
		_ = godotenv.Load() // Chỉ tải .env nếu không chạy trong Docker
	}


	helpers.Redis = redis.NewClient(&redis.Options{
		// Addr:         "172.26.168.7:6379",
		Addr:         "redis:6379",
		PoolSize:     20,
		MinIdleConns: 5,  
	})

	pong, err := helpers.Redis.Ping(helpers.Ctx).Result()

	if err != nil {
		fmt.Println("Không thể kết nối Redis:", err)
		return
	}

	fmt.Println("Kết nối Redis thành công:", pong)
	
	if err != nil {
		panic(err)
	}

	route := app.Init()

	port := os.Getenv("PORT")

	fmt.Println("Port của bạn: ", port)

	ln, err := net.Listen("tcp", "0.0.0.0:" + port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, route)
}