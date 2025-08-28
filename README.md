# 🚀 Plateforme de Conversion Docker vers Kubernetes

Une plateforme web extensible pour convertir des fichiers Docker (docker-compose.yml, Dockerfile) en configurations Kubernetes.

## 📋 État d'avancement

### ✅ Phase 1 : Configuration initiale et structure du projet - TERMINÉE

#### Backend (Go)
- [x] Initialiser le projet Go avec `go mod init`
- [x] Configurer la structure des dossiers selon l'architecture
- [x] Installer les dépendances principales (Gin, YAML parsers, etc.)
- [x] Configurer les variables d'environnement et fichiers de config
- [x] Mettre en place le système de logging

#### Infrastructure
- [x] Configurer Docker pour le développement local
- [x] Créer docker-compose.yml pour l'environnement de dev
- [x] Configurer les fichiers .gitignore
- [x] Créer la documentation README initiale

### ✅ Phase 2 : Développement du backend - TERMINÉE

#### API Core
- [x] Implémenter le serveur HTTP de base avec Gin
- [x] Créer la structure des middlewares (CORS, logging, error handling)
- [x] Implémenter l'endpoint de health check
- [x] Configurer la validation des requêtes

#### Système de conversion modulaire
- [x] Définir l'interface `Converter` 
- [x] Implémenter le registry des convertisseurs
- [x] Créer la factory `GetConverter()`
- [x] Implémenter la gestion des erreurs de conversion

#### Parsers Docker
- [x] Créer le parser pour docker-compose.yml
  - [x] Parser les services
  - [x] Parser les volumes
  - [x] Parser les réseaux
  - [x] Parser les variables d'environnement
- [x] Gestion des formats complexes (ports, environnement, etc.)

#### Générateurs Kubernetes
- [x] Implémenter le générateur de Deployments
- [x] Implémenter le générateur de Services
- [x] Implémenter le générateur de ConfigMaps
- [x] Implémenter le générateur de PersistentVolumes
- [x] Implémenter le générateur d'Ingress
- [x] Créer les utilitaires YAML (formatting, validation)

#### API Endpoints
- [x] POST `/convert` - endpoint principal de conversion
- [x] GET `/converters` - liste des convertisseurs disponibles
- [x] POST `/validate` - validation des fichiers d'entrée
- [x] GET `/health` - health check

## 🏗️ Architecture Implémentée

```
Frontend (Vue.js + TS) ←→ REST API ←→ Backend (Go) ←→ Modules de Conversion ←→ Générateur YAML
```

### Structure du Projet
```
devops.com/
├── backend/                    # API Go
│   ├── api/
│   │   ├── handlers/          # Gestionnaires HTTP
│   │   ├── middleware/        # Middlewares (CORS, logging, etc.)
│   │   └── routes/           # Configuration des routes
│   ├── converters/           # Système modulaire de conversion
│   │   ├── docker/          # Parsers Docker
│   │   └── kubernetes/      # Générateurs Kubernetes
│   ├── config/              # Configuration de l'application
│   ├── utils/               # Utilitaires (YAML, validation)
│   ├── main.go              # Point d'entrée
│   ├── Dockerfile           # Image Docker pour production
│   └── Makefile            # Commandes de développement
├── frontend/               # Interface Vue.js (à développer)
├── docker-compose.yml      # Environnement de développement
└── README.md              # Documentation
```

## 🛠️ Technologies Utilisées

### Backend
- **Go 1.21+** - Langage principal
- **Gin** - Framework web HTTP
- **YAML v3** - Parser YAML pour Docker Compose et Kubernetes
- **Testify** - Framework de tests
- **GoDotEnv** - Gestion des variables d'environnement

### Convertisseurs Implémentés
- **Docker Compose → Kubernetes** : Conversion complète avec support de :
  - Services → Deployments + Services
  - Volumes → PersistentVolumes + PersistentVolumeClaims
  - Variables d'environnement → ConfigMaps
  - Health checks → Probes
  - Ressources → ResourceRequirements
  - Ports → Services avec sélecteurs

## 🚀 Utilisation

### Démarrage en mode développement

```bash
# 1. Se placer dans le dossier backend
cd backend

# 2. Installer les dépendances
make deps

# 3. Copier la configuration
cp .env.example .env

# 4. Lancer en mode développement
make dev
```

### Endpoints API disponibles

#### Health Check
```bash
GET /health/
GET /health/detailed
GET /health/ready
GET /health/live
```

#### Conversion
```bash
# Convertir un fichier
POST /api/v1/convert/
Content-Type: application/json

{
  "type": "docker-compose",
  "content": "version: '3.8'\nservices:\n  web:\n    image: nginx",
  "options": {
    "namespace": "production",
    "serviceType": "LoadBalancer",
    "replicas": 3
  }
}

# Valider un fichier
POST /api/v1/convert/validate
Content-Type: application/json

{
  "type": "docker-compose",
  "content": "version: '3.8'\nservices:\n  web:\n    image: nginx"
}

# Upload et conversion
POST /api/v1/upload/
Content-Type: multipart/form-data

file: [docker-compose.yml]
type: docker-compose
namespace: production
```

#### Information
```bash
# Liste des convertisseurs
GET /api/v1/info/converters

# Version de l'application
GET /version
```

### Exemple de Conversion

#### Entrée (docker-compose.yml)
```yaml
version: '3.8'
services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
    environment:
      - ENV=production
    volumes:
      - ./html:/usr/share/nginx/html:ro
```

#### Sortie (Kubernetes)
```yaml
# web-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      containers:
      - name: web
        image: nginx:latest
        ports:
        - containerPort: 80
        env:
        - name: ENV
          value: production
---
# web-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: web
spec:
  selector:
    app: web
  ports:
  - port: 80
    targetPort: 80
```

## 🧪 Tests

```bash
# Tests unitaires
make test

# Tests avec couverture
make test-coverage

# Tests de race conditions
make test-race
```

## 🐳 Docker

```bash
# Build de l'image
make docker-build

# Lancement avec Docker
make docker-run

# Avec docker-compose
docker-compose up --build
```

## 📝 Fonctionnalités Avancées Implémentées

### Validation Complète
- Validation des formats Docker Compose
- Validation des noms Kubernetes
- Validation des ports et volumes
- Gestion d'erreurs détaillée avec suggestions

### Conversion Intelligente
- Mapping automatique des ressources
- Génération de ConfigMaps pour les variables non-sensibles
- Support des health checks → probes Kubernetes
- Conversion des contraintes de ressources
- Gestion des volumes nommés vs bind mounts

### Sécurité
- Headers de sécurité HTTP
- Rate limiting
- Validation stricte des entrées
- Sandboxing des conversions

### Observabilité
- Logging structuré
- Health checks multiples (liveness, readiness)
- Métriques détaillées
- Traces de requêtes

## 🔄 Prochaines Étapes

### Phase 3 : Frontend Vue.js
- [ ] Interface utilisateur intuitive
- [ ] Upload par drag & drop
- [ ] Prévisualisation des résultats
- [ ] Téléchargement des fichiers générés

### Phase 4 : Fonctionnalités Avancées
- [ ] Support Dockerfile → Kubernetes
- [ ] Convertisseur Terraform → Kubernetes
- [ ] Helm Charts generation
- [ ] Multi-fichiers (compose + env files)

### Phase 5 : Production
- [ ] CI/CD pipeline
- [ ] Monitoring avancé
- [ ] Déploiement Kubernetes
- [ ] Documentation API complète

## 📄 Licence

Ce projet est sous licence MIT.

---

**Statut** : ✅ Backend fonctionnel - Prêt pour le développement frontend
**Version** : 1.0.0
**Dernière mise à jour** : 28 août 2025
