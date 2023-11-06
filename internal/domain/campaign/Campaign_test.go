package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewCampaign_shouldCampaignIsNonNull(t *testing.T) {

	//arrange
	var assert = assert.New(t)
	const aName string = "Test Campaign"
	const aContent string = "Test Content"
	var aContacts = []string{"anderson@gmail", "babi@gmail"}

	//act
	var campaign = newCampaign(aName, aContent, aContacts)

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
	const aName string = "Test Campaign"
	const aContent string = "Test Content"
	var aContacts = []string{"anderson@gmail", "babi@gmail"}

	//act
	var campaign = newCampaign(aName, aContent, aContacts)

	//assert
	assert.NotNil(campaign.ID)

}
