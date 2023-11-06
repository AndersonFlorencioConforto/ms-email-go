package campaign

import (
	"ms-email/internal/dto"
)

type Service struct {
	Repository Repository
}

func (s *Service) execute(dto dto.NewCampaignDTO) (string, error) {
	campaign, _ := newCampaign(dto.Name, dto.Content, dto.Contacts)
	s.Repository.Save(campaign)
	return campaign.ID, nil

}
