#!/bin/bash

# Script de démarrage pour DevOps Converter
# Ce script démarre le backend Go et le frontend Vue.js

echo "🚀 Démarrage de DevOps Converter"
echo "================================="

# Couleurs pour les logs
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Fonction pour afficher des messages colorés
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_header() {
    echo -e "${BLUE}$1${NC}"
}

# Vérifier si Go est installé
if ! command -v go &> /dev/null; then
    print_error "Go n'est pas installé. Veuillez l'installer depuis https://golang.org/"
    exit 1
fi

# Vérifier si Node.js est installé
if ! command -v node &> /dev/null; then
    print_error "Node.js n'est pas installé. Veuillez l'installer depuis https://nodejs.org/"
    exit 1
fi

# Vérifier si npm est installé
if ! command -v npm &> /dev/null; then
    print_error "npm n'est pas installé. Veuillez l'installer avec Node.js"
    exit 1
fi

print_header "🔧 Vérification des dépendances..."

# Aller dans le répertoire backend
cd "$(dirname "$0")/backend" || exit 1

# Vérifier si go.mod existe
if [ ! -f "go.mod" ]; then
    print_warning "go.mod introuvable, initialisation du module Go..."
    go mod init devops-converter
fi

# Installer les dépendances Go
print_status "Installation des dépendances Go..."
go mod tidy

# Aller dans le répertoire frontend
cd "../frontend" || exit 1

# Vérifier si node_modules existe
if [ ! -d "node_modules" ]; then
    print_status "Installation des dépendances Node.js..."
    npm install
fi

print_header "🚀 Démarrage des services..."

# Fonction pour nettoyer les processus à la sortie
cleanup() {
    print_warning "Arrêt des services..."
    kill $BACKEND_PID 2>/dev/null
    kill $FRONTEND_PID 2>/dev/null
    exit 0
}

# Capturer Ctrl+C pour nettoyer
trap cleanup SIGINT

# Démarrer le backend en arrière-plan
print_status "Démarrage du backend Go sur le port 8081..."
cd "../backend"
go run main.go &
BACKEND_PID=$!

# Attendre que le backend démarre
sleep 3

# Vérifier si le backend est démarré
if kill -0 $BACKEND_PID 2>/dev/null; then
    print_status "✅ Backend démarré avec succès (PID: $BACKEND_PID)"
else
    print_error "❌ Échec du démarrage du backend"
    exit 1
fi

# Démarrer le frontend
print_status "Démarrage du frontend Vue.js sur le port 5173..."
cd "../frontend"
npm run dev &
FRONTEND_PID=$!

# Attendre que le frontend démarre
sleep 5

print_header "🎉 Application démarrée avec succès!"
echo ""
print_status "📱 Frontend: http://localhost:5173"
print_status "🔧 Backend API: http://localhost:8081"
print_status "🩺 Health Check: http://localhost:8081/health"
echo ""
print_warning "Appuyez sur Ctrl+C pour arrêter les services"

# Fonction pour vérifier l'état des services
check_services() {
    while true; do
        sleep 30
        
        if ! kill -0 $BACKEND_PID 2>/dev/null; then
            print_error "❌ Backend s'est arrêté de manière inattendue"
            cleanup
        fi
        
        if ! kill -0 $FRONTEND_PID 2>/dev/null; then
            print_error "❌ Frontend s'est arrêté de manière inattendue"
            cleanup
        fi
    done
}

# Surveillance des services en arrière-plan
check_services &
MONITOR_PID=$!

# Attendre indéfiniment
wait
