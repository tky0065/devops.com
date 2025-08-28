package kubernetes

import (
	"gopkg.in/yaml.v3"
)

// KubernetesObject interface commune pour tous les objets Kubernetes
type KubernetesObject interface {
	ToYAML() (string, error)
	GetName() string
	GetKind() string
}

// KubernetesManifest représente un manifest Kubernetes générique
type KubernetesManifest struct {
	APIVersion string                 `yaml:"apiVersion"`
	Kind       string                 `yaml:"kind"`
	Metadata   Metadata               `yaml:"metadata"`
	Spec       map[string]interface{} `yaml:"spec,omitempty"`
	Data       map[string]string      `yaml:"data,omitempty"`       // Pour ConfigMaps et Secrets
	StringData map[string]string      `yaml:"stringData,omitempty"` // Pour Secrets
}

// ToYAML convertit le manifest en YAML
func (m *KubernetesManifest) ToYAML() (string, error) {
	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

// GeneratedFile représente un fichier généré avec métadonnées
type GeneratedFile struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Size     int    `json:"size,omitempty"`
	Encoding string `json:"encoding,omitempty"`
}

// Metadata représente les métadonnées d'un objet Kubernetes
type Metadata struct {
	Name        string            `yaml:"name"`
	Namespace   string            `yaml:"namespace,omitempty"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}

// Deployment représente un Deployment Kubernetes
type Deployment struct {
	APIVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   Metadata       `yaml:"metadata"`
	Spec       DeploymentSpec `yaml:"spec"`
}

// ToYAML convertit le deployment en YAML
func (d *Deployment) ToYAML() (string, error) {
	yamlBytes, err := yaml.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

// GetName retourne le nom du deployment
func (d *Deployment) GetName() string {
	return d.Metadata.Name
}

// GetKind retourne le type d'objet
func (d *Deployment) GetKind() string {
	return d.Kind
}

// DeploymentSpec représente la spec d'un Deployment
type DeploymentSpec struct {
	Replicas int32               `yaml:"replicas,omitempty"`
	Selector *LabelSelector      `yaml:"selector"`
	Template PodTemplateSpec     `yaml:"template"`
	Strategy *DeploymentStrategy `yaml:"strategy,omitempty"`
}

// LabelSelector représente un sélecteur de labels
type LabelSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels,omitempty"`
}

// PodTemplateSpec représente le template d'un Pod
type PodTemplateSpec struct {
	Metadata Metadata `yaml:"metadata"`
	Spec     PodSpec  `yaml:"spec"`
}

// PodSpec représente la spec d'un Pod
type PodSpec struct {
	Containers                    []Container            `yaml:"containers"`
	InitContainers                []Container            `yaml:"initContainers,omitempty"`
	Volumes                       []Volume               `yaml:"volumes,omitempty"`
	RestartPolicy                 string                 `yaml:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds *int64                 `yaml:"terminationGracePeriodSeconds,omitempty"`
	DNSPolicy                     string                 `yaml:"dnsPolicy,omitempty"`
	ServiceAccountName            string                 `yaml:"serviceAccountName,omitempty"`
	SecurityContext               *PodSecurityContext    `yaml:"securityContext,omitempty"`
	ImagePullSecrets              []LocalObjectReference `yaml:"imagePullSecrets,omitempty"`
	Hostname                      string                 `yaml:"hostname,omitempty"`
	Subdomain                     string                 `yaml:"subdomain,omitempty"`
	NodeSelector                  map[string]string      `yaml:"nodeSelector,omitempty"`
	Tolerations                   []Toleration           `yaml:"tolerations,omitempty"`
	Affinity                      *Affinity              `yaml:"affinity,omitempty"`
}

// Container représente un conteneur
type Container struct {
	Name                     string                `yaml:"name"`
	Image                    string                `yaml:"image"`
	Command                  []string              `yaml:"command,omitempty"`
	Args                     []string              `yaml:"args,omitempty"`
	WorkingDir               string                `yaml:"workingDir,omitempty"`
	Ports                    []ContainerPort       `yaml:"ports,omitempty"`
	Env                      []EnvVar              `yaml:"env,omitempty"`
	EnvFrom                  []EnvFromSource       `yaml:"envFrom,omitempty"`
	Resources                *ResourceRequirements `yaml:"resources,omitempty"`
	VolumeMounts             []VolumeMount         `yaml:"volumeMounts,omitempty"`
	LivenessProbe            *Probe                `yaml:"livenessProbe,omitempty"`
	ReadinessProbe           *Probe                `yaml:"readinessProbe,omitempty"`
	StartupProbe             *Probe                `yaml:"startupProbe,omitempty"`
	SecurityContext          *SecurityContext      `yaml:"securityContext,omitempty"`
	ImagePullPolicy          string                `yaml:"imagePullPolicy,omitempty"`
	TerminationMessagePath   string                `yaml:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy string                `yaml:"terminationMessagePolicy,omitempty"`
	Stdin                    bool                  `yaml:"stdin,omitempty"`
	StdinOnce                bool                  `yaml:"stdinOnce,omitempty"`
	TTY                      bool                  `yaml:"tty,omitempty"`
}

// ContainerPort représente un port de conteneur
type ContainerPort struct {
	Name          string `yaml:"name,omitempty"`
	ContainerPort int32  `yaml:"containerPort"`
	Protocol      string `yaml:"protocol,omitempty"`
	HostIP        string `yaml:"hostIP,omitempty"`
	HostPort      int32  `yaml:"hostPort,omitempty"`
}

// EnvVar représente une variable d'environnement
type EnvVar struct {
	Name      string        `yaml:"name"`
	Value     string        `yaml:"value,omitempty"`
	ValueFrom *EnvVarSource `yaml:"valueFrom,omitempty"`
}

// EnvVarSource représente la source d'une variable d'environnement
type EnvVarSource struct {
	FieldRef         *ObjectFieldSelector   `yaml:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelector `yaml:"resourceFieldRef,omitempty"`
	ConfigMapKeyRef  *ConfigMapKeySelector  `yaml:"configMapKeyRef,omitempty"`
	SecretKeyRef     *SecretKeySelector     `yaml:"secretKeyRef,omitempty"`
}

// Service représente un Service Kubernetes
type Service struct {
	APIVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   Metadata    `yaml:"metadata"`
	Spec       ServiceSpec `yaml:"spec"`
}

// ToYAML convertit le service en YAML
func (s *Service) ToYAML() (string, error) {
	yamlBytes, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

// GetName retourne le nom du service
func (s *Service) GetName() string {
	return s.Metadata.Name
}

// GetKind retourne le type d'objet
func (s *Service) GetKind() string {
	return s.Kind
}

// ServiceSpec représente la spec d'un Service
type ServiceSpec struct {
	Type                     string            `yaml:"type,omitempty"`
	Selector                 map[string]string `yaml:"selector,omitempty"`
	Ports                    []ServicePort     `yaml:"ports,omitempty"`
	ClusterIP                string            `yaml:"clusterIP,omitempty"`
	ExternalIPs              []string          `yaml:"externalIPs,omitempty"`
	LoadBalancerIP           string            `yaml:"loadBalancerIP,omitempty"`
	LoadBalancerSourceRanges []string          `yaml:"loadBalancerSourceRanges,omitempty"`
	ExternalName             string            `yaml:"externalName,omitempty"`
	SessionAffinity          string            `yaml:"sessionAffinity,omitempty"`
}

// ServicePort représente un port de Service
type ServicePort struct {
	Name       string `yaml:"name,omitempty"`
	Port       int32  `yaml:"port"`
	TargetPort string `yaml:"targetPort,omitempty"` // peut être un nom ou un numéro
	Protocol   string `yaml:"protocol,omitempty"`
	NodePort   int32  `yaml:"nodePort,omitempty"`
}

// ConfigMap représente une ConfigMap Kubernetes
type ConfigMap struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   Metadata          `yaml:"metadata"`
	Data       map[string]string `yaml:"data,omitempty"`
	BinaryData map[string][]byte `yaml:"binaryData,omitempty"`
}

// ToYAML convertit la configmap en YAML
func (c *ConfigMap) ToYAML() (string, error) {
	yamlBytes, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

// GetName retourne le nom de la configmap
func (c *ConfigMap) GetName() string {
	return c.Metadata.Name
}

// GetKind retourne le type d'objet
func (c *ConfigMap) GetKind() string {
	return c.Kind
}

// PersistentVolume représente un PersistentVolume Kubernetes
type PersistentVolume struct {
	APIVersion string               `yaml:"apiVersion"`
	Kind       string               `yaml:"kind"`
	Metadata   Metadata             `yaml:"metadata"`
	Spec       PersistentVolumeSpec `yaml:"spec"`
}

// ToYAML convertit le PV en YAML
func (pv *PersistentVolume) ToYAML() (string, error) {
	yamlBytes, err := yaml.Marshal(pv)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

// GetName retourne le nom du PV
func (pv *PersistentVolume) GetName() string {
	return pv.Metadata.Name
}

// GetKind retourne le type d'objet
func (pv *PersistentVolume) GetKind() string {
	return pv.Kind
}

// PersistentVolumeSpec représente la spec d'un PersistentVolume
type PersistentVolumeSpec struct {
	Capacity                      map[string]string     `yaml:"capacity"`
	AccessModes                   []string              `yaml:"accessModes"`
	PersistentVolumeReclaimPolicy string                `yaml:"persistentVolumeReclaimPolicy,omitempty"`
	StorageClassName              string                `yaml:"storageClassName,omitempty"`
	MountOptions                  []string              `yaml:"mountOptions,omitempty"`
	VolumeMode                    string                `yaml:"volumeMode,omitempty"`
	NodeAffinity                  *VolumeNodeAffinity   `yaml:"nodeAffinity,omitempty"`
	HostPath                      *HostPathVolumeSource `yaml:"hostPath,omitempty"`
	NFS                           *NFSVolumeSource      `yaml:"nfs,omitempty"`
	ISCSI                         *ISCSIVolumeSource    `yaml:"iscsi,omitempty"`
	Local                         *LocalVolumeSource    `yaml:"local,omitempty"`
}

// PersistentVolumeClaim représente un PersistentVolumeClaim Kubernetes
type PersistentVolumeClaim struct {
	APIVersion string                    `yaml:"apiVersion"`
	Kind       string                    `yaml:"kind"`
	Metadata   Metadata                  `yaml:"metadata"`
	Spec       PersistentVolumeClaimSpec `yaml:"spec"`
}

// PersistentVolumeClaimSpec représente la spec d'un PersistentVolumeClaim
type PersistentVolumeClaimSpec struct {
	AccessModes      []string              `yaml:"accessModes"`
	Resources        *ResourceRequirements `yaml:"resources,omitempty"`
	VolumeName       string                `yaml:"volumeName,omitempty"`
	StorageClassName *string               `yaml:"storageClassName,omitempty"`
	VolumeMode       string                `yaml:"volumeMode,omitempty"`
	Selector         *LabelSelector        `yaml:"selector,omitempty"`
}

// Volume représente un volume dans un Pod
type Volume struct {
	Name                  string                 `yaml:"name"`
	HostPath              *HostPathVolumeSource  `yaml:"hostPath,omitempty"`
	EmptyDir              *EmptyDirVolumeSource  `yaml:"emptyDir,omitempty"`
	ConfigMap             *ConfigMapVolumeSource `yaml:"configMap,omitempty"`
	Secret                *SecretVolumeSource    `yaml:"secret,omitempty"`
	PersistentVolumeClaim *PVCVolumeSource       `yaml:"persistentVolumeClaim,omitempty"`
	NFS                   *NFSVolumeSource       `yaml:"nfs,omitempty"`
}

// VolumeMount représente un montage de volume
type VolumeMount struct {
	Name             string `yaml:"name"`
	MountPath        string `yaml:"mountPath"`
	SubPath          string `yaml:"subPath,omitempty"`
	SubPathExpr      string `yaml:"subPathExpr,omitempty"`
	ReadOnly         bool   `yaml:"readOnly,omitempty"`
	MountPropagation string `yaml:"mountPropagation,omitempty"`
}

// Probe représente une probe de santé
type Probe struct {
	Handler             Handler `yaml:",inline"`
	InitialDelaySeconds int32   `yaml:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32   `yaml:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32   `yaml:"periodSeconds,omitempty"`
	SuccessThreshold    int32   `yaml:"successThreshold,omitempty"`
	FailureThreshold    int32   `yaml:"failureThreshold,omitempty"`
}

// Handler représente un handler de probe
type Handler struct {
	Exec      *ExecAction      `yaml:"exec,omitempty"`
	HTTPGet   *HTTPGetAction   `yaml:"httpGet,omitempty"`
	TCPSocket *TCPSocketAction `yaml:"tcpSocket,omitempty"`
}

// ExecAction représente une action exec
type ExecAction struct {
	Command []string `yaml:"command,omitempty"`
}

// HTTPGetAction représente une action HTTP GET
type HTTPGetAction struct {
	Path        string       `yaml:"path,omitempty"`
	Port        string       `yaml:"port"`
	Host        string       `yaml:"host,omitempty"`
	Scheme      string       `yaml:"scheme,omitempty"`
	HTTPHeaders []HTTPHeader `yaml:"httpHeaders,omitempty"`
}

// HTTPHeader représente un header HTTP
type HTTPHeader struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// TCPSocketAction représente une action TCP Socket
type TCPSocketAction struct {
	Port string `yaml:"port"`
	Host string `yaml:"host,omitempty"`
}

// ResourceRequirements représente les exigences de ressources
type ResourceRequirements struct {
	Limits   map[string]string `yaml:"limits,omitempty"`
	Requests map[string]string `yaml:"requests,omitempty"`
}

// DeploymentStrategy représente la stratégie de déploiement
type DeploymentStrategy struct {
	Type          string                   `yaml:"type,omitempty"`
	RollingUpdate *RollingUpdateDeployment `yaml:"rollingUpdate,omitempty"`
}

// RollingUpdateDeployment représente les paramètres de rolling update
type RollingUpdateDeployment struct {
	MaxUnavailable string `yaml:"maxUnavailable,omitempty"`
	MaxSurge       string `yaml:"maxSurge,omitempty"`
}

// Sources de volumes
type HostPathVolumeSource struct {
	Path string `yaml:"path"`
	Type string `yaml:"type,omitempty"`
}

type EmptyDirVolumeSource struct {
	Medium    string `yaml:"medium,omitempty"`
	SizeLimit string `yaml:"sizeLimit,omitempty"`
}

type ConfigMapVolumeSource struct {
	LocalObjectReference `yaml:",inline"`
	Items                []KeyToPath `yaml:"items,omitempty"`
	DefaultMode          *int32      `yaml:"defaultMode,omitempty"`
	Optional             *bool       `yaml:"optional,omitempty"`
}

type SecretVolumeSource struct {
	SecretName  string      `yaml:"secretName,omitempty"`
	Items       []KeyToPath `yaml:"items,omitempty"`
	DefaultMode *int32      `yaml:"defaultMode,omitempty"`
	Optional    *bool       `yaml:"optional,omitempty"`
}

type PVCVolumeSource struct {
	ClaimName string `yaml:"claimName"`
	ReadOnly  bool   `yaml:"readOnly,omitempty"`
}

type NFSVolumeSource struct {
	Server   string `yaml:"server"`
	Path     string `yaml:"path"`
	ReadOnly bool   `yaml:"readOnly,omitempty"`
}

type ISCSIVolumeSource struct {
	TargetPortal   string   `yaml:"targetPortal"`
	IQN            string   `yaml:"iqn"`
	Lun            int32    `yaml:"lun"`
	ISCSIInterface string   `yaml:"iscsiInterface,omitempty"`
	FSType         string   `yaml:"fsType,omitempty"`
	ReadOnly       bool     `yaml:"readOnly,omitempty"`
	Portals        []string `yaml:"portals,omitempty"`
}

type LocalVolumeSource struct {
	Path   string `yaml:"path"`
	FSType string `yaml:"fsType,omitempty"`
}

// Objets utilitaires
type LocalObjectReference struct {
	Name string `yaml:"name,omitempty"`
}

type KeyToPath struct {
	Key  string `yaml:"key"`
	Path string `yaml:"path"`
	Mode *int32 `yaml:"mode,omitempty"`
}

type ObjectFieldSelector struct {
	APIVersion string `yaml:"apiVersion,omitempty"`
	FieldPath  string `yaml:"fieldPath"`
}

type ResourceFieldSelector struct {
	ContainerName string `yaml:"containerName,omitempty"`
	Resource      string `yaml:"resource"`
	Divisor       string `yaml:"divisor,omitempty"`
}

type ConfigMapKeySelector struct {
	LocalObjectReference `yaml:",inline"`
	Key                  string `yaml:"key"`
	Optional             *bool  `yaml:"optional,omitempty"`
}

type SecretKeySelector struct {
	LocalObjectReference `yaml:",inline"`
	Key                  string `yaml:"key"`
	Optional             *bool  `yaml:"optional,omitempty"`
}

type EnvFromSource struct {
	Prefix       string              `yaml:"prefix,omitempty"`
	ConfigMapRef *ConfigMapEnvSource `yaml:"configMapRef,omitempty"`
	SecretRef    *SecretEnvSource    `yaml:"secretRef,omitempty"`
}

type ConfigMapEnvSource struct {
	LocalObjectReference `yaml:",inline"`
	Optional             *bool `yaml:"optional,omitempty"`
}

type SecretEnvSource struct {
	LocalObjectReference `yaml:",inline"`
	Optional             *bool `yaml:"optional,omitempty"`
}

// Contextes de sécurité
type SecurityContext struct {
	Capabilities             *Capabilities                  `yaml:"capabilities,omitempty"`
	Privileged               *bool                          `yaml:"privileged,omitempty"`
	SELinuxOptions           *SELinuxOptions                `yaml:"seLinuxOptions,omitempty"`
	WindowsOptions           *WindowsSecurityContextOptions `yaml:"windowsOptions,omitempty"`
	RunAsUser                *int64                         `yaml:"runAsUser,omitempty"`
	RunAsGroup               *int64                         `yaml:"runAsGroup,omitempty"`
	RunAsNonRoot             *bool                          `yaml:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem   *bool                          `yaml:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation *bool                          `yaml:"allowPrivilegeEscalation,omitempty"`
	ProcMount                string                         `yaml:"procMount,omitempty"`
	SeccompProfile           *SeccompProfile                `yaml:"seccompProfile,omitempty"`
}

type PodSecurityContext struct {
	SELinuxOptions      *SELinuxOptions                `yaml:"seLinuxOptions,omitempty"`
	WindowsOptions      *WindowsSecurityContextOptions `yaml:"windowsOptions,omitempty"`
	RunAsUser           *int64                         `yaml:"runAsUser,omitempty"`
	RunAsGroup          *int64                         `yaml:"runAsGroup,omitempty"`
	RunAsNonRoot        *bool                          `yaml:"runAsNonRoot,omitempty"`
	SupplementalGroups  []int64                        `yaml:"supplementalGroups,omitempty"`
	FSGroup             *int64                         `yaml:"fsGroup,omitempty"`
	Sysctls             []Sysctl                       `yaml:"sysctls,omitempty"`
	FSGroupChangePolicy string                         `yaml:"fsGroupChangePolicy,omitempty"`
	SeccompProfile      *SeccompProfile                `yaml:"seccompProfile,omitempty"`
}

type Capabilities struct {
	Add  []string `yaml:"add,omitempty"`
	Drop []string `yaml:"drop,omitempty"`
}

type SELinuxOptions struct {
	User  string `yaml:"user,omitempty"`
	Role  string `yaml:"role,omitempty"`
	Type  string `yaml:"type,omitempty"`
	Level string `yaml:"level,omitempty"`
}

type WindowsSecurityContextOptions struct {
	GMSACredentialSpecName string `yaml:"gmsaCredentialSpecName,omitempty"`
	GMSACredentialSpec     string `yaml:"gmsaCredentialSpec,omitempty"`
	RunAsUserName          string `yaml:"runAsUserName,omitempty"`
	HostProcess            *bool  `yaml:"hostProcess,omitempty"`
}

type SeccompProfile struct {
	Type             string `yaml:"type"`
	LocalhostProfile string `yaml:"localhostProfile,omitempty"`
}

type Sysctl struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Affinité et tolérances
type Affinity struct {
	NodeAffinity    *NodeAffinity    `yaml:"nodeAffinity,omitempty"`
	PodAffinity     *PodAffinity     `yaml:"podAffinity,omitempty"`
	PodAntiAffinity *PodAntiAffinity `yaml:"podAntiAffinity,omitempty"`
}

type NodeAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  *NodeSelector             `yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTerm `yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

type PodAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm         `yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

type PodAntiAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm         `yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

type NodeSelector struct {
	NodeSelectorTerms []NodeSelectorTerm `yaml:"nodeSelectorTerms"`
}

type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement `yaml:"matchExpressions,omitempty"`
	MatchFields      []NodeSelectorRequirement `yaml:"matchFields,omitempty"`
}

type NodeSelectorRequirement struct {
	Key      string   `yaml:"key"`
	Operator string   `yaml:"operator"`
	Values   []string `yaml:"values,omitempty"`
}

type PreferredSchedulingTerm struct {
	Weight     int32            `yaml:"weight"`
	Preference NodeSelectorTerm `yaml:"preference"`
}

type PodAffinityTerm struct {
	LabelSelector     *LabelSelector `yaml:"labelSelector,omitempty"`
	Namespaces        []string       `yaml:"namespaces,omitempty"`
	TopologyKey       string         `yaml:"topologyKey"`
	NamespaceSelector *LabelSelector `yaml:"namespaceSelector,omitempty"`
}

type WeightedPodAffinityTerm struct {
	Weight          int32           `yaml:"weight"`
	PodAffinityTerm PodAffinityTerm `yaml:"podAffinityTerm"`
}

type Toleration struct {
	Key               string `yaml:"key,omitempty"`
	Operator          string `yaml:"operator,omitempty"`
	Value             string `yaml:"value,omitempty"`
	Effect            string `yaml:"effect,omitempty"`
	TolerationSeconds *int64 `yaml:"tolerationSeconds,omitempty"`
}

type VolumeNodeAffinity struct {
	Required *NodeSelector `yaml:"required,omitempty"`
}
