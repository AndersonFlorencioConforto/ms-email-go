package campaign

import (
	"fmt"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	const aName string = "Test Campaign"
	const aContent string = "Test Content"
	var aContacts = []string{"anderson@gmail", "babi@gmail"}

	var campaign = newCampaign(aName, aContent, aContacts)

	if campaign.ID != "1" {
		t.Errorf("Expected campaign.ID to be 1, got %s", campaign.ID)
	}
	if campaign.Name != aName {
		t.Errorf("Expected campaign.Name to be %s, got %s", aName, campaign.Name)
	}
	if campaign.Content != aContent {
		t.Errorf("Expected campaign.Content to be %s, got %s", aContent, campaign.Content)
	}

	fmt.Println(campaign.Contacts)
	if len(campaign.Contacts) != len(aContacts) {
		t.Errorf("Expected campaign.Contacts to have %d items, got %d", len(aContacts), len(campaign.Contacts))

	}

}
