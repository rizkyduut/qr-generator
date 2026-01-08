package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rizkyduut/qr-generator/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	r := gin.Default()

	// Apply security headers globally
	if cfg.Security.EnableSecurityHeaders {
		// TODO: Add security headers middleware
	}

	// Set request size limit
	// TODO: Add request size limit middleware

	// Basic health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	// Configure HTTP server with timeouts
	srv := &http.Server{
		Addr:         cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Printf("Security: Security headers=%v", cfg.Security.EnableSecurityHeaders)

	// Graceful shutdown
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
