package handler

import (
	"dididaren/internal/model"
	"dididaren/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmergencyHandler struct {
	service *service.EmergencyService
}

func NewEmergencyHandler(service *service.EmergencyService) *EmergencyHandler {
	return &EmergencyHandler{service: service}
}

// CreateEmergency 创建紧急事件
// @Summary 创建紧急事件
// @Description 创建一个新的紧急事件
// @Tags 紧急事件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body model.CreateEmergencyRequest true "紧急事件信息"
// @Success 200 {object} model.Emergency
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/emergencies [post]
func (h *EmergencyHandler) CreateEmergency(c *gin.Context) {
	var req model.CreateEmergencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	emergency, err := h.service.CreateEmergency(userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emergency)
}

// GetEmergency 获取紧急事件详情
// @Summary 获取紧急事件详情
// @Description 根据ID获取紧急事件详情
// @Tags 紧急事件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "紧急事件ID"
// @Success 200 {object} model.Emergency
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/emergencies/{id} [get]
func (h *EmergencyHandler) GetEmergency(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	emergency, err := h.service.GetEmergencyByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emergency)
}

// ListEmergencies 获取紧急事件列表
// @Summary 获取紧急事件列表
// @Description 获取所有紧急事件列表，支持分页和状态筛选
// @Tags 紧急事件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码，默认1"
// @Param size query int false "每页数量，默认10"
// @Param status query string false "状态筛选"
// @Success 200 {array} model.Emergency
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/emergencies [get]
func (h *EmergencyHandler) ListEmergencies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	status := c.Query("status")

	emergencies, total, err := h.service.ListEmergencies(page, size, status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": emergencies,
	})
}

// CreateHandlingRecord 创建处理记录
// @Summary 创建处理记录
// @Description 为紧急事件创建一条处理记录
// @Tags 紧急事件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "紧急事件ID"
// @Param request body model.CreateHandlingRecordRequest true "处理记录信息"
// @Success 200 {object} model.HandlingRecord
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/emergencies/{id}/records [post]
func (h *EmergencyHandler) CreateHandlingRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var req model.CreateHandlingRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staffID := c.GetUint("user_id")
	record, err := h.service.CreateHandlingRecord(&req, uint(id), staffID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// ListHandlingRecords 获取处理记录列表
// @Summary 获取处理记录列表
// @Description 获取紧急事件的所有处理记录
// @Tags 紧急事件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "紧急事件ID"
// @Success 200 {array} model.HandlingRecord
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/emergencies/{id}/records [get]
func (h *EmergencyHandler) ListHandlingRecords(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	records, err := h.service.ListHandlingRecords(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// AssignStaff 分配安保人员
// @Summary 分配安保人员
// @Description 为紧急事件分配安保人员
// @Tags 紧急事件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "紧急事件ID"
// @Param staff_id query int true "安保人员ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/emergencies/{id}/assign [post]
func (h *EmergencyHandler) AssignStaff(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的事件ID"})
		return
	}

	staffID, err := strconv.ParseUint(c.Query("staff_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的安保人员ID"})
		return
	}

	err = h.service.AssignStaff(uint(id), uint(staffID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "分配成功"})
}

// CompleteEmergency 完成紧急事件
// @Summary 完成紧急事件
// @Description 标记紧急事件为已完成状态
// @Tags 紧急事件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "紧急事件ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/emergencies/{id}/complete [post]
func (h *EmergencyHandler) CompleteEmergency(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	err = h.service.CompleteEmergency(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "完成成功"})
}
