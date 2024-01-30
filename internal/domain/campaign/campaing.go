package campaing

import (
	"time"

	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
	"github.com/rs/xid"
)

const (
	Pending  = "Pending"
	Started  = "Started"
	Done     = "Done"
	Canceled = "Canceled"
	Deleted  = "Deleted"
)

type Contact struct {
	ID         string `validate:"required" json:"id" gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaingId string
}

type Campaing struct {
	ID        string    `validate:"required" json:"id" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" json:"name" gorm:"size:100"`
	Content   string    `validate:"min=5,max=1024" json:"content" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive" json:"contacts"`
	CreatedBy string    `validate:"email" gorm:"size:50"`
	Status    string    `gorm:"size:20"`
	CreatedAt time.Time `validate:"required" json:"createdAt"`
}

func NewCampaing(name, content string, emails []string, createdBy string) (*Campaing, error) {
	contacts := make([]Contact, len(emails))

	for i, email := range emails {
		contacts[i].ID = xid.New().String()
		contacts[i].Email = email
	}

	campaing := &Campaing{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}

	err := internalerrors.ValidateStruct(campaing)
	if err != nil {
		return nil, err
	}

	return campaing, nil
}

func (c *Campaing) Cancel() {
	c.Status = Canceled
}

func (c *Campaing) Delete() {
	c.Status = Deleted
}
