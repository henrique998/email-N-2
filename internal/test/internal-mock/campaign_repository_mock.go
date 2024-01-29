package internalmock

import (
	campaing "github.com/henrique998/email-N/internal/domain/campaign"
	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Create(campaing *campaing.Campaing) error {
	args := r.Called(campaing)

	return args.Error(0)
}

func (r *CampaignRepositoryMock) Update(campaing *campaing.Campaing) error {
	args := r.Called(campaing)

	return args.Error(0)
}

func (r *CampaignRepositoryMock) Get() ([]campaing.Campaing, error) {

	return nil, nil
}

func (r *CampaignRepositoryMock) GetById(campaingId string) (*campaing.Campaing, error) {
	args := r.Called(campaingId)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*campaing.Campaing), nil
}

func (r *CampaignRepositoryMock) Delete(campaign *campaing.Campaing) error {

	return nil
}
