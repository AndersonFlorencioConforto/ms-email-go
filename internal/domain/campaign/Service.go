package campaign

import "ms-email/internal/dto"

type Service struct {
	Repository Repository
}

func (s *Service) CreateCampaign(dto dto.NewCampaignDTO) error {
	campaign, err := newCampaign(dto.Name, dto.Content, dto.Contacts)
	if err != nil {
		return err
	}
	return s.Repository.Save(campaign)

}
