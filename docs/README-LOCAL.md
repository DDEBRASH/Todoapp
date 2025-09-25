# 🏠 Локальная разработка TodoApp 

## 🎭 Навигация (Ваши Путеводные Звезды) 🌟

- **[🏠 Главный README](../README.md)** - Обзор проекта и быстрый старт
- **[🇺🇸 English Documentation](README.md)** - Complete project documentation
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🚀 AWS Deployment](DEPLOYMENT.md)** - Production deployment guide
- **[🔧 API Documentation](README.md#api-endpoints)** - Complete API reference

### 🌍 Language Versions
- **[🇺🇸 English Version](README-LOCAL-EN.md)** - English version of this guide

## 🎯 Цель 

Настроить единую базу данных, которая работает для всех способов запуска:
- 🏠 Локальная разработка (go run) - традиционный путь
- 🐳 Docker - современный способ
- ☁️ AWS - облачное королевство

## 🚀 Быстрая настройка 

### 1. Установите PostgreSQL локально 
```bash
# macOS 
brew install postgresql
brew services start postgresql
```

<!-- ![PostgreSQL Setup](https://via.placeholder.com/600x300/FF6B9D/FFFFFF?text=PostgreSQL+Setup+Anime) -->
<!-- Place your PostgreSQL setup anime image here -->

### 2. Настройте базу данных
```bash
./setup-local-db.sh
```

### 3. Выберите способ запуска

#### Вариант A: Локальная разработка
```bash
# Скопируйте настройки
cp env.local .env

# Запустите приложение
go run main.go
```

#### Вариант B: Docker (подключается к локальной БД)
```bash
# Запустите только приложение
docker-compose -f docker-compose.dev.yml up -d
```

#### Вариант C: Полный Docker (с собственной БД)
```bash
# Запустите все сервисы
docker-compose up -d
```

## 🔧 Настройка для AWS

Когда будете готовы к продакшну:

1. **Создайте RDS в AWS**
2. **Экспортируйте данные**:
   ```bash
   pg_dump -U postgres -d todoapp > backup.sql
   ```
3. **Импортируйте в RDS**:
   ```bash
   psql -h your-rds-endpoint -U postgres -d todoapp < backup.sql
   ```
4. **Обновите переменные окружения** в AWS

## 🎯 Рекомендуемый workflow

1. **Разработка**: Используйте локальную PostgreSQL
2. **Тестирование**: Docker с локальной БД
3. **Продакшн**: AWS RDS
