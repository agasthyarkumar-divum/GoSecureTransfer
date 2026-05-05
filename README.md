# GoSecureTransfer

A secure file transfer system built with **Go backend** and **React frontend**, featuring end-to-end AES-256 encryption, JWT authentication, and PostgreSQL storage.

## ✨ Features

- **🔐 End-to-End Encryption**: Files encrypted with AES-256 before storage
- **🔑 JWT Authentication**: Secure token-based authentication
- **📁 File Management**: Upload, download, and list files
- **⚡ RESTful API**: Built with Go's native `net/http`
- **⚙️ Environment Configuration**: Flexible `.env` file support
- **🎨 Minimalist UI**: Clean and lightweight React frontend
- **💾 PostgreSQL Backend**: Persistent data storage

## 📋 Prerequisites

- **Go** 1.26.1+
- **Node.js** 18+
- **PostgreSQL** 12+

## 🚀 Quick Start

### Backend Setup

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Copy environment configuration:**
   ```bash
   cp .env.example .env
   ```

3. **Configure PostgreSQL:**
   ```bash
   # Start PostgreSQL service
   sudo systemctl start postgresql

   # Create database
   sudo -u postgres psql -c "CREATE DATABASE securevault;"
   ```

4. **Install dependencies:**
   ```bash
   go mod download
   ```

5. **Run the server:**
   ```bash
   go run cmd/server/main.go
   ```

   Server will start on `http://localhost:8080`

### Frontend Setup

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Copy environment configuration:**
   ```bash
   cp .env.example .env
   ```

3. **Install dependencies:**
   ```bash
   npm install
   ```

4. **Run development server:**
   ```bash
   npm run dev
   ```

   App will be available at `http://localhost:5173`

## 📚 Project Structure

```
GoSecureTransfer/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # Application entry point
│   ├── internal/
│   │   ├── auth/                    # JWT authentication
│   │   ├── handlers/                # HTTP request handlers
│   │   ├── db/                      # Database initialization
│   │   ├── storage/                 # In-memory storage
│   │   └── crypto/                  # AES-256 encryption
│   ├── .env                         # Local configuration (git-ignored)
│   ├── .env.example                 # Configuration template
│   └── go.mod                       # Go module definition
├── frontend/
│   ├── src/
│   │   ├── components/              # React components
│   │   ├── utils/                   # Helper utilities
│   │   └── App.jsx                  # Main component
│   ├── .env                         # Local configuration (git-ignored)
│   ├── .env.example                 # Configuration template
│   └── package.json                 # NPM dependencies
└── README.md                        # This file
```

## 🔧 Configuration

### Backend Environment Variables

See [backend/ENV_SETUP.md](backend/ENV_SETUP.md) for detailed configuration options.

**Key variables:**
- `DB_CONN`: PostgreSQL connection string
- `JWT_SECRET`: Secret key for JWT signing (min 32 chars)
- `ENCRYPTION_KEY`: 32-byte key for AES-256 encryption
- `PORT`: Server port (default: 8080)
- `CORS_ORIGIN`: CORS allowed origin (default: *)

### Frontend Environment Variables

See [frontend/ENV_SETUP.md](frontend/ENV_SETUP.md) for detailed configuration options.

**Key variables:**
- `VITE_API_URL`: Backend API URL
- `VITE_API_PORT`: Backend API port

## 📡 API Endpoints

### Authentication

- `POST /login` - User login
  ```json
  {"username": "user", "password": "pass"}
  ```
  Response: `{"token": "jwt_token"}`

- `POST /register` - User registration
  ```json
  {"username": "user", "password": "pass"}
  ```

### File Operations

- `POST /upload` - Upload encrypted file
  - Headers: `Authorization: Bearer <token>`
  - Body: `FormData` with file

- `GET /download?file=filename` - Download and decrypt file
  - Headers: `Authorization: Bearer <token>`

- `GET /files` - List user's files
  - Headers: `Authorization: Bearer <token>`
  - Response: `["file1.txt", "file2.pdf", ...]`

## 🔐 Security

- **Encryption**: AES-256-GCM encryption for all files
- **Authentication**: JWT tokens with 24-hour expiration
- **Password Hashing**: bcrypt password hashing
- **Database**: PostgreSQL with user constraints
- **CORS**: Configurable cross-origin policies

## 📦 Dependencies

### Backend
- `github.com/golang-jwt/jwt/v5` - JWT token generation/validation
- `github.com/lib/pq` - PostgreSQL driver
- `golang.org/x/crypto` - Password hashing & AES encryption
- `github.com/joho/godotenv` - Environment variable loading

### Frontend
- `react` - UI library
- `vite` - Build tool

## 🛠️ Development

### Running Tests
```bash
cd backend
go test ./...
```

### Building for Production

**Backend:**
```bash
cd backend
go build -o server cmd/server/main.go
./server
```

**Frontend:**
```bash
cd frontend
npm run build
npm run preview
```

## 📝 Contributing

1. Clone the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## 📄 License

MIT License - See [LICENSE](LICENSE) file for details

## 🤝 Support

For issues or questions, please open an issue on GitHub.

---

**Built with ❤️ using Go and React**

