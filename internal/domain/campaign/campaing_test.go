package campaing

import (
	"strings"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "campaing x"
	content   = "Body Hi!"
	emails    = []string{"jhondoe@gmail.com", "henrique942@gmail.com", "henrique714@gmail.com"}
	createdBy = "test@test.com.br"
	fake      = faker.New()
)

func Test_NewCampaing_CreateCampaing(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, emails, createdBy)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(emails))
	assert.Equal(campaing.CreatedBy, createdBy)
}

func Test_NewCampaing_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, emails, createdBy)

	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_StatusMustStartWithPending(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, emails, createdBy)

	assert.Equal(Pending, campaing.Status)
}

func Test_NewCampaing_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaing, _ := NewCampaing(name, content, emails, createdBy)

	assert.Greater(campaing.CreatedAt, now)
}

func Test_NewCampaing_MustValidateNameMinLenght(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("", content, emails, createdBy)

	assert.Equal("name is required with min lenght: 5", err.Error())
}

func Test_NewCampaing_MustValidateNameMaxLenght(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(strings.Repeat("a", 25), content, emails, createdBy)

	assert.Equal("name is required with max lenght: 24", err.Error())
}

func Test_NewCampaing_MustValidateContentMinLenght(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", emails, createdBy)

	assert.Equal("content is required with min lenght: 5", err.Error())
}

func Test_NewCampaing_MustValidateContentMaxLenght(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, fake.Lorem().Text(1090), emails, createdBy)

	assert.Equal("content is required with max lenght: 1024", err.Error())
}

func Test_NewCampaing_MustValidateContactsMinLenght(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{}, createdBy)

	assert.Equal("contacts is required with min lenght: 1", err.Error())
}

func Test_NewCampaing_MustValidateContactsFormat(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{"invalid_email"}, createdBy)

	assert.Equal("email isn't valid", err.Error())
}

func Test_NewCampaing_MustValidateCreatedBy(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, emails, "")

	assert.Equal("createdby isn't valid", err.Error())
}
