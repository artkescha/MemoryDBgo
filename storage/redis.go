package storage

import (
	"sync"
	"time"

	"gopkg.in/redis.v3"
)

//РЕАЛИЗАЦИЯ ПОВЕДЕНИЯ ХРАНИЛИЩА

//Врямя жизни ключа
const BOXLIFETIME = 30

var instance *Redis
var once sync.Once

//Связь с Redis
type Redis struct {
	client *redis.Client
}

//Singleton, так-как хранилище всегда в единственном экземпляре
// Инициализирует и возвращает новый DB redis
func NewRedis(cl *redis.Client) *Redis {

	once.Do(func() {
		instance = &Redis{client: cl}
	})
	return instance
}

// Реализация метода интерфейса Set
func (r *Redis) Set(key string, val string) error {
	return r.client.Set(key, val, time.Second*BOXLIFETIME).Err()
}

// Реализация метода интерфейса Delete
func (r *Redis) Delete(key string) error {
	err := r.client.Del(key)
	if err.Val() == 0 {
		return ErrNotFound
	}
	return r.client.Del(key).Err()
}

// Реализация метода интерфейса Get
func (r *Redis) Get(key string) ([]byte, error) {
	str, err := r.client.Get(key).Result()
	if err != nil {
		return nil, ErrNotFound
	}
	return []byte(str), nil
}
