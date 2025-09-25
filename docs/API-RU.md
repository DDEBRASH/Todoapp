# 🔧 Документация API 
## 🎭 Навигация (Ваши Путеводные Звезды) 🌟

- **[🏠 Главный README](../README.md)** - Обзор проекта и быстрый старт
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🏠 Local Development](README-LOCAL.md)** - Local PostgreSQL setup
- **[🚀 AWS Deployment](DEPLOYMENT.md)** - Production deployment guide

### 🌍 Языковые версии
- **[🇺🇸 English Version](API.md)** - English version of this guide

## 🔐 Аутентификация


Все защищенные эндпоинты требуют JWT токен в заголовке Authorization:
```
Authorization: Bearer <your-jwt-token>
```

## 📡 API Эндпоинты 

### 🔑 Эндпоинты аутентификации (Техники Аутентификации)

#### Регистрация пользователя
```http
POST /api/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

**Ответ:**
```json
{
  "id": 1,
  "username": "testuser",
  "email": "test@example.com",
  "message": "Registration successful!"
}
```

#### Вход пользователя
```http
POST /api/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

**Ответ:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Выход пользователя
```http
POST /api/auth/logout
Authorization: Bearer <token>
```

**Ответ:** `204 No Content`

#### Запрос сброса пароля
```http
POST /api/auth/request-password-reset
Content-Type: application/json

{
  "email": "test@example.com"
}
```

**Ответ:**
```json
{
  "message": "Password reset email sent successfully"
}
```

#### Сброс пароля
```http
POST /api/auth/reset-password
Content-Type: application/json

{
  "token": "reset-token-here",
  "password": "newpassword123"
}
```

**Ответ:** `204 No Content`

### 📝 Личные задачи

#### Получить задачи пользователя
```http
GET /api/tasks
Authorization: Bearer <token>
```

**Ответ:**
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

#### Создать задачу
```http
POST /api/tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "New task",
  "deadline": "2024-12-31T23:59:59Z"
}
```

**Ответ:**
```json
{
  "id": 2,
  "title": "New task",
  "done": false,
  "user_id": 1,
  "deadline": "2024-12-31T23:59:59Z"
}
```

#### Обновить задачу
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

**Ответ:**
```json
{
  "id": 1,
  "title": "Updated task title",
  "done": true,
  "user_id": 1,
  "deadline": "2024-12-31T23:59:59Z"
}
```

#### Удалить задачу
```http
DELETE /api/tasks/{id}
Authorization: Bearer <token>
```

**Ответ:** `204 No Content`

### 👥 Командные проекты

#### Создать командный проект
```http
POST /api/team-projects
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "My Team Project"
}
```

**Ответ:**
```json
{
  "id": 1,
  "name": "My Team Project",
  "code": "123456",
  "created_by": 1,
  "created_at": "2024-01-01T00:00:00Z"
}
```

#### Присоединиться к командному проекту
```http
POST /api/team-projects/join
Authorization: Bearer <token>
Content-Type: application/json

{
  "code": "123456"
}
```

**Ответ:**
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

#### Получить командные проекты пользователя
```http
GET /api/team-projects/my
Authorization: Bearer <token>
```

**Ответ:**
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

#### Получить участников командного проекта
```http
GET /api/team-projects/{id}/members
Authorization: Bearer <token>
```

**Ответ:**
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

#### Получить задачи командного проекта
```http
GET /api/team-projects/{id}/tasks
Authorization: Bearer <token>
```

**Ответ:**
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

#### Создать командную задачу
```http
POST /api/team-projects/{id}/tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "New team task",
  "deadline": "2024-12-31T23:59:59Z"
}
```

**Ответ:**
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

#### Обновить командную задачу
```http
PATCH /api/team-projects/{id}/tasks/{taskId}
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Updated team task",
  "done": true
}
```

**Ответ:**
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

#### Удалить командную задачу
```http
DELETE /api/team-projects/{id}/tasks/{taskId}
Authorization: Bearer <token>
```

**Ответ:** `204 No Content`

## 🚨 Ответы об ошибках

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

## 🔧 Тестирование с cURL

### Регистрация и вход
```bash
# Регистрация
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# Вход
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'

# Использование токена для защищенных эндпоинтов
TOKEN="your-jwt-token-here"
curl -X GET http://localhost:3000/api/tasks \
  -H "Authorization: Bearer $TOKEN"
```

## 📊 Коды ответов

| Код | Описание |
|-----|----------|
| 200 | OK - Запрос успешен |
| 201 | Created - Ресурс создан |
| 204 | No Content - Запрос успешен, нет тела ответа |
| 400 | Bad Request - Неверные данные запроса |
| 401 | Unauthorized - Требуется аутентификация |
| 404 | Not Found - Ресурс не найден |
| 500 | Internal Server Error - Ошибка сервера |

## 🔐 Заметки по безопасности

- Все пароли хешируются с использованием bcrypt
- JWT токены истекают через 24 часа
- Реализовано ограничение скорости для эндпоинтов аутентификации
- Выполняется валидация входных данных на всех эндпоинтах
- Защита от SQL инъекций на месте
