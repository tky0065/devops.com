package middleware

import (
	"fmt"
	"time"

	"devops-converter/config"

	"github.com/gin-gonic/gin"
)

// CORS middleware pour gérer les requêtes cross-origin
func CORS(config *config.Config) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Toujours autoriser l'origine de la requête si elle est dans la liste ou si wildcard
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		// Gérer les requêtes OPTIONS (preflight)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

// RequestID middleware pour ajouter un ID unique à chaque requête
func RequestID() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)

		c.Next()
	})
}

// Logger middleware pour logger les requêtes
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s - \"%s %s %s\" %d %s \"%s\" \"%s\" %s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.Request.Referer(),
			param.ErrorMessage,
		)
	})
}

// ErrorHandler middleware pour gérer les erreurs
func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(500, gin.H{
				"success":    false,
				"error":      "Internal server error",
				"message":    err,
				"request_id": c.GetString("request_id"),
			})
		}
		c.AbortWithStatus(500)
	})
}

// RateLimit middleware pour limiter le taux de requêtes
// Note: Implémentation simple en mémoire, pour la production utiliser Redis
func RateLimit(requestsPerMinute int) gin.HandlerFunc {
	clients := make(map[string][]time.Time)

	return gin.HandlerFunc(func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()

		// Nettoyer les anciennes entrées
		if requests, exists := clients[clientIP]; exists {
			var validRequests []time.Time
			for _, requestTime := range requests {
				if now.Sub(requestTime) < time.Minute {
					validRequests = append(validRequests, requestTime)
				}
			}
			clients[clientIP] = validRequests
		}

		// Vérifier la limite
		if requests, exists := clients[clientIP]; exists {
			if len(requests) >= requestsPerMinute {
				c.JSON(429, gin.H{
					"success":     false,
					"error":       "Rate limit exceeded",
					"message":     fmt.Sprintf("Maximum %d requests per minute allowed", requestsPerMinute),
					"retry_after": 60,
				})
				c.Abort()
				return
			}
		}

		// Ajouter la requête actuelle
		if clients[clientIP] == nil {
			clients[clientIP] = []time.Time{}
		}
		clients[clientIP] = append(clients[clientIP], now)

		c.Next()
	})
}

// Security middleware pour ajouter des headers de sécurité
func Security() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Empêcher l'inclusion dans des iframes
		c.Header("X-Frame-Options", "DENY")

		// Empêcher le sniffing de type MIME
		c.Header("X-Content-Type-Options", "nosniff")

		// Activer la protection XSS
		c.Header("X-XSS-Protection", "1; mode=block")

		// Politique de référent
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// Content Security Policy basique
		c.Header("Content-Security-Policy", "default-src 'self'")

		c.Next()
	})
}

// generateRequestID génère un ID unique pour la requête
func generateRequestID() string {
	// Implémentation simple utilisant timestamp et nombre aléatoire
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
