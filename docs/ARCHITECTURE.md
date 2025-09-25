# 🏗️ Architecture Documentation

## 🎭 Navigation (Ваши Путеводные Звезды) 🌟

- **[🏠 Main README](../README.md)** - Project overview and quick start
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🏠 Local Development](README-LOCAL.md)** - Local PostgreSQL setup
- **[🚀 AWS Deployment](DEPLOYMENT.md)** - Production deployment guide
- **[🔧 API Documentation](API.md)** - Complete API reference

### 🌍 Language Versions
- **[🇷🇺 Русская Версия](ARCHITECTURE-RU.md)** - Русская версия этого руководства

## 🎯 System Overview 

TodoApp is built using a clean, modular architecture that separates concerns and makes the codebase maintainable and scalable.


## 🏛️ Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                        Frontend Layer                       │
├─────────────────────────────────────────────────────────────┤
│  HTML/CSS/JavaScript (Vanilla)                             │
│  - Responsive Design                                        │
│  - Multi-language Support (i18n)                           │
│  - Real-time Updates                                        │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                      API Gateway Layer                      │
├─────────────────────────────────────────────────────────────┤
│  Gorilla Mux Router                                         │
│  - Route Registration                                       │
│  - Middleware Chain                                         │
│  - Static File Serving                                      │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Application Layer                        │
├─────────────────────────────────────────────────────────────┤
│  API Handlers                                               │
│  ├── auth.go (Authentication)                              │
│  ├── user.go (User Management)                             │
│  ├── tasks.go (Personal Tasks)                             │
│  └── team_projects.go (Team Collaboration)                 │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                      Business Logic Layer                   │
├─────────────────────────────────────────────────────────────┤
│  - JWT Authentication                                       │
│  - Password Hashing (bcrypt)                               │
│  - Input Validation                                         │
│  - Business Rules                                           │
│  - Email Services (SMTP)                                    │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                      Data Access Layer                      │
├─────────────────────────────────────────────────────────────┤
│  SQLC Generated Code                                        │
│  - Type-safe Database Queries                              │
│  - Connection Pooling                                       │
│  - Transaction Management                                   │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                      Database Layer                         │
├─────────────────────────────────────────────────────────────┤
│  PostgreSQL Database                                        │
│  - ACID Compliance                                          │
│  - Relational Data Model                                    │
│  - Indexes for Performance                                  │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
todoapp/
├── main.go                    # Application entry point
├── api/                       # API handlers
│   ├── auth.go               # Authentication & authorization
│   ├── routers.go            # Route registration
│   ├── tasks.go              # Personal task management
│   ├── team_projects.go      # Team collaboration
│   └── user.go               # User management
├── database/                  # Database layer
│   ├── database.go           # Connection management
│   ├── generated/            # SQLC generated code
│   │   ├── db.go            # Database connection
│   │   ├── models.go        # Data models
│   │   ├── tasks.sql.go     # Task queries
│   │   ├── users.sql.go     # User queries
│   │   └── team_projects.sql.go
│   ├── migrations/           # Database migrations
│   └── queries/              # SQL query definitions
├── i18n/                     # Internationalization
│   ├── i18n.go              # i18n configuration
│   └── errors.go            # Error translations
├── settings/                 # Configuration
│   └── settings.go          # Environment variables
├── static/                   # Frontend files
│   ├── index.html           # Main application
│   ├── script-simple.js     # Frontend logic
│   ├── i18n.js              # Frontend i18n
│   ├── verify-email.html    # Email verification
│   └── reset-password.html  # Password reset
└── docs/                     # Documentation
    ├── README.md            # Main documentation
    ├── README-RU.md         # Russian documentation
    ├── README-LOCAL.md      # Local development
    ├── DEPLOYMENT.md        # AWS deployment
    ├── API.md               # API reference
    └── ARCHITECTURE.md      # This file
```

## 🔄 Data Flow

### 1. Request Processing
```
HTTP Request → Router → Middleware → Handler → Business Logic → Database
```

### 2. Authentication Flow
```
Login Request → Validate Credentials → Generate JWT → Return Token
Protected Request → Validate JWT → Extract User ID → Process Request
```

### 3. Task Management Flow
```
Create Task → Validate Input → Save to Database → Return Task Data
Update Task → Check Ownership → Update Database → Return Updated Task
```

## 🗄️ Database Schema

### Core Tables

#### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    is_blocked BOOLEAN DEFAULT FALSE,
    failed_attempts INT DEFAULT 0,
    last_failed TIMESTAMP DEFAULT NOW(),
    email_verified BOOLEAN DEFAULT FALSE
);
```

#### Tasks Table
```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN DEFAULT false,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    deadline TIMESTAMP
);
```

#### Team Projects Table
```sql
CREATE TABLE team_projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL UNIQUE,
    created_by INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

#### Team Project Members Table
```sql
CREATE TABLE team_project_members (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES team_projects(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(project_id, user_id)
);
```

#### Team Tasks Table
```sql
CREATE TABLE team_tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT DEFAULT '',
    done BOOLEAN DEFAULT false,
    project_id INT NOT NULL REFERENCES team_projects(id) ON DELETE CASCADE,
    created_by INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    deadline TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### Security Tables

#### Password Reset Tokens
```sql
CREATE TABLE password_reset_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
```

#### Email Verification Tokens
```sql
CREATE TABLE email_verification_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
```

## 🔐 Security Architecture

### Authentication & Authorization
- **JWT Tokens**: Stateless authentication
- **Password Hashing**: bcrypt with salt
- **Rate Limiting**: Brute force protection
- **Input Validation**: SQL injection prevention

### Data Protection
- **Environment Variables**: Sensitive data configuration
- **HTTPS**: Encrypted communication (production)
- **CORS**: Cross-origin request handling
- **Input Sanitization**: XSS prevention

## 🚀 Deployment Architecture

### Local Development
```
Developer Machine
├── Go Application (go run)
├── Local PostgreSQL
└── SMTP (Gmail)
```

### Docker Development
```
Docker Compose
├── Application Container
├── PostgreSQL Container
├── Nginx Load Balancer
└── Volume Persistence
```

### AWS Production
```
AWS Cloud
├── ECS Fargate (Application)
├── RDS PostgreSQL (Database)
├── ECR (Container Registry)
├── CloudWatch (Logging)
├── Secrets Manager (Credentials)
└── Application Load Balancer
```

## 📊 Performance Considerations

### Database Optimization
- **Indexes**: On frequently queried columns
- **Connection Pooling**: Efficient database connections
- **Query Optimization**: SQLC generated queries

### Application Optimization
- **Gorilla Mux**: Efficient routing
- **Static File Serving**: Direct file serving
- **JWT Caching**: Token validation optimization

### Scalability
- **Horizontal Scaling**: Multiple application instances
- **Load Balancing**: Nginx/AWS ALB
- **Database Scaling**: Read replicas (future)

## 🔧 Technology Stack

### Backend
- **Go 1.23**: Programming language
- **Gorilla Mux**: HTTP router
- **SQLC**: Type-safe SQL code generation
- **PostgreSQL**: Primary database
- **JWT**: Authentication tokens

### Frontend
- **HTML5**: Markup
- **CSS3**: Styling
- **Vanilla JavaScript**: Client-side logic
- **Responsive Design**: Mobile-first approach

### Infrastructure
- **Docker**: Containerization
- **Docker Compose**: Local development
- **AWS ECS**: Container orchestration
- **AWS RDS**: Managed database
- **Nginx**: Load balancing

### Development Tools
- **Git**: Version control
- **SQLC**: Code generation
- **Postman**: API testing
- **Docker**: Development environment

## 🎯 Design Patterns

### Repository Pattern
- SQLC generated queries act as repositories
- Clean separation between data access and business logic

### Middleware Pattern
- Authentication middleware
- Request logging middleware
- Error handling middleware

### MVC Pattern
- **Model**: Database models (SQLC generated)
- **View**: HTML templates and static files
- **Controller**: API handlers

## 🔄 Development Workflow

1. **Database Changes**: Update SQL queries in `database/queries/`
2. **Code Generation**: Run `sqlc generate`
3. **API Development**: Add handlers in `api/`
4. **Frontend Updates**: Modify files in `static/`
5. **Testing**: Use Postman or curl
6. **Deployment**: Docker or AWS

