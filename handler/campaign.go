package handler

import (
	"golang-dasar/campaign"
	"golang-dasar/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//tangkap parameter dihandler
//handler ke servie
//service yg menentukan repository mana yg di call
//repository(findAll,FindByUserID) ->DB

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	//tangkap parameter query dan ubah jadi int
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get campaigns", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}
