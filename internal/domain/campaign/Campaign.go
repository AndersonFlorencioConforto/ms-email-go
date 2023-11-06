package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

func newCampaign(name string, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))

	for _, email := range emails {
		contacts = append(contacts, Contact{
			Email: email,
		})

	}

	return &Campaign{
		ID:        "1",
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

}
