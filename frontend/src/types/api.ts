// Types pour l'API de conversion
export interface ConversionRequest {
  type: string
  content: string
  options?: ConversionOptions
}

export interface ConversionOptions {
  namespace?: string
  serviceType?: 'ClusterIP' | 'NodePort' | 'LoadBalancer'
  replicas?: number
  [key: string]: any
}

export interface ConversionResponse {
  success: boolean
  files: GeneratedFile[]
  errors?: ApiError[]
  warnings?: ApiWarning[]
  metadata?: ConversionMetadata
  request_id: string
}

export interface GeneratedFile {
  name: string
  content: string
  type: string
  path: string
}

export interface ApiError {
  code: string
  message: string
}

export interface ApiWarning {
  code: string
  message: string
}

export interface ConversionMetadata {
  docker_version?: string
  services_converted: number
  volumes_converted: number
}

export interface ValidationRequest {
  type: string
  content: string
}

export interface ValidationResponse {
  valid: boolean
  message: string
  errors?: string[]
}

export interface ConverterInfo {
  name: string
  description: string
  supported_types: string[]
}

export interface ConvertersResponse {
  converters: ConverterInfo[]
  count: number
  success: boolean
}

export interface HealthResponse {
  status: string
  message: string
  timestamp: string
}

export interface HealthDetailedResponse extends HealthResponse {
  version: string
  environment: string
  uptime: string
  system: {
    go_version: string
    num_goroutine: number
    num_cpu: number
    memory_usage: {
      alloc: number
      total_alloc: number
      sys: number
      num_gc: number
    }
  }
  services: {
    converters: {
      status: string
      message: string
      details: {
        available_converters: number
        converter_list: ConverterInfo[]
      }
    }
  }
}

export interface VersionResponse {
  name: string
  version: string
  go_version: string
  build_time: string
  git_commit: string
  environment: string
}
