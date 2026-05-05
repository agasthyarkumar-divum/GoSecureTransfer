# Frontend Environment Setup

## Overview
The frontend uses Vite environment variables for configuration. These are prefixed with `VITE_` to be bundled at build time.

## Configuration Files

### `.env` (Local Development)
Contains your local API configuration. **Never commit this file** - it's in `.gitignore`.

### `.env.example`
Template file showing all available environment variables. Commit this to version control.

## Setup Instructions

1. **Copy the example file:**
   ```bash
   cp .env.example .env
   ```

2. **Update `.env` with your settings:**
   ```bash
   # Backend API URL (without port)
   VITE_API_URL=http://localhost
   
   # Backend API Port
   VITE_API_PORT=8080
   
   # Environment
   VITE_ENV=development
   ```

## Environment Variables Reference

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `VITE_API_URL` | Backend API base URL | `http://localhost` | `http://localhost` or `https://api.example.com` |
| `VITE_API_PORT` | Backend API port | `8080` | `8080` or `3000` |
| `VITE_ENV` | Environment mode | `development` | `development`, `staging`, `production` |

## Running the Application

### Development Server:
```bash
npm run dev
```
The server will use variables from `.env` file.

### Build for Production:
```bash
npm run build
```

### Environment-Specific Builds:
```bash
# For different environments, create `.env.production`
VITE_API_URL=https://api.example.com
VITE_API_PORT=443
VITE_ENV=production
npm run build
```

## API Configuration in Code

The frontend uses the `utils/api.js` helper to construct API URLs:

```javascript
import { apiCall, getApiUrl } from "../utils/api";

// getApiUrl() returns: `http://localhost:8080`
// apiCall() automatically prepends the API URL
```

## Environment-Specific Files

You can create environment-specific files:

- `.env` - Development (local)
- `.env.local` - Local overrides (in .gitignore)
- `.env.production` - Production builds
- `.env.staging` - Staging environment

Vite will automatically load the appropriate file based on the environment.

## Production Considerations

1. **Use HTTPS:** Always use `https://` in production
2. **Use proper domain:** Replace `http://localhost` with your domain
3. **Set correct port:** Use standard ports (80 for HTTP, 443 for HTTPS)
4. **Security headers:** Ensure your backend sets appropriate CORS headers
5. **Build optimization:** Use `npm run build` for production builds (creates optimized bundle)
