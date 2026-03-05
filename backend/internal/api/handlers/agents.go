package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Agent struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	IPAddress string    `json:"ip_address"`
	Status    string    `json:"status"`
	LastSeen  time.Time `json:"last_seen"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterAgentRequest struct {
	Name      string `json:"name" binding:"required"`
	IPAddress string `json:"ip_address" binding:"required"`
}

func RegisterAgent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterAgentRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		agentUUID := uuid.New().String()

		var id int
		err := db.QueryRow(`
			INSERT INTO agents (uuid, name, ip_address, status, last_seen)
			VALUES ($1, $2, $3, 'active', NOW())
			RETURNING id
		`,
			agentUUID,
			req.Name,
			req.IPAddress,
		).Scan(&id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":   id,
			"uuid": agentUUID,
		})
	}
}

func ListAgents(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT id, uuid, name, ip_address, status, last_seen, created_at, updated_at
			FROM agents
			ORDER BY created_at DESC
		`)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var agents []Agent
		for rows.Next() {
			var agent Agent
			if err := rows.Scan(
				&agent.ID,
				&agent.UUID,
				&agent.Name,
				&agent.IPAddress,
				&agent.Status,
				&agent.LastSeen,
				&agent.CreatedAt,
				&agent.UpdatedAt,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			agents = append(agents, agent)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  agents,
			"count": len(agents),
		})
	}
}

func GetAgent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var agent Agent
		err := db.QueryRow(`
			SELECT id, uuid, name, ip_address, status, last_seen, created_at, updated_at
			FROM agents
			WHERE id = $1
		`, id).Scan(
			&agent.ID,
			&agent.UUID,
			&agent.Name,
			&agent.IPAddress,
			&agent.Status,
			&agent.LastSeen,
			&agent.CreatedAt,
			&agent.UpdatedAt,
		)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "agent not found"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, agent)
	}
}

func DeleteAgent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		result, err := db.Exec("DELETE FROM agents WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "agent not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "agent deleted"})
	}
}
