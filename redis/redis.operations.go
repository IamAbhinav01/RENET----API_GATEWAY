package redis

import (
	"context"
	"encoding/json"
	"log"

	"time"

	"github.com/redis/go-redis/v9"
)

type SessionStore interface {
	Save(ctx context.Context, sessionId string, data map[string]string, expiry time.Duration) error
	Get(ctx context.Context, sessionId string) (map[string]string, error)
	Destroy(ctx context.Context, sessionId string) error
}

type RedisSessionStore struct {
	redis_client *redis.Client
}

func NewSessionStore(redis_client *redis.Client) *RedisSessionStore {
	return &RedisSessionStore{
		redis_client: redis_client,
	}
}


func (store *RedisSessionStore) Save(ctx context.Context, sessionId string, data map[string]string, expiry time.Duration) error {

	bytes, err := json.Marshal(data)

	if err != nil {
		log.Println("Error Happend while Marshalling data")
		return err
	}

	return store.redis_client.Set(ctx, "session:"+sessionId, bytes, expiry).Err()

}

func (store *RedisSessionStore) Get(ctx context.Context, sessionId string) (map[string]string, error) {

	bytes, err := store.redis_client.Get(ctx, "session:"+sessionId).Bytes()

	if err != nil {
		log.Println("Error occured while retrieved data from redis")
		return nil, err
	}

	var data map[string]string
	err = json.Unmarshal(bytes, &data)
	return data, err

}

func (store *RedisSessionStore) Destroy(ctx context.Context, sessionId string) error {

	return store.redis_client.Del(ctx, "session:"+sessionId).Err()

}
