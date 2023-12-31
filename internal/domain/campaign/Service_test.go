package campaign

import (
	"errors"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ms-email/internal/dto"
	"ms-email/internal/internalErrors"
	"testing"
)

var (
	campaignDTO = dto.NewCampaignDTO{
		Name:     "Testttt",
		Content:  "Bodyttt",
		Contacts: []string{"andersonconforto@email.com"},
	}
	mocks   = new(RepositoryMock)
	service = Service{
		Repository: mocks,
	}
	fake = faker.New()
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Execute(dto.NewCampaignDTO{})
	assert.NotNil(err)
	assert.Equal("Name min 5", err.Error())

}

func Test_Create_ValidateDomainError2(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Execute(dto.NewCampaignDTO{
		Name:     fake.Lorem().Text(38),
		Contacts: aContacts,
		Content:  aContent,
	})
	assert.NotNil(err)
	assert.Equal("Name max 24", err.Error())

}

func Test_Create_SaveCampaign(t *testing.T) {

	mocks.On("Save", mock.MatchedBy(
		func(campaign *Campaign) bool {
			if campaignDTO.Name != campaign.Name {
				return false
			} else if campaignDTO.Content != campaign.Content {
				return false
			} else if len(campaignDTO.Contacts) != len(campaign.Contacts) {
				return false
			}
			return true
		})).Return(nil)

	service.Execute(campaignDTO)
	mocks.AssertExpectations(t)

}

func Test_Create_ValidateDataBase(t *testing.T) {
	assert := assert.New(t)
	mocks.On("Save", mock.Anything).Return(errors.New("error"))

	_, err := service.Execute(campaignDTO)
	assert.True(errors.Is(err, internalErrors.ErrInternal))

}
