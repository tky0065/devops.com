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

        return Promise.reject(error)
      }
    )
  }

  // Conversion
  async convert(request: ConversionRequest): Promise<ConversionResponse> {
    // Restructurer la requête pour correspondre à l'API backend
    const apiRequest = {
      type: request.type,
      content: request.content,
      options: request.options || {}
    }
    try {
      const response: AxiosResponse<ConversionResponse> = await this.api.post('/api/v1/convert/', apiRequest)
      return response.data
    } catch (err: unknown) {
      // If server returned JSON (even for 400), surface it so UI can display structured errors
      const e = err as { response?: { data?: unknown } }
      if (e && e.response && e.response.data) {
        return e.response.data as ConversionResponse
      }
      // rethrow otherwise
      throw err
    }
  }

  // Validation
  async validate(request: ValidationRequest): Promise<ValidationResponse> {
    // Restructurer la requête pour correspondre à l'API backend
    const apiRequest = {
      type: request.type,
      content: request.content  // Le backend attend "content" aussi pour la validation
    }
    try {
      const response: AxiosResponse<ValidationResponse> = await this.api.post('/api/v1/convert/validate', apiRequest)
      return response.data
    } catch (err: unknown) {
      const e = err as { response?: { data?: unknown } }
      if (e && e.response && e.response.data) {
        return e.response.data as ValidationResponse
      }
      throw err
    }
  }

  // Upload et conversion
  async uploadAndConvert(file: File, type: string, options?: Record<string, unknown>): Promise<ConversionResponse> {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('type', type)

    if (options) {
      Object.keys(options).forEach(key => {
        const v = options[key]
        formData.append(key, v === undefined || v === null ? '' : String(v))
      })
    }

    try {
      const response: AxiosResponse<ConversionResponse> = await this.api.post('/api/v1/upload/', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
      return response.data
    } catch (err: unknown) {
      const e = err as { response?: { data?: unknown } }
      if (e && e.response && e.response.data) {
        return e.response.data as ConversionResponse
      }
      throw err
    }
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
