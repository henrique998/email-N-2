package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignFindById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaignId := chi.URLParam(r, "id")

	campaign, err := h.CampaignService.FindById(campaignId)
	if err == nil && campaign == nil {
		return nil, http.StatusNotFound, err
	}

	return campaign, http.StatusOK, err
}
