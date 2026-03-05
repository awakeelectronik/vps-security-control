package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SecurityEvent struct {
	ID            int       `json:"id"`
	EventType     string    `json:"event_type"`
	AgentID       int       `json:"agent_id"`
	SourceIP      string    `json:"source_ip"`
	SourceCountry string    `json:"source_country"`
	Username      string    `json:"username"`
	Service       string    `json:"service"`
	Description   string    `json:"description"`
	Severity      string    `json:"severity"`
	Acknowledged  bool      `json:"acknowledged"`
	CreatedAt     time.Time `json:"created_at"`
}

type LoginAttempt struct {
	ID            int       `json:"id"`
	AgentID       int       `json:"agent_id"`
	LoginType     string    `json:"login_type"`
	Username      string    `json:"username"`
	SourceIP      string    `json:"source_ip"`
	SourceCountry string    `json:"source_country"`
	Status        string    `json:"status"`
	Port          int       `json:"port"`
	CreatedAt     time.Time `json:"created_at"`
}

func GetSecurityEvents(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := 50
		offset := 0

		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil {
				limit = parsed
			}
		}

		if o := c.Query("offset"); o != "" {
			if parsed, err := strconv.Atoi(o); err == nil {
				offset = parsed
			}
		}

		rows, err := db.Query(`
			SELECT id, event_type, agent_id, source_ip, source_country, username, service, description, severity, acknowledged, created_at
			FROM security_events
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2
			
		`, limit, offset)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var events []SecurityEvent
		for rows.Next() {
			var event SecurityEvent
			if err := rows.Scan(
				&event.ID,
				&event.EventType,
				&event.AgentID,
				&event.SourceIP,
				&event.SourceCountry,
				&event.Username,
				&event.Service,
				&event.Description,
				&event.Severity,
				&event.Acknowledged,
				&event.CreatedAt,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			events = append(events, event)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   events,
			"count":  len(events),
			"limit":  limit,
			"offset": offset,
		})
	}
}

func GetSecurityEventByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var event SecurityEvent
		err := db.QueryRow(`
			SELECT id, event_type, agent_id, source_ip, source_country, username, service, description, severity, acknowledged, created_at
			FROM security_events
			WHERE id = $1
		`, id).Scan(
			&event.ID,
			&event.EventType,
			&event.AgentID,
			&event.SourceIP,
			&event.SourceCountry,
			&event.Username,
			&event.Service,
			&event.Description,
			&event.Severity,
			&event.Acknowledged,
			&event.CreatedAt,
		)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, event)
	}
}

func GetFailedLogins(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := 100
		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil {
				limit = parsed
			}
		}

		rows, err := db.Query(`
			SELECT id, agent_id, login_type, username, source_ip, source_country, status, port, created_at
			FROM login_attempts
			WHERE status = 'failed'
			ORDER BY created_at DESC
			LIMIT $1
		`, limit)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var attempts []LoginAttempt
		for rows.Next() {
			var attempt LoginAttempt
			if err := rows.Scan(
				&attempt.ID,
				&attempt.AgentID,
				&attempt.LoginType,
				&attempt.Username,
				&attempt.SourceIP,
				&attempt.SourceCountry,
				&attempt.Status,
				&attempt.Port,
				&attempt.CreatedAt,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			attempts = append(attempts, attempt)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  attempts,
			"count": len(attempts),
		})
	}
}

func GetDetectedAttacks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT id, event_type, agent_id, source_ip, source_country, username, service, description, severity, acknowledged, created_at
			FROM security_events
			WHERE event_type IN ('brute_force', 'port_scan', 'dos_attack')
			ORDER BY created_at DESC
			LIMIT 50
		`)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var events []SecurityEvent
		for rows.Next() {
			var event SecurityEvent
			if err := rows.Scan(
				&event.ID,
				&event.EventType,
				&event.AgentID,
				&event.SourceIP,
				&event.SourceCountry,
				&event.Username,
				&event.Service,
				&event.Description,
				&event.Severity,
				&event.Acknowledged,
				&event.CreatedAt,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			events = append(events, event)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  events,
			"count": len(events),
		})
	}
}

func GetThreats(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.Query("ip")

		if ip == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ip parameter required"})
			return
		}

		rows, err := db.Query(`
			SELECT id, event_type, agent_id, source_ip, source_country, username, service, description, severity, acknowledged, created_at
			FROM security_events
			WHERE source_ip = $1
			ORDER BY created_at DESC
		`, ip)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var events []SecurityEvent
		for rows.Next() {
			var event SecurityEvent
			if err := rows.Scan(
				&event.ID,
				&event.EventType,
				&event.AgentID,
				&event.SourceIP,
				&event.SourceCountry,
				&event.Username,
				&event.Service,
				&event.Description,
				&event.Severity,
				&event.Acknowledged,
				&event.CreatedAt,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			events = append(events, event)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  events,
			"count": len(events),
		})
	}
}

func CreateSecurityEvent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event SecurityEvent
		if err := c.BindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var id int
		err := db.QueryRow(`
			INSERT INTO security_events (event_type, agent_id, source_ip, source_country, username, service, description, severity)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING id
		`,
			event.EventType,
			event.AgentID,
			event.SourceIP,
			event.SourceCountry,
			event.Username,
			event.Service,
			event.Description,
			event.Severity,
		).Scan(&id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}
