import { defineStore } from 'pinia'
import { ref } from 'vue'
import apiService from '@/services/api'
import type { HealthDetailedResponse, VersionResponse } from '@/types/api'

export const useAppStore = defineStore('app', () => {
  // État
  const appInfo = ref<VersionResponse | null>(null)
  const healthInfo = ref<HealthDetailedResponse | null>(null)
  const isHealthy = ref(true)
  const notifications = ref<Array<{
    id: string
    type: 'success' | 'error' | 'warning' | 'info'
    title: string
    message: string
    timestamp: Date
  }>>([])

  // Actions
  async function loadAppInfo() {
    try {
      const [version, health] = await Promise.all([
        apiService.getVersion(),
        apiService.getHealthDetailed()
      ])
      
      appInfo.value = version
      healthInfo.value = health
      isHealthy.value = health.status === 'ok'
    } catch (error) {
      isHealthy.value = false
      console.error('Failed to load app info:', error)
    }
  }

  async function checkHealth() {
    try {
      const health = await apiService.getHealthDetailed()
      healthInfo.value = health
      isHealthy.value = health.status === 'ok'
      return health
    } catch (error) {
      isHealthy.value = false
      throw error
    }
  }

  function addNotification(notification: Omit<typeof notifications.value[0], 'id' | 'timestamp'>) {
    const id = Date.now().toString()
    notifications.value.push({
      ...notification,
      id,
      timestamp: new Date()
    })
    
    // Auto-remove après 5 secondes pour les succès et infos
    if (notification.type === 'success' || notification.type === 'info') {
      setTimeout(() => {
        removeNotification(id)
      }, 5000)
    }
  }

  function removeNotification(id: string) {
    const index = notifications.value.findIndex(n => n.id === id)
    if (index > -1) {
      notifications.value.splice(index, 1)
    }
  }

  function clearNotifications() {
    notifications.value = []
  }

  return {
    // État
    appInfo,
    healthInfo,
    isHealthy,
    notifications,
    
    // Actions
    loadAppInfo,
    checkHealth,
    addNotification,
    removeNotification,
    clearNotifications
  }
})
