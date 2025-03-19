package handler

import (
	"dididaren/internal/model"
	"dididaren/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmergencyHandler struct {
	service *service.EmergencyService
}

func NewEmergencyHandler(service *service.EmergencyService) *EmergencyHandler {
	return &EmergencyHandler{
		service: service,
	}
}

// Create 创建紧急事件
func (h *EmergencyHandler) Create(c *gin.Context) {
	var req model.CreateEmergencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserID = c.GetUint("user_id")
	event, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"data":    event,
	})
}

// GetByID 获取事件详情
func (h *EmergencyHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	event, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// UpdateStatus 更新事件状态
func (h *EmergencyHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStatus(id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// GetHistory 获取事件历史
func (h *EmergencyHandler) GetHistory(c *gin.Context) {
	userID := c.GetUint("user_id")
	events, err := h.service.GetHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

// CreateHandlingRecord 创建处理记录
func (h *EmergencyHandler) CreateHandlingRecord(c *gin.Context) {
	var req model.CreateHandlingRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := h.service.CreateHandlingRecord(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"data":    record,
	})
}

// GetHandlingRecords 获取处理记录
func (h *EmergencyHandler) GetHandlingRecords(c *gin.Context) {
	eventID := c.Param("id")
	records, err := h.service.GetHandlingRecords(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": records})
}
