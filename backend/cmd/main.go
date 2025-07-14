package main

import (
	"log"
	"myvault-backend/configs"
	"myvault-backend/internal/handlers"
	"myvault-backend/internal/middleware"
	"myvault-backend/internal/models"
	"myvault-backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// 初始化配置
	cfg := configs.Load()

	// 连接数据库
	db, err := configs.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 连接Redis
	rdb, err := configs.ConnectRedis(cfg)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// 自动迁移数据库表
	if err := models.AutoMigrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 初始化服务
	userService := services.NewUserService(db)
	authService := services.NewAuthService(db, cfg.JWTSecret)
	githubService := services.NewGithubService(cfg.GithubClientID, cfg.GithubClientSecret)
	aiService := services.NewAIService(cfg.OpenAIAPIKey)
	activityService := services.NewActivityService(db, rdb, aiService)

	// 初始化处理器
	authHandler := handlers.NewAuthHandler(authService, userService)
	githubHandler := handlers.NewGithubHandler(githubService, userService)
	activityHandler := handlers.NewActivityHandler(activityService)

	// 设置路由
	router := gin.Default()

	// 中间件
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// 路由组
	api := router.Group("/api")
	{
		// 认证相关
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/github", githubHandler.GithubLogin)
			auth.GET("/github/callback", githubHandler.GithubCallback)
		}

		// 受保护的路由
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(authService))
		{
			protected.GET("/user", authHandler.GetUser)
			protected.PUT("/user", authHandler.UpdateUser)
			
			// 活动相关
			protected.GET("/activities", activityHandler.GetActivities)
			protected.GET("/activities/:id", activityHandler.GetActivity)
			protected.POST("/activities/sync", activityHandler.SyncActivities)
		}
	}

	// 启动服务器
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}