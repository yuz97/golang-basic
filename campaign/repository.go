package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error) //[] ->mengembalikan byk data
	FindByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

// supaya bisa di akses diluar buat instance NewRepository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {

	var campaigns []Campaign //ambil semua data campaign ditandai []Campaign

	if err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {

	var campaigns []Campaign //ambil semua data campaign berdasark userId

	// lakukan filtering pada campaignImages dengan hanya mengambil image isPrimary = 1
	// CampaignImages  => nama field dengan bentuk camel case
	//campaign_images => nama table dengan bentuk snake case
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
