package main

import (
	"testing"

	"devops-converter/converters"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterConverters(t *testing.T) {
	registry := converters.NewRegistry()
	
	err := registerConverters(registry)
	require.NoError(t, err)
	
	availableConverters := registry.GetAvailableConverters()
	assert.NotEmpty(t, availableConverters)
	
	// Vérifier que le convertisseur docker-compose est disponible
	found := false
	for _, converter := range availableConverters {
		if converter.Name == "docker-compose-to-kubernetes" {
			found = true
			break
		}
	}
	assert.True(t, found, "docker-compose-to-kubernetes converter should be available")
}

func TestDockerComposeConverter(t *testing.T) {
	registry := converters.NewRegistry()
	require.NoError(t, registerConverters(registry))
	
	converter, err := registry.GetConverter("docker-compose")
	require.NoError(t, err)
	assert.NotNil(t, converter)
	
	// Test avec un docker-compose simple
	dockerComposeContent := `version: '3.8'
services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
    environment:
      - ENV=production`
	
	// Test de validation
	err = converter.Validate(nil, dockerComposeContent, "docker-compose")
	assert.NoError(t, err)
	
	// Test de conversion
	req := converters.ConversionRequest{
		Type:    "docker-compose",
		Content: dockerComposeContent,
		Options: map[string]interface{}{
			"namespace": "test",
		},
	}
	
	result, err := converter.Convert(nil, req)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotEmpty(t, result.Files)
	
	// Vérifier qu'on a au moins un deployment
	hasDeployment := false
	for _, file := range result.Files {
		if file.Type == "deployment" {
			hasDeployment = true
			break
		}
	}
	assert.True(t, hasDeployment, "Should generate at least one deployment")
}
