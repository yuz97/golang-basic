package campaign

import (
	"golang-dasar/user"
	"time"
)

type Campaign struct {
	ID               int `gorm:"primaryKey"`
	UserID           int
	Title            string
	ShortDescription string
	Description      string //tipe data text tapi digolang tetap string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage //berelasi dengan CampaignImage
	User             user.User       //berelasi dengan User
}

type CampaignImage struct {
	ID         int `gorm:"primaryKey"`
	CampaignID int
	FileName   string
	IsPrimary  bool `gorm:"type:tinyint(1)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time //1 true 0 false

}

// Opsi 2: Tetap pakai gorm.Model, paksa tipe kolom FK

// Paksa kolom campaign_id jadi unsigned lewat tag, tapi field Go tetap int:

// go
// type CampaignImage struct {
// 	gorm.Model

// 	CampaignID int `gorm:"type:bigint unsigned"`
// 	FileName   string
// 	IsPrimary  bool `gorm:"type:tinyint(1)"`
// }
