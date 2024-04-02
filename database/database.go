package database

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB
var Cache *redis.Client
var Ctx = context.Background()
