package handler

import (
	"dididaren/internal/model"
	"dididaren/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingHandler struct {
	service *service.RatingService
}

func NewRatingHandler(service *service.RatingService) *RatingHandler {
	return &RatingHandler{
		service: service,
	}
}

// CreateRating 创建评价
func (h *RatingHandler) CreateRating(c *gin.Context) {
	var rating model.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功", "data": rating})
}

// GetRating 获取评价信息
func (h *RatingHandler) GetRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	rating, err := h.service.GetRating(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rating == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该评价"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetEventRating 获取事件相关的评价
func (h *RatingHandler) GetEventRating(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("event_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的事件ID"})
		return
	}

	rating, err := h.service.GetEventRating(uint(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rating == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该事件的评价"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetStaffRatings 获取安保人员的所有评价
func (h *RatingHandler) GetStaffRatings(c *gin.Context) {
	staffID, err := strconv.ParseUint(c.Param("staff_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的安保人员ID"})
		return
	}

	ratings, err := h.service.GetStaffRatings(uint(staffID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// GetStaffAverageRating 获取安保人员的平均评分
func (h *RatingHandler) GetStaffAverageRating(c *gin.Context) {
	staffID, err := strconv.ParseUint(c.Param("staff_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的安保人员ID"})
		return
	}

	avgRating, err := h.service.GetStaffAverageRating(uint(staffID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"average_rating": avgRating}})
}

// UpdateRating 更新评价
func (h *RatingHandler) UpdateRating(c *gin.Context) {
	var rating model.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "data": rating})
}

// DeleteRating 删除评价
func (h *RatingHandler) DeleteRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.service.DeleteRating(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
