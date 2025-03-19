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
	return &DangerZoneHandler{service: service}
}

// CreateDangerZone 创建危险区域
// @Summary 创建危险区域
// @Description 创建一个新的危险区域
// @Tags 危险区域
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body model.CreateDangerZoneRequest true "危险区域信息"
// @Success 200 {object} model.DangerZone
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/danger-zones [post]
func (h *DangerZoneHandler) CreateDangerZone(c *gin.Context) {
	var req model.CreateDangerZoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateDangerZone(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// GetDangerZone 获取危险区域详情
// @Summary 获取危险区域详情
// @Description 根据ID获取危险区域详情
// @Tags 危险区域
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "危险区域ID"
// @Success 200 {object} model.DangerZone
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/danger-zones/{id} [get]
func (h *DangerZoneHandler) GetDangerZone(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	zone, err := h.service.GetDangerZoneByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, zone)
}

// ListDangerZones 获取危险区域列表
// @Summary 获取危险区域列表
// @Description 获取所有危险区域列表，支持分页
// @Tags 危险区域
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码，默认1"
// @Param size query int false "每页数量，默认10"
// @Success 200 {array} model.DangerZone
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/danger-zones [get]
func (h *DangerZoneHandler) ListDangerZones(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	zones, total, err := h.service.ListDangerZones(page, size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": zones,
	})
}

// UpdateDangerZone 更新危险区域
// @Summary 更新危险区域
// @Description 更新危险区域信息
// @Tags 危险区域
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "危险区域ID"
// @Param request body model.UpdateDangerZoneRequest true "危险区域信息"
// @Success 200 {object} model.DangerZone
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/danger-zones/{id} [put]
func (h *DangerZoneHandler) UpdateDangerZone(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var req model.UpdateDangerZoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateDangerZone(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// DeleteDangerZone 删除危险区域
// @Summary 删除危险区域
// @Description 删除指定的危险区域
// @Tags 危险区域
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "危险区域ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/danger-zones/{id} [delete]
func (h *DangerZoneHandler) DeleteDangerZone(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	err = h.service.DeleteDangerZone(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// CheckLocation 检查位置是否在危险区域内
// @Summary 检查位置是否在危险区域内
// @Description 检查指定位置是否在任何危险区域内
// @Tags 危险区域
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param latitude query number true "纬度"
// @Param longitude query number true "经度"
// @Success 200 {object} CheckLocationResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/danger-zones/check [get]
func (h *DangerZoneHandler) CheckLocation(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的纬度"})
		return
	}

	lng, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的经度"})
		return
	}

	inDangerZone, zones, err := h.service.CheckLocationInDangerZone(lat, lng)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"in_danger_zone": inDangerZone,
		"danger_zones":   zones,
	})
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
