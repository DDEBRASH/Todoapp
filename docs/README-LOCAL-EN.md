# 🏠 Local Development TodoApp

## 🎭 Navigation (Your Guiding Stars) 🌟

- **[🏠 Main README](../README.md)** - Project overview and quick start
- **[🇺🇸 English Documentation](README.md)** - Complete project documentation
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🚀 AWS Deployment](DEPLOYMENT.md)** - Production deployment guide
- **[🔧 API Documentation](README.md#api-endpoints)** - Complete API reference

### 🌍 Language Versions
- **[🇷🇺 Русская Версия](README-LOCAL.md)** - Русская версия этого руководства

## 🎯 Goal

Set up a unified database that works for all launch methods:
- 🏠 Local development (go run) - traditional path
- 🐳 Docker - modern approach
- ☁️ AWS - cloud kingdom

## 🚀 Quick Setup

### 1. Install PostgreSQL locally
```bash
# macOS
brew install postgresql
brew services start postgresql

# Ubuntu
sudo apt install postgresql
sudo systemctl start postgresql
```

<!-- ![PostgreSQL Setup](https://via.placeholder.com/600x300/FF6B9D/FFFFFF?text=PostgreSQL+Setup+Anime) -->
<!-- Place your PostgreSQL setup anime image here -->

### 2. Set up database
```bash
./setup-local-db.sh
```

### 3. Choose launch method

#### Option A: Local Development
```bash
# Copy settings
cp env.local .env

# Run application
go run main.go
```

#### Option B: Docker (connects to local DB)
```bash
# Run only application
docker-compose -f docker-compose.dev.yml up -d
```

#### Option C: Full Docker (with own DB)
```bash
# Run all services
docker-compose up -d
```

## 🔧 AWS Setup

When ready for production:

1. **Create RDS in AWS**
2. **Export data**:
   ```bash
   pg_dump -U postgres -d todoapp > backup.sql
   ```
3. **Import to RDS**:
   ```bash
   psql -h your-rds-endpoint -U postgres -d todoapp < backup.sql
   ```
4. **Update environment variables** in AWS

## 🎯 Recommended workflow

1. **Development**: Use local PostgreSQL
2. **Testing**: Docker with local DB
3. **Production**: AWS RDS
