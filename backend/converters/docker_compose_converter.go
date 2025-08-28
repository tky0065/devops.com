package converters

import (
	"context"
	"fmt"

	"devops-converter/converters/docker"
	"devops-converter/converters/kubernetes"

	"gopkg.in/yaml.v3"
)

// DockerComposeToKubernetesConverter convertit docker-compose vers Kubernetes
type DockerComposeToKubernetesConverter struct {
	name        string
	description string
}

// NewDockerComposeToKubernetesConverter crée un nouveau convertisseur
func NewDockerComposeToKubernetesConverter() Converter {
	return &DockerComposeToKubernetesConverter{
		name:        "docker-compose-to-kubernetes",
		description: "Converts Docker Compose files to Kubernetes manifests",
	}
}

// GetName retourne le nom du convertisseur
func (c *DockerComposeToKubernetesConverter) GetName() string {
	return c.name
}

// GetDescription retourne la description du convertisseur
func (c *DockerComposeToKubernetesConverter) GetDescription() string {
	return c.description
}

// GetSupportedTypes retourne les types supportés
func (c *DockerComposeToKubernetesConverter) GetSupportedTypes() []string {
	return []string{"docker-compose"}
}

// Validate valide le contenu d'entrée
func (c *DockerComposeToKubernetesConverter) Validate(ctx context.Context, content string, contentType string) error {
	if contentType != "docker-compose" {
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	_, err := docker.ParseDockerCompose(content)
	if err != nil {
		return fmt.Errorf("invalid docker-compose file: %w", err)
	}

	return nil
}

// Convert effectue la conversion
func (c *DockerComposeToKubernetesConverter) Convert(ctx context.Context, req ConversionRequest) (*ConversionResult, error) {
	// Valider la requête
	if err := c.Validate(ctx, req.Content, req.Type); err != nil {
		return &ConversionResult{
			Success: false,
			Errors: []ConversionError{
				{
					Code:    "VALIDATION_ERROR",
					Message: err.Error(),
				},
			},
		}, nil
	}

	// Parser le fichier docker-compose
	dockerCompose, err := docker.ParseDockerCompose(req.Content)
	if err != nil {
		return &ConversionResult{
			Success: false,
			Errors: []ConversionError{
				{
					Code:    "PARSE_ERROR",
					Message: fmt.Sprintf("Failed to parse docker-compose file: %v", err),
				},
			},
		}, nil
	}

	// Extraire les options de conversion
	options := c.extractGeneratorOptions(req.Options)

	// Convertir chaque service
	var generatedFiles []GeneratedFile
	var conversionErrors []ConversionError
	var warnings []ConversionWarning

	for serviceName, service := range dockerCompose.Services {
		files, errs, warns := c.convertService(serviceName, service, options)
		generatedFiles = append(generatedFiles, files...)
		conversionErrors = append(conversionErrors, errs...)
		warnings = append(warnings, warns...)
	}

	// Générer les volumes globaux si nécessaire
	if len(dockerCompose.Volumes) > 0 {
		volumeFiles, volumeErrs := c.convertVolumes(dockerCompose.Volumes, options)
		generatedFiles = append(generatedFiles, volumeFiles...)
		conversionErrors = append(conversionErrors, volumeErrs...)
	}

	success := len(conversionErrors) == 0

	return &ConversionResult{
		Success:  success,
		Files:    generatedFiles,
		Errors:   conversionErrors,
		Warnings: warnings,
		Metadata: map[string]interface{}{
			"services_converted": len(dockerCompose.Services),
			"volumes_converted":  len(dockerCompose.Volumes),
			"docker_version":     dockerCompose.Version,
		},
	}, nil
}

// extractGeneratorOptions extrait les options du générateur
func (c *DockerComposeToKubernetesConverter) extractGeneratorOptions(options map[string]interface{}) kubernetes.GeneratorOptions {
	opts := kubernetes.DefaultGeneratorOptions()

	if namespace, ok := options["namespace"].(string); ok {
		opts.Namespace = namespace
	}

	if labels, ok := options["labels"].(map[string]interface{}); ok {
		opts.Labels = make(map[string]string)
		for k, v := range labels {
			opts.Labels[k] = fmt.Sprintf("%v", v)
		}
	}

	if annotations, ok := options["annotations"].(map[string]interface{}); ok {
		opts.Annotations = make(map[string]string)
		for k, v := range annotations {
			opts.Annotations[k] = fmt.Sprintf("%v", v)
		}
	}

	if imagePullPolicy, ok := options["imagePullPolicy"].(string); ok {
		opts.ImagePullPolicy = imagePullPolicy
	}

	if serviceType, ok := options["serviceType"].(string); ok {
		opts.ServiceType = serviceType
	}

	if replicas, ok := options["replicas"].(float64); ok {
		opts.Replicas = int32(replicas)
	}

	return opts
}

// convertService convertit un service Docker Compose vers Kubernetes
func (c *DockerComposeToKubernetesConverter) convertService(serviceName string, service docker.Service, options kubernetes.GeneratorOptions) ([]GeneratedFile, []ConversionError, []ConversionWarning) {
	var files []GeneratedFile
	var errors []ConversionError
	var warnings []ConversionWarning

	// Conversion du service en map[string]interface{} pour le générateur
	serviceData := c.serviceToMap(service)

	// Générer le Deployment
	deployment, err := kubernetes.GenerateDeployment(serviceName, serviceData, options)
	if err != nil {
		errors = append(errors, ConversionError{
			Code:    "DEPLOYMENT_GENERATION_ERROR",
			Message: fmt.Sprintf("Failed to generate deployment for service %s: %v", serviceName, err),
		})
	} else {
		deploymentYAML, err := yaml.Marshal(deployment)
		if err != nil {
			errors = append(errors, ConversionError{
				Code:    "YAML_MARSHAL_ERROR",
				Message: fmt.Sprintf("Failed to marshal deployment for service %s: %v", serviceName, err),
			})
		} else {
			files = append(files, GeneratedFile{
				Name:    fmt.Sprintf("%s-deployment.yaml", serviceName),
				Content: string(deploymentYAML),
				Type:    "deployment",
				Path:    fmt.Sprintf("deployments/%s-deployment.yaml", serviceName),
			})
		}
	}

	// Générer le Service si nécessaire
	kubernetesService, err := kubernetes.GenerateService(serviceName, serviceData, options)
	if err != nil {
		errors = append(errors, ConversionError{
			Code:    "SERVICE_GENERATION_ERROR",
			Message: fmt.Sprintf("Failed to generate service for %s: %v", serviceName, err),
		})
	} else if kubernetesService != nil {
		serviceYAML, err := yaml.Marshal(kubernetesService)
		if err != nil {
			errors = append(errors, ConversionError{
				Code:    "YAML_MARSHAL_ERROR",
				Message: fmt.Sprintf("Failed to marshal service for %s: %v", serviceName, err),
			})
		} else {
			files = append(files, GeneratedFile{
				Name:    fmt.Sprintf("%s-service.yaml", serviceName),
				Content: string(serviceYAML),
				Type:    "service",
				Path:    fmt.Sprintf("services/%s-service.yaml", serviceName),
			})
		}
	}

	// Générer la ConfigMap si nécessaire
	configMap, err := kubernetes.GenerateConfigMap(serviceName, serviceData, options)
	if err != nil {
		warnings = append(warnings, ConversionWarning{
			Code:    "CONFIGMAP_GENERATION_WARNING",
			Message: fmt.Sprintf("Failed to generate configmap for %s: %v", serviceName, err),
		})
	} else if configMap != nil {
		configMapYAML, err := yaml.Marshal(configMap)
		if err != nil {
			warnings = append(warnings, ConversionWarning{
				Code:    "YAML_MARSHAL_WARNING",
				Message: fmt.Sprintf("Failed to marshal configmap for %s: %v", serviceName, err),
			})
		} else {
			files = append(files, GeneratedFile{
				Name:    fmt.Sprintf("%s-configmap.yaml", serviceName),
				Content: string(configMapYAML),
				Type:    "configmap",
				Path:    fmt.Sprintf("configmaps/%s-configmap.yaml", serviceName),
			})
		}
	}

	// Générer les PVCs si nécessaire
	pvcs, err := kubernetes.GeneratePersistentVolumeClaim(serviceName, serviceData, options)
	if err != nil {
		warnings = append(warnings, ConversionWarning{
			Code:    "PVC_GENERATION_WARNING",
			Message: fmt.Sprintf("Failed to generate PVCs for %s: %v", serviceName, err),
		})
	} else {
		for i, pvc := range pvcs {
			pvcYAML, err := yaml.Marshal(pvc)
			if err != nil {
				warnings = append(warnings, ConversionWarning{
					Code:    "YAML_MARSHAL_WARNING",
					Message: fmt.Sprintf("Failed to marshal PVC %d for %s: %v", i, serviceName, err),
				})
				continue
			}
			files = append(files, GeneratedFile{
				Name:    fmt.Sprintf("%s-pvc-%d.yaml", serviceName, i),
				Content: string(pvcYAML),
				Type:    "persistentvolumeclaim",
				Path:    fmt.Sprintf("pvcs/%s-pvc-%d.yaml", serviceName, i),
			})
		}
	}

	// Ajouter des avertissements pour les fonctionnalités non supportées
	warnings = append(warnings, c.checkUnsupportedFeatures(serviceName, service)...)

	return files, errors, warnings
}

// serviceToMap convertit un service Docker en map pour le générateur
func (c *DockerComposeToKubernetesConverter) serviceToMap(service docker.Service) map[string]interface{} {
	// Conversion manuelle pour simplifier
	result := make(map[string]interface{})

	if service.Image != "" {
		result["image"] = service.Image
	}

	if len(service.Ports) > 0 {
		ports := make([]interface{}, len(service.Ports))
		for i, port := range service.Ports {
			ports[i] = port
		}
		result["ports"] = ports
	}

	if service.Environment != nil {
		result["environment"] = service.Environment
	}

	if len(service.Volumes) > 0 {
		volumes := make([]interface{}, len(service.Volumes))
		for i, volume := range service.Volumes {
			volumes[i] = volume
		}
		result["volumes"] = volumes
	}

	if service.Command != nil {
		result["command"] = service.Command
	}

	if service.Entrypoint != nil {
		result["entrypoint"] = service.Entrypoint
	}

	if service.WorkingDir != "" {
		result["working_dir"] = service.WorkingDir
	}

	if service.User != "" {
		result["user"] = service.User
	}

	if service.Privileged {
		result["privileged"] = service.Privileged
	}

	if service.ReadOnly {
		result["read_only"] = service.ReadOnly
	}

	if service.HealthCheck != nil {
		healthcheck := make(map[string]interface{})
		if service.HealthCheck.Test != nil {
			healthcheck["test"] = service.HealthCheck.Test
		}
		if service.HealthCheck.Interval > 0 {
			healthcheck["interval"] = service.HealthCheck.Interval.String()
		}
		if service.HealthCheck.Timeout > 0 {
			healthcheck["timeout"] = service.HealthCheck.Timeout.String()
		}
		if service.HealthCheck.Retries > 0 {
			healthcheck["retries"] = service.HealthCheck.Retries
		}
		result["healthcheck"] = healthcheck
	}

	if service.Deploy != nil {
		deploy := make(map[string]interface{})
		if service.Deploy.Resources != nil {
			resources := make(map[string]interface{})
			if service.Deploy.Resources.Limits != nil {
				limits := make(map[string]interface{})
				if service.Deploy.Resources.Limits.Memory != "" {
					limits["memory"] = service.Deploy.Resources.Limits.Memory
				}
				if service.Deploy.Resources.Limits.CPUs != "" {
					limits["cpus"] = service.Deploy.Resources.Limits.CPUs
				}
				resources["limits"] = limits
			}
			if service.Deploy.Resources.Reservations != nil {
				reservations := make(map[string]interface{})
				if service.Deploy.Resources.Reservations.Memory != "" {
					reservations["memory"] = service.Deploy.Resources.Reservations.Memory
				}
				if service.Deploy.Resources.Reservations.CPUs != "" {
					reservations["cpus"] = service.Deploy.Resources.Reservations.CPUs
				}
				resources["reservations"] = reservations
			}
			deploy["resources"] = resources
		}
		result["deploy"] = deploy
	}

	return result
}

// convertVolumes convertit les volumes globaux
func (c *DockerComposeToKubernetesConverter) convertVolumes(volumes map[string]docker.Volume, options kubernetes.GeneratorOptions) ([]GeneratedFile, []ConversionError) {
	var files []GeneratedFile
	var errors []ConversionError

	for volumeName, volume := range volumes {
		// Créer un PersistentVolume pour chaque volume nommé
		pv := &kubernetes.PersistentVolume{
			APIVersion: "v1",
			Kind:       "PersistentVolume",
			Metadata: kubernetes.Metadata{
				Name:   volumeName,
				Labels: options.Labels,
			},
			Spec: kubernetes.PersistentVolumeSpec{
				Capacity: map[string]string{
					"storage": "1Gi", // Taille par défaut
				},
				AccessModes: []string{"ReadWriteOnce"},
				PersistentVolumeReclaimPolicy: "Retain",
				HostPath: &kubernetes.HostPathVolumeSource{
					Path: fmt.Sprintf("/mnt/data/%s", volumeName),
				},
			},
		}

		// Override avec les options du volume si disponibles
		if volume.Driver != "" && volume.Driver != "local" {
			errors = append(errors, ConversionError{
				Code:    "UNSUPPORTED_VOLUME_DRIVER",
				Message: fmt.Sprintf("Volume driver '%s' is not supported for volume '%s'", volume.Driver, volumeName),
			})
			continue
		}

		pvYAML, err := yaml.Marshal(pv)
		if err != nil {
			errors = append(errors, ConversionError{
				Code:    "YAML_MARSHAL_ERROR",
				Message: fmt.Sprintf("Failed to marshal PersistentVolume for volume %s: %v", volumeName, err),
			})
			continue
		}

		files = append(files, GeneratedFile{
			Name:    fmt.Sprintf("%s-pv.yaml", volumeName),
			Content: string(pvYAML),
			Type:    "persistentvolume",
			Path:    fmt.Sprintf("volumes/%s-pv.yaml", volumeName),
		})
	}

	return files, errors
}

// checkUnsupportedFeatures vérifie les fonctionnalités non supportées
func (c *DockerComposeToKubernetesConverter) checkUnsupportedFeatures(serviceName string, service docker.Service) []ConversionWarning {
	var warnings []ConversionWarning

	// Networks personnalisés
	if service.Networks != nil {
		if networks, ok := service.Networks.(map[string]docker.NetworkConfig); ok && len(networks) > 0 {
			warnings = append(warnings, ConversionWarning{
				Code:    "UNSUPPORTED_NETWORKS",
				Message: fmt.Sprintf("Custom networks for service %s will be converted to default Kubernetes networking", serviceName),
			})
		}
	}

	// Dependencies
	if service.DependsOn != nil {
		if deps, ok := service.DependsOn.(map[string]docker.DependencyConfig); ok && len(deps) > 0 {
			warnings = append(warnings, ConversionWarning{
				Code:    "UNSUPPORTED_DEPENDS_ON",
				Message: fmt.Sprintf("Dependencies for service %s are not directly supported in Kubernetes", serviceName),
				Suggestion: "Consider using init containers or readiness probes",
			})
		}
	}

	// External links
	if service.PidMode != "" && service.PidMode != "none" {
		warnings = append(warnings, ConversionWarning{
			Code:    "UNSUPPORTED_PID_MODE",
			Message: fmt.Sprintf("PID mode '%s' for service %s is not supported", service.PidMode, serviceName),
		})
	}

	// IPC mode
	if service.IpcMode != "" && service.IpcMode != "none" {
		warnings = append(warnings, ConversionWarning{
			Code:    "UNSUPPORTED_IPC_MODE",
			Message: fmt.Sprintf("IPC mode '%s' for service %s is not supported", service.IpcMode, serviceName),
		})
	}

	// SHM size
	if service.ShmSize != "" {
		warnings = append(warnings, ConversionWarning{
			Code:    "UNSUPPORTED_SHM_SIZE",
			Message: fmt.Sprintf("SHM size configuration for service %s requires manual setup in Kubernetes", serviceName),
		})
	}

	return warnings
}
