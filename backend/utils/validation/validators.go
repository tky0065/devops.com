package validation

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidationError représente une erreur de validation
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}

// Validator interface pour les validateurs
type Validator interface {
	Validate(value interface{}) []ValidationError
}

// StringValidator validateur pour les chaînes de caractères
type StringValidator struct {
	Required   bool
	MinLength  int
	MaxLength  int
	Pattern    *regexp.Regexp
	AllowEmpty bool
}

func (v StringValidator) Validate(value interface{}) []ValidationError {
	var errors []ValidationError
	
	str, ok := value.(string)
	if !ok {
		if value != nil {
			errors = append(errors, ValidationError{
				Code:    "INVALID_TYPE",
				Message: "expected string value",
			})
		}
		str = ""
	}
	
	// Required validation
	if v.Required && str == "" {
		errors = append(errors, ValidationError{
			Code:    "REQUIRED",
			Message: "field is required",
		})
		return errors
	}
	
	// Si la chaîne est vide et que c'est autorisé, pas d'autres validations
	if str == "" && v.AllowEmpty {
		return errors
	}
	
	// MinLength validation
	if v.MinLength > 0 && len(str) < v.MinLength {
		errors = append(errors, ValidationError{
			Code:    "MIN_LENGTH",
			Message: fmt.Sprintf("minimum length is %d, got %d", v.MinLength, len(str)),
		})
	}
	
	// MaxLength validation
	if v.MaxLength > 0 && len(str) > v.MaxLength {
		errors = append(errors, ValidationError{
			Code:    "MAX_LENGTH",
			Message: fmt.Sprintf("maximum length is %d, got %d", v.MaxLength, len(str)),
		})
	}
	
	// Pattern validation
	if v.Pattern != nil && !v.Pattern.MatchString(str) {
		errors = append(errors, ValidationError{
			Code:    "INVALID_PATTERN",
			Message: fmt.Sprintf("value does not match required pattern: %s", v.Pattern.String()),
		})
	}
	
	return errors
}

// IntValidator validateur pour les entiers
type IntValidator struct {
	Required bool
	Min      *int
	Max      *int
}

func (v IntValidator) Validate(value interface{}) []ValidationError {
	var errors []ValidationError
	
	if value == nil {
		if v.Required {
			errors = append(errors, ValidationError{
				Code:    "REQUIRED",
				Message: "field is required",
			})
		}
		return errors
	}
	
	var intVal int
	switch val := value.(type) {
	case int:
		intVal = val
	case int32:
		intVal = int(val)
	case int64:
		intVal = int(val)
	case float64:
		intVal = int(val)
	default:
		errors = append(errors, ValidationError{
			Code:    "INVALID_TYPE",
			Message: "expected integer value",
		})
		return errors
	}
	
	// Min validation
	if v.Min != nil && intVal < *v.Min {
		errors = append(errors, ValidationError{
			Code:    "MIN_VALUE",
			Message: fmt.Sprintf("minimum value is %d, got %d", *v.Min, intVal),
		})
	}
	
	// Max validation
	if v.Max != nil && intVal > *v.Max {
		errors = append(errors, ValidationError{
			Code:    "MAX_VALUE",
			Message: fmt.Sprintf("maximum value is %d, got %d", *v.Max, intVal),
		})
	}
	
	return errors
}

// ArrayValidator validateur pour les tableaux
type ArrayValidator struct {
	Required    bool
	MinLength   int
	MaxLength   int
	ItemValidator Validator
}

func (v ArrayValidator) Validate(value interface{}) []ValidationError {
	var errors []ValidationError
	
	if value == nil {
		if v.Required {
			errors = append(errors, ValidationError{
				Code:    "REQUIRED",
				Message: "field is required",
			})
		}
		return errors
	}
	
	var array []interface{}
	switch val := value.(type) {
	case []interface{}:
		array = val
	case []string:
		for _, s := range val {
			array = append(array, s)
		}
	default:
		errors = append(errors, ValidationError{
			Code:    "INVALID_TYPE",
			Message: "expected array value",
		})
		return errors
	}
	
	// MinLength validation
	if v.MinLength > 0 && len(array) < v.MinLength {
		errors = append(errors, ValidationError{
			Code:    "MIN_LENGTH",
			Message: fmt.Sprintf("minimum length is %d, got %d", v.MinLength, len(array)),
		})
	}
	
	// MaxLength validation
	if v.MaxLength > 0 && len(array) > v.MaxLength {
		errors = append(errors, ValidationError{
			Code:    "MAX_LENGTH",
			Message: fmt.Sprintf("maximum length is %d, got %d", v.MaxLength, len(array)),
		})
	}
	
	// Validate each item
	if v.ItemValidator != nil {
		for i, item := range array {
			itemErrors := v.ItemValidator.Validate(item)
			for _, err := range itemErrors {
				err.Field = fmt.Sprintf("[%d]", i)
				errors = append(errors, err)
			}
		}
	}
	
	return errors
}

// ValidateKubernetesName valide un nom Kubernetes
func ValidateKubernetesName(name string) []ValidationError {
	validator := StringValidator{
		Required:  true,
		MinLength: 1,
		MaxLength: 253,
		Pattern:   regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`),
	}
	
	errors := validator.Validate(name)
	
	// Validation supplémentaire pour les noms Kubernetes
	if len(errors) == 0 {
		if strings.HasPrefix(name, "-") || strings.HasSuffix(name, "-") {
			errors = append(errors, ValidationError{
				Code:    "INVALID_NAME",
				Message: "name cannot start or end with a hyphen",
			})
		}
	}
	
	return errors
}

// ValidateDockerImage valide un nom d'image Docker
func ValidateDockerImage(image string) []ValidationError {
	var errors []ValidationError
	
	if image == "" {
		errors = append(errors, ValidationError{
			Code:    "REQUIRED",
			Message: "image name is required",
		})
		return errors
	}
	
	// Pattern de base pour les images Docker
	imagePattern := regexp.MustCompile(`^[a-z0-9]+(?:[._-][a-z0-9]+)*(?:/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-zA-Z0-9]+(?:[._-][a-zA-Z0-9]+)*)?$`)
	
	if !imagePattern.MatchString(image) {
		errors = append(errors, ValidationError{
			Code:    "INVALID_IMAGE_NAME",
			Message: "invalid Docker image name format",
		})
	}
	
	return errors
}

// ValidatePortMapping valide un mapping de port
func ValidatePortMapping(port string) []ValidationError {
	var errors []ValidationError
	
	if port == "" {
		errors = append(errors, ValidationError{
			Code:    "REQUIRED",
			Message: "port mapping is required",
		})
		return errors
	}
	
	// Patterns pour différents formats de ports
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`^\d+$`),                           // 8080
		regexp.MustCompile(`^\d+:\d+$`),                       // 8080:80
		regexp.MustCompile(`^\d+:\d+/(tcp|udp)$`),            // 8080:80/tcp
		regexp.MustCompile(`^[\d.]+:\d+:\d+$`),               // 127.0.0.1:8080:80
		regexp.MustCompile(`^[\d.]+:\d+:\d+/(tcp|udp)$`),     // 127.0.0.1:8080:80/tcp
	}
	
	valid := false
	for _, pattern := range patterns {
		if pattern.MatchString(port) {
			valid = true
			break
		}
	}
	
	if !valid {
		errors = append(errors, ValidationError{
			Code:    "INVALID_PORT_FORMAT",
			Message: "invalid port mapping format",
		})
	}
	
	return errors
}

// ValidateEnvironmentVariable valide une variable d'environnement
func ValidateEnvironmentVariable(envVar string) []ValidationError {
	var errors []ValidationError
	
	if envVar == "" {
		errors = append(errors, ValidationError{
			Code:    "REQUIRED",
			Message: "environment variable is required",
		})
		return errors
	}
	
	// Pattern pour les variables d'environnement (KEY=value ou KEY)
	envPattern := regexp.MustCompile(`^[A-Z_][A-Z0-9_]*(?:=.*)?$`)
	
	if !envPattern.MatchString(envVar) {
		errors = append(errors, ValidationError{
			Code:    "INVALID_ENV_FORMAT",
			Message: "invalid environment variable format (should be KEY=value or KEY)",
		})
	}
	
	return errors
}

// ValidateVolumeMapping valide un mapping de volume
func ValidateVolumeMapping(volume string) []ValidationError {
	var errors []ValidationError
	
	if volume == "" {
		errors = append(errors, ValidationError{
			Code:    "REQUIRED",
			Message: "volume mapping is required",
		})
		return errors
	}
	
	parts := strings.Split(volume, ":")
	
	if len(parts) < 2 || len(parts) > 3 {
		errors = append(errors, ValidationError{
			Code:    "INVALID_VOLUME_FORMAT",
			Message: "volume mapping should be in format host_path:container_path[:options]",
		})
		return errors
	}
	
	// Valider le chemin hôte
	hostPath := parts[0]
	if hostPath == "" {
		errors = append(errors, ValidationError{
			Code:    "INVALID_HOST_PATH",
			Message: "host path cannot be empty",
		})
	}
	
	// Valider le chemin conteneur
	containerPath := parts[1]
	if containerPath == "" {
		errors = append(errors, ValidationError{
			Code:    "INVALID_CONTAINER_PATH",
			Message: "container path cannot be empty",
		})
	}
	
	// Si présentes, valider les options
	if len(parts) == 3 {
		options := parts[2]
		validOptions := []string{"ro", "rw", "z", "Z"}
		
		optionParts := strings.Split(options, ",")
		for _, opt := range optionParts {
			opt = strings.TrimSpace(opt)
			valid := false
			for _, validOpt := range validOptions {
				if opt == validOpt {
					valid = true
					break
				}
			}
			if !valid {
				errors = append(errors, ValidationError{
					Code:    "INVALID_VOLUME_OPTION",
					Message: fmt.Sprintf("invalid volume option: %s", opt),
				})
			}
		}
	}
	
	return errors
}

// ValidateDockerComposeVersion valide une version docker-compose
func ValidateDockerComposeVersion(version string) []ValidationError {
	var errors []ValidationError
	
	if version == "" {
		errors = append(errors, ValidationError{
			Code:    "REQUIRED",
			Message: "docker-compose version is required",
		})
		return errors
	}
	
	// Versions supportées
	supportedVersions := []string{"3.0", "3.1", "3.2", "3.3", "3.4", "3.5", "3.6", "3.7", "3.8", "3.9"}
	
	valid := false
	for _, supportedVersion := range supportedVersions {
		if version == supportedVersion {
			valid = true
			break
		}
	}
	
	if !valid {
		errors = append(errors, ValidationError{
			Code:    "UNSUPPORTED_VERSION",
			Message: fmt.Sprintf("unsupported docker-compose version: %s", version),
		})
	}
	
	return errors
}
