package handlers

import (
	"OLTPStorage/storage"
	"fmt"
	"net/http"
)

//Обработчик Get
func Get(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}
		val, err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(w, "NotFound!", http.StatusNotFound)
			return

		} else if err != nil {

			http.Error(w, fmt.Sprintf("error getting value from DB: %s", err), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}
