# 🔧 API Documentation

## 🎭 Navigation (Ваши Путеводные Звезды) 🌟

- **[🏠 Main README](../README.md)** - Project overview and quick start
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🏠 Local Development](README-LOCAL.md)** - Local PostgreSQL setup
- **[🚀 AWS Deployment](DEPLOYMENT.md)** - Production deployment guide

### 🌍 Language Versions
- **[🇷🇺 Русская Версия](API-RU.md)** - Русская версия этого руководства

## 🔐 Authentication


All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

## 📡 API Endpoints 

### 🔑 Authentication Endpoints

#### Register User
```http
POST /api/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "id": 1,
  "username": "testuser",
  "email": "test@example.com",
  "message": "Registration successful!"
}
```

#### Login User
```http
POST /api/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Logout User
```http
POST /api/auth/logout
Authorization: Bearer <token>
```

**Response:** `204 No Content`

#### Request Password Reset
```http
POST /api/auth/request-password-reset
Content-Type: application/json

{
  "email": "test@example.com"
}
```

**Response:**
```json
{
  "message": "Password reset email sent successfully"
}
```

#### Reset Password
```http
POST /api/auth/reset-password
Content-Type: application/json

{
  "token": "reset-token-here",
  "password": "newpassword123"
}
```

**Response:** `204 No Content`

### 📝 Personal Tasks

#### Get User Tasks
```http
GET /api/tasks
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "title": "Complete project documentation",
    "done": false,
    "user_id": 1,
    "deadline": "2024-12-31T23:59:59Z"
  }
]
```

#### Create Task
```http
POST /api/tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "New task",
  "deadline": "2024-12-31T23:59:59Z"
}
```

**Response:**
```json
{
  "id": 2,
  "title": "New task",
  "done": false,
  "user_id": 1,
  "deadline": "2024-12-31T23:59:59Z"
}
```

#### Update Task
```http
PATCH /api/tasks/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Updated task title",
  "done": true,
  "deadline": "2024-12-31T23:59:59Z"
}
```

**Response:**
```json
{
  "id": 1,
  "title": "Updated task title",
  "done": true,
  "user_id": 1,
  "deadline": "2024-12-31T23:59:59Z"
}
```

#### Delete Task
```http
DELETE /api/tasks/{id}
Authorization: Bearer <token>
```

**Response:** `204 No Content`

### 👥 Team Projects

#### Create Team Project
```http
POST /api/team-projects
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "My Team Project"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "My Team Project",
  "code": "123456",
  "created_by": 1,
  "created_at": "2024-01-01T00:00:00Z"
}
```

#### Join Team Project
```http
POST /api/team-projects/join
Authorization: Bearer <token>
Content-Type: application/json

{
  "code": "123456"
}
```

**Response:**
```json
{
  "message": "Successfully joined project",
  "project": {
    "id": 1,
    "name": "My Team Project",
    "code": "123456"
  }
}
```

#### Get User Team Projects
```http
GET /api/team-projects/my
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "My Team Project",
    "code": "123456",
    "created_by": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "joined_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Team Project Members
```http
GET /api/team-projects/{id}/members
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "username": "user1",
    "email": "user1@example.com",
    "joined_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Team Project Tasks
```http
GET /api/team-projects/{id}/tasks
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": 1,
    "title": "Team task",
    "done": false,
    "project_id": 1,
    "created_by": 1,
    "created_by_username": "user1",
    "deadline": "2024-12-31T23:59:59Z"
  }
]
```

#### Create Team Task
```http
POST /api/team-projects/{id}/tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "New team task",
  "deadline": "2024-12-31T23:59:59Z"
}
```

**Response:**
```json
{
  "id": 2,
  "title": "New team task",
  "done": false,
  "project_id": 1,
  "created_by": 1,
  "deadline": "2024-12-31T23:59:59Z"
}
```

#### Update Team Task
```http
PATCH /api/team-projects/{id}/tasks/{taskId}
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Updated team task",
  "done": true
}
```

**Response:**
```json
{
  "id": 1,
  "title": "Updated team task",
  "done": true,
  "project_id": 1,
  "created_by": 1,
  "deadline": "2024-12-31T23:59:59Z"
}
```

#### Delete Team Task
```http
DELETE /api/team-projects/{id}/tasks/{taskId}
Authorization: Bearer <token>
```

**Response:** `204 No Content`

## 🚨 Error Responses

### 400 Bad Request
```json
{
  "error": "bad json"
}
```

### 401 Unauthorized
```json
{
  "error": "Unauthorized"
}
```

### 404 Not Found
```json
{
  "error": "user not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "database error"
}
```

## 🔧 Testing with cURL

### Register and Login
```bash
# Register
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# Login
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'

# Use token for protected endpoints
TOKEN="your-jwt-token-here"
curl -X GET http://localhost:3000/api/tasks \
  -H "Authorization: Bearer $TOKEN"
```

## 📊 Response Codes

| Code | Description |
|------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created |
| 204 | No Content - Request successful, no response body |
| 400 | Bad Request - Invalid request data |
| 401 | Unauthorized - Authentication required |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

## 🔐 Security Notes

- All passwords are hashed using bcrypt
- JWT tokens expire after 24 hours
- Rate limiting is implemented for authentication endpoints
- Input validation is performed on all endpoints
- SQL injection protection is in place
