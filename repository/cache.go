package repository

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	redigo "github.com/gomodule/redigo/redis"
)

type CacheRepository interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	GetAll() ([]string, error)
	Delete(header string) error
}

type cacheRepository struct {
	rdb *redis.Client
	rdg redigo.Conn
}

func (c *cacheRepository) Set(key string, value []byte) error {
	_, err := c.rdb.Set(key, value, 10*time.Minute).Result()

	if err != nil {
		return fmt.Errorf("error to set cache %v", err.Error())
	}
	return nil
}

func (c *cacheRepository) Get(key string) ([]byte, error) {

	result, err := c.rdb.Get(key).Result()

	if err != nil {
		return []byte{}, fmt.Errorf("error to get cache %s", err.Error())
	}

	return []byte(result), nil
}

func (c *cacheRepository) GetAll() ([]string, error) {
	var result []string
	keys, err := redigo.Strings(c.rdg.Do("KEYS", "*"))
	if err != nil {
		fmt.Println(err)
	}

	for _, key := range keys {
		result = append(result, string(key))
	}

	return result, nil
}

func (c *cacheRepository) Delete(header string) error {
	_, err := c.rdb.Del(header).Result()

	if err != nil {
		return err
	}
	return nil
}

func NewCache(rdb *redis.Client, rdg redigo.Conn) CacheRepository {
	return &cacheRepository{
		rdb: rdb,
		rdg: rdg,
	}
}
