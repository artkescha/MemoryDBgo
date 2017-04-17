package handlers

import (
	"OLTPStorage/storage"
	"io/ioutil"
	"net/http"
)

//Обработчик Set
func Set(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading value body", http.StatusBadRequest)
			return

		}
		//val = []byte("sasasas")
		err = db.Set(key, string(val))
		if err != nil {

			http.Error(w, "Error Set value", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}
