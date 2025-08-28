<template>
  <!-- Conteneur de notifications avec meilleur positioning -->
  <Teleport to="body">
    <div
      v-if="notifications.length > 0"
      class="fixed top-20 right-4 z-50 space-y-3 max-w-sm w-full"
    >
      <TransitionGroup
        name="notification"
        tag="div"
        class="space-y-3"
      >
        <div
          v-for="notification in notifications"
          :key="notification.id"
          :class="[
            'bg-white/95 backdrop-blur-sm shadow-2xl rounded-xl border border-gray-200/50',
            'transform transition-all duration-500 ease-out',
            'hover:shadow-3xl hover:scale-105'
          ]"
        >
          <!-- Barre de couleur en haut -->
          <div :class="[
            'h-1 rounded-t-xl',
            notification.type === 'success' ? 'bg-gradient-to-r from-green-400 to-emerald-500' :
            notification.type === 'error' ? 'bg-gradient-to-r from-red-400 to-rose-500' :
            notification.type === 'warning' ? 'bg-gradient-to-r from-yellow-400 to-amber-500' :
            'bg-gradient-to-r from-blue-400 to-indigo-500'
          ]"></div>
          
          <div class="p-4">
            <div class="flex items-start">
              <div class="flex-shrink-0">
                <!-- Success Icon avec animation -->
                <div v-if="notification.type === 'success'" 
                     class="flex items-center justify-center w-10 h-10 bg-green-100 rounded-full animate-bounce">
                  <svg class="h-6 w-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </div>
                
                <!-- Error Icon avec animation -->
                <div v-else-if="notification.type === 'error'" 
                     class="flex items-center justify-center w-10 h-10 bg-red-100 rounded-full animate-pulse">
                  <svg class="h-6 w-6 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </div>
                
                <!-- Warning Icon avec animation -->
                <div v-else-if="notification.type === 'warning'" 
                     class="flex items-center justify-center w-10 h-10 bg-yellow-100 rounded-full animate-pulse">
                  <svg class="h-6 w-6 text-yellow-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4.5c-.77-.833-2.694-.833-3.464 0L3.34 16.5c-.77.833.192 2.5 1.732 2.5z" />
                  </svg>
                </div>
                
                <!-- Info Icon avec animation -->
                <div v-else 
                     class="flex items-center justify-center w-10 h-10 bg-blue-100 rounded-full animate-pulse">
                  <svg class="h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
              </div>
              
              <div class="ml-4 flex-1">
                <h4 class="text-sm font-bold text-gray-900 mb-1">
                  {{ notification.title }}
                </h4>
                <p class="text-sm text-gray-600 leading-relaxed">
                  {{ notification.message }}
                </p>
                
                <!-- Timestamp -->
                <p class="text-xs text-gray-400 mt-2">
                  {{ formatTime(notification.timestamp) }}
                </p>
              </div>
              
              <div class="ml-4 flex-shrink-0">
                <button
                  @click="$emit('remove', notification.id)"
                  class="group inline-flex items-center justify-center w-8 h-8 rounded-full text-gray-400 hover:text-gray-600 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-colors duration-200"
                  title="Fermer"
                >
                  <svg class="h-4 w-4 group-hover:scale-110 transition-transform" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
          
          <!-- Progress bar pour auto-dismiss -->
          <div v-if="notification.autoDismiss" 
               class="h-1 bg-gray-200 rounded-b-xl overflow-hidden">
            <div class="h-full bg-gradient-to-r from-gray-400 to-gray-600 rounded-b-xl animate-shrink"></div>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
interface Notification {
  id: string
  type: 'success' | 'error' | 'warning' | 'info'
  title: string
  message: string
  timestamp: Date
  autoDismiss?: boolean
}

defineProps<{
  notifications: Notification[]
}>()

defineEmits<{
  remove: [id: string]
}>()

// Format time function
function formatTime(timestamp: Date): string {
  const now = new Date()
  const diff = now.getTime() - timestamp.getTime()
  const seconds = Math.floor(diff / 1000)
  
  if (seconds < 60) {
    return 'À l\'instant'
  } else if (seconds < 3600) {
    const minutes = Math.floor(seconds / 60)
    return `Il y a ${minutes} min`
  } else {
    const hours = Math.floor(seconds / 3600)
    return `Il y a ${hours}h`
  }
}
</script>

<style scoped>
/* Animations pour les notifications */
.notification-enter-active,
.notification-leave-active {
  transition: all 0.5s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(100%) scale(0.95);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(100%) scale(0.95);
}

.notification-move {
  transition: transform 0.5s ease;
}

/* Animation pour la barre de progression */
@keyframes shrink {
  from {
    width: 100%;
  }
  to {
    width: 0%;
  }
}

.animate-shrink {
  animation: shrink 5s linear forwards;
}

/* Ombres personnalisées */
.shadow-3xl {
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
}
</style>
