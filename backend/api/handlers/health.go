package handlers

import (
	"net/http"
	"runtime"
	"time"

	"devops-converter/config"
	"devops-converter/converters"

	"github.com/gin-gonic/gin"
)

// HealthHandler gère les vérifications de santé
type HealthHandler struct {
	config   *config.Config
	registry converters.ConverterRegistry
	startTime time.Time
}

// NewHealthHandler crée un nouveau handler de santé
func NewHealthHandler(cfg *config.Config, registry converters.ConverterRegistry) *HealthHandler {
	return &HealthHandler{
		config:    cfg,
		registry:  registry,
		startTime: time.Now(),
	}
}

// HealthResponse structure de la réponse de santé
type HealthResponse struct {
	Status      string                 `json:"status"`
	Timestamp   time.Time              `json:"timestamp"`
	Version     string                 `json:"version"`
	Environment string                 `json:"environment"`
	Uptime      string                 `json:"uptime"`
	System      SystemInfo             `json:"system"`
	Services    map[string]ServiceInfo `json:"services"`
}

// SystemInfo informations système
type SystemInfo struct {
	GoVersion    string `json:"go_version"`
	NumGoroutine int    `json:"num_goroutine"`
	NumCPU       int    `json:"num_cpu"`
	MemoryUsage  MemoryInfo `json:"memory_usage"`
}

// MemoryInfo informations mémoire
type MemoryInfo struct {
	Alloc      uint64 `json:"alloc"`       // bytes allocated
	TotalAlloc uint64 `json:"total_alloc"` // bytes allocated (lifetime)
	Sys        uint64 `json:"sys"`         // bytes obtained from system
	NumGC      uint32 `json:"num_gc"`      // number of garbage collections
}

// ServiceInfo informations sur un service
type ServiceInfo struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// Health endpoint de vérification de santé simple
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now(),
		"message":   "Service is healthy",
	})
}

// HealthDetailed endpoint de vérification de santé détaillée
func (h *HealthHandler) HealthDetailed(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	uptime := time.Since(h.startTime)

	response := HealthResponse{
		Status:      "ok",
		Timestamp:   time.Now(),
		Version:     h.config.App.Version,
		Environment: h.config.App.Environment,
		Uptime:      uptime.String(),
		System: SystemInfo{
			GoVersion:    runtime.Version(),
			NumGoroutine: runtime.NumGoroutine(),
			NumCPU:       runtime.NumCPU(),
			MemoryUsage: MemoryInfo{
				Alloc:      m.Alloc,
				TotalAlloc: m.TotalAlloc,
				Sys:        m.Sys,
				NumGC:      m.NumGC,
			},
		},
		Services: h.checkServices(),
	}

	// Déterminer le statut global
	overallStatus := "ok"
	for _, service := range response.Services {
		if service.Status != "ok" {
			overallStatus = "degraded"
			if service.Status == "error" {
				overallStatus = "error"
				break
			}
		}
	}
	response.Status = overallStatus

	// Code de statut HTTP basé sur la santé
	statusCode := http.StatusOK
	if overallStatus == "error" {
		statusCode = http.StatusServiceUnavailable
	} else if overallStatus == "degraded" {
		statusCode = http.StatusPartialContent
	}

	c.JSON(statusCode, response)
}

// checkServices vérifie l'état des services
func (h *HealthHandler) checkServices() map[string]ServiceInfo {
	services := make(map[string]ServiceInfo)

	// Vérifier le registre des convertisseurs
	convertersInfo := h.checkConverters()
	services["converters"] = convertersInfo

	// Ajouter d'autres vérifications de services ici
	// Par exemple : base de données, services externes, etc.

	return services
}

// checkConverters vérifie l'état des convertisseurs
func (h *HealthHandler) checkConverters() ServiceInfo {
	converters := h.registry.GetAvailableConverters()
	
	if len(converters) == 0 {
		return ServiceInfo{
			Status:  "error",
			Message: "No converters available",
		}
	}

	// Vérifier que les convertisseurs de base sont disponibles
	requiredConverters := []string{"docker-compose-to-kubernetes"}
	availableConverters := make(map[string]bool)
	
	for _, converter := range converters {
		availableConverters[converter.Name] = true
	}

	missingConverters := []string{}
	for _, required := range requiredConverters {
		if !availableConverters[required] {
			missingConverters = append(missingConverters, required)
		}
	}

	if len(missingConverters) > 0 {
		return ServiceInfo{
			Status:  "degraded",
			Message: "Some required converters are missing",
			Details: map[string]interface{}{
				"missing_converters": missingConverters,
				"available_count":    len(converters),
			},
		}
	}

	return ServiceInfo{
		Status:  "ok",
		Message: "All converters are available",
		Details: map[string]interface{}{
			"available_converters": len(converters),
			"converter_list":       converters,
		},
	}
}

// Ready endpoint de vérification de disponibilité (readiness probe)
func (h *HealthHandler) Ready(c *gin.Context) {
	// Vérifier que l'application est prête à recevoir du trafic
	converters := h.registry.GetAvailableConverters()
	
	if len(converters) == 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "not ready",
			"message": "No converters available",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ready",
		"message": "Service is ready to accept traffic",
	})
}

// Live endpoint de vérification de vivacité (liveness probe)
func (h *HealthHandler) Live(c *gin.Context) {
	// Vérification basique que l'application fonctionne
	c.JSON(http.StatusOK, gin.H{
		"status":    "alive",
		"timestamp": time.Now(),
		"uptime":    time.Since(h.startTime).String(),
	})
}

// Version endpoint pour obtenir les informations de version
func (h *HealthHandler) Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":        h.config.App.Name,
		"version":     h.config.App.Version,
		"environment": h.config.App.Environment,
		"go_version":  runtime.Version(),
		"build_time":  "", // À ajouter lors du build
		"git_commit":  "", // À ajouter lors du build
	})
}
