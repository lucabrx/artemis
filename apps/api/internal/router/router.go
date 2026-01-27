package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lukabrkovic/artemis/docs"
	"github.com/lukabrkovic/artemis/internal/cache"
	"github.com/lukabrkovic/artemis/internal/config"
	"github.com/lukabrkovic/artemis/internal/handler"
	"github.com/lukabrkovic/artemis/internal/middleware"
	"github.com/lukabrkovic/artemis/internal/service"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/lukabrkovic/artemis/pkg/storage"
	"github.com/lukabrkovic/artemis/pkg/token"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Store       *store.Store
	Cache       *cache.Cache
	Storage     *storage.MinIO
	TokenMaker  token.Maker
	TokenConfig config.TokenConfig
	Logger      zerolog.Logger
}

func New(cfg Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.Logger(cfg.Logger))
	router.Use(middleware.ErrorHandler(cfg.Logger))
	router.Use(middleware.CORS())

	authService := service.NewAuthService(cfg.Store, cfg.Cache, cfg.TokenMaker, cfg.TokenConfig)
	userService := service.NewUserService(cfg.Store, cfg.Cache, cfg.Storage)
	workspaceService := service.NewWorkspaceService(cfg.Store, cfg.Storage)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	workspaceHandler := handler.NewWorkspaceHandler(workspaceService)

	router.GET("/health", handler.Health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		RegisterAuthRoutes(api, authHandler)
		RegisterUserRoutes(api, userHandler, cfg.TokenMaker)
		RegisterWorkspaceRoutes(api, workspaceHandler, cfg.TokenMaker)
	}

	return router
}
