package campaign

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model

	UserID           int
	Title            string
	ShortDescription string
	Description      string //tipe data text tapi digolang tetap string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CampaignImages   []CampaignImage //berelasi
}

type CampaignImage struct {
	gorm.Model

	CampaignID int
	FileName   string
	IsPrimary  bool `gorm:"type:tinyint(1)"` //1 true 0 false

}
