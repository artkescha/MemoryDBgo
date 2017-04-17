package handlers

import (
	"OLTPStorage/storage"
	"fmt"
	"net/http"
)

//Обработчик Delete
func Delete(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}
		err := db.Delete(key)
		if err == storage.ErrNotFound {
			http.Error(w, "NotFound!", http.StatusNotFound)
			return

		} else if err != nil {

			http.Error(w, fmt.Sprintf("error delete value from DB: %s", err), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
	})
}
