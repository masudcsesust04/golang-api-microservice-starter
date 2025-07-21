package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/masudcsesust04/golang-jwt-auth/internal/config"
	"github.com/masudcsesust04/golang-jwt-auth/internal/handlers"
	"github.com/masudcsesust04/golang-jwt-auth/internal/models"
	"github.com/masudcsesust04/golang-jwt-auth/internal/utils"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	if config.AppConfig.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	err := config.InitDB(config.AppConfig.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer config.DbConn.Close()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(&models.User{})

	// Initialize JWT middleware
	utils.SetJWTSecrectKey(config.AppConfig.JWTSecret)

	// Setup router
	router := mux.NewRouter()

	// user routes
	router.Handle("/users", utils.JWTMiddleware(userHandler.GetUsers)).Methods("GET")
	router.Handle("/users/{id}", utils.JWTMiddleware(userHandler.GetUser)).Methods("GET")
	router.Handle("/users/{id}", utils.JWTMiddleware(userHandler.UpdateUser)).Methods("PUT")
	router.Handle("/users/{id}", utils.JWTMiddleware(userHandler.DeleteUser)).Methods("DELETE")

	// Start server
	addr := ":" + config.AppConfig.ServerPort
	log.Printf("User service starting server on %s", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Run the server in a goroutine so that it doesn't block the graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Block until a signal is received
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
