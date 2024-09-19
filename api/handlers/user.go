package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "auth-athlevo/api/docs"
	t "auth-athlevo/api/token"
	auth "auth-athlevo/genproto/auth"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Get user profiles
// @Description Retrieve user profiles with optional filters
// @Tags users
// @Accept json
// @Produce json
// @Param id query string false "Filter by user ID"
// @Param username query string false "Filter by username"
// @Param full_name query string false "Filter by full name"
// @Param email query string false "Filter by email"
// @Param gym_id query string false "Filter by gym ID"
// @Param phone_number query string false "Filter by phone number"
// @Security BearerAuth
// @Success 200 {object} auth.UserRepeated
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/get-profiles [get]
func (h *Handlers) GetProfile(c *gin.Context) {
	req := &auth.GetByIdReq{
		Id:          c.Query("id"),
		Username:    c.Query("username"),
		FullName:    c.Query("full_name"),
		Email:       c.Query("email"),
		GymId:       c.Query("gym_id"),
		PhoneNumber: c.Query("phone_number"),
	}

	profiles, err := h.User.GetProfile(c, req)
	if err != nil {
		log.Println("Error getting profiles:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, profiles)
}

// EditProfile godoc
// @Summary Edit user profile
// @Description Update the profile of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body auth.EditProfileReqBpdy true "Updated profile details"
// @Success 200 {object} string "Profile updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/profiles [put]
func (h *Handlers) EditProfile(c *gin.Context) {
	var req auth.EditProfileReqBpdy
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := h.User.EditProfile(c, &auth.UserRes{
		Id:          req.Id,
		Username:    req.Username,
		FullName:    req.FullName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth,
		GymId:       req.GymId,
	})
	if err != nil {
		log.Println("Error editing profile:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// input, err := json.Marshal(req)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
	// 	return
	// }

	// err = h.Producer.ProduceMessages("upd-user", input)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Profile for user %s updated successfully", req.Id)})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Update the password of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param password body auth.ChangePasswordReqBody true "Updated password details"
// @Success 200 {object} string "Password updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/passwords [put]
func (h *Handlers) ChangePassword(c *gin.Context) {
	userID := getuserId(c)

	var body auth.ChangePasswordReqBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	password, err := t.HashPassword(body.NewPassword)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	body.NewPassword = password

	req := &auth.ChangePasswordReq{
		Id:              userID,
		CurrentPassword: body.CurrentPassword,
		NewPassword:     body.NewPassword,
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	err = h.Producer.ProduceMessages("upd-pass", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// GetSetting godoc
// @Summary Get user settings
// @Description Retrieve the settings of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.Setting
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/setting [get]
func (h *Handlers) GetSetting(c *gin.Context) {
	userID := getuserId(c)

	req := &auth.GetById{
		Id: userID,
	}

	setting, err := h.User.GetSetting(c, req)
	if err != nil {
		log.Println("Error getting setting:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, setting)
}

// EditSetting godoc
// @Summary Edit user settings
// @Description Update the settings of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param setting body auth.Setting true "Updated setting details"
// @Success 200 {object} string "Setting updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/setting [put]
func (h *Handlers) EditSetting(c *gin.Context) {
	userID := getuserId(c)

	var body auth.Setting
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &auth.SettingReq{
		Id:           userID,
		PrivacyLevel: body.PrivacyLevel,
		Notification: body.Notification,
		Language:     body.Language,
		Theme:        body.Theme,
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	err = h.Producer.ProduceMessages("upd-setting", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Setting for user %s updated successfully", req.Id)})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} string "User deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/delete/{id} [delete]
func (h *Handlers) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	req := &auth.GetById{
		Id: userID,
	}

	_, err := h.User.DeleteUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s deleted successfully", req.Id)})
}
