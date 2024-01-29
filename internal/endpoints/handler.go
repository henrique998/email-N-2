package endpoints

import campaing "github.com/henrique998/email-N/internal/domain/campaign"

type Handler struct {
	CampaignService campaing.Service
}
