package routes

import (
	"rest/controllers"
	"rest/middlewares"
	"rest/utils"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	e.Use(middleware.Recover())

	rateLimiterConfig := middlewares.RateLimiterConfig{
		Rate:      10,
		Burst:     30,
		ExpiresIn: 3 * time.Minute,
	}

	rateLimiterMiddleware := rateLimiterConfig.Init()

	e.Use(rateLimiterMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SCRET_KEY"),
		ExpiresDuration: 1,
	}

	customValidator := middlewares.InitValidator()

	e.Validator = customValidator

	authMiddlewareConfig := jwtConfig.Init()

	e.Pre(middleware.RemoveTrailingSlash())

	// Auth
	authController := controllers.InitAuthController(&jwtConfig)

	auth := e.Group("/api/v1")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)

	// Posts
	postController := controllers.InitPostController()

	posts := e.Group("/api/v1/posts", echojwt.WithConfig(authMiddlewareConfig))
	posts.GET("", postController.GetPosts)
	posts.GET("/:id", postController.GetPostByID)
	posts.POST("", postController.AddPost)
	posts.PUT("/:id", postController.UpdatePost)
	posts.DELETE("/:id", postController.DeletePost)

	// Posts
	categoryController := controllers.InitCategoryController()

	categories := e.Group("/api/v1/categories", echojwt.WithConfig(authMiddlewareConfig))
	categories.Use(middlewares.VerifyToken)
	categories.GET("", categoryController.GetCategories)
	categories.POST("", categoryController.AddCategory)
	categories.DELETE("/:id", categoryController.DeleteCategory)
}
