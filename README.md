# DevOps Converter

Petit projet qui convertit des fichiers Docker Compose en manifestes Kubernetes.

Ce README décrit rapidement l'architecture, comment lancer l'application en local, exécuter les tests Playwright et résoudre les erreurs courantes (notamment les 400 lors de conversions).

## Architecture

- backend: API en Go (Gin) exposant les endpoints de conversion et d'upload.
- converters: logique de conversion (docker -> kubernetes).
- frontend: UI Vue 3 + Pinia + Vite, contient des tests Playwright dans `frontend/tests`.

### Ports par défaut

- Backend: `http://localhost:8081`
- Frontend (Vite dev): `http://localhost:5173` (Vite) — Playwright config utilise `5174` par défaut, vérifiez `frontend/playwright.config.ts` si nécessaire.

## Prérequis

- Go 1.20+ (ou version compatible avec go.mod)
- Node.js 16+ / npm
- Playwright (pour tests E2E)
## Lancer localement (développement)

Ouvrir deux terminaux (ou utilisez des onglets) :

1. Backend

```bash
# depuis le dossier du repo
cd backend
go run main.go
# écoute sur localhost:8081
```

1. Frontend

```bash
cd frontend
npm install
npm run dev
# par défaut Vite sert sur 5173
```

Notes:

- Si vous utilisez un `VITE_API_BASE_URL` différent, exportez la variable avant de lancer le frontend, p.ex.:

```bash
export VITE_API_BASE_URL=http://localhost:8081
```

## Endpoints API (essentiels)

- POST /api/v1/convert/ — corps JSON attendu:

```bash
GET /health/
GET /health/detailed
GET /health/ready
GET /health/live
```

- POST /api/v1/convert/validate — JSON similaire à convert mais pour validation

- POST /api/v1/upload/ — attend `multipart/form-data` (champ `file`, `type`, options en champs supplémentaires)

Réponses: la réponse renvoie généralement un JSON `ConversionResponse` contenant `success`, `files[]`, `errors[]`, `warnings[]`, `metadata`.

## Exécuter les tests Playwright (E2E)

Playwright est configuré dans `frontend/playwright.config.ts`.

```bash
cd frontend
npx playwright test
```

### Test de régression: conversion /api/v1/convert/ (cas 400)

Un test Playwright `frontend/tests/convert-400.spec.ts` est ajouté pour reproduire le cas où un mapping de volume à une seule partie (par ex. `/app/node_modules`) pouvait déclencher un HTTP 400.

Pour lancer uniquement ce test:

```bash
cd frontend
npx playwright test tests/convert-400.spec.ts -p chromium
```

Le test vérifie que l'API ne répond pas avec un statut 400 pour cet exemple.


Si Playwright ne trouve pas le serveur frontend, vérifiez le `baseURL` et `webServer.url` dans `playwright.config.ts`. Le config par défaut pointe vers `http://localhost:5174`.

## Tests unitaires / backend

Pour lancer les tests Go (backend) :

```bash
cd backend
go test ./...
```
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

## Debugging rapide — 400 Bad Request sur `/api/v1/convert/`

Symptôme: la requête XHR POST vers `/api/v1/convert/` retourne HTTP 400 et l'UI affiche une erreur Axios.

Causes fréquentes et pistes:

- Payload mal formé (vérifier que le frontend envoie JSON et que `Content-Type: application/json` est présent). Le frontend fournit `convert()` qui envoie {type, content, options}.
```bash
# Liste des convertisseurs
GET /api/v1/info/converters

# Version de l'application
GET /version
```

- Confusion entre `upload` (multipart/form-data) et `convert` (JSON). Utiliser le bon endpoint.

- Volume mappings invalides dans le docker-compose (ex: une entrée comme `/app/node_modules` peut auparavant causer un échec de parsing). Le backend contient une validation/parse des mappings de volumes; les formats valides attendus sont:

  - `hostPath:containerPath[:mode]` (ex. `./data:/var/lib/data:ro`)

  - `named-volume:containerPath[:mode]`

  - Le backend a été récemment ajusté pour accepter aussi un mapping à une seule partie (par ex. `/app/node_modules`) en le traitant comme hostPath->containerPath.

- Regardez les logs backend (console) — les erreurs de binding/validation s'affichent et contiennent `request_id` permettant de tracer la requête.

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
Si vous voyez encore un 400, copiez le body retourné par l'API (le frontend a été modifié pour exposer le JSON d'erreur) et partagez l'objet `errors[]` renvoyé.

## Conseils pour contribuer

- Fork -> branch -> PR vers `main`

- Respectez les linters TypeScript et exécutez les tests Playwright et Go localement.

## Ressources utiles

- Fichiers importants:

  - `backend/main.go` — démarrage du service

  - `backend/api/handlers/convert.go` — handler conversion

  - `frontend/src/services/api.ts` — wrapper HTTP côté client

  - `frontend/playwright.config.ts` — config des tests E2E

## Licence

MIT — voir fichier LICENSE si présent.

---


## 📋 État d'avancement

### ✅ Phase 1 : Configuration initiale et structure du projet - TERMINÉE

#### Backend (Go)

- [x] Initialiser le projet Go avec `go mod init`
- [x] Configurer la structure des dossiers selon l'architecture
- [x] Installer les dépendances principales (Gin, YAML parsers, etc.)
- [x] Configurer les variables d'environnement et fichiers de config
- [x] Mettre en place le système de logging
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

```text
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

## Déploiement en production avec HTTPS (via Caddy)

Si votre VPS a déjà Traefik (ou un autre reverse-proxy) qui gère TLS, la configuration par défaut du `docker-compose.yml` utilise le réseau Docker externe `traefik` et des labels pour déclarer les routes:

1. DNS: ajoutez un enregistrement A pour `devops.enokdev.com` pointant vers l'IP publique de votre serveur (VPS).

2. Ouvrez/autorisez les ports 80 et 443 sur le serveur (firewall/cloud provider). Traefik doit pouvoir joindre Let's Encrypt via ces ports.

3. Lancer les services:

```bash
# depuis la racine du repo
docker compose up --build -d

# vérifier l'état
docker compose ps

# logs si nécessaire
docker compose logs -f backend
docker compose logs -f frontend
```

4. Vérifier TLS et accessibilité:

```bash
# vérifier que Traefik a obtenu un cert
curl -I https://devops.enokdev.com

# health checks
curl -I https://devops.enokdev.com/api/health
```

Remarques:

- Le `docker-compose.yml` attache les services au réseau Docker `traefik` (external: true). Assurez-vous que ce réseau existe sur l'hôte (généralement créé par Traefik lors de son déploiement). Vous pouvez vérifier avec `docker network ls`.
- Si le réseau `traefik` n'existe pas, créez-le manuellement :

```bash
docker network create traefik
```

- Les labels configurés sur les services exposent `devops.enokdev.com` pour le frontend et `/api` pour le backend. Adaptez les règles si vous voulez sous-domaines séparés (ex: `api.devops.enokdev.com`).

### Construire l'image backend pour l'architecture du VPS

Si votre VPS est arm64 (par ex. AWS Graviton, Apple Silicon), le binaire compilé pour amd64 provoquera `exec format error`. Deux approches :

1) Construire localement avec BuildKit / buildx (multi-arch) :

```bash
# activer buildx si nécessaire
docker buildx create --use --name mybuilder || true
docker buildx inspect --bootstrap

# builder respectera TARGETARCH automatiquement
docker buildx build --platform linux/amd64,linux/arm64 -t devops-converter-backend:latest --push ./backend
```

2) Reconstruire directement sur le VPS (simple) :

```bash
# sur le VPS, depuis la racine du repo
docker compose build --no-cache backend
docker compose up -d
```

3) Vérifier l'architecture du VPS :

```bash
uname -m
# x86_64 => amd64 ; aarch64 => arm64
```

Si vous voulez, je peux générer un `docker buildx` script dans le repo pour automatiser la publication multi-arch sur un registry.


