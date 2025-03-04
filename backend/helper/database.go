package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/ariefsn/upwork/env"
	"github.com/ariefsn/upwork/logger"
	"github.com/ariefsn/upwork/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient(address string) (client *mongo.Client, cancel context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		logger.Fatal(err, models.M{
			"func": "helper.MongoClient",
		})
	}

	return
}

func RedisClient(env env.EnvDb) *redis.Client {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://:%s@%s:%s/%d", env.Password, env.Host, env.Port, env.DbIndex))
	if err != nil {
		logger.Fatal(err, models.M{
			"func": "helper.RedisClient.opt",
		})
	}

	rdb := redis.NewClient(opt)

	ping := rdb.Ping(context.Background())

	if ping.Err() != nil {
		logger.Fatal(ping.Err(), models.M{
			"func": "helper.RedisClient",
		})
	}

	return rdb
}
