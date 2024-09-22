package handlers

import (
    "context"
    "log"
    "net/http"

    "auth-athlevo/genproto/auth"

    "github.com/gin-gonic/gin"
)

// GetPersonalAccessCount godoc
// @Summary Get the access count of a user
// @Description Get the number of times a user accessed the gym within the specified date range
// @Tags dashboard
// @Accept json
// @Produce json
// @Param req body auth.AccessCountReq true "Access Count Request"
// @Success 200 {object} auth.AccessCountRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/access-count [get]
func (h *Handlers) GetPersonalAccessCount(c *gin.Context) {
    var req auth.AccessCountReq
    if err := c.BindJSON(&req); err != nil {
        log.Printf("failed to bind JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    res, err := h.Dashboard.GetPersonalAccessCount(context.Background(), &req)
    if err != nil {
        log.Printf("failed to get personal access count: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, res)
}

// GetTotalPersonalBookingRevenue godoc
// @Summary Get total booking revenue for a user
// @Description Get the total booking revenue within the specified date range for a user
// @Tags dashboard
// @Accept json
// @Produce json
// @Param req body auth.TotalRevenueReq true "Total Booking Revenue Request"
// @Success 200 {object} auth.TotalRevenueRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/booking-revenue [get]
func (h *Handlers) GetTotalPersonalBookingRevenue(c *gin.Context) {
    var req auth.TotalRevenueReq
    if err := c.BindJSON(&req); err != nil {
        log.Printf("failed to bind JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    res, err := h.Dashboard.GetTotalPersonalBookingRevenue(context.Background(), &req)
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
// @Param req body auth.SubscriptionCountReq true "Subscription Access Count Request"
// @Success 200 {object} auth.SubscriptionCountRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/subscription/access-count [get]
func (h *Handlers) GetAccessCountBySubscriptionID(c *gin.Context) {
    var req auth.SubscriptionCountReq
    if err := c.BindJSON(&req); err != nil {
        log.Printf("failed to bind JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    res, err := h.Dashboard.GetAccessCountBySubscriptionID(context.Background(), &req)
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
// @Param req body auth.TotalRevenueBySubscriptionReq true "Total Booking Revenue By Subscription Request"
// @Success 200 {object} auth.TotalRevenueBySubscriptionRes
// @Failure 400 {object} string "Invalid Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dashboard/subscription/booking-revenue [get]
func (h *Handlers) GetBookingRevenueBySubscriptionID(c *gin.Context) {
    var req auth.TotalRevenueBySubscriptionReq
    if err := c.BindJSON(&req); err != nil {
        log.Printf("failed to bind JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    res, err := h.Dashboard.GetBookingRevenueBySubscriptionID(context.Background(), &req)
    if err != nil {
        log.Printf("failed to get booking revenue by subscription: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, res)
}
