package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amoung"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

// single format campaign
func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Title:            campaign.Title,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}
	//biarkan image jadi kosong kalo gak ada image
	formatter.ImageURL = ""

	//jika ada gambar
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return formatter
}

// multiple format campaigns
func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	//cara 1
	//jika campaigns tidak ada
	// if len(campaigns) == 0 {
	// 	return []CampaignFormatter{} //return data kosong berupa [] jangan null
	// }

	//cara 2
	campaignsFormatter := []CampaignFormatter{} //nelai default []

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}
