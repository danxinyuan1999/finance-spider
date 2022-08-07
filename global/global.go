package global

import (
	"github.com/go-redis/redis/v9"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	RBD     *redis.Client
	HttpReq *resty.Client
)
