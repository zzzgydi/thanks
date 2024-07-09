package common

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/zzzgydi/thanks/common/initializer"
)

var RDB *redis.Client

func initRedis() error {
	dsn := viper.GetString("REDIS_DSN")
	if dsn == "" {
		return fmt.Errorf("redis dsn error")
	}

	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return fmt.Errorf("redis parse url error: %s", err)
	}

	RDB = redis.NewClient(opt)

	_, err = RDB.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("redis connect error: %s", err)
	}

	return nil
}

func init() {
	initializer.Register("redis", initRedis)
}
