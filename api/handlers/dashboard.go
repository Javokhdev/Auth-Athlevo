package handlers

import (
	"context"
	"log"
	"net/http"

	"auth-athlevo/genproto/auth"

	"github.com/gin-gonic/gin"
)

// // GetPersonalAccessCount godoc
// // @Summary Get the access count of a gym
// // @Description Get the number of times a gym accessed the gym within the specified date range
// // @Tags dashboard
// // @Accept json
// // @Produce json
// // @Security BearerAuth
// // @Param gym_id query string true "gym ID"
// // @Param start_date query string true "Start Date"
// // @Param end_date query string true "End Date"
// // @Success 200 {object} auth.AccessCountRes
// // @Failure 400 {object} string "Invalid Request"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /dashboard/access-count [get]
// func (h *Handlers) GetPersonalAccessCount(c *gin.Context) {
// 	gymID := c.Query("gym_id")
// 	startDate := c.Query("start_date")
// 	endDate := c.Query("end_date")

// 	if gymID == "" || startDate == "" || endDate == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
// 		return
// 	}

// 	req := auth.AccessCountReq{
// 		GymId:     gymID,
// 		StartDate: startDate,
// 		EndDate:   endDate,
// 	}

// 	res, err := h.Dashboard.GetDailyPersonalAccessCount(context.Background(), &req)
// 	if err != nil {
// 		log.Printf("failed to get personal access count: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// // GetTotalPersonalBookingRevenue godoc
// // @Summary Get total booking revenue for a gym
// // @Description Get the total booking revenue within the specified date range for a gym
// // @Tags dashboard
// // @Accept json
// // @Produce json
// // @Security BearerAuth
// // @Param gym_id query string true "gym ID"
// // @Param start_date query string true "Start Date"
// // @Param end_date query string true "End Date"
// // @Success 200 {object} auth.BookingRevenueRes
// // @Failure 400 {object} string "Invalid Request"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /dashboard/booking-revenue [get]
// func (h *Handlers) GetTotalPersonalBookingRevenue(c *gin.Context) {
// 	gymId := c.Query("gym_id")
// 	startDate := c.Query("start_date")
// 	endDate := c.Query("end_date")

// 	if gymId == "" || startDate == "" || endDate == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
// 		return
// 	}

// 	req := auth.BookingRevenueReq{
// 		GymId:     gymId,
// 		StartDate: startDate,
// 		EndDate:   endDate,
// 	}

// 	res, err := h.Dashboard.GetDailyPersonalBookingRevenueByDay(context.Background(), &req)
// 	if err != nil {
// 		log.Printf("failed to get booking revenue: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// // GetAccessCountBySubscriptionID godoc
// // @Summary Get the access count by subscription ID
// // @Description Get the number of times a user accessed the gym based on subscription ID within a date range
// // @Tags dashboard
// // @Accept json
// // @Produce json
// // @Security BearerAuth
// // @Param subscription_id query string true "Subscription ID"
// // @Param start_date query string true "Start Date"
// // @Param end_date query string true "End Date"
// // @Success 200 {object} auth.SubscriptionCountRes
// // @Failure 400 {object} string "Invalid Request"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /dashboard/subscription/access-count [get]
// func (h *Handlers) GetAccessCountBySubscriptionID(c *gin.Context) {
// 	subscriptionID := c.Query("subscription_id")
// 	startDate := c.Query("start_date")
// 	endDate := c.Query("end_date")

// 	if subscriptionID == "" || startDate == "" || endDate == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
// 		return
// 	}

// 	req := auth.SubscriptionCountReq{
// 		SubscriptionID: subscriptionID,
// 		StartDate:      startDate,
// 		EndDate:        endDate,
// 	}

// 	res, err := h.Dashboard.GetDailyAccessCountBySubscriptionID(context.Background(), &req)
// 	if err != nil {
// 		log.Printf("failed to get access count by subscription: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// // GetBookingRevenueBySubscriptionID godoc
// // @Summary Get booking revenue by subscription ID
// // @Description Get total booking revenue based on subscription ID within a date range
// // @Tags dashboard
// // @Accept json
// // @Produce json
// // @Security BearerAuth
// // @Param subscription_id query string true "Subscription ID"
// // @Param start_date query string true "Start Date"
// // @Param end_date query string true "End Date"
// // @Success 200 {object} auth.BookingRevenueBySubscriptionRes
// // @Failure 400 {object} string "Invalid Request"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /dashboard/subscription/booking-revenue [get]
// func (h *Handlers) GetBookingRevenueBySubscriptionID(c *gin.Context) {
// 	subscriptionID := c.Query("subscription_id")
// 	startDate := c.Query("start_date")
// 	endDate := c.Query("end_date")

// 	if subscriptionID == "" || startDate == "" || endDate == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
// 		return
// 	}

// 	req := auth.BookingRevenueBySubscriptionReq{
// 		SubscriptionID: subscriptionID,
// 		StartDate:      startDate,
// 		EndDate:        endDate,
// 	}

// 	res, err := h.Dashboard.GetDailyBookingRevenueBySubscriptionID(context.Background(), &req)
// 	if err != nil {
// 		log.Printf("failed to get booking revenue by subscription: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// TotalMen godoc
// @Summary Get the total number of men in a gym
// @Description Get the total count of male members for a specified gym
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param gym_id query string true "Gym ID"
// @Success 200 {object} auth.TotalMenRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/total-men [get]
func (h *Handlers) TotalMen(c *gin.Context) {
	gymID := c.Query("gym_id")

	if gymID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required gym_id query parameter"})
		return
	}

	req := auth.TotalMenReq{
		GymId: gymID,
	}

	res, err := h.Dashboard.TotalMen(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get total men: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// TotalWomen godoc
// @Summary Get the total number of women in a gym
// @Description Get the total count of female members for a specified gym
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param gym_id query string true "Gym ID"
// @Success 200 {object} auth.TotalWomenRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/total-women [get]
func (h *Handlers) TotalWomen(c *gin.Context) {
	gymID := c.Query("gym_id")

	if gymID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required gym_id query parameter"})
		return
	}

	req := auth.TotalWomenReq{
		GymId: gymID,
	}

	res, err := h.Dashboard.TotalWomen(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get total women: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}


// TotalMembers godoc
// @Summary Get the total number of members in a gym
// @Description Get the total count of members for a specified gym
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param gym_id query string true "Gym ID"
// @Success 200 {object} auth.TotalMembersRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/total-members [get]
func (h *Handlers) TotalMembers(c *gin.Context) {
	gymID := c.Query("gym_id")

	if gymID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required gym_id query parameter"})
		return
	}

	req := auth.TotalMembersReq{
		GymId: gymID,
	}

	res, err := h.Dashboard.TotalMembers(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get total members: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}


// TotalAmount godoc
// @Summary Get the total revenue amount for a gym
// @Description Get the total revenue amount for a specified gym
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param gym_id query string true "Gym ID"
// @Success 200 {object} auth.TotalAmountRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/total-amount [get]
func (h *Handlers) TotalAmount(c *gin.Context) {
	gymID := c.Query("gym_id")

	if gymID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required gym_id query parameter"})
		return
	}

	req := auth.TotalAmountReq{
		GymId: gymID,
	}

	res, err := h.Dashboard.TotalAmount(context.Background(), &req)
	if err != nil {
		log.Printf("failed to get total amount: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CompareCurrentAndPreviousMonthRevenue godoc
// @Summary Compare current and previous
// @Description Compare current and previous
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.RevenueReq
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/compare-month-revenue [get]
func (h *Handlers) CompareCurrentAndPreviousMonthRevenue(c *gin.Context) {

	res, err := h.Dashboard.CompareCurrentAndPreviousMonthRevenue(context.Background(), &auth.Void{})
	if err != nil {
		log.Printf("failed to compare current and previous", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetMonthlyRevenueStats godoc
// @Summary Get monthly revenue
// @Description Get monthly revenue
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.MonthlyRevenueRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/monthly-renue-stats [get]
func (h *Handlers) GetMonthlyRevenueStats(c *gin.Context) {

	res, err := h.Dashboard.GetMonthlyRevenueStats(context.Background(), &auth.Void{})
	if err != nil {
		log.Printf("failed to get monthly revenue", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetGenderCounts godoc
// @Summary Get the total number of males and females in a gym within a specified date range
// @Description Get the total count of gender members for a specified gym and date range
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param gym_id query string true "Gym ID"
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Success 200 {object} auth.GenderCountsRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/gender [get]
func (h *Handlers) GetGenderCounts(c *gin.Context) {
    gymID := c.Query("gym_id")
    startDate := c.Query("start_date")
    endDate := c.Query("end_date")

    if gymID == "" || startDate == "" || endDate == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
        return
    }

    req := auth.TotalGenderReq{
        GymId:     gymID,
        StartDate: startDate,
        EndDate:   endDate,
    }

    res, err := h.Dashboard.GetGenderCounts(context.Background(), &req)
    if err != nil {
        log.Printf("failed to get total gender: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, res)
}

// GetRevenueByTariff godoc
// @Summary Get the revenue by Tariff
// @Description Get the revenue by Tariff
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.TariffRevenueRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/revenue-by-subscription [get]
func (h *Handlers) GetRevenueByTariff(c *gin.Context) {

    res, err := h.Dashboard.GetRevenueByTariff(context.Background(), &auth.Void{})
    if err != nil {
        log.Printf("failed to get revenue by tariff: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, res)
}

// GetUsersByTariff godoc
// @Summary Get users by tariff
// @Description Get users by tariff
// @Tags dashboard
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.TariffUsersRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/users-by-subscription [get]
func (h *Handlers) GetUsersByTariff(c *gin.Context) {

    res, err := h.Dashboard.GetUsersByTariff(context.Background(), &auth.Void{})
    if err != nil {
        log.Printf("failed to get users by tariff: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, res)
}
