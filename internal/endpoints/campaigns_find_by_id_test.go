package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/henrique998/email-N/internal/contracts"
	internalMock "github.com/henrique998/email-N/internal/test/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignFindById_Should_Return_Campaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contracts.CampaignResponseDTO{
		ID:      "34",
		Name:    "test",
		Content: "Hi there!",
		Status:  "Pending",
	}
	service := new(internalMock.CampaignServiceMock)
	service.On("FindById", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignFindById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaign.ID, response.(*contracts.CampaignResponseDTO).ID)
	assert.Equal(campaign.Name, response.(*contracts.CampaignResponseDTO).Name)
	assert.Equal(campaign.Status, "Pending")
}

func Test_CampaignFindById_Should_Return_Error_When_Something_Wrong(t *testing.T) {
	assert := assert.New(t)

	service := new(internalMock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("FindById", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignFindById(rr, req)

	assert.Equal(errExpected.Error(), errReturned.Error())
}
