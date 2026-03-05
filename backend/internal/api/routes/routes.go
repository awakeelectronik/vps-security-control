package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"vps-security-control/backend/internal/api/handlers"
	"vps-security-control/backend/internal/config"
)

func SetupRoutes(router *gin.Engine, db *sql.DB, cfg *config.Config) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := router.Group("/api/v1")
	{
		// Security events endpoints
		security := api.Group("/security")
		{
			security.GET("/events", handlers.GetSecurityEvents(db))
			security.GET("/events/:id", handlers.GetSecurityEventByID(db))
			security.GET("/failed-logins", handlers.GetFailedLogins(db))
			security.GET("/attacks", handlers.GetDetectedAttacks(db))
			security.GET("/threats", handlers.GetThreats(db))
			security.POST("/events", handlers.CreateSecurityEvent(db))
		}

		// Service health endpoints
		health := api.Group("/health")
		{
			health.GET("/nginx", handlers.GetNginxHealth(db))
			health.GET("/mysql", handlers.GetMySQLHealth(db))
			health.GET("/system", handlers.GetSystemHealth(db))
		}

		// Agent management endpoints
		agents := api.Group("/agents")
		{
			agents.POST("/register", handlers.RegisterAgent(db))
			agents.GET("", handlers.ListAgents(db))
			agents.GET("/:id", handlers.GetAgent(db))
			agents.DELETE("/:id", handlers.DeleteAgent(db))
		}
	}
}
