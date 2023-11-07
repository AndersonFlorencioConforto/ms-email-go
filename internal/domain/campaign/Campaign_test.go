package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	aName    string = "Test Campaign"
	aContent string = "Test Content"
)

var (
	aContacts = []string{"anderson@gmail.com", "babi@gmail.com"}
)

func Test_NewCampaign_shouldCampaignIsNonNull(t *testing.T) {

	//arrange
	var assert = assert.New(t)

	//act
	var campaign, _ = newCampaign(aName, aContent, aContacts)

	//assert
	assert.NotNilf(campaign.ID, "campaign.ID should not be nil")
	assert.Equal(aName, campaign.Name)
	assert.Equal(aContent, campaign.Content)
	assert.Equal(len(aContacts), len(campaign.Contacts))
	assert.Equal(campaign.CreatedAt, time.Now())

	//if campaign.ID != "1" {
	//	t.Errorf("Expected campaign.ID to be 1, got %s", campaign.ID)
	//}
	//if campaign.Name != aName {
	//	t.Errorf("Expected campaign.Name to be %s, got %s", aName, campaign.Name)
	//}
	//if campaign.Content != aContent {
	//	t.Errorf("Expected campaign.Content to be %s, got %s", aContent, campaign.Content)
	//}
	//
	//fmt.Println(campaign.Contacts)
	//if len(campaign.Contacts) != len(aContacts) {
	//	t.Errorf("Expected campaign.Contacts to have %d items, got %d", len(aContacts), len(campaign.Contacts))
	//
	//}

}

func Test_NewCampaign_IDIsNotNull(t *testing.T) {
	var assert = assert.New(t)

	//act
	var campaign, _ = newCampaign(aName, aContent, aContacts)

	//assert
	assert.NotNil(campaign.ID)
}

func Test_givenAValidCampaignWithNameNull_whenValidateName_thenThrowError(t *testing.T) {
	var assert = assert.New(t)

	//act
	var _, error = newCampaign("", aContent, aContacts)

	//assert
	assert.Equal("Name min 5", error.Error())

}

func Test_givenAValidCampaignWithContentNull_whenValidateContent_thenThrowError(t *testing.T) {
	var assert = assert.New(t)

	//act
	var _, error = newCampaign(aName, "", aContacts)

	//assert
	assert.Equal("Content min 5", error.Error())

}

func Test_givenAValidCampaignWithContactsNull_whenValidateContact_thenThrowError(t *testing.T) {
	var assert = assert.New(t)

	//act
	var _, error = newCampaign(aName, aContent, nil)

	//assert
	assert.Equal("Contacts min 1", error.Error())

}
