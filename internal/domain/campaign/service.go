package campaing

import (
	"errors"

	"github.com/henrique998/email-N/internal/contracts"
	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
)

type Service interface {
	Create(newCampaing contracts.NewCampaingDTO) (string, error)
	FindById(campaingId string) (*contracts.CampaignResponseDTO, error)
	Delete(campaingId string) error
}

type ServiceImp struct {
	Repo Repository
}

func (s *ServiceImp) Create(newCampaing contracts.NewCampaingDTO) (string, error) {
	campaing, err := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails, newCampaing.CreatedBy)
	if err != nil {
		return "", err
	}

	err = s.Repo.Create(campaing)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaing.ID, nil
}

func (s *ServiceImp) FindById(campaingId string) (*contracts.CampaignResponseDTO, error) {
	campaign, err := s.Repo.GetById(campaingId)
	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}

	return &contracts.CampaignResponseDTO{
		ID:            campaign.ID,
		Name:          campaign.Name,
		Content:       campaign.Content,
		Status:        campaign.Status,
		AmoutOfEmails: len(campaign.Contacts),
		CreatedBy:     campaign.CreatedBy,
	}, nil
}

func (s *ServiceImp) Delete(campaingId string) error {
	campaign, err := s.Repo.GetById(campaingId)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("campaign status invalid")
	}

	campaign.Delete()

	err = s.Repo.Delete(campaign)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
