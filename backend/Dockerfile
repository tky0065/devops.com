# Multi-stage build pour optimiser la taille de l'image

# Stage 1: Build
FROM golang:1.24.5-alpine AS builder

# build-time target architecture (populated by BuildKit / docker buildx)
ARG TARGETARCH

# Installer les dépendances nécessaires
RUN apk add --no-cache git ca-certificates tzdata

# Créer un utilisateur non-root
RUN adduser -D -s /bin/sh -u 1001 appuser

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers de dépendances
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le code source
COPY . .

# Build de l'application avec optimisations
# Use TARGETARCH if provided (e.g. arm64/amd64) so the produced binary matches the host
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH:-amd64} go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o devops-converter .

# Stage 2: Runtime
FROM scratch

# Copier les certificats CA
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copier les informations de timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copier l'utilisateur
COPY --from=builder /etc/passwd /etc/passwd

# Copier le binaire
COPY --from=builder /app/devops-converter /devops-converter

# Créer le répertoire pour les logs
USER appuser

# Exposer le port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/devops-converter", "--health-check"]

# Variables d'environnement par défaut
ENV APP_ENV=production
ENV SERVER_HOST=0.0.0.0
ENV SERVER_PORT=8080

# Point d'entrée
ENTRYPOINT ["/devops-converter"]
