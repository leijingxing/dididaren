package main

import (
	"dididaren/internal/handler"
	"dididaren/internal/middleware"
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"dididaren/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/dididaren?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&model.User{}, &model.Emergency{}, &model.HandlingRecord{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 初始化依赖
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	emergencyRepo := repository.NewEmergencyRepository(db)
	emergencyService := service.NewEmergencyService(emergencyRepo)
	emergencyHandler := handler.NewEmergencyHandler(emergencyService)

	// 创建路由
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// 公开路由
	public := r.Group("/api/v1")
	{
		// 用户相关
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	// 需要认证的路由
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.Auth())
	{
		// 用户相关
		authorized.GET("/user/profile", userHandler.GetProfile)
		authorized.PUT("/user/profile", userHandler.UpdateProfile)

		// 紧急事件相关
		authorized.POST("/emergency", emergencyHandler.Create)
		authorized.GET("/emergency/:id", emergencyHandler.GetByID)
		authorized.PUT("/emergency/:id/status", emergencyHandler.UpdateStatus)
		authorized.GET("/emergency/history", emergencyHandler.GetHistory)
		authorized.POST("/emergency/record", emergencyHandler.CreateHandlingRecord)
		authorized.GET("/emergency/:id/records", emergencyHandler.GetHandlingRecords)
	}

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
