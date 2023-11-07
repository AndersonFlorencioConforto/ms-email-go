package campaign

import (
	"github.com/rs/xid"
	"ms-email/internal/internalErrors"
	"time"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedAt time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func newCampaign(name string, content string, emails []string) (*Campaign, error) {
	//if name == "" || name == " " {
	//	return nil, errors.New("name is required")
	//}
	//if content == "" || content == " " {
	//	return nil, errors.New("content is required")
	//
	//}
	//if emails == nil || len(emails) == 0 {
	//	return nil, errors.New("contacts is required")
	//}

	contacts := make([]Contact, 0)

	for _, email := range emails {
		contacts = append(contacts, Contact{
			Email: email,
		})

	}

	//for index, email := range emails {
	//	contacts[index] = Contact{Email: email}
	//}

	var campaign = &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}
	err := internalErrors.ValidStruct(campaign)

	if err == nil {
		return campaign, nil

	}

	return nil, err

}
