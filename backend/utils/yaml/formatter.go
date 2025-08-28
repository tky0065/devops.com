package yaml

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// FormatYAML formate un contenu YAML avec les bonnes indentations
func FormatYAML(content string) (string, error) {
	var data interface{}
	
	if err := yaml.Unmarshal([]byte(content), &data); err != nil {
		return "", fmt.Errorf("failed to parse YAML: %w", err)
	}

	formatted, err := yaml.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to format YAML: %w", err)
	}

	return string(formatted), nil
}

// ValidateYAML valide qu'un contenu est un YAML valide
func ValidateYAML(content string) error {
	var data interface{}
	return yaml.Unmarshal([]byte(content), &data)
}

// MergeYAMLFiles combine plusieurs fichiers YAML en un seul avec des séparateurs
func MergeYAMLFiles(files []string) string {
	var merged strings.Builder
	
	for i, file := range files {
		if i > 0 {
			merged.WriteString("\n---\n")
		}
		merged.WriteString(strings.TrimSpace(file))
	}
	
	return merged.String()
}

// SplitYAMLDocuments divise un fichier YAML multi-documents
func SplitYAMLDocuments(content string) []string {
	documents := strings.Split(content, "---")
	var result []string
	
	for _, doc := range documents {
		doc = strings.TrimSpace(doc)
		if doc != "" {
			result = append(result, doc)
		}
	}
	
	return result
}

// AddYAMLHeader ajoute un header commenté à un fichier YAML
func AddYAMLHeader(content string, header string) string {
	var result strings.Builder
	
	// Ajouter le header comme commentaire
	lines := strings.Split(header, "\n")
	for _, line := range lines {
		if line != "" {
			result.WriteString("# ")
			result.WriteString(line)
		}
		result.WriteString("\n")
	}
	
	result.WriteString("\n")
	result.WriteString(content)
	
	return result.String()
}

// ConvertToYAML convertit une structure Go en YAML
func ConvertToYAML(data interface{}) (string, error) {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to convert to YAML: %w", err)
	}
	
	return string(yamlData), nil
}

// ParseFromYAML parse un YAML vers une structure Go
func ParseFromYAML(content string, target interface{}) error {
	if err := yaml.Unmarshal([]byte(content), target); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}
	
	return nil
}

// CleanEmptyFields supprime les champs vides d'un YAML
func CleanEmptyFields(content string) (string, error) {
	var data map[string]interface{}
	
	if err := yaml.Unmarshal([]byte(content), &data); err != nil {
		return "", fmt.Errorf("failed to parse YAML: %w", err)
	}
	
	cleaned := removeEmptyFields(data)
	
	cleanedYAML, err := yaml.Marshal(cleaned)
	if err != nil {
		return "", fmt.Errorf("failed to marshal cleaned YAML: %w", err)
	}
	
	return string(cleanedYAML), nil
}

// removeEmptyFields supprime récursivement les champs vides
func removeEmptyFields(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			if !isEmpty(value) {
				cleaned := removeEmptyFields(value)
				if !isEmpty(cleaned) {
					result[key] = cleaned
				}
			}
		}
		return result
	case []interface{}:
		var result []interface{}
		for _, item := range v {
			if !isEmpty(item) {
				cleaned := removeEmptyFields(item)
				if !isEmpty(cleaned) {
					result = append(result, cleaned)
				}
			}
		}
		return result
	default:
		return v
	}
}

// isEmpty vérifie si une valeur est considérée comme vide
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	
	switch v := value.(type) {
	case string:
		return v == ""
	case []interface{}:
		return len(v) == 0
	case map[string]interface{}:
		return len(v) == 0
	case int, int32, int64:
		return v == 0
	case float32, float64:
		return v == 0
	case bool:
		return false // bool n'est jamais considéré comme vide
	default:
		return false
	}
}
