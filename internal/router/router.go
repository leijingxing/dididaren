package router

import (
	"dididaren/internal/handler"
	"dididaren/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handler.UserHandler,
	emergencyHandler *handler.EmergencyHandler,
	securityHandler *handler.SecurityHandler,
	dangerZoneHandler *handler.DangerZoneHandler,
	ratingHandler *handler.RatingHandler,
	systemConfigHandler *handler.SystemConfigHandler,
) *gin.Engine {
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// 公开路由
	public := r.Group("/api/v1")
	{
		// 用户相关
		public.POST("/users/register", userHandler.Register)
		public.POST("/users/login", userHandler.Login)
		public.POST("/users/verify-code", userHandler.VerifyCode)

		// 系统配置
		public.GET("/configs", systemConfigHandler.GetAllConfigs)
		public.GET("/configs/:key", systemConfigHandler.GetConfigValue)
	}

	// 需要认证的路由
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.Auth())
	{
		// 用户相关
		authorized.GET("/users/profile", userHandler.GetProfile)
		authorized.PUT("/users/profile", userHandler.UpdateProfile)
		authorized.POST("/users/emergency-contacts", userHandler.AddEmergencyContact)
		authorized.GET("/users/emergency-contacts", userHandler.GetEmergencyContacts)
		authorized.DELETE("/users/emergency-contacts/:id", userHandler.DeleteEmergencyContact)

		// 紧急事件相关
		authorized.POST("/emergencies", emergencyHandler.CreateEmergency)
		authorized.GET("/emergencies/:id", emergencyHandler.GetEmergency)
		authorized.PUT("/emergencies/:id/status", emergencyHandler.UpdateStatus)
		authorized.GET("/emergencies/history", emergencyHandler.GetHistory)
		authorized.POST("/emergencies/:id/handling-records", emergencyHandler.CreateHandlingRecord)
		authorized.GET("/emergencies/:id/handling-records", emergencyHandler.GetHandlingRecords)

		// 安保人员相关
		authorized.POST("/security/apply", securityHandler.ApplySecurityStaff)
		authorized.PUT("/security/location", securityHandler.UpdateLocation)
		authorized.POST("/security/events/:id/accept", securityHandler.AcceptEvent)
		authorized.PUT("/security/events/:id/complete", securityHandler.CompleteEvent)
		authorized.GET("/security/profile", securityHandler.GetStaffInfo)

		// 危险区域相关
		authorized.POST("/danger-zones", dangerZoneHandler.CreateDangerZone)
		authorized.GET("/danger-zones/:id", dangerZoneHandler.GetDangerZone)
		authorized.PUT("/danger-zones/:id", dangerZoneHandler.UpdateDangerZone)
		authorized.DELETE("/danger-zones/:id", dangerZoneHandler.DeleteDangerZone)
		authorized.GET("/danger-zones/nearby", dangerZoneHandler.GetNearbyZones)
		authorized.GET("/danger-zones/active", dangerZoneHandler.GetAllActiveZones)
		authorized.PUT("/danger-zones/:id/heat-level", dangerZoneHandler.UpdateHeatLevel)

		// 评价相关
		authorized.POST("/ratings", ratingHandler.CreateRating)
		authorized.GET("/ratings/:id", ratingHandler.GetRating)
		authorized.GET("/ratings/event/:event_id", ratingHandler.GetEventRating)
		authorized.GET("/ratings/staff/:staff_id", ratingHandler.GetStaffRatings)
		authorized.GET("/ratings/staff/:staff_id/average", ratingHandler.GetStaffAverageRating)
		authorized.PUT("/ratings/:id", ratingHandler.UpdateRating)
		authorized.DELETE("/ratings/:id", ratingHandler.DeleteRating)
	}

	// 管理员路由
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.Auth(), middleware.Admin())
	{
		// 系统配置管理
		admin.POST("/configs", systemConfigHandler.CreateConfig)
		admin.PUT("/configs", systemConfigHandler.UpdateConfig)
		admin.DELETE("/configs/:key", systemConfigHandler.DeleteConfig)
		admin.PUT("/configs/:key/value", systemConfigHandler.UpdateConfigValue)
	}

	return r
}
