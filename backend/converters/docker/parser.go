package docker

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseDockerCompose parse un fichier docker-compose.yml
func ParseDockerCompose(content string) (*DockerCompose, error) {
	var compose DockerCompose
	
	if err := yaml.Unmarshal([]byte(content), &compose); err != nil {
		return nil, fmt.Errorf("failed to parse docker-compose file: %w", err)
	}

	// Valider la version
	if compose.Version == "" {
		compose.Version = "3.8" // Version par défaut
	}

	// Normaliser les services
	if err := normalizeServices(&compose); err != nil {
		return nil, fmt.Errorf("failed to normalize services: %w", err)
	}

	return &compose, nil
}

// normalizeServices normalise la structure des services
func normalizeServices(compose *DockerCompose) error {
	for serviceName, service := range compose.Services {
		// Normaliser les ports
		normalizedPorts, err := normalizePorts(service.Ports)
		if err != nil {
			return fmt.Errorf("failed to normalize ports for service %s: %w", serviceName, err)
		}
		service.Ports = normalizedPorts

		// Normaliser les variables d'environnement
		normalizedEnv, err := normalizeEnvironment(service.Environment)
		if err != nil {
			return fmt.Errorf("failed to normalize environment for service %s: %w", serviceName, err)
		}
		service.Environment = normalizedEnv

		// Normaliser les réseaux
		normalizedNetworks, err := normalizeNetworks(service.Networks)
		if err != nil {
			return fmt.Errorf("failed to normalize networks for service %s: %w", serviceName, err)
		}
		service.Networks = normalizedNetworks

		// Normaliser depends_on
		normalizedDeps, err := normalizeDependsOn(service.DependsOn)
		if err != nil {
			return fmt.Errorf("failed to normalize depends_on for service %s: %w", serviceName, err)
		}
		service.DependsOn = normalizedDeps

		// Normaliser les commandes
		service.Command = normalizeCommand(service.Command)
		service.Entrypoint = normalizeCommand(service.Entrypoint)

		// Mettre à jour le service dans la map
		compose.Services[serviceName] = service
	}

	return nil
}

// normalizePorts normalise les définitions de ports
func normalizePorts(ports []string) ([]string, error) {
	var normalized []string
	
	for _, port := range ports {
		// Valider le format du port
		if strings.Contains(port, ":") {
			parts := strings.Split(port, ":")
			if len(parts) < 2 || len(parts) > 3 {
				return nil, fmt.Errorf("invalid port format: %s", port)
			}
			
			// Valider que les parties sont des nombres valides
			for i, part := range parts {
				if part == "" {
					continue
				}
				if i > 0 { // Skip protocol part for host:container:protocol format
					if _, err := strconv.Atoi(part); err != nil {
						return nil, fmt.Errorf("invalid port number: %s", part)
					}
				}
			}
		} else {
			// Port simple
			if _, err := strconv.Atoi(port); err != nil {
				return nil, fmt.Errorf("invalid port number: %s", port)
			}
		}
		
		normalized = append(normalized, port)
	}
	
	return normalized, nil
}

// normalizeEnvironment normalise les variables d'environnement
func normalizeEnvironment(env interface{}) (map[string]string, error) {
	if env == nil {
		return nil, nil
	}

	switch e := env.(type) {
	case map[string]interface{}:
		result := make(map[string]string)
		for k, v := range e {
			result[k] = fmt.Sprintf("%v", v)
		}
		return result, nil
	case map[string]string:
		return e, nil
	case []interface{}:
		result := make(map[string]string)
		for _, item := range e {
			str, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("invalid environment variable format: %v", item)
			}
			parts := strings.SplitN(str, "=", 2)
			if len(parts) == 2 {
				result[parts[0]] = parts[1]
			} else {
				result[parts[0]] = ""
			}
		}
		return result, nil
	case []string:
		result := make(map[string]string)
		for _, item := range e {
			parts := strings.SplitN(item, "=", 2)
			if len(parts) == 2 {
				result[parts[0]] = parts[1]
			} else {
				result[parts[0]] = ""
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported environment format: %T", env)
	}
}

// normalizeNetworks normalise la configuration des réseaux
func normalizeNetworks(networks interface{}) (map[string]NetworkConfig, error) {
	if networks == nil {
		return nil, nil
	}

	switch n := networks.(type) {
	case []interface{}:
		result := make(map[string]NetworkConfig)
		for _, item := range n {
			if str, ok := item.(string); ok {
				result[str] = NetworkConfig{}
			}
		}
		return result, nil
	case []string:
		result := make(map[string]NetworkConfig)
		for _, item := range n {
			result[item] = NetworkConfig{}
		}
		return result, nil
	case map[string]interface{}:
		result := make(map[string]NetworkConfig)
		for name, config := range n {
			if config == nil {
				result[name] = NetworkConfig{}
				continue
			}
			
			configMap, ok := config.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid network config for %s", name)
			}
			
			netConfig := NetworkConfig{}
			if aliases, ok := configMap["aliases"].([]interface{}); ok {
				for _, alias := range aliases {
					if str, ok := alias.(string); ok {
						netConfig.Aliases = append(netConfig.Aliases, str)
					}
				}
			}
			if ipv4, ok := configMap["ipv4_address"].(string); ok {
				netConfig.Ipv4Address = ipv4
			}
			if ipv6, ok := configMap["ipv6_address"].(string); ok {
				netConfig.Ipv6Address = ipv6
			}
			
			result[name] = netConfig
		}
		return result, nil
	case map[string]NetworkConfig:
		return n, nil
	default:
		return nil, fmt.Errorf("unsupported networks format: %T", networks)
	}
}

// normalizeDependsOn normalise la configuration des dépendances
func normalizeDependsOn(deps interface{}) (map[string]DependencyConfig, error) {
	if deps == nil {
		return nil, nil
	}

	switch d := deps.(type) {
	case []interface{}:
		result := make(map[string]DependencyConfig)
		for _, item := range d {
			if str, ok := item.(string); ok {
				result[str] = DependencyConfig{}
			}
		}
		return result, nil
	case []string:
		result := make(map[string]DependencyConfig)
		for _, item := range d {
			result[item] = DependencyConfig{}
		}
		return result, nil
	case map[string]interface{}:
		result := make(map[string]DependencyConfig)
		for name, config := range d {
			if config == nil {
				result[name] = DependencyConfig{}
				continue
			}
			
			if configMap, ok := config.(map[string]interface{}); ok {
				depConfig := DependencyConfig{}
				if condition, ok := configMap["condition"].(string); ok {
					depConfig.Condition = condition
				}
				result[name] = depConfig
			} else {
				result[name] = DependencyConfig{}
			}
		}
		return result, nil
	case map[string]DependencyConfig:
		return d, nil
	default:
		return nil, fmt.Errorf("unsupported depends_on format: %T", deps)
	}
}

// normalizeCommand normalise les commandes et entrypoints
func normalizeCommand(cmd interface{}) []string {
	if cmd == nil {
		return nil
	}

	switch c := cmd.(type) {
	case string:
		// Commande simple en string - la diviser en arguments
		return strings.Fields(c)
	case []interface{}:
		var result []string
		for _, item := range c {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	case []string:
		return c
	default:
		return nil
	}
}
