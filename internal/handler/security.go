package handler

import (
	"dididaren/internal/model"
	"dididaren/internal/service"
	"dididaren/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SecurityHandler struct {
	service *service.SecurityService
}

func NewSecurityHandler(service *service.SecurityService) *SecurityHandler {
	return &SecurityHandler{
		service: service,
	}
}

// CreateStaff godoc
// @Summary      创建安保人员
// @Description  创建新的安保人员
// @Tags         安保人员
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        staff  body      model.CreateStaffRequest  true  "安保人员信息"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]interface{}
// @Failure      401    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /security/staff [post]
func (h *SecurityHandler) CreateStaff(c *gin.Context) {
	var req model.CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := h.service.CreateStaff(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"data":    staff,
	})
}

// GetStaff godoc
// @Summary      获取安保人员
// @Description  根据ID获取安保人员详情
// @Tags         安保人员
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      uint  true  "安保人员ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /security/staff/{id} [get]
func (h *SecurityHandler) GetStaff(c *gin.Context) {
	id := c.Param("id")
	staff, err := h.service.GetStaffByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": staff})
}

// ListStaffs godoc
// @Summary      获取安保人员列表
// @Description  获取安保人员列表，支持分页和筛选
// @Tags         安保人员
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page     query     int     false  "页码"  default(1)
// @Param        size     query     int     false  "每页数量"  default(10)
// @Param        status   query     string  false  "状态"
// @Success      200      {object}  map[string]interface{}
// @Failure      401      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /security/staff [get]
func (h *SecurityHandler) ListStaffs(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	status := c.Query("status")

	staffs, total, err := h.service.ListStaffs(page, size, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":  staffs,
			"total": total,
		},
	})
}

// UpdateStaffStatus godoc
// @Summary      更新安保人员状态
// @Description  更新安保人员的工作状态
// @Tags         安保人员
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      string  true  "安保人员ID"
// @Param        status query     string  true  "状态"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]interface{}
// @Failure      401    {object}  map[string]interface{}
// @Failure      404    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /security/staff/{id}/status [put]
func (h *SecurityHandler) UpdateStaffStatus(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")

	if err := h.service.UpdateStaffStatus(id, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// CreateRating godoc
// @Summary      创建评价
// @Description  为安保人员创建评价
// @Tags         安保人员
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        rating  body      model.CreateRatingRequest  true  "评价信息"
// @Success      200     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]interface{}
// @Failure      401     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /security/rating [post]
func (h *SecurityHandler) CreateRating(c *gin.Context) {
	var req model.CreateRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating, err := h.service.CreateRating(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"data":    rating,
	})
}

// ListRatings godoc
// @Summary      获取评价列表
// @Description  获取安保人员的评价列表
// @Tags         安保人员
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        staff_id  query     uint  true  "安保人员ID"
// @Success      200       {object}  map[string]interface{}
// @Failure      401       {object}  map[string]interface{}
// @Failure      500       {object}  map[string]interface{}
// @Router       /security/rating [get]
func (h *SecurityHandler) ListRatings(c *gin.Context) {
	staffID := c.Query("staff_id")
	ratings, err := h.service.ListRatings(staffID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// ApplySecurityStaff 申请成为安保人员
func (h *SecurityHandler) ApplySecurityStaff(c *gin.Context) {
	var req struct {
		Name   string `json:"name" binding:"required"`
		Phone  string `json:"phone" binding:"required"`
		IDCard string `json:"id_card" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID := c.GetUint("user_id")
	err := h.service.ApplySecurityStaff(userID, req.Name, req.Phone, req.IDCard)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// UpdateLocation 更新位置
func (h *SecurityHandler) UpdateLocation(c *gin.Context) {
	var req struct {
		Lat float64 `json:"lat" binding:"required"`
		Lng float64 `json:"lng" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID := c.GetUint("user_id")
	err := h.service.UpdateLocation(userID, req.Lat, req.Lng)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// AcceptEvent 接受事件
func (h *SecurityHandler) AcceptEvent(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID := c.GetUint("user_id")
	err = h.service.AcceptEvent(userID, uint(eventID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// CompleteEvent 完成事件
func (h *SecurityHandler) CompleteEvent(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID := c.GetUint("user_id")
	err = h.service.CompleteEvent(userID, uint(eventID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}

// GetStaffInfo 获取安保人员信息
func (h *SecurityHandler) GetStaffInfo(c *gin.Context) {
	userID := c.GetUint("user_id")
	staff, err := h.service.GetStaffInfo(userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, staff)
}
