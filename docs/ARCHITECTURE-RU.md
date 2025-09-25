# 🏗️ Документация по архитектуре (Архитектура Вашего Цифрового Додзё) ⛩️

> *"Понимание архитектуры - это как изучение планов замка перед штурмом!"* ⚔️✨

## 🎭 Навигация (Ваши Путеводные Звезды) 🌟

- **[🏠 Главный README](../README.md)** - Обзор проекта и быстрый старт
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🏠 Local Development](README-LOCAL.md)** - Local PostgreSQL setup
- **[🚀 AWS Deployment](DEPLOYMENT.md)** - Production deployment guide
- **[🔧 API Documentation](API.md)** - Complete API reference

### 🌍 Языковые версии
- **[🇺🇸 English Version](ARCHITECTURE.md)** - English version of this guide

## 🎯 Обзор системы (Обзор Вашей Цифровой Империи) 🗾

> *"TodoApp построен с использованием чистой, модульной архитектуры - как хорошо спроектированный замок!"* ⛩️

TodoApp построен с использованием чистой, модульной архитектуры, которая разделяет ответственность и делает кодовую базу поддерживаемой и масштабируемой.

<!-- ![Architecture Overview](https://via.placeholder.com/800x400/FF6B9D/FFFFFF?text=Architecture+Overview+Anime) -->
<!-- Place your architecture overview anime image here -->

## 🏛️ Диаграмма архитектуры

```
┌─────────────────────────────────────────────────────────────┐
│                        Слой Frontend                        │
├─────────────────────────────────────────────────────────────┤
│  HTML/CSS/JavaScript (Vanilla)                             │
│  - Адаптивный дизайн                                        │
│  - Поддержка нескольких языков (i18n)                      │
│  - Обновления в реальном времени                            │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Слой API Gateway                         │
├─────────────────────────────────────────────────────────────┤
│  Gorilla Mux Router                                         │
│  - Регистрация маршрутов                                    │
│  - Цепочка middleware                                       │
│  - Обслуживание статических файлов                          │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Слой приложения                          │
├─────────────────────────────────────────────────────────────┤
│  API обработчики                                            │
│  ├── auth.go (Аутентификация)                              │
│  ├── user.go (Управление пользователями)                   │
│  ├── tasks.go (Личные задачи)                              │
│  └── team_projects.go (Командное сотрудничество)           │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Слой бизнес-логики                       │
├─────────────────────────────────────────────────────────────┤
│  - JWT аутентификация                                       │
│  - Хеширование паролей (bcrypt)                            │
│  - Валидация входных данных                                 │
│  - Бизнес-правила                                           │
│  - Email сервисы (SMTP)                                     │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Слой доступа к данным                    │
├─────────────────────────────────────────────────────────────┤
│  SQLC сгенерированный код                                   │
│  - Типобезопасные запросы к БД                             │
│  - Пул соединений                                           │
│  - Управление транзакциями                                  │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    Слой базы данных                         │
├─────────────────────────────────────────────────────────────┤
│  PostgreSQL Database                                        │
│  - Соответствие ACID                                        │
│  - Реляционная модель данных                                │
│  - Индексы для производительности                           │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Структура проекта

```
todoapp/
├── main.go                    # Точка входа приложения
├── api/                       # API обработчики
│   ├── auth.go               # Аутентификация и авторизация
│   ├── routers.go            # Регистрация маршрутов
│   ├── tasks.go              # Управление личными задачами
│   ├── team_projects.go      # Командное сотрудничество
│   └── user.go               # Управление пользователями
├── database/                  # Слой базы данных
│   ├── database.go           # Управление соединениями
│   ├── generated/            # SQLC сгенерированный код
│   │   ├── db.go            # Соединение с БД
│   │   ├── models.go        # Модели данных
│   │   ├── tasks.sql.go     # Запросы задач
│   │   ├── users.sql.go     # Запросы пользователей
│   │   └── team_projects.sql.go
│   ├── migrations/           # Миграции базы данных
│   └── queries/              # Определения SQL запросов
├── i18n/                     # Интернационализация
│   ├── i18n.go              # Конфигурация i18n
│   └── errors.go            # Переводы ошибок
├── settings/                 # Конфигурация
│   └── settings.go          # Переменные окружения
├── static/                   # Файлы frontend
│   ├── index.html           # Основное приложение
│   ├── script-simple.js     # Логика frontend
│   ├── i18n.js              # Frontend i18n
│   ├── verify-email.html    # Подтверждение email
│   └── reset-password.html  # Сброс пароля
└── docs/                     # Документация
    ├── README.md            # Основная документация
    ├── README-RU.md         # Русская документация
    ├── README-LOCAL.md      # Локальная разработка
    ├── DEPLOYMENT.md        # Развертывание AWS
    ├── API.md               # Справочник API
    └── ARCHITECTURE.md      # Этот файл
```

## 🔄 Поток данных

### 1. Обработка запросов
```
HTTP запрос → Router → Middleware → Handler → Бизнес-логика → База данных
```

### 2. Поток аутентификации
```
Запрос входа → Валидация учетных данных → Генерация JWT → Возврат токена
Защищенный запрос → Валидация JWT → Извлечение ID пользователя → Обработка запроса
```

### 3. Поток управления задачами
```
Создание задачи → Валидация входных данных → Сохранение в БД → Возврат данных задачи
Обновление задачи → Проверка владения → Обновление БД → Возврат обновленной задачи
```

## 🗄️ Схема базы данных

### Основные таблицы

#### Таблица пользователей
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

#### Таблица задач
```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN DEFAULT false,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    deadline TIMESTAMP
);
```

#### Таблица командных проектов
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

#### Таблица участников командных проектов
```sql
CREATE TABLE team_project_members (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES team_projects(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(project_id, user_id)
);
```

#### Таблица командных задач
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

### Таблицы безопасности

#### Токены сброса пароля
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

#### Токены подтверждения email
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

## 🔐 Архитектура безопасности

### Аутентификация и авторизация
- **JWT токены**: Безсостоятельная аутентификация
- **Хеширование паролей**: bcrypt с солью
- **Ограничение скорости**: Защита от брутфорса
- **Валидация входных данных**: Предотвращение SQL инъекций

### Защита данных
- **Переменные окружения**: Конфигурация чувствительных данных
- **HTTPS**: Шифрованная связь (продакшн)
- **CORS**: Обработка кросс-доменных запросов
- **Санитизация входных данных**: Предотвращение XSS

## 🚀 Архитектура развертывания

### Локальная разработка
```
Машина разработчика
├── Go приложение (go run)
├── Локальный PostgreSQL
└── SMTP (Gmail)
```

### Docker разработка
```
Docker Compose
├── Контейнер приложения
├── Контейнер PostgreSQL
├── Nginx балансировщик нагрузки
└── Постоянное хранилище
```

### AWS продакшн
```
AWS Cloud
├── ECS Fargate (Приложение)
├── RDS PostgreSQL (База данных)
├── ECR (Реестр контейнеров)
├── CloudWatch (Логирование)
├── Secrets Manager (Учетные данные)
└── Application Load Balancer
```

## 📊 Соображения производительности

### Оптимизация базы данных
- **Индексы**: На часто запрашиваемых столбцах
- **Пул соединений**: Эффективные соединения с БД
- **Оптимизация запросов**: SQLC сгенерированные запросы

### Оптимизация приложения
- **Gorilla Mux**: Эффективная маршрутизация
- **Обслуживание статических файлов**: Прямое обслуживание файлов
- **Кэширование JWT**: Оптимизация валидации токенов

### Масштабируемость
- **Горизонтальное масштабирование**: Несколько экземпляров приложения
- **Балансировка нагрузки**: Nginx/AWS ALB
- **Масштабирование БД**: Реплики для чтения (будущее)

## 🔧 Технологический стек

### Backend
- **Go 1.23**: Язык программирования
- **Gorilla Mux**: HTTP роутер
- **SQLC**: Генерация типобезопасного SQL кода
- **PostgreSQL**: Основная база данных
- **JWT**: Токены аутентификации

### Frontend
- **HTML5**: Разметка
- **CSS3**: Стилизация
- **Vanilla JavaScript**: Клиентская логика
- **Адаптивный дизайн**: Mobile-first подход

### Инфраструктура
- **Docker**: Контейнеризация
- **Docker Compose**: Локальная разработка
- **AWS ECS**: Оркестрация контейнеров
- **AWS RDS**: Управляемая база данных
- **Nginx**: Балансировка нагрузки

### Инструменты разработки
- **Git**: Контроль версий
- **SQLC**: Генерация кода
- **Postman**: Тестирование API
- **Docker**: Среда разработки

## 🎯 Паттерны проектирования

### Паттерн Repository
- SQLC сгенерированные запросы действуют как репозитории
- Чистое разделение между доступом к данным и бизнес-логикой

### Паттерн Middleware
- Middleware аутентификации
- Middleware логирования запросов
- Middleware обработки ошибок

### Паттерн MVC
- **Model**: Модели базы данных (SQLC сгенерированные)
- **View**: HTML шаблоны и статические файлы
- **Controller**: API обработчики

## 🔄 Рабочий процесс разработки

1. **Изменения БД**: Обновите SQL запросы в `database/queries/`
2. **Генерация кода**: Запустите `sqlc generate`
3. **Разработка API**: Добавьте обработчики в `api/`
4. **Обновления Frontend**: Измените файлы в `static/`
5. **Тестирование**: Используйте Postman или curl
6. **Развертывание**: Docker или AWS
