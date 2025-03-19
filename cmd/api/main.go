package main

import (
	"dididaren/internal/handler"
	"dididaren/internal/middleware"
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"dididaren/internal/service"
	"fmt"
	"log"

	_ "dididaren/docs" // 导入 swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title           滴滴打人 API
// @version         1.0
// @description     滴滴打人服务 API 文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(localhost:3306)/dididaren?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 自动迁移数据库表结构
	err = db.AutoMigrate(
		&model.User{},
		&model.Staff{},
		&model.Emergency{},
		&model.EmergencyContact{},
		&model.HandlingRecord{},
		&model.DangerZone{},
		&model.Rating{},
		&model.SystemConfig{},
	)
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	// 初始化仓库
	userRepo := repository.NewUserRepository(db)
	securityRepo := repository.NewSecurityRepository(db)
	emergencyRepo := repository.NewEmergencyRepository(db)
	dangerZoneRepo := repository.NewDangerZoneRepository(db)
	systemConfigRepo := repository.NewSystemConfigRepository(db)

	// 初始化服务
	userService := service.NewUserService(userRepo)
	securityService := service.NewSecurityService(securityRepo)
	emergencyService := service.NewEmergencyService(emergencyRepo)
	dangerZoneService := service.NewDangerZoneService(dangerZoneRepo)
	systemConfigService := service.NewSystemConfigService(systemConfigRepo)

	// 初始化处理器
	userHandler := handler.NewUserHandler(userService)
	securityHandler := handler.NewSecurityHandler(securityService)
	emergencyHandler := handler.NewEmergencyHandler(emergencyService)
	dangerZoneHandler := handler.NewDangerZoneHandler(dangerZoneService)
	systemConfigHandler := handler.NewSystemConfigHandler(systemConfigService)

	// 创建路由
	r := gin.Default()

	// 添加 CORS 中间件
	r.Use(middleware.CORSMiddleware())

	// 静态文件服务
	r.Static("/web", "./web")
	r.StaticFile("/", "./web/login.html")

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API 路由组
	api := r.Group("/api/v1")
	{
		// 用户相关路由
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.GET("/user/profile", middleware.AuthMiddleware(), userHandler.GetProfile)
		api.PUT("/user/profile", middleware.AuthMiddleware(), userHandler.UpdateProfile)

		// 安保人员相关路由
		api.POST("/security/staff", middleware.AuthMiddleware(), securityHandler.CreateStaff)
		api.GET("/security/staff/:id", middleware.AuthMiddleware(), securityHandler.GetStaff)
		api.GET("/security/staff", middleware.AuthMiddleware(), securityHandler.ListStaffs)
		api.PUT("/security/staff/:id/status", middleware.AuthMiddleware(), securityHandler.UpdateStaffStatus)
		api.POST("/security/rating", middleware.AuthMiddleware(), securityHandler.CreateRating)
		api.GET("/security/rating", middleware.AuthMiddleware(), securityHandler.ListRatings)

		// 紧急事件相关路由
		api.POST("/emergency", middleware.AuthMiddleware(), emergencyHandler.CreateEmergency)
		api.GET("/emergency/:id", middleware.AuthMiddleware(), emergencyHandler.GetEmergency)
		api.GET("/emergency", middleware.AuthMiddleware(), emergencyHandler.ListEmergencies)
		api.POST("/emergency/:id/handling", middleware.AuthMiddleware(), emergencyHandler.CreateHandlingRecord)
		api.GET("/emergency/:id/handling", middleware.AuthMiddleware(), emergencyHandler.ListHandlingRecords)

		// 危险区域相关路由
		api.POST("/danger-zone", middleware.AuthMiddleware(), dangerZoneHandler.CreateDangerZone)
		api.GET("/danger-zone/:id", middleware.AuthMiddleware(), dangerZoneHandler.GetDangerZone)
		api.GET("/danger-zone", middleware.AuthMiddleware(), dangerZoneHandler.ListDangerZones)
		api.PUT("/danger-zone/:id", middleware.AuthMiddleware(), dangerZoneHandler.UpdateDangerZone)
		api.DELETE("/danger-zone/:id", middleware.AuthMiddleware(), dangerZoneHandler.DeleteDangerZone)
		api.GET("/danger-zone/check", middleware.AuthMiddleware(), dangerZoneHandler.CheckLocation)

		// 系统配置相关路由
		api.POST("/system/config", middleware.AuthMiddleware(), systemConfigHandler.CreateConfig)
		api.GET("/system/config/:key", middleware.AuthMiddleware(), systemConfigHandler.GetConfig)
		api.GET("/system/configs", middleware.AuthMiddleware(), systemConfigHandler.ListConfigs)
		api.PUT("/system/config/:key", middleware.AuthMiddleware(), systemConfigHandler.UpdateConfig)
		api.DELETE("/system/config/:key", middleware.AuthMiddleware(), systemConfigHandler.DeleteConfig)
	}

	// 启动服务器
	fmt.Println("Server is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
