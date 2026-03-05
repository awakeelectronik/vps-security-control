# VPS Security Control - Architecture

## System Design

### High-Level Overview

```
┌─────────────────────────────────────────────────────────────────────┐
│                         VPS Servers (Multiple)                      │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐  │
│  │   VPS #1         │  │   VPS #2         │  │   VPS #N         │  │
│  │ ┌──────────────┐ │  │ ┌──────────────┐ │  │ ┌──────────────┐ │  │
│  │ │ Security     │ │  │ │ Security     │ │  │ │ Security     │ │  │
│  │ │ Agent        │ │  │ │ Agent        │ │  │ │ Agent        │ │  │
│  │ │ (Go Binary)  │ │  │ │ (Go Binary)  │ │  │ │ (Go Binary)  │ │  │
│  │ └──────────────┘ │  │ └──────────────┘ │  │ └──────────────┘ │  │
│  │   - auth.log     │  │   - auth.log     │  │   - auth.log     │  │
│  │   - nginx logs   │  │   - nginx logs   │  │   - nginx logs   │  │
│  │   - mysql logs   │  │   - mysql logs   │  │   - mysql logs   │  │
│  │   - syslog       │  │   - syslog       │  │   - syslog       │  │
│  │   - metrics      │  │   - metrics      │  │   - metrics      │  │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘  │
└────────┬──────────────────────┬────────────────────────┬────────────┘
         │                      │                        │
         │ (REST API - TLS)     │ (REST API - TLS)       │ (REST API - TLS)
         │                      │                        │
         └──────────────────────┴────────────────────────┘
                                 │
                    ┌────────────▼────────────┐
                    │  Control Center        │
                    │  ┌────────────────────┐│
                    │  │  API Server (Go)   ││
                    │  │  Port: 8080        ││
                    │  │  ┌────────────────┐││
                    │  │  │ /api/v1 routes │││
                    │  │  │ - security/*   │││
                    │  │  │ - health/*     │││
                    │  │  │ - agents/*     │││
                    │  │  └────────────────┘││
                    │  └────────────────────┘│
                    │  ┌────────────────────┐│
                    │  │ Database (PostgreSQL)││
                    │  │ Port: 5432         ││
                    │  │ - agents           ││
                    │  │ - security_events  ││
                    │  │ - login_attempts   ││
                    │  │ - service_errors   ││
                    │  │ - system_metrics   ││
                    │  └────────────────────┘│
                    │  ┌────────────────────┐│
                    │  │ WebSocket Server   ││
                    │  │ (Real-time updates)││
                    │  └────────────────────┘│
                    └────────┬───────────────┘
                             │
                ┌────────────▼────────────┐
                │  Frontend (Vue.js)     │
                │  Port: 3000            │
                │  ┌────────────────────┐│
                │  │  Dashboard         ││
                │  │  - Events          ││
                │  │  - Agents          ││
                │  │  - Alerts          ││
                │  │  - Metrics         ││
                │  └────────────────────┘│
                └────────────────────────┘
```

## Component Details

### 1. Security Agent (On Each VPS)

**Technology**: Go 1.21+
**Binary Size**: ~15-20 MB

**Responsibilities**:
- Parse `/var/log/auth.log` for failed logins
- Monitor Nginx access/error logs
- Monitor MySQL error logs
- Collect system metrics (CPU, RAM, Disk, Network)
- Detect patterns:
  - Brute force attacks (multiple failed logins)
  - Port scans (connection attempts to random ports)
  - Service errors (spike in error rates)
- Send normalized events to Control Center API

**Communication**:
- HTTPS REST API to Control Center
- Periodic polling (configurable, default: 30s)
- Authentication via agent UUID + secret key

**Data Sent**:
```json
{
  "agent_uuid": "550e8400-e29b-41d4-a716-446655440000",
  "timestamp": "2024-03-04T21:30:45Z",
  "events": [
    {
      "type": "failed_login",
      "source_ip": "203.0.113.10",
      "username": "root",
      "service": "sshd",
      "severity": "high",
      "metadata": {}
    }
  ]
}
```

### 2. Backend API Server

**Technology**: Go 1.21 + Gin Web Framework
**Port**: 8080
**Database**: PostgreSQL 14+

**Modules**:

#### api/handlers
HTTP request handlers:
- `security.go` - Security event endpoints
- `health.go` - Service health checks
- `agents.go` - Agent management

#### security/
Business logic:
- `log_parser.go` - Parse log formats
- `event_processor.go` - Normalize events
- `threat_detector.go` - Pattern detection

#### database/
- Connection pooling
- Query builders
- Migration system

**API Endpoints**:
```
GET    /api/v1/security/events
GET    /api/v1/security/events/:id
GET    /api/v1/security/failed-logins
GET    /api/v1/security/attacks
GET    /api/v1/security/threats?ip=X.X.X.X
POST   /api/v1/security/events

GET    /api/v1/health/nginx
GET    /api/v1/health/mysql
GET    /api/v1/health/system

POST   /api/v1/agents/register
GET    /api/v1/agents
GET    /api/v1/agents/:id
DELETE /api/v1/agents/:id
```

### 3. Database Schema

```sql
-- Agents (VPS servers)
Table: agents
  - id (PK)
  - uuid (unique identifier for agent)
  - name (user-friendly name)
  - ip_address
  - status (active/inactive/error)
  - last_seen
  - created_at, updated_at

-- Events
Table: security_events
  - id (PK)
  - event_type (failed_login, brute_force, port_scan, etc.)
  - agent_id (FK -> agents)
  - source_ip
  - source_country
  - username
  - service (ssh, nginx, mysql, etc.)
  - description
  - severity (critical, high, medium, low)
  - acknowledged
  - created_at

Table: login_attempts
  - id (PK)
  - agent_id (FK -> agents)
  - login_type
  - username
  - source_ip
  - source_country
  - status (success/failed)
  - port
  - created_at

Table: service_errors
  - id (PK)
  - agent_id (FK -> agents)
  - service_name
  - error_type
  - error_message
  - error_count
  - last_occurrence
  - created_at

Table: system_metrics
  - id (PK)
  - agent_id (FK -> agents)
  - cpu_usage
  - memory_usage
  - disk_usage
  - network_in
  - network_out
  - created_at
```

### 4. Frontend Application

**Technology**: Vue 3 + TypeScript + Vite + Tailwind CSS
**Port**: 3000

**Pages**:
- **Dashboard** - Overview of all security events
- **Events** - Detailed event list with filtering
- **Agents** - Register and manage VPS agents

**Features**:
- Real-time event updates via WebSocket
- Event filtering and search
- Agent status monitoring
- Security metrics visualization
- Responsive design for mobile access

## Data Flow

### 1. Event Detection & Submission
```
VPS Agent
  ↓
  Reads: /var/log/auth.log
         /var/log/nginx/error.log
         /var/log/mysql/error.log
  ↓
  Normalizes & Aggregates
  ↓
  POST /api/v1/security/events
  ↓
Control Center API
```

### 2. Pattern Detection
```
New Event Arrives
  ↓
  Threat Detector:
    - Check IP reputation
    - Count failed logins from same IP (24h window)
    - Detect port scan patterns
    - Detect DDoS patterns
  ↓
  Assign Severity:
    - Single failed login → medium
    - 5+ from same IP in 1h → high
    - 50+ from same IP in 1h → critical (brute force)
  ↓
  Store in Database
  ↓
  Emit WebSocket Update
  ↓
Frontend (Real-time)
```

### 3. Frontend Display
```
Browser
  ↓
  Connect WebSocket: ws://localhost:8080/ws
  ↓
  Load Initial Data: GET /api/v1/security/events
  ↓
  Display Dashboard
  ↓
  Receive WebSocket Messages (new events)
  ↓
  Update Dashboard in Real-time
```

## Scalability Considerations

### Current Architecture
- Single backend instance
- PostgreSQL primary database
- Direct WebSocket connections

### For Production Scale

**Horizontal Scaling**:
```
Load Balancer (Nginx/HAProxy)
    ↓
  ┌─────────────────────────────┐
  │                             │
API 1 ─ Cache (Redis) ─ DB Replica 1
API 2                   DB Replica 2
API 3                   DB Primary (write)
  │                             │
  └─────────────────────────────┘
```

**Key Changes Needed**:
1. Add Redis for caching & session management
2. Use database replication for read scaling
3. Implement message queue (RabbitMQ/Kafka) for event processing
4. Separate WebSocket server from API server
5. Use Kubernetes for orchestration

## Security Considerations

### In Development
- No authentication (open API)
- HTTP only
- SQLite or local PostgreSQL

### For Production
1. **TLS/SSL** - All communications encrypted
2. **Authentication** - JWT tokens for API, mTLS for agents
3. **Rate Limiting** - Prevent abuse
4. **Input Validation** - Prevent SQL injection, XSS
5. **Secrets Management** - Use HashiCorp Vault or AWS Secrets Manager
6. **Logging & Auditing** - Track all access
7. **CORS** - Restrict to trusted origins
8. **Database Encryption** - Encrypt at rest
9. **Agent Verification** - Validate agent certificates
10. **Firewall** - Network isolation

## Deployment Options

### Docker Compose (Development)
```bash
docker-compose up -d
```

### Kubernetes (Production)
```yaml
# Helm chart structure
chart/
  templates/
    - api-deployment.yaml
    - frontend-deployment.yaml
    - database-statefulset.yaml
    - ingress.yaml
  values.yaml
```

### VPS/Bare Metal
```bash
# Systemd service for backend
/etc/systemd/system/vps-security-api.service

# Nginx reverse proxy
/etc/nginx/sites-available/vps-security

# SSL certificates
/etc/letsencrypt/live/vps-security.example.com/
```

## Performance Metrics

**Expected Performance**:
- Event processing: < 100ms
- API response time: < 200ms
- Dashboard load: < 500ms
- WebSocket update latency: < 50ms

**Database**:
- Can handle 1,000+ events/second
- Retention: Configurable (default: 90 days)
- Query optimization via indexes

## Monitoring & Observability

**Currently Missing** (Future Additions):
- Application metrics (Prometheus)
- Distributed tracing (Jaeger)
- Log aggregation (ELK Stack)
- APM (Application Performance Monitoring)

## Technology Stack Summary

| Layer | Technology | Version |
|-------|-----------|----------|
| **Agent** | Go | 1.21+ |
| **Backend** | Go + Gin | 1.21+ |
| **Database** | PostgreSQL | 14+ |
| **Frontend** | Vue.js | 3.3+ |
| **Build Tool** | Vite | 4.4+ |
| **Styling** | Tailwind CSS | 3.3+ |
| **Container** | Docker | 20.10+ |
| **Orchestration** | Docker Compose | 2.0+ |

---

For deployment details, see [INSTALLATION.md](./INSTALLATION.md)
For API details, see [docs/API.md](./docs/API.md)
