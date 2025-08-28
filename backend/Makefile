# Variables
APP_NAME=devops-converter
VERSION?=1.0.0
BUILD_DIR=build
BINARY_NAME=$(APP_NAME)
DOCKER_IMAGE=$(APP_NAME):$(VERSION)

# Commandes Go
GO_BUILD=go build
GO_TEST=go test
GO_CLEAN=go clean
GO_MOD=go mod

# Couleurs pour les messages
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

.PHONY: help build test clean deps run dev docker fmt lint security

# Aide par défaut
help: ## Affiche cette aide
	@echo "$(GREEN)Commandes disponibles:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

# Build
build: deps ## Compile l'application
	@echo "$(GREEN)Building $(APP_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "$(GREEN)Build complete: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

build-prod: deps ## Compile l'application pour la production
	@echo "$(GREEN)Building $(APP_NAME) for production...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -ldflags="-w -s" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	@echo "$(GREEN)Production build complete: $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64$(NC)"

# Tests
test: ## Lance les tests unitaires
	@echo "$(GREEN)Running tests...$(NC)"
	@$(GO_TEST) -v ./...

test-coverage: ## Lance les tests avec couverture
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	@$(GO_TEST) -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

test-race: ## Lance les tests avec détection de race conditions
	@echo "$(GREEN)Running tests with race detection...$(NC)"
	@$(GO_TEST) -v -race ./...

# Développement
dev: ## Lance l'application en mode développement
	@echo "$(GREEN)Starting development server...$(NC)"
	@cp .env.example .env 2>/dev/null || true
	@go run . 

run: build ## Compile et lance l'application
	@echo "$(GREEN)Starting $(APP_NAME)...$(NC)"
	@./$(BUILD_DIR)/$(BINARY_NAME)

watch: ## Lance l'application avec rechargement automatique (nécessite entr)
	@echo "$(GREEN)Starting with auto-reload...$(NC)"
	@find . -name "*.go" | entr -r make dev

# Gestion des dépendances
deps: ## Télécharge et vérifie les dépendances
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	@$(GO_MOD) download
	@$(GO_MOD) tidy

deps-update: ## Met à jour les dépendances
	@echo "$(GREEN)Updating dependencies...$(NC)"
	@$(GO_MOD) get -u ./...
	@$(GO_MOD) tidy

deps-vendor: ## Crée le dossier vendor
	@echo "$(GREEN)Creating vendor directory...$(NC)"
	@$(GO_MOD) vendor

# Code quality
fmt: ## Formate le code
	@echo "$(GREEN)Formatting code...$(NC)"
	@go fmt ./...

lint: ## Lance le linter
	@echo "$(GREEN)Running linter...$(NC)"
	@golangci-lint run ./...

security: ## Analyse de sécurité avec gosec
	@echo "$(GREEN)Running security analysis...$(NC)"
	@gosec ./...

# Docker
docker-build: ## Construit l'image Docker
	@echo "$(GREEN)Building Docker image...$(NC)"
	@docker build -t $(DOCKER_IMAGE) .

docker-run: docker-build ## Lance le conteneur Docker
	@echo "$(GREEN)Running Docker container...$(NC)"
	@docker run --rm -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

docker-push: docker-build ## Pousse l'image Docker vers le registre
	@echo "$(GREEN)Pushing Docker image...$(NC)"
	@docker push $(DOCKER_IMAGE)

# Nettoyage
clean: ## Nettoie les fichiers de build
	@echo "$(GREEN)Cleaning up...$(NC)"
	@$(GO_CLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@rm -f $(BINARY_NAME)

clean-all: clean ## Nettoie tout (build + cache)
	@echo "$(GREEN)Deep cleaning...$(NC)"
	@go clean -cache -modcache -i -r

# Installation des outils de développement
install-tools: ## Installe les outils de développement
	@echo "$(GREEN)Installing development tools...$(NC)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# Génération de documentation
docs: ## Génère la documentation
	@echo "$(GREEN)Generating documentation...$(NC)"
	@godoc -http=:6060 &
	@echo "$(GREEN)Documentation available at http://localhost:6060$(NC)"

# Migration et setup
setup: deps install-tools ## Configuration initiale du projet
	@echo "$(GREEN)Setting up project...$(NC)"
	@cp .env.example .env 2>/dev/null || true
	@echo "$(GREEN)Project setup complete!$(NC)"
	@echo "$(YELLOW)Edit .env file with your configuration$(NC)"

# Version et release
version: ## Affiche la version
	@echo "$(GREEN)$(APP_NAME) version $(VERSION)$(NC)"

release: test build-prod ## Prépare une release
	@echo "$(GREEN)Preparing release $(VERSION)...$(NC)"
	@tar -czf $(BUILD_DIR)/$(APP_NAME)-$(VERSION)-linux-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-linux-amd64
	@echo "$(GREEN)Release package ready: $(BUILD_DIR)/$(APP_NAME)-$(VERSION)-linux-amd64.tar.gz$(NC)"

# Initialisation frontend (si nécessaire)
frontend-init: ## Initialise le projet frontend
	@echo "$(GREEN)Initializing frontend...$(NC)"
	@cd ../frontend && npm init vue@latest . --typescript --router --pinia

# Build complet
all: clean deps test build ## Clean, test et build
	@echo "$(GREEN)All tasks completed successfully!$(NC)"
