package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/henrique998/email-N/internal/contracts"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contracts.NewCampaingDTO
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	return map[string]string{"campaignId": id}, http.StatusCreated, err
}
