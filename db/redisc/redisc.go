package redisc

import(
	"github.com/go-redis/redis"
	"log"
)

// Redis holds a reference to the active client.
type Redis struct{
	*redis.Client
}
// InitRedisNoPw initialises the Redis connection when there is no password given.
func InitRedisNoPw(redisHost string) Redis  {
	client := redis.NewClient(&redis.Options{
		Addr: redisHost,
		DB: 0,
	})

	log.Printf("Redis: Connected to %s", client.ClientGetName())

	return Redis{
		Client: client,
	}
}

// InitRedis initialises the Redis connection.
func InitRedis(redisHost string, password string) Redis {
	client := redis.NewClient(&redis.Options{
		Addr: redisHost,
		Password: password,
		DB: 0,
	})

	log.Printf("Redis: Connected to %s", client.ClientGetName())

	return Redis{
		Client: client,
	}
}