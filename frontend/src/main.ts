import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import './assets/main.css'

// Import des composants pour les routes
import ConversionPage from './components/ConversionPage.vue'

// Configuration du router
const routes = [
  { path: '/', name: 'home', component: ConversionPage },
  { path: '/convert', name: 'convert', component: ConversionPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Configuration de l'app
const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.mount('#app')
