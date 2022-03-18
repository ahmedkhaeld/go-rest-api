package cache

import (
	"encoding/json"
	"github.com/ahmedkhaeld/rest-api/entity"
	"github.com/go-redis/redis/v7"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

// getClient creates a new redis client
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		// settings for the database
		Addr:     cache.host, // address host of database
		Password: "",
		DB:       cache.db, // database index
	})
}

func (cache *redisCache) Set(key string, post *entity.Post) {
	client := cache.getClient()

	// serialize Post object to JSON
	marshalledJSON, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	client.Set(key, marshalledJSON, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *entity.Post {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}

	return &post
}
