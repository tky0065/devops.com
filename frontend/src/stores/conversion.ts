import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import apiService from '@/services/api'
import type {
  ConversionRequest,
  ConversionResponse,
  ValidationRequest,
  ValidationResponse,
  ConverterInfo
} from '@/types/api'

export const useConversionStore = defineStore('conversion', () => {
  // État
  const isLoading = ref(false)
  const currentConversion = ref<ConversionResponse | null>(null)
  const validationResult = ref<ValidationResponse | null>(null)
  const availableConverters = ref<ConverterInfo[]>([])
  const error = ref<string | null>(null)

  // Actions
  async function convert(request: ConversionRequest) {
    isLoading.value = true
    error.value = null

    try {
      const response = await apiService.convert(request)
      currentConversion.value = response
      return response
    } catch (err: any) {
      error.value = err.message || 'Erreur lors de la conversion'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function validate(request: ValidationRequest) {
    isLoading.value = true
    error.value = null

    try {
      const response = await apiService.validate(request)
      validationResult.value = response
      return response
    } catch (err: any) {
      error.value = err.message || 'Erreur lors de la validation'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function uploadAndConvert(file: File, type: string, options?: any) {
    isLoading.value = true
    error.value = null

    try {
      const response = await apiService.uploadAndConvert(file, type, options)
      currentConversion.value = response
      return response
    } catch (err: any) {
      error.value = err.message || 'Erreur lors du téléchargement et conversion'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function loadConverters() {
    try {
      const response = await apiService.getConverters()
      availableConverters.value = response.converters
      return response
    } catch (err: any) {
      error.value = err.message || 'Erreur lors du chargement des convertisseurs'
      throw err
    }
  }

  function clearConversion() {
    currentConversion.value = null
    validationResult.value = null
    error.value = null
  }

  function clearError() {
    error.value = null
  }

  // Getters calculés
  const hasResults = computed(() => currentConversion.value !== null)
  const isSuccess = computed(() => currentConversion.value?.success === true)
  const hasErrors = computed(() => currentConversion.value?.errors && currentConversion.value.errors.length > 0)
  const hasWarnings = computed(() => currentConversion.value?.warnings && currentConversion.value.warnings.length > 0)
  const generatedFiles = computed(() => currentConversion.value?.files || [])

  return {
    // État
    isLoading,
    currentConversion,
    validationResult,
    availableConverters,
    error,

    // Actions
    convert,
    validate,
    uploadAndConvert,
    loadConverters,
    clearConversion,
    clearError,

    // Getters
    hasResults,
    isSuccess,
    hasErrors,
    hasWarnings,
    generatedFiles
  }
})
