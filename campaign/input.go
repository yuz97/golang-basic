package campaign

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"` //ambil input ID dari uri
}
