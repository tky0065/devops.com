<template>
  <div id="app" class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-100">
    <!-- Header professionnel -->
    <header class="bg-white/80 backdrop-blur-md shadow-lg border-b border-slate-200/60 sticky top-0 z-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-4">
          <!-- Logo et branding -->
          <div class="flex items-center space-x-4">
            <div class="relative">
              <div class="flex items-center justify-center w-12 h-12 bg-gradient-to-br from-blue-600 via-blue-700 to-indigo-700 rounded-xl shadow-lg transform rotate-3 hover:rotate-0 transition-transform duration-300">
                <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
              <div class="absolute -top-1 -right-1 w-4 h-4 bg-gradient-to-br from-green-400 to-green-500 rounded-full border-2 border-white animate-pulse"></div>
            </div>
            <div class="flex flex-col">
              <h1 class="text-2xl font-bold bg-gradient-to-r from-slate-900 via-blue-900 to-indigo-900 bg-clip-text text-transparent">
                DevOps Converter
              </h1>
              <p class="text-sm text-slate-600 font-medium">Docker Compose → Kubernetes</p>
            </div>
            <div v-if="appStore.appInfo" class="hidden sm:flex items-center space-x-2">
              <div class="px-3 py-1.5 text-xs font-semibold bg-gradient-to-r from-blue-500 to-indigo-500 text-white rounded-full shadow-sm">
                v{{ appStore.appInfo.version }}
              </div>
              <div class="px-3 py-1.5 text-xs font-semibold bg-gradient-to-r from-emerald-500 to-teal-500 text-white rounded-full shadow-sm">
                Production Ready
              </div>
            </div>
          </div>
          
          <!-- Navigation et status -->
          <div class="flex items-center space-x-6">
            <!-- Indicateur de performance -->
            <div class="hidden md:flex items-center space-x-4 text-sm">
              <div class="flex items-center space-x-2">
                <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                <span class="text-slate-600 font-medium">99.9% Uptime</span>
              </div>
              <div class="w-px h-4 bg-slate-300"></div>
              <div class="flex items-center space-x-2">
                <svg class="w-4 h-4 text-slate-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
                <span class="text-slate-600 font-medium">&lt;200ms</span>
              </div>
            </div>

            <!-- Status de santé amélioré -->
            <div class="status-badge" 
                 data-testid="health-status"
                 :class="appStore.isHealthy ? 'online' : 'offline'">
              <div class="status-dot" :class="appStore.isHealthy ? 'online' : 'offline'"></div>
              <span class="font-semibold">
                {{ appStore.isHealthy ? 'En ligne' : 'Hors ligne' }}
              </span>
            </div>
            
            <!-- Actions -->
            <div class="flex items-center space-x-3">
              <!-- Documentation -->
              <button class="p-2 text-slate-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-all duration-200 group" title="Documentation">
                <svg class="w-5 h-5 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                </svg>
              </button>

              <!-- Support -->
              <button class="p-2 text-slate-600 hover:text-green-600 hover:bg-green-50 rounded-lg transition-all duration-200 group" title="Support">
                <svg class="w-5 h-5 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
              </button>
            
              <!-- GitHub avec design amélioré -->
              <a
                href="https://github.com/tky0065/devops-converter"
                target="_blank"
                rel="noopener noreferrer"
                class="flex items-center space-x-2 px-4 py-2 text-slate-700 hover:text-slate-900 bg-white hover:bg-slate-50 border border-slate-200 hover:border-slate-300 rounded-lg transition-all duration-200 shadow-sm hover:shadow-md group"
                title="Voir sur GitHub"
              >
                <svg class="h-5 w-5 group-hover:scale-110 transition-transform" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z" clip-rule="evenodd" />
                </svg>
                <span class="text-sm font-semibold hidden sm:inline">GitHub</span>
              </a>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- Contenu principal avec espacement professionnel -->
    <main class="relative">
      <router-view />
    </main>

    <!-- Notifications avec positionnement amélioré -->
    <NotificationContainer
      :notifications="appStore.notifications"
      @remove="appStore.removeNotification"
    />

    <!-- Footer invisible pour l'espace -->
    <div class="h-16"></div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useAppStore } from '@/stores/app'
import NotificationContainer from '@/components/NotificationContainer.vue'

const appStore = useAppStore()

onMounted(async () => {
  // Charger les informations de l'application au démarrage
  try {
    await appStore.loadAppInfo()
  } catch (error) {
    console.error('Failed to load app info:', error)
    appStore.addNotification({
      type: 'warning',
      title: 'Connexion',
      message: 'Impossible de se connecter au serveur backend'
    })
  }
})
</script>

<style scoped>
#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
}
</style>
