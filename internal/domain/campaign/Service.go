package campaign

import (
	"ms-email/internal/dto"
)

type Service struct {
	Repository Repository
}

func (s *Service) execute(dto dto.NewCampaignDTO) (string, error) {
	campaign, err := newCampaign(dto.Name, dto.Content, dto.Contacts)

	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)

	if err != nil {
		return "", err
	}
	return campaign.ID, nil

}
