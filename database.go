package main

import (
	"github.com/gomodule/redigo/redis"
)

func newDbPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   5,
		MaxActive: 200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddress+":"+redisPort)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

func writeKeyValue(dbPool *redis.Pool, key string, value string) error {
	conn := dbPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}
