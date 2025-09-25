# TodoApp - Complete Task Management System



A comprehensive web application built with Go for task management and team collaboration.

## 🎬 Project Overview Video

> *"Watch the journey of a true task management warrior!"* ⚔️

<video width="800" height="450" controls>
  <source src="docs/images/overview.mp4" type="video/mp4">
</video>

## 🎭 Navigation & Documentation

![TodoApp Demo](https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExc3FldWFyNWJyaDJnNGIwbDkxcDM1NnM2amtmMDMzN2ZkdjJsczNhZCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/l2Je5Mochl7iyWHoQ/giphy.gif)


### ⚡ Quick Start Guides
- **[🏠 Local Development](docs/README-LOCAL.md)** - Set up local PostgreSQL and development environment
- **[☁️ AWS Deployment](docs/DEPLOYMENT.md)** - Deploy to AWS ECS with RDS
- **[🔧 API Documentation](docs/API.md)** - Complete API reference with examples
- **[🏗️ Architecture](docs/ARCHITECTURE.md)** - System design and technical details


### 🎮 Configuration Files
- **[🌐 Public Templates](public/)** - Safe configuration templates with placeholders
- **[🔒 Private Config](private/)** - Your personal configuration (not in git)


## 🎯 What is TodoApp? 

![TodoApp Demo](https://media1.tenor.com/m/Qr-JcZEAOekAAAAd/jaggydohwhift.gif)

TodoApp is a full-featured task management system that supports:

- **🔐 User Authentication** - JWT-based auth with email verification (like a digital ninja scroll!)
- **📝 Personal Tasks** - Create, manage, and track individual tasks (your personal quest log!)
- **👥 Team Projects** - Collaborate on shared projects with unique codes (guild missions!)
- **🌍 Multi-language Support** - Russian, English, Mongolian (universal communication!)
- **🐳 Docker Ready** - Containerized for easy deployment (sealed in a scroll!)
- **☁️ AWS Compatible** - Ready for production deployment (scales like a dragon!)

## 🏯 Project Structure (Your Digital Dojo) ⛩️

```
todoapp/                     # 🎌 Your main training ground
├── docs/                    # 📚 scrolls of knowledge
│   ├── README.md           # 📖 Main documentation (English)
│   ├── README-RU.md        # 🎎 Russian documentation  
│   ├── README-LOCAL.md     # 🏠 Local development guide
│   └── DEPLOYMENT.md       # ☁️ AWS deployment guide
├── private/                 # 🔒 Ur dirty secrets
│   ├── DEPLOYMENT-PERSONAL.md
│   ├── ecs-task-definition.personal.json
│   └── env.local.personal
├── public/                  # 🌐 Public
│   ├── docker-compose.yml
│   ├── ecs-task-definition.json
│   └── env.local
├── api/                     # ⚔️ API handlers
├── database/                # 🗄️ Database layer
├── static/                  # 🎨 Frontend files
└── main.go                  # 🚀 Application entry point
```

## ⚡ Quick Start (Begin Your Journey!) 🌟
![TodoApp Demo](https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExcm9jNHp2YmJ0ZnpwMWNnbzZpNTJzM2lubWk3dnhlMzl2dDJteDRuMyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/gUnRTJ0zqHJRe/giphy.gif)
### 🏠 Option 1: Local Development (The Traditional Path)
```bash
# 1. Set up local PostgreSQL (prepare your training ground)
./setup-local-db.sh

# 2. Copy environment template (gather your tools)
cp public/env.local .env

# 3. Run application (begin your quest!)
go run main.go
```

### 🐳 Option 2: Docker Development (The Modern Way)
```bash
# 1. Copy environment template (prepare your scrolls)
cp public/env.local .env

# 2. Start with Docker (summon your containers!)
docker-compose -f public/docker-compose.dev.yml up -d
```

### ⚔️ Option 3: Full Docker Stack (The Complete Arsenal)
```bash
# 1. Copy environment template (gather all your equipment)
cp public/env.local .env

# 2. Start all services (unleash the full power!)
docker-compose up -d
```

## 🔐 Security Features (Your Digital Armor) ⚔️
![](https://media1.tenor.com/m/cBMIWdZM-5MAAAAd/tampa-bay-rays-security-guard-security-guard.gif)

- **🔑 JWT Authentication** - Secure token-based auth (
- **🔒 Password Hashing** - bcrypt encryption 
- **📧 Email Verification** - Account activation via email 
- **🛡️ Brute Force Protection** - Account lockout after failed attempts (your digital shield!)
- **✅ Input Validation** - SQL injection prevention (your code's immune system!)
- **🌐 CORS Configuration** - Cross-origin request handling 

## 🌐 API Endpoints (Your Digital Techniques) ⚔️

![](https://remimercier.com/media/2017/what-is-an-api-remi-mercier.gif)

### 🔐 Authentication 
- `POST /api/auth/register` - User registration 
- `POST /api/auth/login` - User login 
- `POST /api/auth/logout` - User logout 
- `POST /api/auth/request-password-reset` - Password reset request 
- `POST /api/auth/reset-password` - Password reset 

### 📝 Personal Tasks 
- `GET /api/tasks` - Get user tasks 
- `POST /api/tasks` - Create task 
- `PATCH /api/tasks/{id}` - Update task 
- `DELETE /api/tasks/{id}` - Delete task 

### 👥 Team Projects 
- `POST /api/team-projects` - Create team project 
- `POST /api/team-projects/join` - Join project 
- `GET /api/team-projects/my` - Get user projects 
- `GET /api/team-projects/{id}/tasks` - Get project tasks 

## 🛠️ Technology Stack (Your Arsenal) ⚔️


- **⚔️ Backend**: Go 1.23, Gorilla Mux, SQLC 
- **🗄️ Database**: PostgreSQL 
- **🎨 Frontend**: Vanilla HTML/CSS/JavaScript 
- **🔑 Authentication**: JWT 
- **📧 Email**: SMTP (Gmail)
- **🐳 Containerization**: Docker 
- **☁️ Cloud**: AWS (ECS, RDS, ECR) 

## 📖 Documentation Links  📚

![](https://media1.tenor.com/m/_V8TTKAXYB0AAAAd/spongebob-squarepants-sunglasses.gif)

| Document | Description | Language |
|----------|-------------|----------|
| [📖 Main Documentation](docs/README.md) | Complete project overview | English |
| [🎎 Русская Документация](docs/README-RU.md) | Полный обзор проекта | Русский |
| [🏠 Local Setup](docs/README-LOCAL.md) | Local development guide | English |
| [☁️ AWS Deployment](docs/DEPLOYMENT.md) | Production deployment | English |

## 📄 License 

MIT License - see LICENSE file for details


## 🌟 Ready to Start Your Journey? Choose Your Path! ⚔️


- 🏠 **[Local Development](docs/README-LOCAL.md)** - For development 
- 🚀 **[AWS Deployment](docs/DEPLOYMENT.md)** - For production 
- 📚 **[Full Documentation](docs/README.md)** - Complete guide 

![](https://media1.tenor.com/m/o9HiEwdmmtAAAAAC/anime-diy.gif)