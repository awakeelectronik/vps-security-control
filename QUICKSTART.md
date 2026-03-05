# VPS Security Control - Quick Start

## 🚀 Start in 5 Minutes

### Requirements
- Docker & Docker Compose (that's it!)

### Step 1: Clone & Enter Directory
```bash
git clone https://github.com/awakeelectronik/vps-security-control.git
cd vps-security-control
git checkout feature/security-module
```

### Step 2: Start the Stack
```bash
docker-compose up -d
```

### Step 3: Wait for Services
```bash
# Check status
docker-compose ps

# Should see:
# - vps-security-db     (postgres)
# - vps-security-api    (backend)
# - vps-security-web    (frontend)
```

### Step 4: Access the Dashboard

**Frontend**: http://localhost:3000

![Dashboard Screenshot]

## 📊 What You See

### Dashboard Tab
- Failed login attempts (real-time)
- Detected attacks
- Active agents
- System health score
- Recent security events timeline

### Events Tab
- Detailed list of all security events
- Filter and search capabilities
- Event severity levels (Critical, High, Medium, Low)
- Source IP geolocation
- Timestamps

### Agents Tab
- Register new VPS agents
- View agent status (Active/Offline)
- Agent IP addresses
- Agent UUIDs

## 🔌 Register Your First VPS Agent

### In the Dashboard (Agents Tab)

1. Click **Agents** in navigation
2. Enter Agent Name: `production-vps-1`
3. Enter IP Address: `192.168.1.100`
4. Click **Register**

You'll get a UUID to install the agent on your VPS.

## 🧪 Test with Sample Data

```bash
# Seed the database with test data
docker-compose exec backend go run cmd/seed/main.go
```

This will populate:
- 20 sample security events
- 30 login attempts (some failed)
- 3 service error logs
- System metrics

Refresh your dashboard to see the data!

## 🔗 API Endpoints

### Test the API
```bash
# Get all security events
curl http://localhost:8080/api/v1/security/events

# Get failed logins
curl http://localhost:8080/api/v1/security/failed-logins

# Get detected attacks
curl http://localhost:8080/api/v1/security/attacks

# List agents
curl http://localhost:8080/api/v1/agents

# System health
curl http://localhost:8080/api/v1/health/system
```

## 📝 Check Logs

```bash
# Backend logs
docker-compose logs -f backend

# Frontend logs
docker-compose logs -f frontend

# Database logs
docker-compose logs -f postgres

# All logs
docker-compose logs -f
```

## 🛑 Stop Everything

```bash
# Stop services
docker-compose down

# Stop and remove data (WARNING: deletes database)
docker-compose down -v
```

## 🐛 Troubleshooting

### Port 3000 or 8080 already in use?

Edit `docker-compose.yml`:
```yaml
frontend:
  ports:
    - "3001:3000"  # Change 3000 to 3001

backend:
  ports:
    - "8081:8080"  # Change 8080 to 8081
```

Then: `docker-compose down && docker-compose up -d`

### Database connection error?

```bash
# Check if postgres is healthy
docker-compose ps

# If not healthy, check logs
docker-compose logs postgres

# Restart postgres
docker-compose restart postgres
```

### Frontend shows blank page?

```bash
# Clear browser cache
# Hard refresh: Ctrl+Shift+R (or Cmd+Shift+R on Mac)

# Or restart frontend
docker-compose restart frontend
```

## 📚 Next Steps

1. **Install Agent on VPS**: (Agent source coming soon)
   ```bash
   # On your VPS:
   curl -O https://your-control-center/agent/vps-agent
   chmod +x vps-agent
   ./vps-agent --register https://your-control-center
   ```

2. **Configure Real Notifications**:
   - Enable Slack webhooks
   - Set up email alerts
   - Configure alerting rules

3. **Set Up SSL/TLS**:
   - Generate certificates
   - Configure nginx reverse proxy
   - Enable HTTPS

4. **Deploy to Production**:
   - Follow [INSTALLATION.md](./INSTALLATION.md)
   - Set strong JWT secret
   - Configure database backups

## 💡 Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│  VPS Servers (Agent)                                   │
│  - Parse /var/log/auth.log                             │
│  - Monitor Nginx errors                                │
│  - Monitor MySQL errors                                │
│  - Collect system metrics                              │
│  - Send events to Control Center                       │
└────────────┬────────────────────────────────────────────┘
             │ (REST API)
             ▼
┌─────────────────────────────────────────────────────────┐
│  Backend (Go + Gin) - localhost:8080                    │
│  - Process security events                             │
│  - Detect patterns (brute force, port scans)           │
│  - Store in PostgreSQL                                 │
│  - Provide REST API                                    │
└────────────┬────────────────────────────────────────────┘
             │ (GraphQL/REST)
             ▼
┌─────────────────────────────────────────────────────────┐
│  Frontend (Vue.js) - localhost:3000                     │
│  - Real-time dashboard                                 │
│  - Event visualization                                 │
│  - Agent management                                    │
│  - Threat alerts                                       │
└─────────────────────────────────────────────────────────┘
```

## 🤝 Contributing

Want to contribute?

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit: `git commit -m 'Add amazing feature'`
4. Push: `git push origin feature/amazing-feature`
5. Open a Pull Request

## 📄 License

MIT License - See LICENSE file

## 📞 Support

- GitHub Issues: [Open an issue](https://github.com/awakeelectronik/vps-security-control/issues)
- Documentation: Check [README.md](./README.md)
- Install Guide: See [INSTALLATION.md](./INSTALLATION.md)

---

**Happy monitoring!** 🛡️
