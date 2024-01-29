package contracts

type CampaignResponseDTO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Content       string `json:"content"`
	Status        string `json:"status"`
	AmoutOfEmails int    `json:"amountOfEmails"`
}
