package handler

import (
	"fmt"
	"golang-dasar/auth"
	"golang-dasar/helper"
	"golang-dasar/user"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	//map input dari user ke struct  RegisterUserInput
	//struct diatas kita parsing  sebagai parameter service

	var input user.RegisterUserInput

	//validation
	err := c.ShouldBindJSON(&input)
	if err != nil {

		//format error biar gampang terbaca
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//buat token
	token, err := h.authService.GenerateToken(int(newUser.ID))
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) LoginUser(c *gin.Context) {
	//user memasukkan input email dan password
	//input ditangkap handler
	//mapping dari input user ke input struct
	//input struct passing ke service
	//diservice mencari bantuan repository user dengan email yg sesuai
	//cocokan password

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login account failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.Login(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("user not found", http.StatusUnauthorized, "error", errMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	//buat token
	token, err := h.authService.GenerateToken(int(userData.ID))
	if err != nil {
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(userData, token)

	response := helper.APIResponse("Successfully login!", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	//ada input email dr user
	//input email di mapping ke struct input
	//struct input di passing ke service
	//service manggil repository -> email sudah ada atau belum
	//repository query ke database

	var input user.CheckEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("email chacking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	//input dr user
	//simpan gambar di folder "images/"
	//service memanggil repo
	//upload file bawaan gin

	file, err := c.FormFile("avatar") //nama json pada field image
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload an avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//batasi ukuran file
	const maxFileSize = 1 * 1024 * 1024 // 1 MB dalam hitungan Bytes
	if file.Size > maxFileSize {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("file size too large, maximum is 1MB", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1

	//simpan file dengan generate nama file agar tidak tumpang tindih
	// path := "images/" + file.Filename
	timeToStr := strconv.Itoa(int(time.Now().UnixNano())) //ubah waktu ke string
	ext := filepath.Ext(file.Filename)                    //ambil extention saja
	path := fmt.Sprintf("images/%d-%s%s", userID, timeToStr, ext)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload an avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload an avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
