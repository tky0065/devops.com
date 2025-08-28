package kubernetes

import (
	"fmt"
	"strconv"
	"strings"
)

// Generator interface pour générer des manifests Kubernetes
type Generator interface {
	GenerateManifests(serviceName string, service interface{}, options GeneratorOptions) ([]KubernetesManifest, error)
}

// GeneratorOptions options de génération
type GeneratorOptions struct {
	Namespace       string            `json:"namespace"`
	Labels          map[string]string `json:"labels"`
	Annotations     map[string]string `json:"annotations"`
	ImagePullPolicy string            `json:"imagePullPolicy"`
	ServiceType     string            `json:"serviceType"`
	Replicas        int32             `json:"replicas"`
}

// DefaultGeneratorOptions retourne les options par défaut
func DefaultGeneratorOptions() GeneratorOptions {
	return GeneratorOptions{
		Namespace:       "default",
		Labels:          make(map[string]string),
		Annotations:     make(map[string]string),
		ImagePullPolicy: "IfNotPresent",
		ServiceType:     "ClusterIP",
		Replicas:        1,
	}
}

// GenerateDeployment génère un Deployment Kubernetes à partir d'un service Docker Compose
func GenerateDeployment(serviceName string, service interface{}, options GeneratorOptions) (*Deployment, error) {
	serviceMap, ok := service.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid service format for %s", serviceName)
	}

	deployment := &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      serviceName,
			Namespace: options.Namespace,
			Labels:    mergeLabels(options.Labels, map[string]string{"app": serviceName}),
		},
		Spec: DeploymentSpec{
			Replicas: options.Replicas,
			Selector: &LabelSelector{
				MatchLabels: map[string]string{"app": serviceName},
			},
			Template: PodTemplateSpec{
				Metadata: Metadata{
					Labels: map[string]string{"app": serviceName},
				},
				Spec: PodSpec{
					RestartPolicy: "Always",
				},
			},
		},
	}

	// Générer le conteneur principal
	container, err := generateContainer(serviceName, serviceMap, options)
	if err != nil {
		return nil, fmt.Errorf("failed to generate container for %s: %w", serviceName, err)
	}

	deployment.Spec.Template.Spec.Containers = []Container{*container}

	// Générer les volumes si nécessaire
	volumes, volumeMounts, err := generateVolumes(serviceMap)
	if err != nil {
		return nil, fmt.Errorf("failed to generate volumes for %s: %w", serviceName, err)
	}

	if len(volumes) > 0 {
		deployment.Spec.Template.Spec.Volumes = volumes
		deployment.Spec.Template.Spec.Containers[0].VolumeMounts = volumeMounts
	}

	return deployment, nil
}

// generateContainer génère un conteneur Kubernetes
func generateContainer(serviceName string, service map[string]interface{}, options GeneratorOptions) (*Container, error) {
	container := &Container{
		Name:            serviceName,
		ImagePullPolicy: options.ImagePullPolicy,
	}

	// Image
	if image, ok := service["image"].(string); ok {
		container.Image = image
	} else {
		return nil, fmt.Errorf("no image specified for service %s", serviceName)
	}

	// Commande et arguments
	if command, ok := service["command"]; ok {
		container.Command = normalizeStringSlice(command)
	}

	if entrypoint, ok := service["entrypoint"]; ok {
		container.Args = normalizeStringSlice(entrypoint)
	}

	// Working directory
	if workingDir, ok := service["working_dir"].(string); ok {
		container.WorkingDir = workingDir
	}

	// Ports
	if ports, ok := service["ports"]; ok {
		containerPorts, err := generateContainerPorts(ports)
		if err != nil {
			return nil, fmt.Errorf("failed to generate ports: %w", err)
		}
		container.Ports = containerPorts
	}

	// Variables d'environnement
	if env, ok := service["environment"]; ok {
		envVars, err := generateEnvVars(env)
		if err != nil {
			return nil, fmt.Errorf("failed to generate environment variables: %w", err)
		}
		container.Env = envVars
	}

	// Health check
	if healthcheck, ok := service["healthcheck"]; ok {
		probes, err := generateProbes(healthcheck)
		if err != nil {
			return nil, fmt.Errorf("failed to generate health check: %w", err)
		}
		if probes.liveness != nil {
			container.LivenessProbe = probes.liveness
		}
		if probes.readiness != nil {
			container.ReadinessProbe = probes.readiness
		}
	}

	// Ressources
	if resources, ok := service["deploy"].(map[string]interface{}); ok {
		if resourcesConfig, ok := resources["resources"].(map[string]interface{}); ok {
			resourceReqs, err := generateResourceRequirements(resourcesConfig)
			if err != nil {
				return nil, fmt.Errorf("failed to generate resource requirements: %w", err)
			}
			container.Resources = resourceReqs
		}
	}

	// Security context
	securityContext, err := generateSecurityContext(service)
	if err != nil {
		return nil, fmt.Errorf("failed to generate security context: %w", err)
	}
	if securityContext != nil {
		container.SecurityContext = securityContext
	}

	return container, nil
}

// generateContainerPorts génère les ports de conteneur
func generateContainerPorts(ports interface{}) ([]ContainerPort, error) {
	var containerPorts []ContainerPort

	portSlice, ok := ports.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid ports format")
	}

	for _, port := range portSlice {
		portStr, ok := port.(string)
		if !ok {
			continue
		}

		containerPort, err := parsePortMapping(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid port format %s: %w", portStr, err)
		}

		containerPorts = append(containerPorts, *containerPort)
	}

	return containerPorts, nil
}

// parsePortMapping parse une mapping de port Docker Compose
func parsePortMapping(portMapping string) (*ContainerPort, error) {
	parts := strings.Split(portMapping, ":")
	
	var containerPortNum int32
	var protocol = "TCP"

	switch len(parts) {
	case 1:
		// Format: "8080"
		port, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid port number: %s", parts[0])
		}
		containerPortNum = int32(port)
	case 2:
		// Format: "8080:80"
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid container port number: %s", parts[1])
		}
		containerPortNum = int32(port)
	case 3:
		// Format: "127.0.0.1:8080:80" ou "8080:80:tcp"
		if strings.Contains(parts[2], "tcp") || strings.Contains(parts[2], "udp") {
			protocol = strings.ToUpper(parts[2])
			port, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid container port number: %s", parts[1])
			}
			containerPortNum = int32(port)
		} else {
			port, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, fmt.Errorf("invalid container port number: %s", parts[2])
			}
			containerPortNum = int32(port)
		}
	default:
		return nil, fmt.Errorf("invalid port mapping format: %s", portMapping)
	}

	return &ContainerPort{
		ContainerPort: containerPortNum,
		Protocol:      protocol,
	}, nil
}

// generateEnvVars génère les variables d'environnement
func generateEnvVars(env interface{}) ([]EnvVar, error) {
	var envVars []EnvVar

	switch e := env.(type) {
	case map[string]interface{}:
		for key, value := range e {
			envVars = append(envVars, EnvVar{
				Name:  key,
				Value: fmt.Sprintf("%v", value),
			})
		}
	case map[string]string:
		for key, value := range e {
			envVars = append(envVars, EnvVar{
				Name:  key,
				Value: value,
			})
		}
	case []interface{}:
		for _, item := range e {
			envStr, ok := item.(string)
			if !ok {
				continue
			}
			parts := strings.SplitN(envStr, "=", 2)
			if len(parts) == 2 {
				envVars = append(envVars, EnvVar{
					Name:  parts[0],
					Value: parts[1],
				})
			} else {
				envVars = append(envVars, EnvVar{
					Name:  parts[0],
					Value: "",
				})
			}
		}
	default:
		return nil, fmt.Errorf("unsupported environment format: %T", env)
	}

	return envVars, nil
}

// generateVolumes génère les volumes et volume mounts
func generateVolumes(service map[string]interface{}) ([]Volume, []VolumeMount, error) {
	var volumes []Volume
	var volumeMounts []VolumeMount

	volumesConfig, ok := service["volumes"]
	if !ok {
		return volumes, volumeMounts, nil
	}

	volumeSlice, ok := volumesConfig.([]interface{})
	if !ok {
		return volumes, volumeMounts, nil
	}

	for i, vol := range volumeSlice {
		volStr, ok := vol.(string)
		if !ok {
			continue
		}

		volume, volumeMount, err := parseVolumeMapping(volStr, i)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid volume mapping %s: %w", volStr, err)
		}

		volumes = append(volumes, *volume)
		volumeMounts = append(volumeMounts, *volumeMount)
	}

	return volumes, volumeMounts, nil
}

// parseVolumeMapping parse un mapping de volume Docker Compose
func parseVolumeMapping(volumeMapping string, index int) (*Volume, *VolumeMount, error) {
	parts := strings.Split(volumeMapping, ":")
	
	if len(parts) < 2 {
		return nil, nil, fmt.Errorf("invalid volume mapping format: %s", volumeMapping)
	}

	hostPath := parts[0]
	containerPath := parts[1]
	readOnly := false

	if len(parts) > 2 && parts[2] == "ro" {
		readOnly = true
	}

	volumeName := fmt.Sprintf("volume-%d", index)

	volume := &Volume{
		Name: volumeName,
		HostPath: &HostPathVolumeSource{
			Path: hostPath,
		},
	}

	volumeMount := &VolumeMount{
		Name:      volumeName,
		MountPath: containerPath,
		ReadOnly:  readOnly,
	}

	return volume, volumeMount, nil
}

// generateProbes génère les health check probes
type probeSet struct {
	liveness  *Probe
	readiness *Probe
}

func generateProbes(healthcheck interface{}) (*probeSet, error) {
	healthMap, ok := healthcheck.(map[string]interface{})
	if !ok {
		return &probeSet{}, nil
	}

	var probes probeSet

	// Test command
	if test, ok := healthMap["test"]; ok {
		testCmd := normalizeStringSlice(test)
		if len(testCmd) > 0 {
			probe := &Probe{
				Handler: Handler{
					Exec: &ExecAction{
						Command: testCmd,
					},
				},
			}

			// Interval, timeout, etc.
			if interval, ok := healthMap["interval"].(string); ok {
				// Convertir duration string en secondes
				if seconds, err := parseDurationToSeconds(interval); err == nil {
					probe.PeriodSeconds = seconds
				}
			}

			if timeout, ok := healthMap["timeout"].(string); ok {
				if seconds, err := parseDurationToSeconds(timeout); err == nil {
					probe.TimeoutSeconds = seconds
				}
			}

			if retries, ok := healthMap["retries"].(int); ok {
				probe.FailureThreshold = int32(retries)
			}

			// Utiliser comme liveness et readiness probe
			probes.liveness = probe
			probes.readiness = &(*probe) // copie
		}
	}

	return &probes, nil
}

// generateResourceRequirements génère les exigences de ressources
func generateResourceRequirements(resources map[string]interface{}) (*ResourceRequirements, error) {
	reqs := &ResourceRequirements{
		Limits:   make(map[string]string),
		Requests: make(map[string]string),
	}

	if limits, ok := resources["limits"].(map[string]interface{}); ok {
		if memory, ok := limits["memory"].(string); ok {
			reqs.Limits["memory"] = memory
		}
		if cpus, ok := limits["cpus"].(string); ok {
			reqs.Limits["cpu"] = cpus
		}
	}

	if reservations, ok := resources["reservations"].(map[string]interface{}); ok {
		if memory, ok := reservations["memory"].(string); ok {
			reqs.Requests["memory"] = memory
		}
		if cpus, ok := reservations["cpus"].(string); ok {
			reqs.Requests["cpu"] = cpus
		}
	}

	// Retourner nil si aucune limite/demande n'est définie
	if len(reqs.Limits) == 0 && len(reqs.Requests) == 0 {
		return nil, nil
	}

	return reqs, nil
}

// generateSecurityContext génère le contexte de sécurité
func generateSecurityContext(service map[string]interface{}) (*SecurityContext, error) {
	var securityContext *SecurityContext

	// User
	if user, ok := service["user"].(string); ok {
		if securityContext == nil {
			securityContext = &SecurityContext{}
		}
		if userID, err := strconv.ParseInt(user, 10, 64); err == nil {
			securityContext.RunAsUser = &userID
		}
	}

	// Privileged
	if privileged, ok := service["privileged"].(bool); ok && privileged {
		if securityContext == nil {
			securityContext = &SecurityContext{}
		}
		securityContext.Privileged = &privileged
	}

	// Read-only root filesystem
	if readOnly, ok := service["read_only"].(bool); ok && readOnly {
		if securityContext == nil {
			securityContext = &SecurityContext{}
		}
		securityContext.ReadOnlyRootFilesystem = &readOnly
	}

	return securityContext, nil
}

// Fonctions utilitaires

func normalizeStringSlice(input interface{}) []string {
	switch v := input.(type) {
	case string:
		return strings.Fields(v)
	case []interface{}:
		var result []string
		for _, item := range v {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	case []string:
		return v
	default:
		return nil
	}
}

func mergeLabels(base, additional map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range base {
		result[k] = v
	}
	for k, v := range additional {
		result[k] = v
	}
	return result
}

func parseDurationToSeconds(duration string) (int32, error) {
	// Simple parser pour les durées comme "30s", "1m", etc.
	duration = strings.TrimSpace(duration)
	if duration == "" {
		return 0, fmt.Errorf("empty duration")
	}

	var multiplier int32 = 1
	var numStr string

	if strings.HasSuffix(duration, "s") {
		multiplier = 1
		numStr = strings.TrimSuffix(duration, "s")
	} else if strings.HasSuffix(duration, "m") {
		multiplier = 60
		numStr = strings.TrimSuffix(duration, "m")
	} else if strings.HasSuffix(duration, "h") {
		multiplier = 3600
		numStr = strings.TrimSuffix(duration, "h")
	} else {
		numStr = duration
	}

	num, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid duration format: %s", duration)
	}

	return int32(num) * multiplier, nil
}
