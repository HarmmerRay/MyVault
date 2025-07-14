package handlers

import (
	"net/http"
	"strconv"
	"myvault-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	activityService ActivityService
}

type ActivityService interface {
	GetUserActivities(userID uint, limit int, offset int) ([]models.Activity, error)
	GetActivityByID(userID, activityID uint) (*models.Activity, error)
	SyncActivities(userID uint, force bool) error
	GetTodayActivity(userID uint) (*models.Activity, error)
}

func NewActivityHandler(activityService ActivityService) *ActivityHandler {
	return &ActivityHandler{
		activityService: activityService,
	}
}

func (h *ActivityHandler) GetActivities(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// 获取查询参数
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	activities, err := h.activityService.GetUserActivities(userID.(uint), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get activities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"activities": activities,
		"total":      len(activities),
	})
}

func (h *ActivityHandler) GetActivity(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseUint(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	activity, err := h.activityService.GetActivityByID(userID.(uint), uint(activityID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activity": activity})
}

func (h *ActivityHandler) SyncActivities(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.SyncRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.activityService.SyncActivities(userID.(uint), req.Force)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync activities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activities synced successfully"})
}

func (h *ActivityHandler) GetTodayActivity(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	activity, err := h.activityService.GetTodayActivity(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get today's activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activity": activity})
}