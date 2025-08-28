package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"devops-converter/api/routes"
	"devops-converter/config"
	"devops-converter/converters"
)

func main() {
	// Charger la configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Afficher les informations de démarrage
	fmt.Printf("Starting %s v%s in %s mode\n", cfg.App.Name, cfg.App.Version, cfg.App.Environment)

	// Créer le registre des convertisseurs
	registry := converters.NewRegistry()

	// Enregistrer les convertisseurs disponibles
	if err := registerConverters(registry); err != nil {
		log.Fatalf("Failed to register converters: %v", err)
	}

	// Configurer les routes
	router := routes.SetupRoutes(cfg, registry)

	// Configurer le serveur HTTP
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	// Canal pour capturer les signaux d'arrêt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Démarrer le serveur dans une goroutine
	go func() {
		fmt.Printf("Server starting on %s\n", server.Addr)
		
		var err error
		if cfg.Server.TLS.Enabled {
			fmt.Println("Starting HTTPS server...")
			err = server.ListenAndServeTLS(cfg.Server.TLS.CertFile, cfg.Server.TLS.KeyFile)
		} else {
			fmt.Println("Starting HTTP server...")
			err = server.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %v\n", err)
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Attendre le signal d'arrêt
	<-quit
	fmt.Println("\nShutting down server...")

	// Arrêt gracieux avec timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exited")
}

// registerConverters enregistre tous les convertisseurs disponibles
func registerConverters(registry converters.ConverterRegistry) error {
	// Enregistrer le convertisseur Docker Compose vers Kubernetes
	dockerComposeConverter := converters.NewDockerComposeToKubernetesConverter()
	if err := registry.Register(dockerComposeConverter); err != nil {
		return fmt.Errorf("failed to register docker-compose converter: %w", err)
	}

	// Ici, on pourrait ajouter d'autres convertisseurs :
	// - Dockerfile vers Kubernetes
	// - Terraform vers Kubernetes
	// - Helm Charts, etc.

	fmt.Printf("Registered %d converter(s)\n", len(registry.GetAvailableConverters()))
	
	return nil
}
