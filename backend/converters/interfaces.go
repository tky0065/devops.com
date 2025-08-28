package converters

import "context"

// ConversionRequest représente une demande de conversion
type ConversionRequest struct {
	Type      string                 `json:"type" binding:"required"`      // docker-compose, dockerfile, etc.
	Content   string                 `json:"content" binding:"required"`   // Contenu du fichier
	Options   map[string]interface{} `json:"options,omitempty"`            // Options de conversion
	Filename  string                 `json:"filename,omitempty"`           // Nom du fichier original
}

// ConversionResult représente le résultat d'une conversion
type ConversionResult struct {
	Success     bool                   `json:"success"`
	Files       []GeneratedFile        `json:"files,omitempty"`
	Errors      []ConversionError      `json:"errors,omitempty"`
	Warnings    []ConversionWarning    `json:"warnings,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// GeneratedFile représente un fichier généré par la conversion
type GeneratedFile struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Type     string `json:"type"`     // deployment, service, configmap, etc.
	Path     string `json:"path"`     // Chemin relatif suggéré
}

// ConversionError représente une erreur de conversion
type ConversionError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Line        int    `json:"line,omitempty"`
	Column      int    `json:"column,omitempty"`
	Field       string `json:"field,omitempty"`
	Suggestion  string `json:"suggestion,omitempty"`
}

// ConversionWarning représente un avertissement de conversion
type ConversionWarning struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Line       int    `json:"line,omitempty"`
	Field      string `json:"field,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

// Converter interface principale pour tous les convertisseurs
type Converter interface {
	// GetSupportedTypes retourne les types de fichiers supportés par ce convertisseur
	GetSupportedTypes() []string
	
	// Convert effectue la conversion du contenu d'entrée
	Convert(ctx context.Context, req ConversionRequest) (*ConversionResult, error)
	
	// Validate valide le contenu d'entrée sans effectuer la conversion
	Validate(ctx context.Context, content string, contentType string) error
	
	// GetName retourne le nom du convertisseur
	GetName() string
	
	// GetDescription retourne la description du convertisseur
	GetDescription() string
}

// ConverterRegistry interface pour le registre des convertisseurs
type ConverterRegistry interface {
	// Register enregistre un nouveau convertisseur
	Register(converter Converter) error
	
	// GetConverter retourne un convertisseur pour le type donné
	GetConverter(contentType string) (Converter, error)
	
	// GetAvailableConverters retourne la liste de tous les convertisseurs disponibles
	GetAvailableConverters() []ConverterInfo
}

// ConverterInfo représente les informations sur un convertisseur disponible
type ConverterInfo struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	SupportedTypes  []string `json:"supported_types"`
	Version         string   `json:"version,omitempty"`
}
