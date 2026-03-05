# VPS Security Control Center

A comprehensive security monitoring dashboard for VPS servers. Monitor failed SSH login attempts, detect port scans, track service errors (Nginx, MySQL), and manage security events from a centralized web interface.

## Features

✅ **Failed Login Monitoring** - Real-time tracking of SSH authentication failures  
✅ **Attack Detection** - Pattern recognition for brute force and port scan attempts  
✅ **Service Health** - Monitor Nginx, MySQL, and system services  
✅ **Centralized Dashboard** - Web-based security events visualization  
✅ **Multi-Server Support** - Manage multiple VPS from one control center  
✅ **IP Geolocation** - Identify suspicious geographical patterns  
✅ **Alert System** - Real-time notifications for critical events  
✅ **Docker Ready** - Easy deployment with Docker Compose

## Architecture

```
┌─────────────────┐
│   VPS Server    │
│  (Agent)        │
│  - Auth logs    │
│  - Nginx logs   │
│  - MySQL logs   │
└────────┬────────┘
         │ (REST API / gRPC)
         ↓
┌─────────────────────────────────────┐
│   Backend (Go + Gin)                │
│  - Log parser                       │
│  - Event aggregator                 │
│  - Security analyzer                │
│  - SQLite/PostgreSQL storage        │
└────────┬────────────────────────────┘
         │ (GraphQL / REST)
         ↓
┌─────────────────────────────────────┐
│   Frontend (Vue.js)                 │
│  - Security events dashboard        │
│  - Attack patterns visualization    │
│  - Real-time alerts                 │
│  - System health metrics            │
└─────────────────────────────────────┘
```

## Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.21+
- Node.js 18+
- PostgreSQL or SQLite

### Local Development

```bash
# Clone repository
git clone https://github.com/awakeelectronik/vps-security-control.git
cd vps-security-control

# Checkout feature branch
git checkout feature/security-module

# Start with Docker Compose
docker-compose up -d

# Backend will be available at http://localhost:8080
# Frontend will be available at http://localhost:3000
```

### Manual Setup (Development)

**Backend:**
```bash
cd backend
go mod download
cp .env.example .env
go run cmd/api/main.go
```

**Frontend:**
```bash
cd frontend
npm install
npm run dev
```

## API Endpoints

### Security Events
- `GET /api/v1/security/events` - List all security events
- `GET /api/v1/security/events/:id` - Get event details
- `GET /api/v1/security/failed-logins` - Failed SSH attempts
- `GET /api/v1/security/attacks` - Detected attacks
- `GET /api/v1/security/threats?ip=<ip>` - Threats from specific IP

### Service Health
- `GET /api/v1/health/nginx` - Nginx service status
- `GET /api/v1/health/mysql` - MySQL service status
- `GET /api/v1/health/system` - System resources (CPU, RAM, Disk)

### Agent Management
- `POST /api/v1/agents/register` - Register new VPS agent
- `GET /api/v1/agents` - List registered agents
- `DELETE /api/v1/agents/:id` - Remove agent

## Agent Installation

On your VPS:

```bash
# Download and run the security agent
curl -O https://your-server.com/agent/vps-agent-linux-x64
chmod +x vps-agent-linux-x64

# Configure
sudo VPS_SECURITY_KEY=your_key \
      VPS_CONTROL_CENTER=https://your-server.com \
      VPS_AGENT_NAME=production-vps-1 \
      ./vps-agent-linux-x64

# Or use systemd
sudo systemctl start vps-security-agent
```

## Configuration

### Backend (.env)
```env
DATABASE_URL=postgresql://user:pass@localhost/vps_security
JWT_SECRET=your-secret-key
PORT=8080
LOG_LEVEL=info
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
VITE_WS_URL=ws://localhost:8080/ws
```

## Dashboard Features

1. **Security Events Timeline** - Real-time feed of login attempts, errors, alerts
2. **Attack Map** - Geographic visualization of attack sources
3. **Failed Logins** - Detailed list of SSH authentication failures with:
   - Source IP
   - Username attempts
   - Timestamp
   - Geolocation
   - Brute force pattern detection
4. **Service Status** - Health indicators for Nginx, MySQL, system
5. **Error Logs** - Aggregated errors from Nginx, MySQL, system logs
6. **Threat Score** - Overall security posture of the VPS

## Technology Stack

**Backend:**
- Go 1.21+
- Gin Web Framework
- PostgreSQL / SQLite
- JWT Authentication
- WebSocket for real-time updates

**Frontend:**
- Vue 3
- TypeScript
- Vite
- Tailwind CSS
- Chart.js for visualizations
- Socket.io for WebSocket

**Deployment:**
- Docker / Docker Compose
- Optional: Kubernetes
- Nginx reverse proxy

## Project Structure

```
vps-security-control/
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   │   └── main.go
│   │   └── agent/
│   │       └── main.go
│   ├── internal/
│   │   ├── api/
│   │   │   ├── handlers/
│   │   │   ├── middleware/
│   │   │   └── routes.go
│   │   ├── security/
│   │   │   ├── log_parser.go
│   │   │   ├── event_processor.go
│   │   │   └── threat_detector.go
│   │   ├── models/
│   │   ├── database/
│   │   └── config/
│   ├── go.mod
│   ├── go.sum
│   ├── Dockerfile
│   └── .env.example
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── services/
│   │   ├── App.vue
│   │   └── main.ts
│   ├── vite.config.ts
│   ├── package.json
│   └── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

## Development

### Building
```bash
make build
```

### Testing
```bash
make test
```

### Run with Hot Reload
```bash
make dev
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit changes (`git commit -am 'Add new feature'`)
4. Push to branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

MIT License - see LICENSE file

## Roadmap

- [ ] Slack/Telegram notifications integration
- [ ] Machine learning anomaly detection
- [ ] 2FA support for dashboard
- [ ] Advanced threat intelligence feeds
- [ ] Custom alert rules builder
- [ ] Kubernetes cluster monitoring
- [ ] Mobile app (React Native)

## Support

For issues, questions, or suggestions: [GitHub Issues](https://github.com/awakeelectronik/vps-security-control/issues)
