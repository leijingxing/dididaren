package handler

import (
	"dididaren/internal/model"
	"dididaren/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DangerZoneHandler struct {
	service *service.DangerZoneService
}

func NewDangerZoneHandler(service *service.DangerZoneService) *DangerZoneHandler {
	return &DangerZoneHandler{
		service: service,
	}
}

// CreateDangerZone 创建危险区域
func (h *DangerZoneHandler) CreateDangerZone(c *gin.Context) {
	var zone model.DangerZone
	if err := c.ShouldBindJSON(&zone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateDangerZone(&zone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功", "data": zone})
}

// GetDangerZone 获取危险区域信息
func (h *DangerZoneHandler) GetDangerZone(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	zone, err := h.service.GetDangerZone(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if zone == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该危险区域"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": zone})
}

// UpdateDangerZone 更新危险区域信息
func (h *DangerZoneHandler) UpdateDangerZone(c *gin.Context) {
	var zone model.DangerZone
	if err := c.ShouldBindJSON(&zone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateDangerZone(&zone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "data": zone})
}

// DeleteDangerZone 删除危险区域
func (h *DangerZoneHandler) DeleteDangerZone(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := h.service.DeleteDangerZone(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetNearbyZones 获取附近的危险区域
func (h *DangerZoneHandler) GetNearbyZones(c *gin.Context) {
	latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的纬度"})
		return
	}

	longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的经度"})
		return
	}

	radius, err := strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的半径"})
		return
	}

	zones, err := h.service.GetNearbyZones(latitude, longitude, radius)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": zones})
}

// GetAllActiveZones 获取所有活跃的危险区域
func (h *DangerZoneHandler) GetAllActiveZones(c *gin.Context) {
	zones, err := h.service.GetAllActiveZones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": zones})
}

// UpdateHeatLevel 更新危险区域的热度等级
func (h *DangerZoneHandler) UpdateHeatLevel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	heatLevel, err := strconv.Atoi(c.Query("heat_level"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的热度等级"})
		return
	}

	if err := h.service.UpdateHeatLevel(uint(id), heatLevel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}
