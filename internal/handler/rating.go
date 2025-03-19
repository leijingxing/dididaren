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
	return &RatingHandler{service: service}
}

// CreateRating 创建评价
// @Summary 创建评价
// @Description 创建一条新的评价
// @Tags 评价
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body model.CreateRatingRequest true "评价信息"
// @Success 200 {object} model.Rating
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/ratings [post]
func (h *RatingHandler) CreateRating(c *gin.Context) {
	var req model.CreateRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating, err := h.service.CreateRating(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rating)
}

// GetRating 获取评价详情
// @Summary 获取评价详情
// @Description 根据ID获取评价详情
// @Tags 评价
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "评价ID"
// @Success 200 {object} model.Rating
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/ratings/{id} [get]
func (h *RatingHandler) GetRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	rating, err := h.service.GetRatingByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rating)
}

// ListRatings 获取评价列表
// @Summary 获取评价列表
// @Description 获取安保人员的所有评价
// @Tags 评价
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param staff_id query int true "安保人员ID"
// @Success 200 {array} model.Rating
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/ratings [get]
func (h *RatingHandler) ListRatings(c *gin.Context) {
	staffID, err := strconv.ParseUint(c.Query("staff_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的安保人员ID"})
		return
	}

	ratings, err := h.service.ListRatings(uint(staffID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ratings)
}

// UpdateRating 更新评价
// @Summary 更新评价
// @Description 更新评价信息
// @Tags 评价
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "评价ID"
// @Param request body model.CreateRatingRequest true "评价信息"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/ratings/{id} [put]
func (h *RatingHandler) UpdateRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var req model.CreateRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateRating(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// DeleteRating 删除评价
// @Summary 删除评价
// @Description 删除指定的评价
// @Tags 评价
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "评价ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/ratings/{id} [delete]
func (h *RatingHandler) DeleteRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	err = h.service.DeleteRating(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
