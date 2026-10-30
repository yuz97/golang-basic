package campaign

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	//cek userID
	if userID != 0 {
		//jika ada
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	//jika tidak ada userId tampilkan semua
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {

	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Title = input.Title
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID // ambil userId utk kebutuhan jwt

	//buat unique slug
	strSlug := fmt.Sprintf("%s %s", input.Title, strconv.Itoa(input.User.ID))
	campaign.Slug = slug.Make(strSlug)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *service) UpdateCampaign(input GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	//find by ID
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	//cegah agar campaign hanya dibuat user dengan authorization sama (owner)
	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner of the campaign!")
	}

	//update data
	campaign.Title = inputData.Title
	campaign.ShortDescription = inputData.Title
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	//update slug
	strSlug := fmt.Sprintf("%s %s", inputData.Title, strconv.Itoa(inputData.User.ID))
	campaign.Slug = slug.Make(strSlug)

	// save kedatabase
	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}
