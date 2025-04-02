package app

import (
	"admin-v1/app/configs"
	"admin-v1/app/helpers"
	"admin-v1/app/initRedis"
	"admin-v1/app/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))

	helpers.GormDB = configs.GormConnection()

	r := routes.InitRoute()

	isExist, redisErr := helpers.Redis.Exists(helpers.Ctx, "role:1").Result()

	if redisErr != nil {
		fmt.Println("loi khi khoi tao quyen cho redis")
	}

	if isExist == 0 {
		initRedis.InitRolePermission()
	}
// initRedis.InitRolePermission()
	return r
}