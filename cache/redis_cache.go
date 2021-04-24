package cache

import (
	"encoding/json"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) EventCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (c *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.host,
		Password: "",
		DB:       c.db,
	})
}

func (c *redisCache) Set(key string, value *entity.Event) {
	client := c.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, json, c.expires*time.Second)
}

func (c *redisCache) Get(key string) *entity.Event {
	client := c.getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	post := entity.Event{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	return &post
}
