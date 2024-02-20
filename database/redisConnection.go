package database

import (
	"fmt"
	"list/config"
	"log"
	"strconv"

	"github.com/go-redis/redis"
	redigo "github.com/gomodule/redigo/redis"
)

func GetConnectionRedis(cnf *config.Config) *redis.Client {
	Db, err := strconv.Atoi(cnf.Redis.Database)

	if err != nil {
		log.Fatalf("something wrong in redis connection %s", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Redis.Address,
		Password: cnf.Redis.Password,
		DB:       Db,
	})

	_, err = rdb.Ping().Result()

	if err != nil {
		log.Fatal("Redis connection was refused")
	}

	return rdb
}

func GetConnectionDial(cnf *config.Config) redigo.Conn {
	conn, err := redigo.Dial("tcp", cnf.Redis.Address)

	if err != nil {
		fmt.Println(err, "SINI KENA1")
	}

	keys, err := redigo.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		fmt.Println(err, "SINI KENA2")

	}

	for _, key := range keys {
		fmt.Println(key)
	}
	return conn

}
