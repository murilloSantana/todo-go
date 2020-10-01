package todo

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bateu a nave"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
