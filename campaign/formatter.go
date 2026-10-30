package campaign

import "strings"

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

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary" gorm:"type:tinyint(1)"`
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Title            string                   `json:"title"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"` //perks byk data jd slice/array
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	formatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		Title:            campaign.Title,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		UserID:           campaign.UserID,
	}

	//biarkan image jadi kosong kalo gak ada image
	formatter.ImageURL = ""

	//jika ada gambar
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	//perks
	var perks []string
	//split atau potong berdasarkan tanda koma dan loop perks
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	formatter.Perks = perks

	//user
	campaignUserFormatter := CampaignUserFormatter{
		Name:     campaign.User.Name,
		ImageURL: campaign.User.AvatarFileName,
	}

	formatter.User = campaignUserFormatter

	//images
	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName
		campaignImageFormatter.IsPrimary = image.IsPrimary

		images = append(images, campaignImageFormatter)
	}

	formatter.Images = images

	return formatter
}
