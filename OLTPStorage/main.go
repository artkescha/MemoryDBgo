// OLTPStorage project main.go
package main

import (
	"OLTPStorage/handlers"
	"OLTPStorage/storage"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"gopkg.in/redis.v3"
)

func main() {

	conf, err := GetConfig()
	if err != nil {
		log.Printf("Error getting config [%s]", err)
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	//считываем данные по настройке Redis
	var db *storage.Redis
	switch conf.DBType {
	case "redis":
		redisOpts := &redis.Options{
			Addr:     conf.RedisHost,
			Password: conf.RedisPass,
			DB:       conf.RedisDB,
		}
		rawRedisClient := redis.NewClient(redisOpts)
		db = storage.NewRedis(rawRedisClient)
	default:
		log.Fatal("Error: no available DB type %s", conf.DBType)
		os.Exit(1)
	}

	//создание и настройка сервера
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:           ":8081",
		Handler:        mux,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	mux.Handle("/get", handlers.Get(db))
	mux.Handle("/set", handlers.Set(db))
	mux.Handle("/del", handlers.Delete(db))

	log.Printf("serving on port 8081")

	errs := s.ListenAndServe()
	log.Fatal(errs)
}
