package campaing_test

import (
	"errors"
	"testing"

	"github.com/henrique998/email-N/internal/contracts"
	campaing "github.com/henrique998/email-N/internal/domain/campaign"
	internalerrors "github.com/henrique998/email-N/internal/internalErrors"
	internalMock "github.com/henrique998/email-N/internal/test/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaing = contracts.NewCampaingDTO{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"jhondoe@gmail.com", "henrique@gmail.com"},
	}
	service = campaing.ServiceImp{}
)

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("Create", mock.Anything).Return(nil)
	service.Repo = repository

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contracts.NewCampaingDTO{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Save_CreateCampaing(t *testing.T) {
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("Create", mock.MatchedBy(func(campaing *campaing.Campaing) bool {
		if campaing.Name != newCampaing.Name {
			return false
		}

		if campaing.Content != newCampaing.Content {
			return false
		}

		if len(campaing.Contacts) != len(newCampaing.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repo = repository

	service.Create(newCampaing)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("Create", mock.Anything).Return(errors.New("error to save on database"))
	service.Repo = repository

	_, err := service.Create(newCampaing)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_FindById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaing.NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("GetById", mock.MatchedBy(func(campaignId string) bool {
		return campaignId == campaign.ID
	})).Return(campaign, nil)
	service.Repo = repository

	campaignReturned, _ := service.Repo.GetById(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_FindById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaing.NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("GetById", mock.Anything).Return(nil, errors.New("Something wrong"))
	service.Repo = repository

	_, err := service.FindById(campaign.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNotFound_When_Campaign_Not_Exists(t *testing.T) {
	assert := assert.New(t)

	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repo = repository

	err := service.Delete("Invalid-Id")

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_ReturnStatusInvalid_When_Campaign_Has_Status_Not_Equals_Pending(t *testing.T) {
	assert := assert.New(t)

	campaign := &campaing.Campaing{ID: "1", Status: campaing.Started}
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("GetById", mock.Anything).Return(campaign, nil)
	service.Repo = repository

	err := service.Delete(campaign.ID)

	assert.Equal("Campaign status invalid", err.Error())
}

func Test_Delete_ReturnInternalError_When_Delete_Has_Problem(t *testing.T) {
	assert := assert.New(t)

	campaignFound, _ := campaing.NewCampaing("Test", "Body", []string{"test@test.com.br"})
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("GetById", mock.Anything).Return(campaignFound, nil)
	repository.On("Delete", mock.MatchedBy(func(campaign *campaing.Campaing) bool {
		return campaignFound == campaign
	})).Return(errors.New("error to delete campaign"))
	service.Repo = repository

	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNil_When_Delete_Has_Success(t *testing.T) {
	assert := assert.New(t)

	campaignFound, _ := campaing.NewCampaing("Test", "Body", []string{"test@test.com.br"})
	repository := new(internalMock.CampaignRepositoryMock)
	repository.On("GetById", mock.Anything).Return(campaignFound, nil)
	repository.On("Delete", mock.MatchedBy(func(campaign *campaing.Campaing) bool {
		return campaignFound == campaign
	})).Return(nil)
	service.Repo = repository

	err := service.Delete(campaignFound.ID)

	assert.Nil(err)
}
