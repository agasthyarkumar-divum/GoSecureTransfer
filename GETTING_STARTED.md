# Getting Started - GoSecureTransfer

## System Requirements

- **OS**: Linux, macOS, or Windows (with WSL)
- **Go**: 1.26.1 or later
- **Node.js**: 18.0.0 or later
- **PostgreSQL**: 12 or later
- **npm**: 8.0.0 or later

## Installation Steps

### 1. Clone the Repository

```bash
git clone <repository-url>
cd GoSecureTransfer
```

### 2. Backend Setup

#### Install Go Dependencies

```bash
cd backend
go mod download
```

#### Configure PostgreSQL

```bash
# Start PostgreSQL service (Linux/macOS)
sudo systemctl start postgresql
# OR on macOS with Homebrew
brew services start postgresql

# Create database and user
sudo -u postgres psql <<EOF
CREATE DATABASE securevault;
CREATE USER gosecure WITH PASSWORD 'secure_password';
GRANT ALL PRIVILEGES ON DATABASE securevault TO gosecure;
\q
EOF
```

#### Configure Environment

```bash
# Copy example configuration
cp .env.example .env

# Edit .env with your PostgreSQL credentials
nano .env
```

**Update these values in `.env`:**
```bash
DB_CONN=user=gosecure password=secure_password dbname=securevault sslmode=disable
JWT_SECRET=your-32-character-secret-key-here
ENCRYPTION_KEY=your-32-character-encryption-key
PORT=8080
CORS_ORIGIN=http://localhost:5173
```

#### Start Backend Server

```bash
go run cmd/server/main.go
```

You should see:
```
2026/05/05 16:45:00 ⚠️ No .env file found, using system environment variables
2026/05/05 16:45:00 ✅ Connected to PostgreSQL
2026/05/05 16:45:00 ✅ Tables ready
2026/05/05 16:45:00 🚀 Server running on :8080
```

### 3. Frontend Setup

#### Install Node Dependencies

```bash
cd frontend
npm install
```

#### Configure Environment

```bash
# Copy example configuration
cp .env.example .env

# Default values should work for local development
# cat .env
```

**Default `.env` values for development:**
```bash
VITE_API_URL=http://localhost
VITE_API_PORT=8080
VITE_ENV=development
```

#### Start Development Server

```bash
npm run dev
```

You should see:
```
VITE v8.0.10 ready in 234 ms

➜ Local: http://localhost:5173/
```

## Usage

### 1. Create an Account

1. Open `http://localhost:5173` in your browser
2. Click "Create one now" on the login page
3. Enter a username and password (min 6 characters)
4. Click "Create Account"

### 2. Upload Files

1. Sign in with your credentials
2. Drag and drop files onto the upload area, or click to select
3. Files are automatically encrypted with AES-256
4. Your files are associated with your account

### 3. Download Files

1. In the "Your Files" section, you'll see all your uploaded files
2. Click the "Download" button to decrypt and download
3. Files are decrypted client-side before download

## Production Deployment

### Backend

1. **Build the binary:**
   ```bash
   cd backend
   CGO_ENABLED=1 go build -o server cmd/server/main.go
   ```

2. **Set production environment variables:**
   ```bash
   export DB_CONN="your-production-db-connection"
   export JWT_SECRET="long-random-secret-key-minimum-32-characters"
   export ENCRYPTION_KEY="32-character-encryption-key"
   export PORT="8080"
   export CORS_ORIGIN="https://yourdomain.com"
   ```

3. **Run the server:**
   ```bash
   ./server
   ```

### Frontend

1. **Build for production:**
   ```bash
   cd frontend
   npm run build
   ```

2. **Deploy the `dist/` directory:**
   - Deploy to Vercel, Netlify, or any static hosting service
   - Or serve with a web server (nginx, Apache, etc.)

3. **Update `.env.production`:**
   ```bash
   VITE_API_URL=https://api.yourdomain.com
   VITE_API_PORT=443
   VITE_ENV=production
   ```

## Security Best Practices

### Secrets Management

1. **Generate Strong Secrets:**
   ```bash
   # Generate JWT secret
   openssl rand -base64 32

   # Generate encryption key
   openssl rand -base64 32
   ```

2. **Use a Secrets Manager:**
   - AWS Secrets Manager
   - HashiCorp Vault
   - Azure Key Vault
   - Environment variables (for deployment)

### Database Security

1. Use strong PostgreSQL passwords
2. Enable SSL connections (`sslmode=require`)
3. Use a dedicated database user with minimal privileges
4. Regularly backup your database

### Frontend Security

1. Always use HTTPS in production
2. Set proper CORS headers on backend
3. Use Content Security Policy (CSP) headers
4. Implement rate limiting

## Troubleshooting

### Backend Issues

**"DB ping error: password authentication failed"**
- Check PostgreSQL credentials in `.env`
- Ensure PostgreSQL is running: `sudo systemctl status postgresql`
- Verify user exists: `sudo -u postgres psql -l`

**"ENCRYPTION_KEY must be exactly 32 bytes"**
- Generate a new key: `openssl rand -base64 32`
- Update `.env` file

### Frontend Issues

**"Connection refused / API not reachable"**
- Ensure backend is running on port 8080
- Check `VITE_API_URL` in `.env`
- Check browser console for CORS errors

**"Module not found" errors**
- Run `npm install` again
- Clear node_modules: `rm -rf node_modules && npm install`

## Development Tips

### Hot Reload

**Backend:** Use `air` for auto-reload:
```bash
go install github.com/cosmtrek/air@latest
air
```

**Frontend:** Vite has built-in hot reload - just save and refresh!

### Testing API Endpoints

```bash
# Login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"test123"}'

# Upload file
curl -X POST http://localhost:8080/upload \
  -H "Authorization: Bearer <token>" \
  -F "file=@path/to/file"

# List files
curl -X GET http://localhost:8080/files \
  -H "Authorization: Bearer <token>"
```

## Support & Contributing

For issues, questions, or contributions, please:
1. Check existing issues
2. Create a new issue with detailed description
3. Submit pull requests with improvements

---

**Happy secure file transferring! 🔐**
