package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		w.Write([]byte("hello world" + param))
	})
	router.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "name")
		w.Write([]byte("hello world" + param))
	})
	http.ListenAndServe(":3000", router)
}
