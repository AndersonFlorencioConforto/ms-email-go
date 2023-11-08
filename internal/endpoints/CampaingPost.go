package endpoints

import (
	"github.com/go-chi/render"
	"ms-email/internal/dto"
	"net/http"
)

func (h *Handler) CampaignPost(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {

	var req dto.NewCampaignDTO
	render.DecodeJSON(request.Body, &req)
	id, err := h.CampaignService.Execute(req)

	response := map[string]string{
		"id": id,
	}
	return response, http.StatusCreated, err

}
