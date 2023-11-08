package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"ms-email/internal/domain/campaign"
	"ms-email/internal/dto"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/campaigns", func(writer http.ResponseWriter, request *http.Request) {
		var service = campaign.Service{}

		var req dto.NewCampaignDTO
		render.DecodeJSON(request.Body, &req)
		id, err := service.Execute(req)

		if err != nil {
			badRequest := http.StatusBadRequest
			render.Status(request, badRequest)
			render.JSON(writer, request, map[string]interface{}{
				"error": err.Error(),
				"code":  badRequest,
			})
			return
		}
		response := map[string]string{
			"id": id,
		}
		render.Status(request, http.StatusCreated)
		render.JSON(writer, request, response)

	})
	//router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	param := r.URL.Query().Get("name")
	//	w.Write([]byte("hello world" + param))
	//})
	//router.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
	//	param := chi.URLParam(r, "name")
	//	w.Write([]byte("hello world" + param))
	//})
	//router.Get("/json", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	var m = map[string]string{
	//		"message": "sucess",
	//	}
	//	marshal, _ := json.Marshal(m)
	//	w.Write(marshal)
	//})
	//router.Get("/json/chi", func(w http.ResponseWriter, r *http.Request) {
	//	var m = map[string]string{
	//		"message": "sucess",
	//	}
	//	render.JSON(w, r, m)
	//
	//})
	//router.Post("/post", func(w http.ResponseWriter, r *http.Request) {
	//
	//	var product Product
	//	render.DecodeJSON(r.Body, &product)
	//	render.JSON(w, r, product)
	//
	//})
	http.ListenAndServe(":3000", router)
}
