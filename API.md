# API Documentation - GoSecureTransfer

## Base URL

```
http://localhost:8080
```

## Authentication

All protected endpoints require a JWT token in the `Authorization` header:

```
Authorization: Bearer <jwt_token>
```

## Endpoints

### Authentication Endpoints

#### Register User

Create a new user account.

- **URL:** `/register`
- **Method:** `POST`
- **Auth Required:** No
- **Content-Type:** `application/json`

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "secure_password_123"
}
```

**Success Response (200):**
```
User registered successfully
```

**Error Response (400):**
```
Username already exists
```

---

#### Login User

Authenticate a user and receive a JWT token.

- **URL:** `/login`
- **Method:** `POST`
- **Auth Required:** No
- **Content-Type:** `application/json`

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "secure_password_123"
}
```

**Success Response (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Error Response (401):**
```
Invalid password
```

---

### File Endpoints

#### Upload File

Upload and encrypt a file.

- **URL:** `/upload`
- **Method:** `POST`
- **Auth Required:** Yes
- **Content-Type:** `multipart/form-data`

**Request:**
```
Headers:
  Authorization: Bearer <token>

Body:
  Form Data with file field containing the file to upload
```

**Success Response (200):**
```
File uploaded and encrypted successfully
```

**Error Response (401):**
```
Unauthorized
```

**Error Response (400):**
```
Error reading file
```

---

#### Download File

Download and decrypt a file.

- **URL:** `/download?file=<filename>`
- **Method:** `GET`
- **Auth Required:** Yes
- **Query Parameters:**
  - `file` (required): Name of the file to download

**Request:**
```
Headers:
  Authorization: Bearer <token>
```

**Success Response (200):**
```
Binary file content (decrypted)
```

**Headers:**
```
Content-Disposition: attachment; filename=<filename>
```

**Error Response (401):**
```
Unauthorized
```

**Error Response (403):**
```
Forbidden
```

**Error Response (404):**
```
File not found
```

---

#### List Files

Get a list of all files uploaded by the authenticated user.

- **URL:** `/files`
- **Method:** `GET`
- **Auth Required:** Yes

**Request:**
```
Headers:
  Authorization: Bearer <token>
```

**Success Response (200):**
```json
[
  "document.pdf",
  "image.jpg",
  "archive.zip"
]
```

**Empty Response (200):**
```json
[]
```

**Error Response (401):**
```
Unauthorized
```

---

## Error Codes

| Code | Message | Description |
|------|---------|-------------|
| 200 | OK | Successful request |
| 400 | Bad Request | Invalid request data |
| 401 | Unauthorized | Authentication failed or missing |
| 403 | Forbidden | Access denied to resource |
| 404 | Not Found | Resource not found |
| 500 | Internal Server Error | Server error |

---

## Examples

### Example: Complete Flow

```bash
# 1. Register
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'

# 2. Login
TOKEN=$(curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }' | jq -r '.token')

echo "Token: $TOKEN"

# 3. Upload file
curl -X POST http://localhost:8080/upload \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@/path/to/file.pdf"

# 4. List files
curl -X GET http://localhost:8080/files \
  -H "Authorization: Bearer $TOKEN"

# 5. Download file
curl -X GET "http://localhost:8080/download?file=file.pdf" \
  -H "Authorization: Bearer $TOKEN" \
  -o downloaded_file.pdf
```

---

## Rate Limiting

Currently, no built-in rate limiting. Recommended to add:

```go
import "github.com/didip/tollbooth"

limiter := tollbooth.NewLimiter(1, nil)
http.Handle("/login", tollbooth.LimitHandler(limiter, LoginHandler))
```

---

## CORS Configuration

CORS is configured via the `CORS_ORIGIN` environment variable:

```
CORS_ORIGIN=http://localhost:5173
```

All endpoints return these headers when configured:
```
Access-Control-Allow-Origin: <CORS_ORIGIN>
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization
```

---

## Security Notes

1. **JWT Expiration:** Tokens expire after 24 hours
2. **Password Hashing:** Passwords are hashed using bcrypt
3. **File Encryption:** Files are encrypted with AES-256-GCM
4. **HTTPS:** Always use HTTPS in production
5. **Token Storage:** Store tokens in secure, httpOnly cookies when possible

---

## Pagination (Future)

Currently not implemented. Recommendation:

```
GET /files?page=1&limit=20
```

Response:
```json
{
  "files": ["file1.pdf", "file2.pdf"],
  "total": 50,
  "page": 1,
  "limit": 20
}
```

---

## Versioning

API version: `v1` (implicit, no version prefix in URLs)

Future versions might use:
```
/api/v1/login
/api/v2/login
```

---

**For issues or feature requests, please open a GitHub issue.**
