package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"devops-converter/converters"

	"github.com/gin-gonic/gin"
)

// ConvertHandler gère les requêtes de conversion
type ConvertHandler struct {
	registry converters.ConverterRegistry
}

// NewConvertHandler crée un nouveau handler de conversion
func NewConvertHandler(registry converters.ConverterRegistry) *ConvertHandler {
	return &ConvertHandler{
		registry: registry,
	}
}

// ConvertRequest structure de la requête de conversion
type ConvertRequest struct {
	Type     string                 `json:"type" binding:"required"`
	Content  string                 `json:"content" binding:"required"`
	Options  map[string]interface{} `json:"options,omitempty"`
	Filename string                 `json:"filename,omitempty"`
}

// ConvertResponse structure de la réponse de conversion
type ConvertResponse struct {
	Success   bool                           `json:"success"`
	Files     []converters.GeneratedFile     `json:"files,omitempty"`
	Errors    []converters.ConversionError   `json:"errors,omitempty"`
	Warnings  []converters.ConversionWarning `json:"warnings,omitempty"`
	Metadata  map[string]interface{}         `json:"metadata,omitempty"`
	RequestID string                         `json:"request_id,omitempty"`
}

// Convert endpoint principal de conversion
func (h *ConvertHandler) Convert(c *gin.Context) {
	var req ConvertRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	// Vérifier la taille du contenu
	if len(req.Content) > 10*1024*1024 { // 10MB max
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"success": false,
			"error":   "File too large",
			"details": "Maximum file size is 10MB",
		})
		return
	}

	// Obtenir le convertisseur approprié
	converter, err := h.registry.GetConverter(req.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Unsupported conversion type",
			"details": err.Error(),
		})
		return
	}

	// Créer la requête de conversion
	conversionReq := converters.ConversionRequest{
		Type:     req.Type,
		Content:  req.Content,
		Options:  req.Options,
		Filename: req.Filename,
	}

	// Effectuer la conversion
	result, err := converter.Convert(c.Request.Context(), conversionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Conversion failed",
			"details": err.Error(),
		})
		return
	}

	// Préparer la réponse
	response := ConvertResponse{
		Success:   result.Success,
		Files:     result.Files,
		Errors:    result.Errors,
		Warnings:  result.Warnings,
		Metadata:  result.Metadata,
		RequestID: c.GetString("request_id"),
	}

	// Déterminer le code de statut
	statusCode := http.StatusOK
	if !result.Success {
		if len(result.Errors) > 0 {
			statusCode = http.StatusBadRequest
		}
	}

	c.JSON(statusCode, response)
}

// ValidateRequest structure pour la validation
type ValidateRequest struct {
	Type    string `json:"type" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// ValidateResponse structure de la réponse de validation
type ValidateResponse struct {
	Valid   bool     `json:"valid"`
	Errors  []string `json:"errors,omitempty"`
	Message string   `json:"message"`
}

// Validate endpoint de validation des fichiers
func (h *ConvertHandler) Validate(c *gin.Context) {
	var req ValidateRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"valid":   false,
			"message": "Invalid request format",
			"errors":  []string{err.Error()},
		})
		return
	}

	// Obtenir le convertisseur approprié
	converter, err := h.registry.GetConverter(req.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"valid":   false,
			"message": "Unsupported file type",
			"errors":  []string{err.Error()},
		})
		return
	}

	// Valider le contenu
	if err := converter.Validate(c.Request.Context(), req.Content, req.Type); err != nil {
		c.JSON(http.StatusOK, ValidateResponse{
			Valid:   false,
			Message: "Validation failed",
			Errors:  []string{err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, ValidateResponse{
		Valid:   true,
		Message: "File is valid",
	})
}

// GetConverters retourne la liste des convertisseurs disponibles
func (h *ConvertHandler) GetConverters(c *gin.Context) {
	converters := h.registry.GetAvailableConverters()
	
	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"converters": converters,
		"count":      len(converters),
	})
}

// UploadHandler gère l'upload de fichiers
type UploadHandler struct {
	registry converters.ConverterRegistry
}

// NewUploadHandler crée un nouveau handler d'upload
func NewUploadHandler(registry converters.ConverterRegistry) *UploadHandler {
	return &UploadHandler{
		registry: registry,
	}
}

// UploadAndConvert upload et convertit un fichier
func (h *UploadHandler) UploadAndConvert(c *gin.Context) {
	// Récupérer le fichier uploadé
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "No file uploaded",
			"details": err.Error(),
		})
		return
	}

	// Vérifier la taille du fichier
	if file.Size > 10*1024*1024 { // 10MB max
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"success": false,
			"error":   "File too large",
			"details": "Maximum file size is 10MB",
		})
		return
	}

	// Ouvrir le fichier
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to open uploaded file",
			"details": err.Error(),
		})
		return
	}
	defer src.Close()

	// Lire le contenu
	content := make([]byte, file.Size)
	if _, err := src.Read(content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to read uploaded file",
			"details": err.Error(),
		})
		return
	}

	// Déterminer le type de fichier
	fileType := c.PostForm("type")
	if fileType == "" {
		// Essayer de déterminer le type à partir du nom de fichier
		fileType = determineFileType(file.Filename)
	}

	if fileType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Could not determine file type",
			"details": "Please specify the file type",
		})
		return
	}

	// Obtenir les options de conversion
	options := make(map[string]interface{})
	
	if namespace := c.PostForm("namespace"); namespace != "" {
		options["namespace"] = namespace
	}
	
	if serviceType := c.PostForm("serviceType"); serviceType != "" {
		options["serviceType"] = serviceType
	}
	
	if replicas := c.PostForm("replicas"); replicas != "" {
		if r, err := strconv.Atoi(replicas); err == nil {
			options["replicas"] = r
		}
	}

	// Obtenir le convertisseur
	converter, err := h.registry.GetConverter(fileType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Unsupported file type",
			"details": err.Error(),
		})
		return
	}

	// Créer la requête de conversion
	conversionReq := converters.ConversionRequest{
		Type:     fileType,
		Content:  string(content),
		Options:  options,
		Filename: file.Filename,
	}

	// Effectuer la conversion
	result, err := converter.Convert(c.Request.Context(), conversionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Conversion failed",
			"details": err.Error(),
		})
		return
	}

	// Préparer la réponse
	response := ConvertResponse{
		Success:   result.Success,
		Files:     result.Files,
		Errors:    result.Errors,
		Warnings:  result.Warnings,
		Metadata:  result.Metadata,
		RequestID: c.GetString("request_id"),
	}

	// Ajouter les métadonnées du fichier uploadé
	if response.Metadata == nil {
		response.Metadata = make(map[string]interface{})
	}
	response.Metadata["uploaded_file"] = map[string]interface{}{
		"filename": file.Filename,
		"size":     file.Size,
		"type":     fileType,
	}

	statusCode := http.StatusOK
	if !result.Success {
		statusCode = http.StatusBadRequest
	}

	c.JSON(statusCode, response)
}

// determineFileType détermine le type de fichier à partir du nom
func determineFileType(filename string) string {
	if filename == "" {
		return ""
	}
	
	// Convertir en minuscules pour la comparaison
	lower := strings.ToLower(filename)
	
	if strings.HasSuffix(lower, "docker-compose.yml") || 
	   strings.HasSuffix(lower, "docker-compose.yaml") ||
	   strings.Contains(lower, "compose") {
		return "docker-compose"
	}
	
	if strings.HasSuffix(lower, "dockerfile") {
		return "dockerfile"
	}
	
	return ""
}
