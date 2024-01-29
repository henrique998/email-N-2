package endpoints

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
	"gorm.io/gorm"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpoint EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpoint(w, r)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, map[string]string{"error": err.Error()})
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, map[string]string{"error": err.Error()})
			}

			return
		}

		render.Status(r, status)

		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
