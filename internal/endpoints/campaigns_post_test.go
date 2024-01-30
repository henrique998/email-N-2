package endpoints

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/henrique998/email-N/internal/contracts"
	internalMock "github.com/henrique998/email-N/internal/test/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(body contracts.NewCampaingDTO, createdBy string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	ctx := context.WithValue(req.Context(), "email", createdBy)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	return req, rr
}

func Test_CampaignsPost_Should_Save_New_Campaign(t *testing.T) {
	assert := assert.New(t)
	createdByExpected := "test@test.com.br"
	body := contracts.NewCampaingDTO{
		Name:    "test",
		Content: "Hi there!",
		Emails:  []string{"test@test.com"},
	}
	service := new(internalMock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contracts.NewCampaingDTO) bool {
		if request.Name == body.Name &&
			request.Content == body.Content &&
			request.CreatedBy == createdByExpected {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaignService: service}

	req, rr := setup(body, createdByExpected)

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignsPost_Should_Inform_Error_When_Exists(t *testing.T) {
	assert := assert.New(t)
	body := contracts.NewCampaingDTO{
		Name:    "test",
		Content: "Hi there!",
		Emails:  []string{"test@test.com"},
	}
	service := new(internalMock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	req, rr := setup(body, "test@test.com.br")

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
