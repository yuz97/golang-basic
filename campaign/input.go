package campaign

import "golang-dasar/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"` //ambil input ID dari uri
}

type CreateCampaignInput struct {
	Title            string    `json:"title" binding:"required"`
	ShortDescription string    `json:"short_description" binding:"required"`
	Description      string    `json:"description" binding:"required"`
	GoalAmount       int       `json:"goal_amount" binding:"required"`
	Perks            string    `json:"perks" binding:"required"`
	User             user.User `json:"-"` //berdasarkan user yg create campaign jd butuh jwt gak usah kirim data ke json
}
