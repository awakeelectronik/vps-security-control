package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceHealth struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Port   int    `json:"port"`
	Errors int    `json:"errors"`
}

type SystemHealth struct {
	CPUUsage     float64 `json:"cpu_usage"`
	MemoryUsage  float64 `json:"memory_usage"`
	DiskUsage    float64 `json:"disk_usage"`
	UptimeHours  int     `json:"uptime_hours"`
}

func GetNginxHealth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement actual Nginx status check via agent
		health := ServiceHealth{
			Name:   "Nginx",
			Status: "running",
			Port:   80,
			Errors: 0,
		}
		c.JSON(http.StatusOK, health)
	}
}

func GetMySQLHealth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Test database connection
		err := db.Ping()
		status := "running"
		if err != nil {
			status = "error"
		}

		health := ServiceHealth{
			Name:   "MySQL",
			Status: status,
			Port:   3306,
			Errors: 0,
		}
		c.JSON(http.StatusOK, health)
	}
}

func GetSystemHealth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement actual system metrics from agent
		health := SystemHealth{
			CPUUsage:    35.2,
			MemoryUsage: 62.8,
			DiskUsage:   45.1,
			UptimeHours: 168,
		}
		c.JSON(http.StatusOK, health)
	}
}
