package main

import (
	"time"

	"github.com/ntec-io/Nostradamus/internal/config"
	"github.com/ntec-io/Nostradamus/internal/logger"
	"github.com/ntec-io/Nostradamus/internal/redis"
	"github.com/ntec-io/Nostradamus/pkg/fifaindex"
)

var redisClient redis.Client

func init() {
	logger.Log().Info("Starting Nostradamus")

	// Config setup
	cfg, err := config.ReadConfig("config/config.yml")
	if err != nil {
		logger.Log().Panic(err)
	}

	// Redis setup
	redisClient, err = redis.NewClient(cfg.RedisPassword)
	if err != nil {
		logger.Log().Panic(err)
	}
}

func main() {
	// Updating Date IDs
	err := redisClient.SetDateIDs(fifaindex.GetAllDateIDs())
	if err != nil {
		logger.Log().Panic(err)
	}

	t, _ := time.Parse(redis.TimeLayout, "10-10-2020")
	id, _ := (redisClient.GetLastDateID(t))
	link, _ := fifaindex.GetPlayerLink("Cristiano Ronaldo")

	fifaindexPlayer, _ := fifaindex.GetPlayer(link, id)

	redisClient.SavePlayer(redis.Player{
		ID:            0,
		FifaindexLink: link,
		Name:          "CR7",
		Stats:         fifaindexPlayer,
		DateID:        id,
	})

}
