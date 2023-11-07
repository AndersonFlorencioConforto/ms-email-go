package campaign

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ms-email/internal/dto"
	"ms-email/internal/internalErrors"
	"testing"
)

var (
	campaignDTO = dto.NewCampaignDTO{
		Name:     "Test",
		Content:  "Body",
		Contacts: []string{"andersonconforto@email"},
	}
	mocks   = new(RepositoryMock)
	service = Service{
		Repository: mocks,
	}
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

	campaignDTO.Name = ""
	_, err := service.execute(campaignDTO)
	assert.NotNil(err)
	assert.Equal("name is required", err.Error())

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

	service.execute(campaignDTO)
	mocks.AssertExpectations(t)

}

func Test_Create_ValidateDataBase(t *testing.T) {
	assert := assert.New(t)
	mocks.On("Save", mock.Anything).Return(errors.New("error"))

	_, err := service.execute(campaignDTO)
	assert.True(errors.Is(err, internalErrors.ErrInternal))

}
