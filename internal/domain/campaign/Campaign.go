package campaign

import (
	"errors"
	"github.com/rs/xid"
	"time"
)

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

func newCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" || name == " " {
		return nil, errors.New("name is required")
	}
	if content == "" || content == " " {
		return nil, errors.New("content is required")

	}
	if emails == nil || len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]Contact, 0)

	for _, email := range emails {
		contacts = append(contacts, Contact{
			Email: email,
		})

	}

	//for index, email := range emails {
	//	contacts[index] = Contact{Email: email}
	//}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil

}
