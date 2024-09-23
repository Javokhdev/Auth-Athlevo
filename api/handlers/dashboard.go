package handlers

import (
	"context"
	"log"
	"net/http"

	"auth-athlevo/genproto/auth"

	"github.com/gin-gonic/gin"
)

// GetPersonalAccessCount godoc
// @Summary Get the access count of a gym
// @Description Get the number of times a gym accessed the gym within the specified date range
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param gym_id query string true "gym ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} auth.AccessCountRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/access-count [get]
func (h *Handlers) GetPersonalAccessCount(c *gin.Context) {
	gymID := c.Query("gym_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if gymID == "" || startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
		return
	}

	req := auth.AccessCountReq{
		GymId:     gymID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	res, err := h.Dashboard.GetDailyPersonalAccessCount(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get personal access count: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTotalPersonalBookingRevenue godoc
// @Summary Get total booking revenue for a gym
// @Description Get the total booking revenue within the specified date range for a gym
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param gym_id query string true "gym ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} auth.BookingRevenueRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/booking-revenue [get]
func (h *Handlers) GetTotalPersonalBookingRevenue(c *gin.Context) {
	gymId := c.Query("gym_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if gymId == "" || startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
		return
	}

	req := auth.BookingRevenueReq{
		GymId:     gymId,
		StartDate: startDate,
		EndDate:   endDate,
	}

	res, err := h.Dashboard.GetDailyPersonalBookingRevenueByDay(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get booking revenue: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAccessCountBySubscriptionID godoc
// @Summary Get the access count by subscription ID
// @Description Get the number of times a user accessed the gym based on subscription ID within a date range
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param subscription_id query string true "Subscription ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} auth.SubscriptionCountRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/subscription/access-count [get]
func (h *Handlers) GetAccessCountBySubscriptionID(c *gin.Context) {
	subscriptionID := c.Query("subscription_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if subscriptionID == "" || startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
		return
	}

	req := auth.SubscriptionCountReq{
		SubscriptionID: subscriptionID,
		StartDate:      startDate,
		EndDate:        endDate,
	}

	res, err := h.Dashboard.GetDailyAccessCountBySubscriptionID(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get access count by subscription: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetBookingRevenueBySubscriptionID godoc
// @Summary Get booking revenue by subscription ID
// @Description Get total booking revenue based on subscription ID within a date range
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param subscription_id query string true "Subscription ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} auth.BookingRevenueBySubscriptionRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/subscription/booking-revenue [get]
func (h *Handlers) GetBookingRevenueBySubscriptionID(c *gin.Context) {
	subscriptionID := c.Query("subscription_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if subscriptionID == "" || startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
		return
	}

	req := auth.BookingRevenueBySubscriptionReq{
		SubscriptionID: subscriptionID,
		StartDate:      startDate,
		EndDate:        endDate,
	}

	res, err := h.Dashboard.GetDailyBookingRevenueBySubscriptionID(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get booking revenue by subscription: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
