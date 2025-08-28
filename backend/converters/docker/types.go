package docker

import "time"

// DockerCompose représente la structure d'un fichier docker-compose.yml
type DockerCompose struct {
	Version  string                    `yaml:"version"`
	Services map[string]Service        `yaml:"services"`
	Volumes  map[string]Volume         `yaml:"volumes,omitempty"`
	Networks map[string]Network        `yaml:"networks,omitempty"`
	Configs  map[string]Config         `yaml:"configs,omitempty"`
	Secrets  map[string]Secret         `yaml:"secrets,omitempty"`
}

// Service représente un service dans docker-compose
type Service struct {
	Image         string                 `yaml:"image,omitempty"`
	Build         interface{}            `yaml:"build,omitempty"` // string ou BuildConfig
	ContainerName string                 `yaml:"container_name,omitempty"`
	Ports         []string               `yaml:"ports,omitempty"`
	Expose        []string               `yaml:"expose,omitempty"`
	Environment   interface{}            `yaml:"environment,omitempty"` // []string ou map[string]string
	EnvFile       interface{}            `yaml:"env_file,omitempty"`    // string ou []string
	Volumes       []string               `yaml:"volumes,omitempty"`
	Networks      interface{}            `yaml:"networks,omitempty"` // []string ou map[string]NetworkConfig
	DependsOn     interface{}            `yaml:"depends_on,omitempty"` // []string ou map[string]DependencyConfig
	Command       interface{}            `yaml:"command,omitempty"`    // string ou []string
	Entrypoint    interface{}            `yaml:"entrypoint,omitempty"` // string ou []string
	WorkingDir    string                 `yaml:"working_dir,omitempty"`
	User          string                 `yaml:"user,omitempty"`
	Restart       string                 `yaml:"restart,omitempty"`
	Labels        map[string]string      `yaml:"labels,omitempty"`
	HealthCheck   *HealthCheck           `yaml:"healthcheck,omitempty"`
	Deploy        *DeployConfig          `yaml:"deploy,omitempty"`
	Resources     *ResourcesConfig       `yaml:"resources,omitempty"`
	Logging       *LoggingConfig         `yaml:"logging,omitempty"`
	Tmpfs         []string               `yaml:"tmpfs,omitempty"`
	Ulimits       map[string]interface{} `yaml:"ulimits,omitempty"`
	StdinOpen     bool                   `yaml:"stdin_open,omitempty"`
	Tty           bool                   `yaml:"tty,omitempty"`
	Privileged    bool                   `yaml:"privileged,omitempty"`
	ReadOnly      bool                   `yaml:"read_only,omitempty"`
	ShmSize       string                 `yaml:"shm_size,omitempty"`
	PidMode       string                 `yaml:"pid,omitempty"`
	IpcMode       string                 `yaml:"ipc,omitempty"`
}

// BuildConfig représente la configuration de build
type BuildConfig struct {
	Context      string            `yaml:"context,omitempty"`
	Dockerfile   string            `yaml:"dockerfile,omitempty"`
	Args         map[string]string `yaml:"args,omitempty"`
	Target       string            `yaml:"target,omitempty"`
	CacheFrom    []string          `yaml:"cache_from,omitempty"`
	Labels       map[string]string `yaml:"labels,omitempty"`
	ShmSize      string            `yaml:"shm_size,omitempty"`
}

// NetworkConfig représente la configuration réseau d'un service
type NetworkConfig struct {
	Aliases     []string `yaml:"aliases,omitempty"`
	Ipv4Address string   `yaml:"ipv4_address,omitempty"`
	Ipv6Address string   `yaml:"ipv6_address,omitempty"`
}

// DependencyConfig représente la configuration des dépendances
type DependencyConfig struct {
	Condition string `yaml:"condition,omitempty"`
}

// HealthCheck représente la configuration de health check
type HealthCheck struct {
	Test        interface{}   `yaml:"test,omitempty"`        // string ou []string
	Interval    time.Duration `yaml:"interval,omitempty"`
	Timeout     time.Duration `yaml:"timeout,omitempty"`
	Retries     int           `yaml:"retries,omitempty"`
	StartPeriod time.Duration `yaml:"start_period,omitempty"`
	Disable     bool          `yaml:"disable,omitempty"`
}

// DeployConfig représente la configuration de déploiement
type DeployConfig struct {
	Mode          string                 `yaml:"mode,omitempty"`
	Replicas      int                    `yaml:"replicas,omitempty"`
	Labels        map[string]string      `yaml:"labels,omitempty"`
	UpdateConfig  *UpdateConfig          `yaml:"update_config,omitempty"`
	Resources     *ResourcesConfig       `yaml:"resources,omitempty"`
	RestartPolicy *RestartPolicyConfig   `yaml:"restart_policy,omitempty"`
	Placement     *PlacementConfig       `yaml:"placement,omitempty"`
	EndpointMode  string                 `yaml:"endpoint_mode,omitempty"`
}

// UpdateConfig représente la configuration de mise à jour
type UpdateConfig struct {
	Parallelism     int           `yaml:"parallelism,omitempty"`
	Delay           time.Duration `yaml:"delay,omitempty"`
	FailureAction   string        `yaml:"failure_action,omitempty"`
	Monitor         time.Duration `yaml:"monitor,omitempty"`
	MaxFailureRatio float64       `yaml:"max_failure_ratio,omitempty"`
	Order           string        `yaml:"order,omitempty"`
}

// ResourcesConfig représente la configuration des ressources
type ResourcesConfig struct {
	Limits       *ResourceLimits `yaml:"limits,omitempty"`
	Reservations *ResourceLimits `yaml:"reservations,omitempty"`
}

// ResourceLimits représente les limites de ressources
type ResourceLimits struct {
	CPUs     string `yaml:"cpus,omitempty"`
	Memory   string `yaml:"memory,omitempty"`
	Pids     int    `yaml:"pids,omitempty"`
}

// RestartPolicyConfig représente la politique de redémarrage
type RestartPolicyConfig struct {
	Condition   string        `yaml:"condition,omitempty"`
	Delay       time.Duration `yaml:"delay,omitempty"`
	MaxAttempts int           `yaml:"max_attempts,omitempty"`
	Window      time.Duration `yaml:"window,omitempty"`
}

// PlacementConfig représente la configuration de placement
type PlacementConfig struct {
	Constraints []string               `yaml:"constraints,omitempty"`
	Preferences []map[string]string    `yaml:"preferences,omitempty"`
}

// LoggingConfig représente la configuration de logging
type LoggingConfig struct {
	Driver  string            `yaml:"driver,omitempty"`
	Options map[string]string `yaml:"options,omitempty"`
}

// Volume représente un volume docker
type Volume struct {
	Driver     string            `yaml:"driver,omitempty"`
	DriverOpts map[string]string `yaml:"driver_opts,omitempty"`
	External   interface{}       `yaml:"external,omitempty"` // bool ou ExternalConfig
	Labels     map[string]string `yaml:"labels,omitempty"`
	Name       string            `yaml:"name,omitempty"`
}

// Network représente un réseau docker
type Network struct {
	Driver     string                 `yaml:"driver,omitempty"`
	DriverOpts map[string]string      `yaml:"driver_opts,omitempty"`
	IPAM       *IPAMConfig            `yaml:"ipam,omitempty"`
	External   interface{}            `yaml:"external,omitempty"` // bool ou ExternalConfig
	Internal   bool                   `yaml:"internal,omitempty"`
	Attachable bool                   `yaml:"attachable,omitempty"`
	EnableIPv6 bool                   `yaml:"enable_ipv6,omitempty"`
	Labels     map[string]string      `yaml:"labels,omitempty"`
	Name       string                 `yaml:"name,omitempty"`
}

// IPAMConfig représente la configuration IPAM
type IPAMConfig struct {
	Driver string       `yaml:"driver,omitempty"`
	Config []IPAMPool   `yaml:"config,omitempty"`
	Options map[string]string `yaml:"options,omitempty"`
}

// IPAMPool représente un pool IPAM
type IPAMPool struct {
	Subnet     string `yaml:"subnet,omitempty"`
	IPRange    string `yaml:"ip_range,omitempty"`
	Gateway    string `yaml:"gateway,omitempty"`
	AuxAddress map[string]string `yaml:"aux_addresses,omitempty"`
}

// Config représente une configuration Docker
type Config struct {
	File     string            `yaml:"file,omitempty"`
	External interface{}       `yaml:"external,omitempty"` // bool ou ExternalConfig
	Labels   map[string]string `yaml:"labels,omitempty"`
	Name     string            `yaml:"name,omitempty"`
}

// Secret représente un secret Docker
type Secret struct {
	File     string            `yaml:"file,omitempty"`
	External interface{}       `yaml:"external,omitempty"` // bool ou ExternalConfig
	Labels   map[string]string `yaml:"labels,omitempty"`
	Name     string            `yaml:"name,omitempty"`
}

// ExternalConfig représente la configuration d'une ressource externe
type ExternalConfig struct {
	Name string `yaml:"name,omitempty"`
}
