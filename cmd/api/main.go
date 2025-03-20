package main

import (
	"dididaren/docs"
	"dididaren/internal/handler"
	"dididaren/internal/middleware"
	"dididaren/internal/repository"
	"dididaren/internal/service"
	"dididaren/pkg/config"
	"dididaren/pkg/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          滴滴打人 API
// @version        1.0
// @description    滴滴打人服务端 API 文档
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                      Authorization
// @description              Bearer token authentication

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库连接
	db, err := database.Init(cfg)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 初始化 repositories
	userRepo := repository.NewUserRepository(db)
	securityRepo := repository.NewSecurityRepository(db)
	emergencyRepo := repository.NewEmergencyRepository(db)
	dangerZoneRepo := repository.NewDangerZoneRepository(db)
	systemConfigRepo := repository.NewSystemConfigRepository(db)
	ratingRepo := repository.NewRatingRepository(db)

	// 初始化 services
	userService := service.NewUserService(userRepo)
	securityService := service.NewSecurityService(securityRepo)
	emergencyService := service.NewEmergencyService(emergencyRepo)
	dangerZoneService := service.NewDangerZoneService(dangerZoneRepo)
	systemConfigService := service.NewSystemConfigService(systemConfigRepo)
	ratingService := service.NewRatingService(ratingRepo)

	// 初始化 handlers
	userHandler := handler.NewUserHandler(userService)
	securityHandler := handler.NewSecurityHandler(securityService)
	emergencyHandler := handler.NewEmergencyHandler(emergencyService)
	dangerZoneHandler := handler.NewDangerZoneHandler(dangerZoneService)
	systemConfigHandler := handler.NewSystemConfigHandler(systemConfigService)
	ratingHandler := handler.NewRatingHandler(ratingService)

	// 初始化路由
	r := gin.Default()

	// 配置 swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API 路由组
	api := r.Group("/api/v1")
	{
		// 用户相关路由
		api.POST("/users/register", userHandler.Register)
		api.POST("/users/login", userHandler.Login)

		// 需要认证的路由组
		auth := api.Group("/", middleware.Auth())
		{
			// 用户相关
			auth.GET("/users/info", userHandler.GetUserInfo)
			auth.PUT("/users/info", userHandler.UpdateUserInfo)
			auth.PUT("/users/password", userHandler.UpdatePassword)

			// 安保人员相关
			auth.POST("/security/staff", securityHandler.CreateStaff)
			auth.GET("/security/staff/:id", securityHandler.GetStaff)
			auth.GET("/security/staff", securityHandler.ListStaffs)
			auth.PUT("/security/staff/:id/status", securityHandler.UpdateStaffStatus)
			auth.POST("/security/ratings", securityHandler.CreateRating)
			auth.GET("/security/ratings", securityHandler.ListRatings)
			auth.PUT("/security/staff/location", securityHandler.UpdateLocation)
			auth.GET("/security/staff/info", securityHandler.GetStaffInfo)
			auth.POST("/security/staff/apply", securityHandler.ApplySecurityStaff)
			auth.POST("/security/staff/accept-event", securityHandler.AcceptEvent)
			auth.POST("/security/staff/complete-event", securityHandler.CompleteEvent)

			// 紧急事件相关
			auth.POST("/emergency", emergencyHandler.Create)
			auth.GET("/emergency/:id", emergencyHandler.GetByID)
			auth.GET("/emergency", emergencyHandler.List)
			auth.PUT("/emergency/:id", emergencyHandler.Update)
			auth.DELETE("/emergency/:id", emergencyHandler.Delete)
			auth.POST("/emergency/:id/handling", emergencyHandler.CreateHandlingRecord)
			auth.GET("/emergency/:id/handling", emergencyHandler.ListHandlingRecords)
			auth.PUT("/emergency/:id/status", emergencyHandler.UpdateStatus)

			// 危险区域相关
			auth.POST("/danger-zones", dangerZoneHandler.Create)
			auth.GET("/danger-zones/:id", dangerZoneHandler.GetByID)
			auth.GET("/danger-zones", dangerZoneHandler.List)
			auth.PUT("/danger-zones/:id", dangerZoneHandler.Update)
			auth.DELETE("/danger-zones/:id", dangerZoneHandler.Delete)
			auth.GET("/danger-zones/nearby", dangerZoneHandler.GetNearbyZones)
			auth.GET("/danger-zones/active", dangerZoneHandler.GetAllActiveZones)
			auth.PUT("/danger-zones/:id/heat-level", dangerZoneHandler.UpdateHeatLevel)

			// 系统配置相关
			auth.POST("/system/configs", systemConfigHandler.Create)
			auth.GET("/system/configs/:id", systemConfigHandler.GetByID)
			auth.GET("/system/configs", systemConfigHandler.List)
			auth.PUT("/system/configs/:id", systemConfigHandler.Update)
			auth.DELETE("/system/configs/:id", systemConfigHandler.Delete)
			auth.GET("/system/configs/:key/value", systemConfigHandler.GetValue)
			auth.PUT("/system/configs/:key/value", systemConfigHandler.UpdateValue)

			// 评价相关
			auth.POST("/ratings", ratingHandler.CreateRating)
			auth.GET("/ratings/:id", ratingHandler.GetRating)
			auth.GET("/ratings", ratingHandler.ListRatings)
			auth.PUT("/ratings/:id", ratingHandler.UpdateRating)
			auth.DELETE("/ratings/:id", ratingHandler.DeleteRating)
		}
	}

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("服务器启动在 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
