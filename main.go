// OLTPStorage project main.go
package main

import (
	"OLTPStorage/handlers"
	"OLTPStorage/storage"
	"log"
	"net/http"
	"time"
)

func main() {

	db := storage.CreateMemoryDB()
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:           ":8081",
		Handler:        mux,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	db.StartCleaning(0)

	mux.Handle("/get", handlers.Get(db))
	mux.Handle("/set", handlers.Set(db))
	mux.Handle("/del", handlers.Delete(db))

	log.Printf("serving on port 8081")

	err := s.ListenAndServe()
	log.Fatal(err)
}
