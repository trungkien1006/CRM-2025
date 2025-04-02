package helpers

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

var Redis *redis.Client

var Ctx context.Context