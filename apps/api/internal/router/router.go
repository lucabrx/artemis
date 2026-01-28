package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lukabrkovic/artemis/docs"
	"github.com/lukabrkovic/artemis/internal/audit"
	"github.com/lukabrkovic/artemis/internal/cache"
	"github.com/lukabrkovic/artemis/internal/config"
	"github.com/lukabrkovic/artemis/internal/events"
	"github.com/lukabrkovic/artemis/internal/handler"
	"github.com/lukabrkovic/artemis/internal/middleware"
	"github.com/lukabrkovic/artemis/internal/service"
	"github.com/lukabrkovic/artemis/internal/store"
	"github.com/lukabrkovic/artemis/pkg/storage"
	"github.com/lukabrkovic/artemis/pkg/token"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Store                   *store.Store
	Cache                   *cache.Cache
	Storage                 *storage.MinIO
	TokenMaker              token.Maker
	TokenConfig             config.TokenConfig
	Logger                  zerolog.Logger
	Environment             string
	EnableOpenAPIValidation bool
	MaxRequestSize          int64
	EventBus                *events.Bus
	AuditLogger             *audit.Logger
}

func New(cfg Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID(cfg.Logger))
	router.Use(middleware.PrometheusMetrics())
	router.Use(middleware.NewStructuredLogger(cfg.Logger).Middleware())

	if cfg.EnableOpenAPIValidation {
		validator, err := middleware.NewOpenAPIValidator("./docs/swagger.yaml", cfg.Logger)
		if err != nil {
			cfg.Logger.Warn().Err(err).Msg("failed to load OpenAPI validator, continuing without validation")
		} else {
			router.Use(validator.Validate())
			cfg.Logger.Info().Msg("OpenAPI validation enabled")
		}
	}

	router.Use(middleware.ErrorHandler(cfg.Logger))
	router.Use(middleware.CORS())
	router.Use(middleware.RequestSizeLimiter(cfg.MaxRequestSize))
	router.Use(middleware.RequireContentLength(cfg.MaxRequestSize))
	router.Use(middleware.RateLimiterByIP(1000, 60))

	if cfg.AuditLogger != nil {
		router.Use(middleware.NewAuditMiddleware(cfg.AuditLogger).Middleware())
	}

	authService := service.NewAuthService(cfg.Store, cfg.Cache, cfg.TokenMaker, cfg.TokenConfig, cfg.EventBus, cfg.Logger)
	userService := service.NewUserService(cfg.Store, cfg.Cache, cfg.Storage, cfg.Logger)
	workspaceService := service.NewWorkspaceService(cfg.Store, cfg.Storage, cfg.EventBus, cfg.Logger)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	workspaceHandler := handler.NewWorkspaceHandler(workspaceService)

	router.GET("/health", handler.Health)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		RegisterAuthRoutes(api, authHandler)
		RegisterUserRoutes(api, userHandler, cfg.TokenMaker)
		RegisterWorkspaceRoutes(api, workspaceHandler, cfg.TokenMaker)
	}

	return router
}
