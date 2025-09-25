# TodoApp - Complete Task Management System 


## 🎭 Navigation (Your Guiding Stars) 🌟

- **[🏠 Main README](../README.md)** - Project overview and quick start
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🏠 Local Development](README-LOCAL.md)** - Local PostgreSQL setup
- **[🚀 AWS Deployment](DEPLOYMENT.md)** - Production deployment guide
- **[🔧 API Documentation](#api-endpoints)** - Complete API reference
- **[🏗️ Architecture](#technology-stack)** - Technical details

## Overview 

TodoApp is a comprehensive web application built with Go that provides task management capabilities for both individual users and team collaboration. The application features user authentication, email verification, password reset functionality, personal task management, and team project collaboration with real-time task sharing.

<!-- ![Anime Hero](https://via.placeholder.com/800x200/FF6B9D/FFFFFF?text=TodoApp+Anime+Style) -->
<!-- Place your anime GIF here -->

## Key Features

- **User Authentication & Security**
  - JWT-based authentication
  - User registration with email verification
  - Password reset via email
  - Brute force protection
  - Secure password hashing with bcrypt

- **Task Management**
  - Create, read, update, and delete personal tasks
  - Set task deadlines with visual indicators
  - Mark tasks as completed
  - Responsive design for mobile and desktop

- **Team Collaboration**
  - Create team projects with unique 6-digit codes
  - Join existing projects using project codes
  - Shared task management within projects
  - Real-time collaboration features

- **Internationalization**
  - Multi-language support (Russian, English, Mongolian)
  - Dynamic language switching
  - Localized error messages

- **Deployment Ready**
  - Docker containerization
  - AWS ECS deployment configuration
  - Load balancing with Nginx
  - PostgreSQL database with migrations

## Technology Stack

### Backend
- **Language**: Go 1.23
- **Web Framework**: Gorilla Mux
- **Database**: PostgreSQL
- **ORM**: SQLC for type-safe database queries
- **Authentication**: JWT tokens
- **Email**: SMTP integration (Gmail)

### Frontend
- **HTML/CSS/JavaScript**: Vanilla implementation
- **Responsive Design**: Mobile-first approach
- **Internationalization**: Custom i18n system

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Cloud Platform**: AWS (ECS, RDS, ECR)
- **Load Balancer**: Nginx
- **Database**: PostgreSQL with connection pooling

## Project Structure

```
todoapp/
├── api/                          # API handlers and routes
│   ├── auth.go                   # Authentication middleware and password reset
│   ├── routers.go                # Route registration
│   ├── tasks.go                  # Personal task management
│   ├── team_projects.go          # Team project functionality
│   └── user.go                   # User management
├── database/                     # Database layer
│   ├── generated/                # SQLC generated code
│   │   ├── db.go                 # Database connection
│   │   ├── models.go             # Data models
│   │   ├── tasks.sql.go          # Task queries
│   │   ├── users.sql.go          # User queries
│   │   ├── team_projects.sql.go  # Team project queries
│   │   ├── email_verification.sql.go
│   │   └── password_resets.sql.go
│   ├── migrations/               # Database migrations
│   │   ├── 0001_create_tasks.up.sql
│   │   ├── 0002_create_users.up.sql
│   │   ├── 0003_user_id_tasks.up.sql
│   │   ├── 0004_security_users.up.sql
│   │   ├── 0005_email_password.up.sql
│   │   ├── 0005_user_blocker.up.sql
│   │   ├── 0006_email_verification.up.sql
│   │   ├── 0007_create_team_projects.up.sql
│   │   └── 0008_add_deadline_to_tasks.up.sql
│   ├── queries/                  # SQL query definitions
│   │   ├── tasks.sql
│   │   ├── users.sql
│   │   ├── team_projects.sql
│   │   ├── email_verification.sql
│   │   └── password_resets.sql
│   ├── database.go               # Database initialization
│   └── schema.sql                # Database schema
├── i18n/                         # Internationalization
│   ├── i18n.go                   # i18n configuration
│   └── errors.go                 # Error message translations
├── settings/                     # Application settings
│   └── settings.go               # Environment variable handling
├── static/                       # Frontend files
│   ├── index.html                # Main application interface
│   ├── script-simple.js          # Frontend JavaScript
│   ├── i18n.js                   # Frontend internationalization
│   ├── verify-email.html         # Email verification page
│   └── reset-password.html       # Password reset page
├── docker-compose.yml            # Local development setup
├── Dockerfile                    # Application containerization
├── ecs-task-definition.json      # AWS ECS configuration
├── ecs-trust-policy.json         # IAM trust policy
├── setup_online_db.sql           # Database initialization script
├── sqlc.yaml                     # SQLC configuration
├── go.mod                        # Go module dependencies
└── main.go                       # Application entry point
```

## Database Schema

### Users Table
- `id`: Primary key (SERIAL)
- `username`: Unique username (TEXT)
- `email`: Unique email address (TEXT)
- `password_hash`: Bcrypt hashed password (TEXT)
- `is_blocked`: Account lock status (BOOLEAN)
- `failed_attempts`: Login failure counter (INT)
- `last_failed`: Last failed login timestamp (TIMESTAMP)
- `email_verified`: Email verification status (BOOLEAN)

### Tasks Table
- `id`: Primary key (SERIAL)
- `title`: Task description (TEXT)
- `done`: Completion status (BOOLEAN)
- `user_id`: Foreign key to users (INT)
- `deadline`: Optional task deadline (TIMESTAMP)

### Team Projects Table
- `id`: Primary key (SERIAL)
- `name`: Project name (TEXT)
- `code`: Unique 6-digit project code (TEXT)
- `created_by`: Project creator user ID (INT)
- `created_at`: Creation timestamp (TIMESTAMP)

### Team Project Members Table
- `id`: Primary key (SERIAL)
- `project_id`: Foreign key to team_projects (INT)
- `user_id`: Foreign key to users (INT)
- `joined_at`: Join timestamp (TIMESTAMP)

### Team Tasks Table
- `id`: Primary key (SERIAL)
- `title`: Task description (TEXT)
- `done`: Completion status (BOOLEAN)
- `project_id`: Foreign key to team_projects (INT)
- `created_by`: Task creator user ID (INT)
- `deadline`: Optional task deadline (TIMESTAMP)

### Security Tables
- `password_reset_tokens`: Password reset functionality
- `email_verification_tokens`: Email verification system

## API Endpoints

### Authentication
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - User login
- `POST /api/auth/logout` - User logout
- `POST /api/auth/request-password-reset` - Request password reset
- `POST /api/auth/reset-password` - Reset password with token

### Personal Tasks
- `GET /api/tasks` - Get user's tasks
- `POST /api/tasks` - Create new task
- `PATCH /api/tasks/{id}` - Update task
- `DELETE /api/tasks/{id}` - Delete task

### Team Projects
- `POST /api/team-projects` - Create team project
- `POST /api/team-projects/join` - Join existing project
- `GET /api/team-projects/my` - Get user's team projects
- `GET /api/team-projects/{id}/members` - Get project members
- `GET /api/team-projects/{id}/tasks` - Get project tasks
- `POST /api/team-projects/{id}/tasks` - Create project task
- `PATCH /api/team-projects/{id}/tasks/{taskId}` - Update project task
- `DELETE /api/team-projects/{id}/tasks/{taskId}` - Delete project task

## Quick Start

### Prerequisites
- Docker and Docker Compose
- Git

### Local Development Setup

1. **Clone the repository:**
```bash
git clone <repository-url>
cd todoapp
```

2. **Create environment file:**
```bash
cp env.example .env
# Edit .env with your configuration
```

3. **Start the application:**
```bash
docker-compose up -d
```

4. **Access the application:**
```
http://localhost:3000
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Application port | `7540` |
| `APP_BASE_URL` | Base URL for email links | `http://localhost:7540` |
| `PG_HOST` | PostgreSQL host | `localhost` |
| `PG_PORT` | PostgreSQL port | `5432` |
| `PG_USER` | Database user | `postgres` |
| `PG_PASS` | Database password | - |
| `PG_DB` | Database name | `todoapp` |
| `PG_SSLMODE` | SSL mode | `disable` |
| `JWT_SECRET` | JWT signing secret | - |
| `ADMIN_TOKEN` | Admin access token | - |
| `SMTP_HOST` | SMTP server | `smtp.gmail.com` |
| `SMTP_PORT` | SMTP port | `587` |
| `SMTP_USER` | SMTP username | - |
| `SMTP_PASS` | SMTP password | - |
| `SMTP_FROM` | Email sender address | - |
| `LOGIN_MAX_FAIL` | Max login failures before lockout | `4` |

### SMTP Configuration (Gmail)

1. Enable 2-factor authentication on your Gmail account
2. Generate an App Password
3. Use the App Password in `SMTP_PASS` environment variable

## AWS Deployment

### Prerequisites
- AWS CLI configured
- Docker installed
- AWS account with appropriate permissions

### Step-by-Step Deployment

1. **Create ECR Repository:**
```bash
aws ecr create-repository --repository-name todoapp --region us-east-1
```

2. **Authenticate Docker with ECR:**
```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com
```

3. **Build and Push Docker Image:**
```bash
docker build -t todoapp .
docker tag todoapp:latest YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
docker push YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
```

4. **Create RDS Database:**
```bash
aws rds create-db-instance \
  --db-instance-identifier todoapp-db \
  --db-instance-class db.t3.micro \
  --engine postgres \
  --master-username postgres \
  --master-user-password YOUR_DB_PASSWORD \
  --allocated-storage 20 \
  --region us-east-1
```

5. **Wait for Database:**
```bash
aws rds wait db-instance-available --db-instance-identifier todoapp-db --region us-east-1
```

6. **Initialize Database:**
```bash
DB_ENDPOINT=$(aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1 --query 'DBInstances[0].Endpoint.Address' --output text)
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d postgres -c "CREATE DATABASE todoapp;" --set=sslmode=require
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d todoapp -f setup_online_db.sql --set=sslmode=require
```

7. **Create IAM Role:**
```bash
aws iam create-role --role-name ecsTaskExecutionRole --assume-role-policy-document file://ecs-trust-policy.json
aws iam attach-role-policy --role-name ecsTaskExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
```

8. **Create Secrets:**
```bash
aws secretsmanager create-secret --name todoapp/db-password --secret-string "YOUR_DB_PASSWORD" --region us-east-1
aws secretsmanager create-secret --name todoapp/smtp-password --secret-string "your_smtp_password" --region us-east-1
```

9. **Create ECS Cluster:**
```bash
aws ecs create-cluster --cluster-name todoapp-cluster --region us-east-1
```

10. **Create CloudWatch Log Group:**
```bash
aws logs create-log-group --log-group-name /ecs/todoapp --region us-east-1
```

11. **Register Task Definition:**
```bash
aws ecs register-task-definition --cli-input-json file://ecs-task-definition.json --region us-east-1
```

12. **Create ECS Service:**
```bash
aws ecs create-service \
  --cluster todoapp-cluster \
  --service-name todoapp-service \
  --task-definition todoapp \
  --desired-count 1 \
  --launch-type FARGATE \
  --network-configuration "awsvpcConfiguration={subnets=[subnet-0545de7b031651d31],securityGroups=[sg-054c63b7a75086273],assignPublicIp=ENABLED}" \
  --region us-east-1
```

### Get Application URL

```bash
TASK_ARN=$(aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1 --query 'taskArns[0]' --output text)
NETWORK_INTERFACE_ID=$(aws ecs describe-tasks --cluster todoapp-cluster --tasks $TASK_ARN --region us-east-1 --query 'tasks[0].attachments[0].details[?name==`networkInterfaceId`].value' --output text)
PUBLIC_IP=$(aws ec2 describe-network-interfaces --network-interface-ids $NETWORK_INTERFACE_ID --region us-east-1 --query 'NetworkInterfaces[0].Association.PublicIp' --output text)
echo "Application URL: http://$PUBLIC_IP:8080"
```

## Security Features

### Authentication & Authorization
- JWT token-based authentication
- Secure password hashing with bcrypt
- Token expiration and refresh mechanisms
- Role-based access control

### Protection Mechanisms
- Brute force protection with account lockout
- Rate limiting on authentication endpoints
- Input validation and sanitization
- SQL injection prevention through parameterized queries
- CORS configuration for cross-origin requests

### Data Security
- Environment variable configuration for secrets
- AWS Secrets Manager integration
- SSL/TLS encryption for database connections
- Secure session management

## Monitoring & Logging

### Application Logs
- Structured logging with request/response details
- Error tracking and debugging information
- Performance metrics logging

### AWS CloudWatch Integration
- Centralized log aggregation
- Log retention policies
- Real-time log streaming
- Custom metrics and alarms

### Health Checks
- Database connectivity monitoring
- Application health endpoints
- Container health status
- Service availability monitoring

## Troubleshooting

### Common Issues

1. **Application won't start:**
   - Check environment variables
   - Verify database connectivity
   - Review application logs

2. **Database connection errors:**
   - Verify database credentials
   - Check network connectivity
   - Ensure database is running

3. **Email functionality not working:**
   - Verify SMTP credentials
   - Check Gmail App Password setup
   - Review email service logs

4. **Authentication issues:**
   - Verify JWT secret configuration
   - Check token expiration settings
   - Review authentication middleware logs

### Debug Commands

```bash
# Check application logs
docker-compose logs -f todoapp

# Check database connectivity
docker-compose exec postgres psql -U postgres -d todoapp -c "\dt"

# Verify environment variables
docker-compose exec todoapp env | grep -E "(PG_|JWT_|SMTP_)"

# Test database connection
docker-compose exec todoapp psql -h postgres -U postgres -d todoapp -c "SELECT version();"
```

## Development

### Adding New Features

1. **Database Changes:**
   - Create migration files in `database/migrations/`
   - Update SQL queries in `database/queries/`
   - Regenerate code with `sqlc generate`

2. **API Endpoints:**
   - Add handlers in appropriate `api/` files
   - Register routes in `routers.go`
   - Update frontend JavaScript as needed

3. **Frontend Changes:**
   - Modify HTML/CSS in `static/` directory
   - Update JavaScript functionality
   - Add new translations to i18n system

### Code Generation

```bash
# Generate database code
sqlc generate

# Download dependencies
go mod download

# Run tests
go test ./...
```

## License

MIT License - see LICENSE file for details
