<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50">
    <!-- Hero Section -->
    <div class="max-w-7xl mx-auto px-4 pt-12 pb-8">
      <div class="text-center animate-fade-in">
        <!-- Logo interactif -->
        <div class="relative inline-block mb-8">
          <div class="flex items-center justify-center w-20 h-20 bg-gradient-to-br from-blue-600 via-indigo-600 to-purple-700 rounded-2xl shadow-2xl transform hover:scale-105 transition-all duration-300 animate-bounce-in">
            <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
          </div>
          <div class="absolute -top-2 -right-2 w-6 h-6 bg-gradient-to-r from-green-400 to-emerald-500 rounded-full border-3 border-white animate-pulse"></div>
        </div>

        <!-- Titre et description -->
        <h1 class="text-5xl lg:text-6xl font-bold bg-gradient-to-r from-slate-900 via-blue-900 to-indigo-900 bg-clip-text text-transparent mb-6 leading-tight">
          DevOps Transformation
          <span class="block text-3xl lg:text-4xl mt-2 bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">
            Docker → Kubernetes
          </span>
        </h1>
        
        <p class="text-xl lg:text-2xl text-slate-600 max-w-4xl mx-auto leading-relaxed mb-8">
          Transformez vos configurations Docker Compose en manifestes Kubernetes 
          <span class="font-semibold bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">
            professionnels et optimisés
          </span>
        </p>
        
        <!-- KPIs professionnels -->
        <div class="flex flex-wrap justify-center gap-8 mt-12">
          <div class="group text-center transform hover:scale-105 transition-all duration-300">
            <div class="w-16 h-16 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl shadow-lg mx-auto mb-3 flex items-center justify-center group-hover:shadow-xl transition-shadow">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"/>
              </svg>
            </div>
            <div class="text-3xl font-bold text-blue-600">10,000+</div>
            <div class="text-sm font-medium text-slate-500 uppercase tracking-wider">Conversions</div>
          </div>
          
          <div class="group text-center transform hover:scale-105 transition-all duration-300">
            <div class="w-16 h-16 bg-gradient-to-br from-emerald-500 to-emerald-600 rounded-xl shadow-lg mx-auto mb-3 flex items-center justify-center group-hover:shadow-xl transition-shadow">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <div class="text-3xl font-bold text-emerald-600">99.99%</div>
            <div class="text-sm font-medium text-slate-500 uppercase tracking-wider">Fiabilité</div>
          </div>
          
          <div class="group text-center transform hover:scale-105 transition-all duration-300">
            <div class="w-16 h-16 bg-gradient-to-br from-purple-500 to-purple-600 rounded-xl shadow-lg mx-auto mb-3 flex items-center justify-center group-hover:shadow-xl transition-shadow">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
              </svg>
            </div>
            <div class="text-3xl font-bold text-purple-600">&lt;150ms</div>
            <div class="text-sm font-medium text-slate-500 uppercase tracking-wider">Latence</div>
          </div>
          
          <div class="group text-center transform hover:scale-105 transition-all duration-300">
            <div class="w-16 h-16 bg-gradient-to-br from-amber-500 to-amber-600 rounded-xl shadow-lg mx-auto mb-3 flex items-center justify-center group-hover:shadow-xl transition-shadow">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
              </svg>
            </div>
            <div class="text-3xl font-bold text-amber-600">SOC2</div>
            <div class="text-sm font-medium text-slate-500 uppercase tracking-wider">Sécurité</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Zone de conversion professionnelle -->
    <div class="max-w-7xl mx-auto px-4 pb-12">
      <div class="grid grid-cols-1 xl:grid-cols-12 gap-8">
        
        <!-- Zone d'entrée (7 colonnes) -->
        <div class="xl:col-span-7 space-y-6 animate-slide-in-up">
          <div class="card card-body">
            <div class="card-header pb-6">
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 bg-gradient-to-br from-blue-100 to-indigo-100 rounded-xl flex items-center justify-center">
                    <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                    </svg>
                  </div>
                  <div>
                    <h2 class="text-2xl font-bold text-slate-900">Configuration d'entrée</h2>
                    <p class="text-slate-600">Importez votre configuration Docker Compose</p>
                  </div>
                </div>
                <div class="status-badge online">
                  <div class="status-dot online"></div>
                  <span class="font-semibold">Prêt</span>
                </div>
              </div>
            </div>
            
            <!-- Sélecteur de type avec design professionnel -->
            <div class="form-group">
              <label class="form-label">
                <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"/>
                </svg>
                Format de fichier
              </label>
              <select v-model="selectedType" class="form-input form-select">
                <option value="docker-compose">🐳 Docker Compose (YAML)</option>
                <option value="dockerfile" disabled>🏗️ Dockerfile (Bientôt disponible)</option>
                <option value="helm" disabled>⚙️ Helm Chart (Bientôt disponible)</option>
              </select>
            </div>

            <!-- Navigation onglets redesignée -->
            <div class="nav-tabs">
              <button
                v-for="tab in inputTabs"
                :key="tab.id"
                @click="activeInputTab = tab.id"
                :class="['nav-tab', { 'active': activeInputTab === tab.id }]"
              >
                <svg v-if="tab.id === 'text'" class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/>
                </svg>
                <svg v-if="tab.id === 'upload'" class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"/>
                </svg>
                {{ tab.label }}
              </button>
            </div>

            <!-- Zone de texte améliorée -->
            <div v-if="activeInputTab === 'text'" class="form-group">
              <label class="form-label">
                <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/>
                </svg>
                Contenu Docker Compose
              </label>
              <textarea
                v-model="inputContent"
                class="form-input form-textarea h-80"
                placeholder="version: '3.8'
services:
  web:
    image: nginx:latest
    ports:
      - '80:80'
    environment:
      - ENV=production
  
  db:
    image: postgres:13
    environment:
      - POSTGRES_DB=myapp"
              ></textarea>
              <div class="flex items-center justify-between mt-2 text-sm text-slate-500">
                <span>{{ inputContent.length }} caractères</span>
                <div class="flex items-center space-x-4">
                  <button @click="inputContent = ''" class="text-blue-600 hover:text-blue-700 font-medium">
                    Effacer
                  </button>
                  <button @click="loadExample" class="text-blue-600 hover:text-blue-700 font-medium">
                    Charger un exemple
                  </button>
                </div>
              </div>
            </div>

            <!-- Zone d'upload redesignée -->
            <div v-if="activeInputTab === 'upload'" class="form-group">
              <label class="form-label">
                <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"/>
                </svg>
                Fichier Docker Compose
              </label>
              <div
                @drop="handleDrop"
                @dragover.prevent
                @dragenter.prevent
                class="drop-zone"
                :class="{ 'drag-over': isDragOver }"
              >
                <input
                  ref="fileInput"
                  type="file"
                  @change="handleFileSelect"
                  accept=".yml,.yaml"
                  class="hidden"
                >
                
                <div v-if="!selectedFile" class="text-center">
                  <div class="w-16 h-16 bg-gradient-to-br from-blue-100 to-indigo-100 rounded-xl mx-auto mb-4 flex items-center justify-center">
                    <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"/>
                    </svg>
                  </div>
                  <h3 class="text-lg font-semibold text-slate-900 mb-2">Importez votre fichier</h3>
                  <p class="text-slate-600 mb-4">
                    Glissez-déposez votre fichier YAML ou 
                    <button @click="fileInput?.click()" class="font-semibold text-blue-600 hover:text-blue-700">
                      parcourez vos fichiers
                    </button>
                  </p>
                  <div class="flex items-center justify-center space-x-4 text-sm text-slate-500">
                    <span class="flex items-center">
                      <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                      </svg>
                      YAML uniquement
                    </span>
                    <span class="flex items-center">
                      <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                      </svg>
                      Max 10MB
                    </span>
                  </div>
                </div>
                
                <div v-else class="flex items-center justify-between bg-white p-4 rounded-lg border">
                  <div class="flex items-center space-x-3">
                    <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                      <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                      </svg>
                    </div>
                    <div>
                      <p class="font-semibold text-slate-900">{{ selectedFile.name }}</p>
                      <p class="text-sm text-slate-500">{{ formatFileSize(selectedFile.size) }}</p>
                    </div>
                  </div>
                  <button
                    @click="clearFile"
                    class="text-red-600 hover:text-red-700 p-2 rounded-lg hover:bg-red-50 transition-colors"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                    </svg>
                  </button>
                </div>
              </div>
            </div>

            <!-- Options de conversion avancées -->
            <div class="form-group">
              <label class="form-label">
                <svg class="w-4 h-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                </svg>
                Configuration Kubernetes
              </label>
              <div class="grid grid-cols-1 lg:grid-cols-3 gap-4">
                <div>
                  <label class="block text-sm font-medium text-slate-600 mb-2">Namespace</label>
                  <input
                    v-model="conversionOptions.namespace"
                    type="text"
                    class="form-input"
                    placeholder="default"
                  >
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-slate-600 mb-2">Type de service</label>
                  <select v-model="conversionOptions.serviceType" class="form-input form-select">
                    <option value="ClusterIP">🔒 ClusterIP (Interne)</option>
                    <option value="NodePort">🌐 NodePort (Externe)</option>
                    <option value="LoadBalancer">⚖️ LoadBalancer (Cloud)</option>
                  </select>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-slate-600 mb-2">Replicas</label>
                  <input
                    v-model.number="conversionOptions.replicas"
                    type="number"
                    min="1"
                    max="20"
                    class="form-input"
                    placeholder="1"
                  >
                </div>
              </div>
            </div>

            <!-- Boutons d'action professionnels -->
            <div class="flex flex-col sm:flex-row gap-3 pt-6 border-t">
              <button
                @click="validateInput"
                :disabled="!hasValidInput || conversionStore.isLoading"
                class="btn btn-secondary sm:w-auto"
              >
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                <span v-if="conversionStore.isLoading">Validation en cours...</span>
                <span v-else>Valider la syntaxe</span>
              </button>
              
              <button
                @click="convertInput"
                :disabled="!hasValidInput || conversionStore.isLoading"
                class="btn btn-primary flex-1"
              >
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>
                <span v-if="conversionStore.isLoading">Conversion en cours...</span>
                <span v-else>Convertir vers Kubernetes</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Zone de résultats (5 colonnes) -->
        <div class="xl:col-span-5 space-y-6 animate-slide-in-right">
          <!-- Résultats de validation -->
          <div v-if="conversionStore.validationResult" class="card card-body animate-bounce-in">
            <div class="flex items-center justify-between mb-4">
              <h3 class="text-xl font-bold text-slate-900">Résultat de validation</h3>
              <div :class="[
                'status-badge',
                conversionStore.validationResult.valid ? 'online' : 'offline'
              ]">
                <div :class="[
                  'status-dot',
                  conversionStore.validationResult.valid ? 'online' : 'offline'
                ]"></div>
                <span class="font-semibold">
                  {{ conversionStore.validationResult.valid ? 'Valide' : 'Invalide' }}
                </span>
              </div>
            </div>
            
            <div :class="[
              'p-4 rounded-xl border-l-4',
              conversionStore.validationResult.valid
                ? 'bg-emerald-50 border-emerald-500 text-emerald-800'
                : 'bg-red-50 border-red-500 text-red-800'
            ]">
              <div class="flex items-start">
                <svg v-if="conversionStore.validationResult.valid" class="w-5 h-5 text-emerald-500 mr-3 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <svg v-else class="w-5 h-5 text-red-500 mr-3 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                </svg>
                <div class="flex-1">
                  <p class="font-semibold mb-1">{{ conversionStore.validationResult.message }}</p>
                  <ul v-if="conversionStore.validationResult.errors?.length" class="space-y-1 text-sm">
                    <li v-for="error in conversionStore.validationResult.errors" :key="error" class="flex items-start">
                      <span class="w-1 h-1 bg-current rounded-full mt-2 mr-2 flex-shrink-0"></span>
                      {{ error }}
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>

          <!-- Résultats de conversion -->
          <div v-if="conversionStore.hasResults" class="card card-body animate-bounce-in">
            <div class="flex items-center justify-between mb-6">
              <h3 class="text-xl font-bold text-slate-900">Manifestes générés</h3>
              <div :class="[
                'status-badge',
                conversionStore.isSuccess ? 'online' : 'offline'
              ]">
                <div :class="[
                  'status-dot',
                  conversionStore.isSuccess ? 'online' : 'offline'
                ]"></div>
                <span class="font-semibold">
                  {{ conversionStore.isSuccess ? 'Succès' : 'Échec' }}
                </span>
              </div>
            </div>
            
            <!-- Statut de conversion -->
            <div :class="[
              'p-4 rounded-xl border-l-4 mb-6',
              conversionStore.isSuccess
                ? 'bg-emerald-50 border-emerald-500 text-emerald-800'
                : 'bg-red-50 border-red-500 text-red-800'
            ]">
              <div class="flex items-center">
                <svg v-if="conversionStore.isSuccess" class="w-5 h-5 text-emerald-500 mr-3" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <svg v-else class="w-5 h-5 text-red-500 mr-3" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                </svg>
                <span class="font-semibold">
                  {{ conversionStore.isSuccess 
                      ? `${conversionStore.generatedFiles.length} fichier(s) généré(s) avec succès` 
                      : 'Erreur lors de la conversion' 
                  }}
                </span>
              </div>
            </div>

            <!-- Actions globales -->
            <div v-if="conversionStore.generatedFiles.length > 0" class="flex items-center justify-between mb-6">
              <div class="text-sm text-slate-600">
                <span class="font-medium">{{ conversionStore.generatedFiles.length }}</span> fichiers prêts
              </div>
              <button
                @click="downloadAllFiles"
                class="btn btn-success btn-sm"
              >
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-4-4m4 4l4-4m-6 0h.01M19 4v8a2 2 0 01-2 2H7a2 2 0 01-2-2V4a2 2 0 012-2h10a2 2 0 012 2z"/>
                </svg>
                Télécharger tout
              </button>
            </div>
            
            <!-- Liste des fichiers -->
            <div v-if="conversionStore.generatedFiles.length > 0" class="space-y-4">
              <div
                v-for="(file, index) in conversionStore.generatedFiles"
                :key="file.name"
                class="border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition-colors"
              >
                <div class="bg-slate-50 px-4 py-3 border-b border-slate-200">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                      <div :class="[
                        'w-8 h-8 rounded-lg flex items-center justify-center text-white text-sm font-bold',
                        getFileTypeColor(file.type)
                      ]">
                        {{ index + 1 }}
                      </div>
                      <div>
                        <p class="font-semibold text-slate-900">{{ file.name }}</p>
                        <p class="text-sm text-slate-500 capitalize">{{ file.type }}</p>
                      </div>
                    </div>
                    <div class="flex items-center space-x-2">
                      <button
                        @click="copyToClipboard(file.content)"
                        class="p-2 text-slate-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
                        title="Copier"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                        </svg>
                      </button>
                      <button
                        @click="downloadFile(file)"
                        class="p-2 text-slate-600 hover:text-green-600 hover:bg-green-50 rounded-lg transition-colors"
                        title="Télécharger"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-4-4m4 4l4-4m-6 0h.01M19 4v8a2 2 0 01-2 2H7a2 2 0 01-2-2V4a2 2 0 012-2h10a2 2 0 012 2z"/>
                        </svg>
                      </button>
                    </div>
                  </div>
                </div>
                
                <div class="p-4">
                  <pre class="code-block text-xs overflow-x-auto"><code>{{ file.content }}</code></pre>
                </div>
              </div>
            </div>

            <!-- Métadonnées -->
            <div v-if="conversionStore.currentConversion?.metadata" class="mt-6 pt-6 border-t border-slate-200">
              <h4 class="font-semibold text-slate-900 mb-3">Statistiques de conversion</h4>
              <div class="grid grid-cols-2 gap-4 text-sm">
                <div class="flex items-center space-x-2">
                  <div class="w-2 h-2 bg-blue-500 rounded-full"></div>
                  <span class="text-slate-600">Services:</span>
                  <span class="font-semibold">{{ conversionStore.currentConversion.metadata.services_converted }}</span>
                </div>
                <div class="flex items-center space-x-2">
                  <div class="w-2 h-2 bg-green-500 rounded-full"></div>
                  <span class="text-slate-600">Volumes:</span>
                  <span class="font-semibold">{{ conversionStore.currentConversion.metadata.volumes_converted }}</span>
                </div>
                <div v-if="conversionStore.currentConversion.metadata.docker_version" class="flex items-center space-x-2 col-span-2">
                  <div class="w-2 h-2 bg-purple-500 rounded-full"></div>
                  <span class="text-slate-600">Version Docker:</span>
                  <span class="font-semibold">{{ conversionStore.currentConversion.metadata.docker_version }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Zone vide avec design amélioré -->
          <div v-if="!conversionStore.hasResults && !conversionStore.validationResult" class="card card-body text-center">
            <div class="py-12">
              <div class="w-16 h-16 bg-gradient-to-br from-slate-100 to-slate-200 rounded-xl mx-auto mb-4 flex items-center justify-center">
                <svg class="w-8 h-8 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"/>
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-slate-900 mb-2">Prêt pour la conversion</h3>
              <p class="text-slate-600 mb-6">Les résultats de validation et de conversion apparaîtront ici</p>
              <div class="text-xs text-slate-500 space-y-1">
                <p>✨ Conversion instantanée</p>
                <p>🔒 Données sécurisées</p>
                <p>📋 Manifestes optimisés</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, nextTick } from 'vue'
import { useConversionStore } from '@/stores/conversion'
import { useAppStore } from '@/stores/app'
import type { ConversionOptions, GeneratedFile } from '@/types/api'

// Stores
const conversionStore = useConversionStore()
const appStore = useAppStore()

// Refs
const inputContent = ref('')
const selectedFile = ref<File | null>(null)
const selectedType = ref('docker-compose')
const activeInputTab = ref('text')
const isDragOver = ref(false)
const fileInput = ref<HTMLInputElement>()

// Configuration de conversion
const conversionOptions = ref<ConversionOptions>({
  namespace: 'default',
  serviceType: 'ClusterIP',
  replicas: 1
})

// Configuration des onglets
const inputTabs = ref([
  { id: 'text', label: 'Saisie manuelle' },
  { id: 'upload', label: 'Import de fichier' }
])

// Computed
const hasValidInput = computed(() => {
  if (activeInputTab.value === 'text') {
    return inputContent.value.trim().length > 0
  }
  return selectedFile.value !== null
})

// Méthodes de fichier
const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    const file = target.files[0]
    if (file.type === 'application/x-yaml' || file.name.endsWith('.yml') || file.name.endsWith('.yaml')) {
      selectedFile.value = file
      readFileContent(file)
    } else {
      appStore.addNotification({
        type: 'error',
        title: 'Erreur de fichier',
        message: 'Veuillez sélectionner un fichier YAML valide (.yml ou .yaml)'
      })
    }
  }
}

const handleDrop = (event: DragEvent) => {
  event.preventDefault()
  isDragOver.value = false
  
  if (event.dataTransfer?.files && event.dataTransfer.files.length > 0) {
    const file = event.dataTransfer.files[0]
    if (file.type === 'application/x-yaml' || file.name.endsWith('.yml') || file.name.endsWith('.yaml')) {
      selectedFile.value = file
      readFileContent(file)
    } else {
      appStore.addNotification({
        type: 'error',
        title: 'Erreur de fichier',
        message: 'Veuillez déposer un fichier YAML valide (.yml ou .yaml)'
      })
    }
  }
}

const readFileContent = async (file: File) => {
  try {
    const text = await file.text()
    inputContent.value = text
    
    appStore.addNotification({
      type: 'success',
      title: 'Fichier chargé',
      message: `Fichier "${file.name}" chargé avec succès`
    })
  } catch (error) {
    console.error('Erreur lors de la lecture du fichier:', error)
    appStore.addNotification({
      type: 'error',
      title: 'Erreur de lecture',
      message: 'Erreur lors de la lecture du fichier'
    })
  }
}

const clearFile = () => {
  selectedFile.value = null
  if (fileInput.value) {
    fileInput.value.value = ''
  }
  inputContent.value = ''
}

const formatFileSize = (bytes: number): string => {
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  if (bytes === 0) return '0 Bytes'
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return Math.round(bytes / Math.pow(1024, i) * 100) / 100 + ' ' + sizes[i]
}

// Méthodes de conversion
const validateInput = async () => {
  if (!hasValidInput.value) {
    appStore.addNotification({
      type: 'warning',
      title: 'Configuration manquante',
      message: 'Veuillez saisir ou importer une configuration Docker Compose'
    })
    return
  }

  try {
    await conversionStore.validate({
      content: inputContent.value,
      type: selectedType.value
    })
    
    if (conversionStore.validationResult?.valid) {
      appStore.addNotification({
        type: 'success',
        title: 'Validation réussie',
        message: 'Configuration validée avec succès !'
      })
    } else {
      appStore.addNotification({
        type: 'error',
        title: 'Validation échouée',
        message: 'Configuration invalide. Veuillez corriger les erreurs.'
      })
    }
  } catch (error) {
    console.error('Erreur de validation:', error)
    appStore.addNotification({
      type: 'error',
      title: 'Erreur de validation',
      message: 'Erreur lors de la validation'
    })
  }
}

const convertInput = async () => {
  if (!hasValidInput.value) {
    appStore.addNotification({
      type: 'warning',
      title: 'Configuration manquante',
      message: 'Veuillez saisir ou importer une configuration Docker Compose'
    })
    return
  }

  try {
    await conversionStore.convert({
      content: inputContent.value,
      type: selectedType.value,
      options: conversionOptions.value
    })
    
    if (conversionStore.isSuccess) {
      appStore.addNotification({
        type: 'success',
        title: 'Conversion réussie',
        message: `Conversion réussie ! ${conversionStore.generatedFiles.length} fichier(s) généré(s)`
      })
    } else {
      appStore.addNotification({
        type: 'error',
        title: 'Conversion échouée',
        message: 'Erreur lors de la conversion'
      })
    }
  } catch (error) {
    console.error('Erreur de conversion:', error)
    appStore.addNotification({
      type: 'error',
      title: 'Erreur de conversion',
      message: 'Erreur lors de la conversion'
    })
  }
}

// Méthodes utilitaires
const loadExample = () => {
  inputContent.value = `version: '3.8'

services:
  # Service web frontend
  web:
    image: nginx:1.21-alpine
    container_name: web-server
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
      - ./ssl:/etc/nginx/ssl:ro
    environment:
      - NGINX_HOST=localhost
      - NGINX_PORT=80
    depends_on:
      - api
    restart: unless-stopped
    networks:
      - frontend
      - backend

  # API Backend
  api:
    image: node:16-alpine
    container_name: api-server
    build:
      context: ./api
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    volumes:
      - ./api:/app
      - /app/node_modules
    environment:
      - NODE_ENV=production
      - DATABASE_URL=postgresql://user:password@db:5432/myapp
      - REDIS_URL=redis://redis:6379
      - JWT_SECRET=your-super-secret-jwt-key
    depends_on:
      - db
      - redis
    restart: unless-stopped
    networks:
      - backend

  # Base de données PostgreSQL
  db:
    image: postgres:13-alpine
    container_name: postgres-db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    environment:
      - POSTGRES_DB=myapp
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
    restart: unless-stopped
    networks:
      - backend

  # Cache Redis
  redis:
    image: redis:6-alpine
    container_name: redis-cache
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    command: redis-server --appendonly yes
    restart: unless-stopped
    networks:
      - backend

volumes:
  postgres_data:
    driver: local
  redis_data:
    driver: local

networks:
  frontend:
    driver: bridge
  backend:
    driver: bridge
    internal: true`

  appStore.addNotification({
    type: 'info',
    title: 'Exemple chargé',
    message: 'Exemple de configuration chargé'
  })
}

const getFileTypeColor = (type: string): string => {
  const colors: Record<string, string> = {
    deployment: 'bg-gradient-to-br from-blue-500 to-blue-600',
    service: 'bg-gradient-to-br from-green-500 to-green-600',
    configmap: 'bg-gradient-to-br from-purple-500 to-purple-600',
    secret: 'bg-gradient-to-br from-red-500 to-red-600',
    persistentvolume: 'bg-gradient-to-br from-amber-500 to-amber-600',
    persistentvolumeclaim: 'bg-gradient-to-br from-orange-500 to-orange-600',
    ingress: 'bg-gradient-to-br from-indigo-500 to-indigo-600',
    namespace: 'bg-gradient-to-br from-teal-500 to-teal-600'
  }
  return colors[type.toLowerCase()] || 'bg-gradient-to-br from-slate-500 to-slate-600'
}

// Méthodes de téléchargement
const downloadFile = (file: GeneratedFile) => {
  const blob = new Blob([file.content], { type: 'text/yaml' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = file.name
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
  
  appStore.addNotification({
    type: 'success',
    title: 'Fichier téléchargé',
    message: `Fichier "${file.name}" téléchargé`
  })
}

const downloadAllFiles = () => {
  if (conversionStore.generatedFiles.length === 0) return
  
  // Télécharger les fichiers un par un
  conversionStore.generatedFiles.forEach((file, index) => {
    setTimeout(() => downloadFile(file), index * 200)
  })
  
  appStore.addNotification({
    type: 'success',
    title: 'Téléchargement démarré',
    message: `Téléchargement de ${conversionStore.generatedFiles.length} fichier(s) démarré`
  })
}

const copyToClipboard = async (content: string) => {
  try {
    await navigator.clipboard.writeText(content)
    appStore.addNotification({
      type: 'success',
      title: 'Copié',
      message: 'Contenu copié dans le presse-papiers'
    })
  } catch (error) {
    console.error('Erreur lors de la copie:', error)
    appStore.addNotification({
      type: 'error',
      title: 'Erreur de copie',
      message: 'Erreur lors de la copie'
    })
  }
}

// Lifecycle
onMounted(() => {
  // Animation d'entrée
  nextTick(() => {
    const elements = document.querySelectorAll('.animate-fade-in, .animate-slide-in-up, .animate-slide-in-right, .animate-bounce-in')
    elements.forEach((el, index) => {
      setTimeout(() => {
        el.classList.add('animate-in')
      }, index * 100)
    })
  })
})
</script>

<style scoped>
/* Animations personnalisées */
@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slide-in-up {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slide-in-right {
  from {
    opacity: 0;
    transform: translateX(30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes bounce-in {
  0% {
    opacity: 0;
    transform: scale(0.8) translateY(20px);
  }
  60% {
    opacity: 1;
    transform: scale(1.05) translateY(0);
  }
  100% {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

.animate-fade-in {
  animation: fade-in 0.6s ease-out;
}

.animate-slide-in-up {
  animation: slide-in-up 0.8s ease-out;
}

.animate-slide-in-right {
  animation: slide-in-right 0.8s ease-out;
}

.animate-bounce-in {
  animation: bounce-in 0.8s ease-out;
}

/* États des zones de drop */
.drop-zone {
  border: 2px dashed rgb(203 213 225);
  border-radius: 0.75rem;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
}

.drop-zone:hover {
  border-color: rgb(96 165 250);
  background-color: rgba(59 130 246, 0.05);
}

.drop-zone.drag-over {
  border-color: rgb(59 130 246);
  background-color: rgb(239 246 255);
  transform: scale(1.02);
}

/* Code blocks améliorés */
.code-block {
  background-color: rgb(15 23 42);
  color: rgb(241 245 249);
  padding: 1rem;
  border-radius: 0.5rem;
  border: 1px solid rgb(51 65 85);
  overflow-x: auto;
  scrollbar-width: thin;
  scrollbar-color: rgb(71 85 105) rgb(30 41 59);
}

.code-block code {
  font-family: ui-monospace, SFMono-Regular, 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', monospace;
  font-size: 0.875rem;
  line-height: 1.625;
}

/* Navigation onglets */
.nav-tabs {
  display: flex;
  gap: 0.25rem;
  background-color: rgb(241 245 249);
  padding: 0.25rem;
  border-radius: 0.75rem;
  margin-bottom: 1.5rem;
}

.nav-tab {
  display: flex;
  align-items: center;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  font-weight: 500;
  font-size: 0.875rem;
  transition: all 0.2s ease;
  flex: 1;
  justify-content: center;
}

.nav-tab:not(.active) {
  color: rgb(71 85 105);
}

.nav-tab:not(.active):hover {
  color: rgb(15 23 42);
  background-color: rgba(255, 255, 255, 0.6);
}

.nav-tab.active {
  background-color: white;
  color: rgb(37 99 235);
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  font-weight: 600;
}

/* Responsive improvements */
@media (max-width: 768px) {
  .nav-tabs {
    flex-direction: column;
    gap: 0.25rem;
  }
  
  .nav-tab {
    text-align: left;
    justify-content: flex-start;
  }
}

/* Status indicators avec animations */
.status-dot {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 50%;
}

.status-dot.online {
  animation: pulse-green 2s infinite;
}

.status-dot.offline {
  animation: pulse-red 2s infinite;
}

@keyframes pulse-green {
  0%, 100% {
    background-color: rgb(34 197 94);
    opacity: 1;
  }
  50% {
    background-color: rgb(74 222 128);
    opacity: 0.8;
  }
}

@keyframes pulse-red {
  0%, 100% {
    background-color: rgb(239 68 68);
    opacity: 1;
  }
  50% {
    background-color: rgb(248 113 113);
    opacity: 0.8;
  }
}

/* Amélioration des transitions */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Scrollbar personnalisée pour WebKit */
.code-block::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

.code-block::-webkit-scrollbar-thumb {
  background-color: rgb(71 85 105);
  border-radius: 3px;
}

.code-block::-webkit-scrollbar-track {
  background-color: rgb(30 41 59);
}
</style>
