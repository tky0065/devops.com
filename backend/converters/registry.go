package converters

import (
	"fmt"
	"sync"
)

// registry implémentation du ConverterRegistry
type registry struct {
	converters map[string]Converter
	mu         sync.RWMutex
}

// NewRegistry crée une nouvelle instance du registre des convertisseurs
func NewRegistry() ConverterRegistry {
	return &registry{
		converters: make(map[string]Converter),
	}
}

// Register enregistre un nouveau convertisseur
func (r *registry) Register(converter Converter) error {
	if converter == nil {
		return fmt.Errorf("converter cannot be nil")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	name := converter.GetName()
	if name == "" {
		return fmt.Errorf("converter name cannot be empty")
	}

	// Vérifier si le convertisseur existe déjà
	if _, exists := r.converters[name]; exists {
		return fmt.Errorf("converter with name '%s' already registered", name)
	}

	r.converters[name] = converter
	return nil
}

// GetConverter retourne un convertisseur pour le type donné
func (r *registry) GetConverter(contentType string) (Converter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Chercher un convertisseur qui supporte ce type de contenu
	for _, converter := range r.converters {
		supportedTypes := converter.GetSupportedTypes()
		for _, supportedType := range supportedTypes {
			if supportedType == contentType {
				return converter, nil
			}
		}
	}

	return nil, fmt.Errorf("no converter found for content type: %s", contentType)
}

// GetAvailableConverters retourne la liste de tous les convertisseurs disponibles
func (r *registry) GetAvailableConverters() []ConverterInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var converters []ConverterInfo
	for _, converter := range r.converters {
		info := ConverterInfo{
			Name:           converter.GetName(),
			Description:    converter.GetDescription(),
			SupportedTypes: converter.GetSupportedTypes(),
		}
		converters = append(converters, info)
	}

	return converters
}

// GetConverterByName retourne un convertisseur par son nom
func (r *registry) GetConverterByName(name string) (Converter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	converter, exists := r.converters[name]
	if !exists {
		return nil, fmt.Errorf("converter with name '%s' not found", name)
	}

	return converter, nil
}
