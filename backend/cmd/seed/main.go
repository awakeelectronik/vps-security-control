package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"vps-security-control/backend/internal/config"
	"vps-security-control/backend/internal/database"
)

func main() {
	godotenv.Load()

	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Insert sample agent
	_, err = db.Exec(`
		INSERT INTO agents (uuid, name, ip_address, status, last_seen)
		VALUES ('550e8400-e29b-41d4-a716-446655440000', 'production-vps-1', '192.168.1.100', 'active', NOW())
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		log.Fatalf("Failed to insert agent: %v", err)
	}

	// Insert sample security events
	eventTypes := []string{"failed_login", "brute_force", "port_scan", "service_error"}
	severities := []string{"critical", "high", "medium", "low"}

	for i := 0; i < 20; i++ {
		eventType := eventTypes[i%len(eventTypes)]
		severity := severities[i%len(severities)]

		_, err := db.Exec(`
			INSERT INTO security_events (event_type, agent_id, source_ip, source_country, username, service, description, severity)
			VALUES ($1, 1, $2, $3, $4, $5, $6, $7)
		`,
			eventType,
			fmt.Sprintf("192.168.1.%d", 50+i),
			"US",
			"admin",
			"ssh",
			fmt.Sprintf("Sample %s event", eventType),
			severity,
		)
		if err != nil {
			log.Fatalf("Failed to insert security event: %v", err)
		}
	}

	// Insert sample login attempts
	for i := 0; i < 30; i++ {
		status := "failed"
		if i%5 == 0 {
			status = "success"
		}

		_, err := db.Exec(`
			INSERT INTO login_attempts (agent_id, login_type, username, source_ip, source_country, status, port)
			VALUES (1, 'ssh', $1, $2, $3, $4, 22)
		`,
			"root",
			fmt.Sprintf("203.0.113.%d", 1+i),
			"CN",
			status,
		)
		if err != nil {
			log.Fatalf("Failed to insert login attempt: %v", err)
		}
	}

	// Insert sample service errors
	_, err = db.Exec(`
		INSERT INTO service_errors (agent_id, service_name, error_type, error_message, error_count, last_occurrence)
		VALUES
		(1, 'nginx', '502', 'Bad Gateway - upstream server timeout', 5, NOW()),
		(1, 'mysql', 'connection_timeout', 'Connection timeout to MySQL server', 2, NOW()),
		(1, 'sshd', 'auth_failure', 'Too many authentication failures', 10, NOW())
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		log.Fatalf("Failed to insert service errors: %v", err)
	}

	// Insert sample system metrics
	_, err = db.Exec(`
		INSERT INTO system_metrics (agent_id, cpu_usage, memory_usage, disk_usage, network_in, network_out)
		VALUES (1, 45.2, 62.8, 71.3, 1024000000, 512000000)
	`)
	if err != nil {
		log.Fatalf("Failed to insert system metrics: %v", err)
	}

	fmt.Println("\u2705 Database seeded with sample data successfully")
}
