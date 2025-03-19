package handler

import (
	"dididaren/internal/model"
	"dididaren/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemConfigHandler struct {
	service *service.SystemConfigService
}

func NewSystemConfigHandler(service *service.SystemConfigService) *SystemConfigHandler {
	return &SystemConfigHandler{
		service: service,
	}
}

// GetConfig 获取配置信息
func (h *SystemConfigHandler) GetConfig(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "配置键不能为空"})
		return
	}

	config, err := h.service.GetConfig(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if config == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该配置"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": config})
}

// CreateConfig 创建配置
func (h *SystemConfigHandler) CreateConfig(c *gin.Context) {
	var config model.SystemConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateConfig(&config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功", "data": config})
}

// UpdateConfig 更新配置
func (h *SystemConfigHandler) UpdateConfig(c *gin.Context) {
	var config model.SystemConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateConfig(&config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "data": config})
}

// DeleteConfig 删除配置
func (h *SystemConfigHandler) DeleteConfig(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "配置键不能为空"})
		return
	}

	if err := h.service.DeleteConfig(key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetAllConfigs 获取所有配置
func (h *SystemConfigHandler) GetAllConfigs(c *gin.Context) {
	configs, err := h.service.GetAllConfigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": configs})
}

// GetConfigValue 获取配置值
func (h *SystemConfigHandler) GetConfigValue(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "配置键不能为空"})
		return
	}

	value, err := h.service.GetConfigValue(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"value": value}})
}

// UpdateConfigValue 更新配置值
func (h *SystemConfigHandler) UpdateConfigValue(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "配置键不能为空"})
		return
	}

	var req struct {
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateConfigValue(key, req.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}
