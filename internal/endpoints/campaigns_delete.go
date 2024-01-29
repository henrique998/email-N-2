package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignsDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaignId := chi.URLParam(r, "id")

	err := h.CampaignService.Delete(campaignId)

	return nil, 200, err
}
