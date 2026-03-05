package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func RunMigrations(db *sql.DB) error {
	// Create tables
	schema := `
		CREATE TABLE IF NOT EXISTS agents (
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			ip_address VARCHAR(45) NOT NULL,
			status VARCHAR(50) DEFAULT 'active',
			last_seen TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS security_events (
			id SERIAL PRIMARY KEY,
			event_type VARCHAR(50) NOT NULL,
			agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE,
			source_ip VARCHAR(45),
			source_country VARCHAR(100),
			username VARCHAR(255),
			service VARCHAR(100),
			description TEXT,
			severity VARCHAR(20),
			acknowledged BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS login_attempts (
			id SERIAL PRIMARY KEY,
			agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE,
			login_type VARCHAR(50),
			username VARCHAR(255),
			source_ip VARCHAR(45),
			source_country VARCHAR(100),
			status VARCHAR(20),
			port INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS service_errors (
			id SERIAL PRIMARY KEY,
			agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE,
			service_name VARCHAR(100),
			error_type VARCHAR(100),
			error_message TEXT,
			error_count INTEGER DEFAULT 1,
			last_occurrence TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS system_metrics (
			id SERIAL PRIMARY KEY,
			agent_id INTEGER REFERENCES agents(id) ON DELETE CASCADE,
			cpu_usage FLOAT,
			memory_usage FLOAT,
			disk_usage FLOAT,
			network_in BIGINT,
			network_out BIGINT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE INDEX IF NOT EXISTS idx_security_events_agent_id ON security_events(agent_id);
		CREATE INDEX IF NOT EXISTS idx_security_events_created_at ON security_events(created_at);
		CREATE INDEX IF NOT EXISTS idx_login_attempts_agent_id ON login_attempts(agent_id);
		CREATE INDEX IF NOT EXISTS idx_login_attempts_source_ip ON login_attempts(source_ip);
		CREATE INDEX IF NOT EXISTS idx_service_errors_agent_id ON service_errors(agent_id);
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
