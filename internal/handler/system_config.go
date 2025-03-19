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

// GetConfig godoc
// @Summary      获取系统配置
// @Description  根据key获取系统配置
// @Tags         系统配置
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        key  path      string  true  "配置key"
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /system/config/{key} [get]
func (h *SystemConfigHandler) GetConfig(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "配置键不能为空"})
		return
	}

	config, err := h.service.GetByKey(key)
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

// CreateConfig godoc
// @Summary      创建系统配置
// @Description  创建新的系统配置项
// @Tags         系统配置
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        config  body      model.CreateConfigRequest  true  "配置信息"
// @Success      200     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]interface{}
// @Failure      401     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /system/config [post]
func (h *SystemConfigHandler) CreateConfig(c *gin.Context) {
	var req model.CreateConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config, err := h.service.Create(req.Key, req.Value, req.Type, req.Remark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"data":    config,
	})
}

// UpdateConfig godoc
// @Summary      更新系统配置
// @Description  更新系统配置信息
// @Tags         系统配置
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        key     path      string                    true  "配置key"
// @Param        config  body      model.UpdateConfigRequest  true  "配置信息"
// @Success      200     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]interface{}
// @Failure      401     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /system/config/{key} [put]
func (h *SystemConfigHandler) UpdateConfig(c *gin.Context) {
	key := c.Param("key")
	var req model.UpdateConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config, err := h.service.Update(key, req.Value, req.Type, req.Remark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"data":    config,
	})
}

// DeleteConfig godoc
// @Summary      删除系统配置
// @Description  删除指定的系统配置
// @Tags         系统配置
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        key  path      string  true  "配置key"
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /system/config/{key} [delete]
func (h *SystemConfigHandler) DeleteConfig(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "配置键不能为空"})
		return
	}

	if err := h.service.Delete(key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// ListConfigs godoc
// @Summary      获取所有系统配置
// @Description  获取所有系统配置列表
// @Tags         系统配置
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /system/configs [get]
func (h *SystemConfigHandler) ListConfigs(c *gin.Context) {
	configs, err := h.service.List()
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

	value, err := h.service.GetValue(key)
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

	if err := h.service.UpdateValue(key, req.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}
