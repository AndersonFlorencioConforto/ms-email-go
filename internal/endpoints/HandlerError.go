package endpoints

import (
	"errors"
	"github.com/go-chi/render"
	"ms-email/internal/internalErrors"
	"net/http"
)

type EndpointFunc func(writer http.ResponseWriter, request *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		obj, status, err := endpointFunc(writer, request)

		if err != nil {
			var errorStatus int
			if errors.Is(err, internalErrors.ErrInternal) {
				errorStatus = http.StatusInternalServerError
				render.Status(request, errorStatus)
			} else {
				errorStatus = http.StatusBadRequest
				render.Status(request, errorStatus)
			}
			render.JSON(writer, request, map[string]interface{}{
				"error": err.Error(),
				"code":  errorStatus,
			})
			return
		}
		render.Status(request, status)
		if obj != nil {
			render.JSON(writer, request, obj)
		}

	})

}
