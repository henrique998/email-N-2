package campaing

type Repository interface {
	Create(campaing *Campaing) error
	Update(campaing *Campaing) error
	Get() ([]Campaing, error)
	GetById(campaignId string) (*Campaing, error)
	Delete(campaing *Campaing) error
}
