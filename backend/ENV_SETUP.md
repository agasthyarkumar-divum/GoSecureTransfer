# Backend Environment Setup

## Overview
The backend uses environment variables for configuration. A `.env` file template is provided via `.env.example`.

## Configuration Files

### `.env` (Local Development)
Contains your actual credentials and settings. **Never commit this file** - it's in `.gitignore`.

### `.env.example`
Template file showing all available environment variables and their descriptions. Commit this to version control.

## Setup Instructions

1. **Copy the example file:**
   ```bash
   cp .env.example .env
   ```

2. **Update `.env` with your settings:**
   ```bash
   # Database Connection
   DB_CONN=user=postgres password=postgres dbname=securevault sslmode=disable
   
   # JWT Secret (use a strong random string in production)
   JWT_SECRET=your-secure-jwt-secret-here
   
   # Encryption Key (must be exactly 32 bytes for AES-256)
   ENCRYPTION_KEY=your-32-byte-encryption-key-here
   
   # Server Configuration
   PORT=8080
   NODE_ENV=development
   CORS_ORIGIN=*
   ```

## Generating Secure Keys

### Generate a Secure JWT Secret:
```bash
openssl rand -base64 32
```

### Generate a Secure Encryption Key (32 bytes):
```bash
openssl rand -base64 32
```

## Environment Variables Reference

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DB_CONN` | PostgreSQL connection string | `user=postgres password=postgres dbname=securevault sslmode=disable` | No |
| `JWT_SECRET` | Secret key for JWT signing | `supersecretkey` | No (but should be set in production) |
| `ENCRYPTION_KEY` | 32-byte key for AES-256 encryption | `12345678901234567890123456789012` | No (but should be set in production) |
| `PORT` | Server port | `8080` | No |
| `NODE_ENV` | Environment mode | `development` | No |
| `CORS_ORIGIN` | CORS allowed origin | `*` | No |

## Running the Application

### Load environment variables and run:
```bash
# Using shell environment variables
export $(cat .env | xargs)
go run cmd/server/main.go

# Or use a tool like `air` for development with auto-reload
air
```

## Production Considerations

1. **Use strong secrets:** Never use default values in production
2. **Use environment-specific configs:** Maintain separate `.env` files for different environments
3. **Rotate credentials:** Regularly update JWT_SECRET and ENCRYPTION_KEY
4. **Restrict CORS:** Set CORS_ORIGIN to specific domains instead of `*`
5. **Use a secrets manager:** Consider using AWS Secrets Manager, HashiCorp Vault, or similar in production
