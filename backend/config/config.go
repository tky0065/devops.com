package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config représente la configuration de l'application
type Config struct {
	Server  ServerConfig  `json:"server"`
	App     AppConfig     `json:"app"`
	Logging LoggingConfig `json:"logging"`
	Cors    CorsConfig    `json:"cors"`
}

// ServerConfig configuration du serveur HTTP
type ServerConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	TLS          struct {
		Enabled  bool   `json:"enabled"`
		CertFile string `json:"cert_file"`
		KeyFile  string `json:"key_file"`
	} `json:"tls"`
}

// AppConfig configuration de l'application
type AppConfig struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
	Debug       bool   `json:"debug"`
	MaxFileSize int64  `json:"max_file_size"` // en bytes
}

// LoggingConfig configuration des logs
type LoggingConfig struct {
	Level      string `json:"level"`
	Format     string `json:"format"` // json ou text
	OutputPath string `json:"output_path"`
}

// CorsConfig configuration CORS
type CorsConfig struct {
	AllowedOrigins   []string `json:"allowed_origins"`
	AllowedMethods   []string `json:"allowed_methods"`
	AllowedHeaders   []string `json:"allowed_headers"`
	ExposedHeaders   []string `json:"exposed_headers"`
	AllowCredentials bool     `json:"allow_credentials"`
	MaxAge           int      `json:"max_age"`
}

// Load charge la configuration depuis les variables d'environnement
func Load() (*Config, error) {
	// Charger le fichier .env s'il existe
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Host:         getEnvString("SERVER_HOST", "localhost"),
			Port:         getEnvInt("SERVER_PORT", 8080),
			ReadTimeout:  getEnvInt("SERVER_READ_TIMEOUT", 30),
			WriteTimeout: getEnvInt("SERVER_WRITE_TIMEOUT", 30),
		},
		App: AppConfig{
			Name:        getEnvString("APP_NAME", "devops-converter"),
			Version:     getEnvString("APP_VERSION", "1.0.0"),
			Environment: getEnvString("APP_ENV", "development"),
			Debug:       getEnvBool("APP_DEBUG", true),
			MaxFileSize: getEnvInt64("APP_MAX_FILE_SIZE", 10*1024*1024), // 10MB par défaut
		},
		Logging: LoggingConfig{
			Level:      getEnvString("LOG_LEVEL", "info"),
			Format:     getEnvString("LOG_FORMAT", "json"),
			OutputPath: getEnvString("LOG_OUTPUT_PATH", "stdout"),
		},
		Cors: CorsConfig{
			AllowedOrigins:   getEnvStringSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:5173", "http://localhost:5174"}),
			AllowedMethods:   getEnvStringSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			AllowedHeaders:   getEnvStringSlice("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin"}),
			ExposedHeaders:   getEnvStringSlice("CORS_EXPOSED_HEADERS", []string{}),
			AllowCredentials: getEnvBool("CORS_ALLOW_CREDENTIALS", true),
			MaxAge:           getEnvInt("CORS_MAX_AGE", 86400),
		},
	}

	// Configuration TLS
	config.Server.TLS.Enabled = getEnvBool("TLS_ENABLED", false)
	config.Server.TLS.CertFile = getEnvString("TLS_CERT_FILE", "")
	config.Server.TLS.KeyFile = getEnvString("TLS_KEY_FILE", "")

	// Valider la configuration
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

// validateConfig valide la configuration
func validateConfig(config *Config) error {
	// Valider le port
	if config.Server.Port < 1 || config.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", config.Server.Port)
	}

	// Valider l'environnement
	validEnvs := []string{"development", "staging", "production"}
	validEnv := false
	for _, env := range validEnvs {
		if config.App.Environment == env {
			validEnv = true
			break
		}
	}
	if !validEnv {
		return fmt.Errorf("invalid environment: %s", config.App.Environment)
	}

	// Valider le niveau de log
	validLevels := []string{"debug", "info", "warn", "error"}
	validLevel := false
	for _, level := range validLevels {
		if config.Logging.Level == level {
			validLevel = true
			break
		}
	}
	if !validLevel {
		return fmt.Errorf("invalid log level: %s", config.Logging.Level)
	}

	// Valider le format de log
	if config.Logging.Format != "json" && config.Logging.Format != "text" {
		return fmt.Errorf("invalid log format: %s", config.Logging.Format)
	}

	// Valider la taille max des fichiers
	if config.App.MaxFileSize <= 0 {
		return fmt.Errorf("invalid max file size: %d", config.App.MaxFileSize)
	}

	// Valider la configuration TLS si activée
	if config.Server.TLS.Enabled {
		if config.Server.TLS.CertFile == "" {
			return fmt.Errorf("TLS cert file is required when TLS is enabled")
		}
		if config.Server.TLS.KeyFile == "" {
			return fmt.Errorf("TLS key file is required when TLS is enabled")
		}
	}

	return nil
}

// Fonctions utilitaires pour lire les variables d'environnement

func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}

func getEnvStringSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		// Supposer que les valeurs sont séparées par des virgules
		return splitAndTrim(value, ",")
	}
	return defaultValue
}

func splitAndTrim(s, sep string) []string {
	var result []string
	parts := splitString(s, sep)
	for _, part := range parts {
		trimmed := trimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func splitString(s, sep string) []string {
	if s == "" {
		return []string{}
	}

	var result []string
	var current string

	for i, char := range s {
		if string(char) == sep {
			result = append(result, current)
			current = ""
		} else {
			current += string(char)
		}

		if i == len(s)-1 {
			result = append(result, current)
		}
	}

	return result
}

func trimSpace(s string) string {
	start := 0
	end := len(s)

	// Trim leading spaces
	for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n' || s[start] == '\r') {
		start++
	}

	// Trim trailing spaces
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n' || s[end-1] == '\r') {
		end--
	}

	return s[start:end]
}
