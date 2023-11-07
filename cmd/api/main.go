package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	router.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var m = map[string]string{
			"message": "sucess",
		}
		marshal, _ := json.Marshal(m)
		w.Write(marshal)
	})
	router.Get("/json/chi", func(w http.ResponseWriter, r *http.Request) {
		var m = map[string]string{
			"message": "sucess",
		}
		render.JSON(w, r, m)

	})
	http.ListenAndServe(":3000", router)
}
