import axios, { type AxiosInstance, type AxiosResponse } from 'axios'
import type {
  ConversionRequest,
  ConversionResponse,
  ValidationRequest,
  ValidationResponse,
  ConvertersResponse,
  HealthResponse,
  HealthDetailedResponse,
  VersionResponse
} from '@/types/api'

class ApiService {
  private api: AxiosInstance

  constructor() {
    this.api = axios.create({
      baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8081',
      timeout: 30000,
      headers: {
        'Content-Type': 'application/json'
      }
    })

    // Intercepteur pour gérer les erreurs globalement
    this.api.interceptors.response.use(
      (response) => response,
      (error) => {
        console.error('API Error:', error)
        return Promise.reject(error)
      }
    )
  }

  // Conversion
  async convert(request: ConversionRequest): Promise<ConversionResponse> {
    const response: AxiosResponse<ConversionResponse> = await this.api.post('/api/v1/convert/', request)
    return response.data
  }

  // Validation
  async validate(request: ValidationRequest): Promise<ValidationResponse> {
    const response: AxiosResponse<ValidationResponse> = await this.api.post('/api/v1/convert/validate', request)
    return response.data
  }

  // Upload et conversion
  async uploadAndConvert(file: File, type: string, options?: any): Promise<ConversionResponse> {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('type', type)
    
    if (options) {
      Object.keys(options).forEach(key => {
        formData.append(key, options[key])
      })
    }

    const response: AxiosResponse<ConversionResponse> = await this.api.post('/api/v1/upload/', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    return response.data
  }

  // Liste des convertisseurs
  async getConverters(): Promise<ConvertersResponse> {
    const response: AxiosResponse<ConvertersResponse> = await this.api.get('/api/v1/info/converters')
    return response.data
  }

  // Health check
  async getHealth(): Promise<HealthResponse> {
    const response: AxiosResponse<HealthResponse> = await this.api.get('/health/')
    return response.data
  }

  // Health check détaillé
  async getHealthDetailed(): Promise<HealthDetailedResponse> {
    const response: AxiosResponse<HealthDetailedResponse> = await this.api.get('/health/detailed')
    return response.data
  }

  // Version
  async getVersion(): Promise<VersionResponse> {
    const response: AxiosResponse<VersionResponse> = await this.api.get('/version')
    return response.data
  }
}

export const apiService = new ApiService()
export default apiService
