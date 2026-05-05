# Documentation Index - GoSecureTransfer

Welcome to GoSecureTransfer! This index helps you navigate all available documentation.

## 📚 Quick Navigation

### Getting Started
- **[GETTING_STARTED.md](GETTING_STARTED.md)** - Complete setup guide for both development and production
  - System requirements
  - Installation steps (backend & frontend)
  - Running locally
  - Troubleshooting

### Core Documentation

- **[README.md](README.md)** - Project overview and features
  - Quick feature overview
  - Project structure
  - API endpoints summary
  - Security highlights

- **[API.md](API.md)** - Complete API reference
  - All endpoints documented
  - Request/response examples
  - Error codes
  - Example curl commands

### Configuration Guides

- **[backend/ENV_SETUP.md](backend/ENV_SETUP.md)** - Backend environment configuration
  - Environment variables reference
  - How to generate secure keys
  - Production considerations

- **[frontend/ENV_SETUP.md](frontend/ENV_SETUP.md)** - Frontend environment configuration
  - Vite environment variables
  - Environment-specific builds
  - Configuration for different environments

### Deployment & Operations

- **[DEPLOYMENT.md](DEPLOYMENT.md)** - Production deployment guide
  - Backend deployment options (VPS, Docker, Cloud)
  - Frontend deployment (Static hosting, Nginx)
  - SSL/TLS setup with Let's Encrypt
  - Monitoring and logging
  - Backup and disaster recovery

### Design & UI

- **[DESIGN.md](DESIGN.md)** - Design system documentation
  - Color palette and typography
  - Component specifications
  - Layout guidelines
  - Accessibility standards
  - Responsive design

## 🗂️ File Structure

```
GoSecureTransfer/
├── README.md                 # Project overview
├── GETTING_STARTED.md       # Setup guide
├── DEPLOYMENT.md            # Production guide
├── API.md                   # API documentation
├── DESIGN.md                # UI/UX guide
├── LICENSE                  # MIT License
│
├── backend/
│   ├── ENV_SETUP.md        # Backend config guide
│   ├── .env.example        # Environment template
│   ├── .env                # Local config (git-ignored)
│   ├── go.mod              # Go dependencies
│   └── cmd/
│       └── server/
│           └── main.go     # Application entry point
│
└── frontend/
    ├── ENV_SETUP.md        # Frontend config guide
    ├── .env.example        # Environment template
    ├── .env                # Local config (git-ignored)
    ├── package.json        # NPM dependencies
    ├── vite.config.js      # Vite configuration
    └── src/
        ├── styles/         # CSS files
        └── components/     # React components
```

## 🚀 Common Tasks

### I want to...

#### Get started locally
1. Read [GETTING_STARTED.md](GETTING_STARTED.md)
2. Set up `.env` files from `.env.example`
3. Start PostgreSQL and create database
4. Run backend: `go run cmd/server/main.go`
5. Run frontend: `npm run dev`

#### Deploy to production
1. Read [DEPLOYMENT.md](DEPLOYMENT.md)
2. Choose deployment option (VPS, Docker, Cloud)
3. Generate secure secrets
4. Configure environment variables
5. Set up SSL/TLS certificates
6. Deploy backend and frontend
7. Set up monitoring and backups

#### Understand API endpoints
1. Read [API.md](API.md)
2. See endpoint specifications
3. Try curl examples
4. Implement client requests

#### Understand the design
1. Read [DESIGN.md](DESIGN.md)
2. Review color palette and typography
3. Check component specifications
4. Review responsive guidelines

#### Configure environment
- Backend: [backend/ENV_SETUP.md](backend/ENV_SETUP.md)
- Frontend: [frontend/ENV_SETUP.md](frontend/ENV_SETUP.md)

## 🔒 Security Topics

### General Security
- [DEPLOYMENT.md - Security Hardening](DEPLOYMENT.md#security-hardening)
- [GETTING_STARTED.md - Security Best Practices](GETTING_STARTED.md#security-best-practices)

### Secrets Management
- [backend/ENV_SETUP.md - Generating Secure Keys](backend/ENV_SETUP.md#generating-secure-keys)

### HTTPS/SSL Setup
- [DEPLOYMENT.md - SSL/TLS Certificates](DEPLOYMENT.md#ssltls-certificates)

## 🛠️ Technology Stack

### Backend
- **Language:** Go 1.26.1+
- **Database:** PostgreSQL 12+
- **Key Libraries:**
  - `github.com/golang-jwt/jwt/v5` - JWT
  - `github.com/lib/pq` - PostgreSQL
  - `golang.org/x/crypto` - Encryption
  - `github.com/joho/godotenv` - Environment loading

### Frontend
- **Framework:** React 19+
- **Build Tool:** Vite 8+
- **Language:** JavaScript (ES6+)
- **Styling:** CSS3

### DevOps
- **Containerization:** Docker (optional)
- **Web Server:** Nginx (recommended)
- **CI/CD:** GitHub Actions (recommended)

## 📊 API Overview

### Authentication
- `POST /register` - Create account
- `POST /login` - Get JWT token

### Files
- `POST /upload` - Upload encrypted file
- `GET /download?file=` - Download file
- `GET /files` - List user's files

Full documentation: [API.md](API.md)

## 🎨 UI Components

### Pages
- **Login** - User authentication
- **Register** - Account creation
- **Dashboard** - Main application
  - Upload Section
  - File List

Full specifications: [DESIGN.md](DESIGN.md)

## 🐛 Troubleshooting

### Backend Issues
- See [GETTING_STARTED.md - Troubleshooting](GETTING_STARTED.md#troubleshooting)
- See [backend/ENV_SETUP.md](backend/ENV_SETUP.md)

### Frontend Issues
- See [GETTING_STARTED.md - Troubleshooting](GETTING_STARTED.md#troubleshooting)
- Check browser console
- Verify `VITE_API_URL` is correct

### Deployment Issues
- See [DEPLOYMENT.md - Troubleshooting Production](DEPLOYMENT.md#troubleshooting-production)

## 📞 Support Resources

### Documentation
- Check the relevant `.md` file for your topic
- Search for keywords in documentation files

### Code Examples
- [API.md - Examples](API.md#examples)
- [GETTING_STARTED.md - Testing API Endpoints](GETTING_STARTED.md#testing-api-endpoints)

### External Resources
- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://react.dev/)
- [Vite Guide](https://vitejs.dev/guide/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

## 📝 Changelog

### Version 1.0 (May 2026)
- ✅ Initial release
- ✅ User authentication with JWT
- ✅ File encryption with AES-256
- ✅ Minimalist light theme UI
- ✅ Environment configuration system
- ✅ Comprehensive documentation

## 📄 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) file for details.

---

## 🎯 Next Steps

1. **New User?** Start with [GETTING_STARTED.md](GETTING_STARTED.md)
2. **Ready to Deploy?** Check [DEPLOYMENT.md](DEPLOYMENT.md)
3. **API Questions?** Reference [API.md](API.md)
4. **UI/Design?** See [DESIGN.md](DESIGN.md)

**Questions or Issues?** Open a GitHub issue or check the troubleshooting sections.

---

**Last Updated:** May 5, 2026  
**Current Version:** 1.0.0
