# VPS Security Control - API Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication

Currently, the API is open for development. In production, add JWT to all requests:

```bash
curl -H "Authorization: Bearer YOUR_JWT_TOKEN" http://localhost:8080/api/v1/security/events
```

## Endpoints

### Security Events

#### Get All Security Events
```
GET /security/events?limit=50&offset=0
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "event_type": "failed_login",
      "agent_id": 1,
      "source_ip": "192.168.1.100",
      "source_country": "US",
      "username": "admin",
      "service": "ssh",
      "description": "Failed SSH login attempt",
      "severity": "high",
      "acknowledged": false,
      "created_at": "2024-03-04T20:30:45Z"
    }
  ],
  "count": 50,
  "limit": 50,
  "offset": 0
}
```

#### Get Event by ID
```
GET /security/events/:id
```

**Response:**
```json
{
  "id": 1,
  "event_type": "failed_login",
  "agent_id": 1,
  "source_ip": "192.168.1.100",
  "source_country": "US",
  "username": "admin",
  "service": "ssh",
  "description": "Failed SSH login attempt",
  "severity": "high",
  "acknowledged": false,
  "created_at": "2024-03-04T20:30:45Z"
}
```

#### Get Failed Logins
```
GET /security/failed-logins?limit=100
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "agent_id": 1,
      "login_type": "ssh",
      "username": "root",
      "source_ip": "203.0.113.10",
      "source_country": "CN",
      "status": "failed",
      "port": 22,
      "created_at": "2024-03-04T20:30:45Z"
    }
  ],
  "count": 100
}
```

#### Get Detected Attacks
```
GET /security/attacks
```

**Response:**
```json
{
  "data": [
    {
      "id": 5,
      "event_type": "brute_force",
      "agent_id": 1,
      "source_ip": "192.168.1.150",
      "source_country": "RU",
      "username": "admin",
      "service": "ssh",
      "description": "Brute force attack detected from 192.168.1.150",
      "severity": "critical",
      "acknowledged": false,
      "created_at": "2024-03-04T21:15:30Z"
    }
  ],
  "count": 1
}
```

#### Get Threats from Specific IP
```
GET /security/threats?ip=192.168.1.100
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "event_type": "failed_login",
      "agent_id": 1,
      "source_ip": "192.168.1.100",
      "source_country": "US",
      "username": "admin",
      "service": "ssh",
      "description": "Failed SSH login attempt",
      "severity": "high",
      "acknowledged": false,
      "created_at": "2024-03-04T20:30:45Z"
    }
  ],
  "count": 1
}
```

#### Create Security Event
```
POST /security/events
```

**Request:**
```json
{
  "event_type": "failed_login",
  "agent_id": 1,
  "source_ip": "192.168.1.200",
  "source_country": "GB",
  "username": "user",
  "service": "ssh",
  "description": "Failed login attempt",
  "severity": "medium"
}
```

**Response:**
```json
{
  "id": 100
}
```

### Service Health

#### Get Nginx Health
```
GET /health/nginx
```

**Response:**
```json
{
  "name": "Nginx",
  "status": "running",
  "port": 80,
  "errors": 0
}
```

#### Get MySQL Health
```
GET /health/mysql
```

**Response:**
```json
{
  "name": "MySQL",
  "status": "running",
  "port": 3306,
  "errors": 0
}
```

#### Get System Health
```
GET /health/system
```

**Response:**
```json
{
  "cpu_usage": 35.2,
  "memory_usage": 62.8,
  "disk_usage": 45.1,
  "uptime_hours": 168
}
```

### Agent Management

#### Register New Agent
```
POST /agents/register
```

**Request:**
```json
{
  "name": "production-vps-1",
  "ip_address": "192.168.1.100"
}
```

**Response:**
```json
{
  "id": 1,
  "uuid": "550e8400-e29b-41d4-a716-446655440000"
}
```

#### List All Agents
```
GET /agents
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "uuid": "550e8400-e29b-41d4-a716-446655440000",
      "name": "production-vps-1",
      "ip_address": "192.168.1.100",
      "status": "active",
      "last_seen": "2024-03-04T21:30:45Z",
      "created_at": "2024-03-01T10:00:00Z",
      "updated_at": "2024-03-04T21:30:45Z"
    }
  ],
  "count": 1
}
```

#### Get Agent by ID
```
GET /agents/:id
```

**Response:**
```json
{
  "id": 1,
  "uuid": "550e8400-e29b-41d4-a716-446655440000",
  "name": "production-vps-1",
  "ip_address": "192.168.1.100",
  "status": "active",
  "last_seen": "2024-03-04T21:30:45Z",
  "created_at": "2024-03-01T10:00:00Z",
  "updated_at": "2024-03-04T21:30:45Z"
}
```

#### Delete Agent
```
DELETE /agents/:id
```

**Response:**
```json
{
  "message": "agent deleted"
}
```

## Error Responses

All errors follow this format:

```json
{
  "error": "error message describing what went wrong"
}
```

### HTTP Status Codes

- `200 OK` - Request successful
- `201 Created` - Resource created
- `400 Bad Request` - Invalid request parameters
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## Rate Limiting

Currently no rate limiting. This will be added in production.

## Pagination

For list endpoints, use `limit` and `offset` parameters:

```
GET /security/events?limit=50&offset=100
```

- `limit`: Max results (default: 50, max: 1000)
- `offset`: Results to skip (default: 0)

## Examples

### Example 1: Get all failed logins in the last 24 hours

```bash
curl "http://localhost:8080/api/v1/security/failed-logins?limit=100"
```

### Example 2: Register a new VPS

```bash
curl -X POST http://localhost:8080/api/v1/agents/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "staging-vps-1",
    "ip_address": "10.0.0.50"
  }'
```

### Example 3: Get all attacks detected

```bash
curl http://localhost:8080/api/v1/security/attacks
```

### Example 4: Get threats from a specific IP

```bash
curl "http://localhost:8080/api/v1/security/threats?ip=203.0.113.10"
```

## Webhooks (Coming Soon)

- POST /webhooks/security-events
- POST /webhooks/alerts
- POST /webhooks/agent-status

---

For more information, check the [README.md](../README.md) or [INSTALLATION.md](../INSTALLATION.md)
