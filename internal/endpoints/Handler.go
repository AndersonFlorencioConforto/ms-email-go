package endpoints

import "ms-email/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
