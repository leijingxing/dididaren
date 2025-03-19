package handler

import (
	"dididaren/internal/service"
	"dididaren/pkg/response"

	"github.com/gin-gonic/gin"
)

type SecurityHandler struct {
	securityService *service.SecurityService
}

func NewSecurityHandler(securityService *service.SecurityService) *SecurityHandler {
	return &SecurityHandler{
		securityService: securityService,
	}
}

// ApplySecurityStaff 申请成为安保人员
func (h *SecurityHandler) ApplySecurityStaff(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req struct {
		CompanyName   string   `json:"company_name" binding:"required"`
		LicenseNumber string   `json:"license_number" binding:"required"`
		CertFiles     []string `json:"cert_files" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	staff, err := h.securityService.ApplySecurityStaff(userID, req.CompanyName, req.LicenseNumber, req.CertFiles)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, staff)
}

// UpdateLocation 更新位置信息
func (h *SecurityHandler) UpdateLocation(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req struct {
		LocationLat  float64 `json:"location_lat" binding:"required"`
		LocationLng  float64 `json:"location_lng" binding:"required"`
		OnlineStatus int8    `json:"online_status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	err := h.securityService.UpdateLocation(userID, req.LocationLat, req.LocationLng, req.OnlineStatus)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// AcceptEvent 接受事件
func (h *SecurityHandler) AcceptEvent(c *gin.Context) {
	userID := c.GetUint("user_id")
	eventID := c.Param("id")
	err := h.securityService.AcceptEvent(userID, eventID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// CompleteEvent 完成事件
func (h *SecurityHandler) CompleteEvent(c *gin.Context) {
	userID := c.GetUint("user_id")
	eventID := c.Param("id")
	var req struct {
		Remark string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	err := h.securityService.CompleteEvent(userID, eventID, req.Remark)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// GetStaffInfo 获取安保人员信息
func (h *SecurityHandler) GetStaffInfo(c *gin.Context) {
	userID := c.GetUint("user_id")
	staff, err := h.securityService.GetStaffInfo(userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, staff)
}
