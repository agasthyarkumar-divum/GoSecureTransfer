# Deployment Guide - GoSecureTransfer

## Pre-Deployment Checklist

- [ ] Environment variables configured securely
- [ ] PostgreSQL database backup created
- [ ] SSL certificates obtained
- [ ] CORS origins configured properly
- [ ] Firewall rules configured
- [ ] Backend and frontend built for production
- [ ] Security headers configured on web server

## Backend Deployment

### Option 1: Traditional Server (VPS/Dedicated)

#### Prerequisites

- SSH access to server
- sudo privileges
- Go 1.26.1+ installed
- PostgreSQL 12+ installed
- systemd available

#### Steps

1. **Prepare the server:**
   ```bash
   # Update system
   sudo apt update && sudo apt upgrade -y
   
   # Install Go (if not already installed)
   wget https://go.dev/dl/go1.26.1.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.26.1.linux-amd64.tar.gz
   
   # Add to PATH
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

2. **Clone repository:**
   ```bash
   cd /opt
   sudo git clone <repository-url>
   cd GoSecureTransfer/backend
   ```

3. **Build the application:**
   ```bash
   CGO_ENABLED=1 go build -o server cmd/server/main.go
   ```

4. **Create systemd service:**
   ```bash
   sudo tee /etc/systemd/system/gosecuretransfer.service > /dev/null <<EOF
   [Unit]
   Description=GoSecureTransfer Backend
   After=network.target
   
   [Service]
   Type=simple
   User=gosecure
   WorkingDirectory=/opt/GoSecureTransfer/backend
   EnvironmentFile=/opt/GoSecureTransfer/backend/.env
   ExecStart=/opt/GoSecureTransfer/backend/server
   Restart=on-failure
   RestartSec=5
   StandardOutput=journal
   StandardError=journal
   
   [Install]
   WantedBy=multi-user.target
   EOF
   ```

5. **Start the service:**
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable gosecuretransfer
   sudo systemctl start gosecuretransfer
   sudo systemctl status gosecuretransfer
   ```

### Option 2: Docker Deployment

#### Dockerfile for Backend

```dockerfile
# Build stage
FROM golang:1.26.1-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o server cmd/server/main.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
```

#### docker-compose.yml

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: gosecure
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: securevault
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  backend:
    build:
      context: ./backend
    environment:
      DB_CONN: user=gosecure password=${DB_PASSWORD} dbname=securevault host=postgres port=5432 sslmode=disable
      JWT_SECRET: ${JWT_SECRET}
      ENCRYPTION_KEY: ${ENCRYPTION_KEY}
      PORT: 8080
      CORS_ORIGIN: ${CORS_ORIGIN}
    ports:
      - "8080:8080"
    depends_on:
      - postgres

volumes:
  postgres_data:
```

Deploy with:
```bash
docker-compose up -d
```

### Option 3: Cloud Platforms

#### Heroku

1. **Create Procfile:**
   ```
   web: cd backend && go run cmd/server/main.go
   ```

2. **Deploy:**
   ```bash
   heroku create gosecuretransfer
   git push heroku main
   ```

#### Railway / Render / Fly.io

Follow their respective guides for Go deployments.

## Frontend Deployment

### Option 1: Static Hosting (Recommended)

#### Build the application:
```bash
cd frontend
npm run build
```

#### Deploy to Vercel

```bash
npm install -g vercel
vercel --prod
```

#### Deploy to Netlify

```bash
npm run build
# Drag and drop the 'dist' folder to Netlify
```

#### Deploy to AWS S3 + CloudFront

```bash
aws s3 sync dist/ s3://your-bucket-name
aws cloudfront create-invalidation --distribution-id YOUR_ID --paths "/*"
```

### Option 2: Traditional Web Server

#### Nginx Configuration

```nginx
server {
    listen 443 ssl http2;
    server_name yourdomain.com;

    ssl_certificate /etc/ssl/certs/yourdomain.crt;
    ssl_certificate_key /etc/ssl/private/yourdomain.key;

    # Security headers
    add_header Strict-Transport-Security "max-age=31536000" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-Frame-Options "DENY" always;

    # Proxy to backend API
    location /api/ {
        proxy_pass http://localhost:8080/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Serve frontend
    root /var/www/gosecuretransfer;
    location / {
        try_files $uri $uri/ /index.html;
    }
}

# Redirect HTTP to HTTPS
server {
    listen 80;
    server_name yourdomain.com;
    return 301 https://$server_name$request_uri;
}
```

## Environment Variables for Production

Create a `.env` file with production values:

```bash
# Database
DB_CONN=user=gosecure password=<strong-password> dbname=securevault host=db.example.com port=5432 sslmode=require

# Secrets (generate with: openssl rand -base64 32)
JWT_SECRET=<generate-strong-secret>
ENCRYPTION_KEY=<generate-32-byte-key>

# Server
PORT=8080
ENV=production
CORS_ORIGIN=https://yourdomain.com
```

## SSL/TLS Certificates

### Using Let's Encrypt with Certbot

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot certonly --nginx -d yourdomain.com
```

### Renewal (automatic with systemd)

```bash
sudo systemctl enable certbot.timer
sudo systemctl start certbot.timer
```

## Monitoring & Logging

### Backend Logs

```bash
# View systemd logs
sudo journalctl -u gosecuretransfer -f

# Or with Docker
docker-compose logs -f backend
```

### Database Backups

```bash
# Create backup
pg_dump -U gosecure securevault > backup.sql

# Restore backup
psql -U gosecure securevault < backup.sql
```

### Prometheus Monitoring

Add metrics to Go server:
```bash
go get github.com/prometheus/client_golang/prometheus
```

## Performance Optimization

### Backend

1. Use connection pooling for database
2. Enable gzip compression
3. Add caching headers
4. Use CDN for static assets

### Frontend

1. Enable code splitting
2. Lazy load components
3. Minify and compress assets
4. Use service workers for offline support

## Security Hardening

### Network

- Use firewall (ufw, iptables)
- Restrict database access to backend only
- Use VPN for SSH access
- Enable rate limiting

### Application

- Rotate secrets regularly
- Keep dependencies updated
- Run security scans
- Implement CORS properly
- Use HTTPS only

### Database

- Enable SSL connections
- Use strong passwords
- Regular backups
- Monitor for unusual activity
- Restrict user privileges

## Troubleshooting Production

### High Memory Usage

```bash
# Check Go version for memory leaks
go version
```

### Database Connection Errors

```bash
# Test connection
psql -h db.example.com -U gosecure -d securevault
```

### SSL Certificate Issues

```bash
# Verify certificate
openssl s_client -connect yourdomain.com:443
```

## Disaster Recovery

1. **Automated Backups:**
   - Daily PostgreSQL backups
   - Off-site backup storage
   - Test restoration regularly

2. **Failover Strategy:**
   - Multi-region deployment
   - Database replication
   - Load balancing

3. **Incident Response:**
   - Monitor error rates
   - Set up alerts
   - Document procedures

---

**For support, create an issue on GitHub or contact the team.**
