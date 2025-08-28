package kubernetes

import (
	"fmt"
	"strconv"
	"strings"
)

// GenerateService génère un Service Kubernetes à partir d'un service Docker Compose
func GenerateService(serviceName string, service interface{}, options GeneratorOptions) (*Service, error) {
	serviceMap, ok := service.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid service format for %s", serviceName)
	}

	// Vérifier si le service expose des ports
	ports, ok := serviceMap["ports"]
	if !ok {
		// Pas de ports exposés, pas de Service nécessaire
		return nil, nil
	}

	portSlice, ok := ports.([]interface{})
	if !ok || len(portSlice) == 0 {
		return nil, nil
	}

	kubernetesService := &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name:      serviceName,
			Namespace: options.Namespace,
			Labels:    mergeLabels(options.Labels, map[string]string{"app": serviceName}),
		},
		Spec: ServiceSpec{
			Type:     options.ServiceType,
			Selector: map[string]string{"app": serviceName},
		},
	}

	// Générer les ports du service
	servicePorts, err := generateServicePorts(portSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to generate service ports for %s: %w", serviceName, err)
	}

	kubernetesService.Spec.Ports = servicePorts

	return kubernetesService, nil
}

// generateServicePorts génère les ports du service à partir des ports Docker Compose
func generateServicePorts(ports []interface{}) ([]ServicePort, error) {
	var servicePorts []ServicePort

	for i, port := range ports {
		portStr, ok := port.(string)
		if !ok {
			continue
		}

		servicePort, err := parseServicePortMapping(portStr, i)
		if err != nil {
			return nil, fmt.Errorf("invalid port format %s: %w", portStr, err)
		}

		servicePorts = append(servicePorts, *servicePort)
	}

	return servicePorts, nil
}

// parseServicePortMapping parse une mapping de port pour créer un ServicePort
func parseServicePortMapping(portMapping string, index int) (*ServicePort, error) {
	parts := strings.Split(portMapping, ":")
	
	var hostPort, containerPort int32
	var protocol = "TCP"
	var portName string

	switch len(parts) {
	case 1:
		// Format: "8080"
		port, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid port number: %s", parts[0])
		}
		hostPort = int32(port)
		containerPort = int32(port)
	case 2:
		// Format: "8080:80"
		hPort, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid host port number: %s", parts[0])
		}
		cPort, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid container port number: %s", parts[1])
		}
		hostPort = int32(hPort)
		containerPort = int32(cPort)
	case 3:
		// Format: "127.0.0.1:8080:80" ou "8080:80:tcp"
		if strings.Contains(parts[2], "tcp") || strings.Contains(parts[2], "udp") {
			protocol = strings.ToUpper(parts[2])
			hPort, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid host port number: %s", parts[0])
			}
			cPort, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid container port number: %s", parts[1])
			}
			hostPort = int32(hPort)
			containerPort = int32(cPort)
		} else {
			// IP:host:container
			hPort, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid host port number: %s", parts[1])
			}
			cPort, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, fmt.Errorf("invalid container port number: %s", parts[2])
			}
			hostPort = int32(hPort)
			containerPort = int32(cPort)
		}
	default:
		return nil, fmt.Errorf("invalid port mapping format: %s", portMapping)
	}

	// Générer un nom de port basé sur le protocole et le numéro
	portName = fmt.Sprintf("%s-%d", strings.ToLower(protocol), containerPort)

	return &ServicePort{
		Name:       portName,
		Port:       hostPort,
		TargetPort: strconv.Itoa(int(containerPort)),
		Protocol:   protocol,
	}, nil
}

// GenerateConfigMap génère une ConfigMap pour les variables d'environnement
func GenerateConfigMap(serviceName string, service interface{}, options GeneratorOptions) (*ConfigMap, error) {
	serviceMap, ok := service.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid service format for %s", serviceName)
	}

	// Vérifier s'il y a des variables d'environnement
	env, ok := serviceMap["environment"]
	if !ok {
		return nil, nil
	}

	envMap, err := normalizeEnvironmentForConfigMap(env)
	if err != nil {
		return nil, fmt.Errorf("failed to normalize environment for %s: %w", serviceName, err)
	}

	if len(envMap) == 0 {
		return nil, nil
	}

	configMap := &ConfigMap{
		APIVersion: "v1",
		Kind:       "ConfigMap",
		Metadata: Metadata{
			Name:      fmt.Sprintf("%s-config", serviceName),
			Namespace: options.Namespace,
			Labels:    mergeLabels(options.Labels, map[string]string{"app": serviceName}),
		},
		Data: envMap,
	}

	return configMap, nil
}

// normalizeEnvironmentForConfigMap normalise les variables d'environnement pour ConfigMap
func normalizeEnvironmentForConfigMap(env interface{}) (map[string]string, error) {
	envMap := make(map[string]string)

	switch e := env.(type) {
	case map[string]interface{}:
		for key, value := range e {
			// Exclure les variables qui contiennent des secrets (mots de passe, clés, etc.)
			if isSecretVariable(key) {
				continue
			}
			envMap[key] = fmt.Sprintf("%v", value)
		}
	case map[string]string:
		for key, value := range e {
			// Exclure les variables qui contiennent des secrets (mots de passe, clés, etc.)
			if isSecretVariable(key) {
				continue
			}
			envMap[key] = value
		}
	case []interface{}:
		for _, item := range e {
			envStr, ok := item.(string)
			if !ok {
				continue
			}
			parts := strings.SplitN(envStr, "=", 2)
			key := parts[0]
			
			// Exclure les variables qui contiennent des secrets
			if isSecretVariable(key) {
				continue
			}
			
			if len(parts) == 2 {
				envMap[key] = parts[1]
			} else {
				envMap[key] = ""
			}
		}
	default:
		return nil, fmt.Errorf("unsupported environment format: %T", env)
	}

	return envMap, nil
}

// isSecretVariable détermine si une variable d'environnement contient des données sensibles
func isSecretVariable(key string) bool {
	key = strings.ToLower(key)
	secretPatterns := []string{
		"password", "passwd", "pwd",
		"secret", "key", "token",
		"api_key", "apikey",
		"private", "credential",
		"auth", "oauth",
	}

	for _, pattern := range secretPatterns {
		if strings.Contains(key, pattern) {
			return true
		}
	}

	return false
}

// GeneratePersistentVolumeClaim génère un PVC pour les volumes nommés
func GeneratePersistentVolumeClaim(serviceName string, service interface{}, options GeneratorOptions) ([]*PersistentVolumeClaim, error) {
	serviceMap, ok := service.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid service format for %s", serviceName)
	}

	volumesConfig, ok := serviceMap["volumes"]
	if !ok {
		return nil, nil
	}

	volumeSlice, ok := volumesConfig.([]interface{})
	if !ok {
		return nil, nil
	}

	var pvcs []*PersistentVolumeClaim

	for _, vol := range volumeSlice {
		volStr, ok := vol.(string)
		if !ok {
			continue
		}

		// Ne créer des PVCs que pour les volumes nommés (pas les bind mounts)
		if !strings.HasPrefix(volStr, "/") && !strings.HasPrefix(volStr, ".") {
			parts := strings.Split(volStr, ":")
			if len(parts) >= 2 {
				volumeName := parts[0]
				
				pvc := &PersistentVolumeClaim{
					APIVersion: "v1",
					Kind:       "PersistentVolumeClaim",
					Metadata: Metadata{
						Name:      fmt.Sprintf("%s-%s", serviceName, volumeName),
						Namespace: options.Namespace,
						Labels:    mergeLabels(options.Labels, map[string]string{"app": serviceName}),
					},
					Spec: PersistentVolumeClaimSpec{
						AccessModes: []string{"ReadWriteOnce"},
						Resources: &ResourceRequirements{
							Requests: map[string]string{
								"storage": "1Gi", // Taille par défaut
							},
						},
					},
				}

				pvcs = append(pvcs, pvc)
			}
		}
	}

	return pvcs, nil
}

// GenerateIngressForService génère un Ingress basique pour un service avec des ports HTTP
func GenerateIngressForService(serviceName string, service interface{}, options GeneratorOptions, host string) (*KubernetesManifest, error) {
	serviceMap, ok := service.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid service format for %s", serviceName)
	}

	// Vérifier s'il y a des ports HTTP
	ports, ok := serviceMap["ports"]
	if !ok {
		return nil, nil
	}

	portSlice, ok := ports.([]interface{})
	if !ok || len(portSlice) == 0 {
		return nil, nil
	}

	// Chercher un port HTTP (80, 8080, 3000, etc.)
	var httpPort int32
	for _, port := range portSlice {
		portStr, ok := port.(string)
		if !ok {
			continue
		}

		parts := strings.Split(portStr, ":")
		var containerPort int
		var err error

		switch len(parts) {
		case 1:
			containerPort, err = strconv.Atoi(parts[0])
		case 2:
			containerPort, err = strconv.Atoi(parts[1])
		case 3:
			if strings.Contains(parts[2], "tcp") || strings.Contains(parts[2], "udp") {
				containerPort, err = strconv.Atoi(parts[1])
			} else {
				containerPort, err = strconv.Atoi(parts[2])
			}
		}

		if err == nil && isHTTPPort(containerPort) {
			httpPort = int32(containerPort)
			break
		}
	}

	if httpPort == 0 {
		return nil, nil // Pas de port HTTP trouvé
	}

	if host == "" {
		host = fmt.Sprintf("%s.local", serviceName)
	}

	ingress := &KubernetesManifest{
		APIVersion: "networking.k8s.io/v1",
		Kind:       "Ingress",
		Metadata: Metadata{
			Name:      fmt.Sprintf("%s-ingress", serviceName),
			Namespace: options.Namespace,
			Labels:    mergeLabels(options.Labels, map[string]string{"app": serviceName}),
		},
		Spec: map[string]interface{}{
			"rules": []map[string]interface{}{
				{
					"host": host,
					"http": map[string]interface{}{
						"paths": []map[string]interface{}{
							{
								"path":     "/",
								"pathType": "Prefix",
								"backend": map[string]interface{}{
									"service": map[string]interface{}{
										"name": serviceName,
										"port": map[string]interface{}{
											"number": httpPort,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return ingress, nil
}

// isHTTPPort détermine si un port est probablement un port HTTP
func isHTTPPort(port int) bool {
	commonHTTPPorts := []int{80, 8080, 3000, 3001, 4000, 5000, 8000, 8888, 9000}
	for _, httpPort := range commonHTTPPorts {
		if port == httpPort {
			return true
		}
	}
	return false
}
