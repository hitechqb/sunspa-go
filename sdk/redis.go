package sdk

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"os"
	"sunspa/common"
	"time"
)

type RedisProvider interface {
	Connect() error

	Ping() error

	// Do redis command
	Do(cmd string, params ...interface{}) (interface{}, error)

	Result(value interface{}) ([]byte, error)

	// Get redis item with key
	GetByKey(key string) (interface{}, error)

	// Get redis item and parse into struct
	GetStruct(out interface{}, cmd string, params ...interface{}) error

	// Get redis item with key and parse into struct
	GetStructByKey(key string, out interface{}) error

	// Set redis item with expire time
	Set(key string, value interface{}, duration time.Duration) error

	// Set redis item from a struct with expire time
	SetObject(key string, value interface{}, duration time.Duration) error

	Remove(keys ...string) error
}

type redisProvider struct {
	pool   *redis.Pool
	logger common.Logger
}

func NewRedisProvider(logger common.Logger) *redisProvider {
	return &redisProvider{
		pool: &redis.Pool{
			MaxIdle:     80,
			MaxActive:   12000,
			IdleTimeout: 10 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", os.Getenv("REDIS_URL"))
			},
		},
		logger: logger,
	}
}

func (r *redisProvider) Connect() error {
	if err := r.Ping(); err != nil {
		return err
	}

	r.logger.Infof(`🎉 Connected to Redis server on "%s" !`, os.Getenv("REDIS_URL"))
	return nil
}

func (r *redisProvider) Ping() error {
	_, err := r.Do("PING")
	if err != nil {
		return err
	}

	return nil
}

// Parse result to byte
func (r *redisProvider) Result(value interface{}) ([]byte, error) {
	if value == nil {
		return nil, redis.ErrNil
	}

	var err error
	return redis.Bytes(value, err)
}

// Do redis command
func (r *redisProvider) Do(cmd string, params ...interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	data, err := conn.Do(cmd, params...)
	if err != nil {
		r.logger.Errorln("[Redis.Error.Params]: ", params)
		r.logger.Errorln("[Redis.Error.Message]: ", err.Error())
	}

	return data, err
}

// Get redis item with key
func (r *redisProvider) GetByKey(key string) (interface{}, error) {
	return r.Do("GET", key)
}

// Get redis item and parse into struct
func (r *redisProvider) GetStruct(out interface{}, cmd string, params ...interface{}) error {
	data, err := r.Do(cmd, params...)
	if err != nil {
		return err
	}

	if data == nil {
		return nil
	}

	bData, err := redis.Bytes(data, err)
	if err != nil {
		return err
	}

	return json.Unmarshal(bData, out)
}

// Get redis item with key and parse into struct
func (r *redisProvider) GetStructByKey(key string, out interface{}) error {
	return r.GetStruct(out, "GET", key)
}

// Set redis item with expire time
func (r *redisProvider) Set(key string, value interface{}, duration time.Duration) error {
	_, err := r.Do("SETEX", key, int64(duration), value)
	if err != nil {
		r.logger.Infoln("[REDIS.Set.Key]: ", key)
		r.logger.Errorln("[REDIS.Set.Error]: ", err.Error())
		return err
	}

	return nil
}

// Set redis item from a struct with expire time
func (r *redisProvider) SetObject(key string, value interface{}, duration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.Set(key, data, duration)
}

// Remove redis Cache by Keys
func (r *redisProvider) Remove(keys ...string) error {
	for _, key := range keys {
		_, err := r.Do("DEL", key)
		if err != nil {
			return err
		}
	}

	return nil
}
