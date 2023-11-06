package campaign

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ms-email/internal/dto"
	"testing"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)

}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	service := Service{}

	campaignDTO := dto.NewCampaignDTO{
		Name:     "Test",
		Content:  "Body",
		Contacts: []string{"andersonconforto@email"},
	}

	id, err := service.execute(campaignDTO)
	assert.NotNil(id)
	assert.Nil(err)

}

func Test_Create_SaveCampaign(t *testing.T) {

	campaignDTO := dto.NewCampaignDTO{
		Name:     "Test",
		Content:  "Body",
		Contacts: []string{"andersonconforto@email"},
	}

	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(
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

	service := Service{
		Repository: repositoryMock,
	}

	service.execute(campaignDTO)
	repositoryMock.AssertExpectations(t)

}
