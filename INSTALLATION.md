# Installation Guide

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker & Docker Compose (easiest way to get started)
- OR
  - Go 1.21+
  - Node.js 18+
  - PostgreSQL 14+

## Quick Start with Docker

### 1. Clone the Repository

```bash
git clone https://github.com/awakeelectronik/vps-security-control.git
cd vps-security-control
git checkout feature/security-module
```

### 2. Configure Environment

```bash
cp .env.example .env
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
```

### 3. Start with Docker Compose

```bash
# Start all services
docker-compose up -d

# Or with logs
docker-compose up
```

### 4. Seed Database (Optional)

```bash
# The database automatically runs migrations on startup
# To seed sample data:
docker-compose exec backend go run cmd/seed/main.go
```

### 5. Access the Application

- **Frontend**: http://localhost:3000
- **API**: http://localhost:8080
- **API Health**: http://localhost:8080/health

## Manual Installation (Development)

### Backend Setup

```bash
cd backend

# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Run migrations
go run cmd/migrate/main.go

# Seed sample data (optional)
go run cmd/seed/main.go

# Start the server
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Create environment file
cp .env.example .env

# Start development server
npm run dev
```

The frontend will be available at `http://localhost:3000`

## Database Setup

### Using Docker (Recommended)

The database is automatically set up when you run `docker-compose up`.

### Manual PostgreSQL Setup

```bash
# Create database
sudo -u postgres createdb vps_security

# Create user
sudo -u postgres createuser -P vps_user

# Grant privileges
sudo -u postgres psql -d vps_security -c "GRANT ALL PRIVILEGES ON DATABASE vps_security TO vps_user;"

# Run migrations
cd backend
DATABASE_URL="postgresql://vps_user:password@localhost/vps_security" go run cmd/migrate/main.go
```

## Configuration

### Backend (.env)

```env
DATABASE_URL=postgresql://user:password@localhost:5432/vps_security
PORT=8080
HOST=0.0.0.0
ENVIRONMENT=development
JWT_SECRET=your-secret-key
LOG_LEVEL=debug
```

### Frontend (.env)

```env
VITE_API_URL=http://localhost:8080/api
VITE_WS_URL=ws://localhost:8080/ws
```

## Troubleshooting

### Database Connection Error

```bash
# Check if PostgreSQL is running
sudo systemctl status postgresql

# Or if using Docker
docker-compose ps

# Check logs
docker-compose logs postgres
```

### Port Already in Use

```bash
# Change ports in docker-compose.yml or .env
# Then restart
docker-compose down
docker-compose up -d
```

### API Connection Error from Frontend

```bash
# Check backend is running
curl http://localhost:8080/health

# Check frontend environment variables
echo $VITE_API_URL

# Restart frontend
npm run dev
```

## Production Deployment

### Update .env for Production

```env
ENVIRONMENT=production
JWT_SECRET=<very-strong-random-secret>
DATABASE_URL=postgresql://user:password@prod-db.example.com:5432/vps_security
LOG_LEVEL=error
```

### Build Production Images

```bash
docker-compose -f docker-compose.yml build

# Push to registry
docker tag vps-security-control_backend:latest your-registry/vps-security-backend:latest
docker push your-registry/vps-security-backend:latest
```

### Deploy with Docker

```bash
# Pull images and start
docker-compose pull
docker-compose up -d

# Check status
docker-compose ps
```

## Next Steps

1. Read [API Documentation](./docs/API.md)
2. Install agent on VPS servers
3. Configure webhooks for alerts
4. Set up monitoring dashboards

## Getting Help

- Check [GitHub Issues](https://github.com/awakeelectronik/vps-security-control/issues)
- Review [README.md](./README.md)
- Check application logs: `docker-compose logs -f backend`
