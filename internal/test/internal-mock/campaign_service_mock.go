package internalmock

import (
	"github.com/henrique998/email-N/internal/contracts"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(newCampaing contracts.NewCampaingDTO) (string, error) {
	args := r.Called(newCampaing)

	return args.String(0), args.Error(1)
}

func (s *CampaignServiceMock) FindById(campaingId string) (*contracts.CampaignResponseDTO, error) {
	args := s.Called(campaingId)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contracts.CampaignResponseDTO), args.Error(1)
}

func (s *CampaignServiceMock) Delete(campaingId string) error {
	return nil
}
