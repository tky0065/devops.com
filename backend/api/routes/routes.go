package routes

import (
	"devops-converter/api/handlers"
	"devops-converter/api/middleware"
	"devops-converter/config"
	"devops-converter/converters"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configure toutes les routes de l'application
func SetupRoutes(cfg *config.Config, registry converters.ConverterRegistry) *gin.Engine {
	// Créer le routeur Gin
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	router := gin.New()

	// Middlewares globaux
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.RequestID())
	router.Use(middleware.CORS(cfg))
	router.Use(middleware.Security())

	// Rate limiting (100 requêtes par minute par défaut)
	router.Use(middleware.RateLimit(100))

	// Créer les handlers
	healthHandler := handlers.NewHealthHandler(cfg, registry)
	convertHandler := handlers.NewConvertHandler(registry)
	uploadHandler := handlers.NewUploadHandler(registry)

	// Routes de santé (sans authentification)
	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/", healthHandler.Health)
		healthGroup.GET("/detailed", healthHandler.HealthDetailed)
		healthGroup.GET("/ready", healthHandler.Ready)
		healthGroup.GET("/live", healthHandler.Live)
	}

	// Route de version
	router.GET("/version", healthHandler.Version)

	// API v1
	v1 := router.Group("/api/v1")
	{
		// Routes de conversion
		conversion := v1.Group("/convert")
		{
			conversion.POST("/", convertHandler.Convert)
			conversion.POST("/validate", convertHandler.Validate)
		}

		// Routes d'upload
		upload := v1.Group("/upload")
		{
			upload.POST("/", uploadHandler.UploadAndConvert)
		}

		// Routes d'information
		info := v1.Group("/info")
		{
			info.GET("/converters", convertHandler.GetConverters)
		}
	}

	// Route pour servir des fichiers statiques (si nécessaire)
	router.Static("/static", "./static")

	// Route de documentation API (si Swagger est intégré)
	router.GET("/docs/*any", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API documentation will be available here",
			"swagger": "/docs/swagger.json",
		})
	})

	// Route par défaut pour les 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "Not Found",
			"message": "The requested resource was not found",
			"path":    c.Request.URL.Path,
		})
	})

	return router
}

// SetupTestRoutes configure les routes pour les tests
func SetupTestRoutes(registry converters.ConverterRegistry) *gin.Engine {
	gin.SetMode(gin.TestMode)
	
	router := gin.New()
	
	// Configuration minimale pour les tests
	cfg := &config.Config{
		App: config.AppConfig{
			Name:        "test-app",
			Version:     "test",
			Environment: "test",
		},
		Cors: config.CorsConfig{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Content-Type", "Authorization"},
		},
	}

	// Middlewares pour les tests
	router.Use(middleware.RequestID())
	router.Use(middleware.CORS(cfg))

	// Handlers
	healthHandler := handlers.NewHealthHandler(cfg, registry)
	convertHandler := handlers.NewConvertHandler(registry)
	uploadHandler := handlers.NewUploadHandler(registry)

	// Routes simplifiées pour les tests
	router.GET("/health", healthHandler.Health)
	router.POST("/convert", convertHandler.Convert)
	router.POST("/validate", convertHandler.Validate)
	router.GET("/converters", convertHandler.GetConverters)
	router.POST("/upload", uploadHandler.UploadAndConvert)

	return router
}
