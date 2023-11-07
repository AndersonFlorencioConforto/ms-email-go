package campaign

import (
	"ms-email/internal/dto"
	"ms-email/internal/internalErrors"
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
		return "", internalErrors.ErrInternal
	}
	return campaign.ID, nil

}
