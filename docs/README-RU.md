# TodoApp - Система управления задачами 


## 🎭 Навигация (Ваши Путеводные Звезды) 🌟

- **[🏠 Главный README](../README.md)** - Обзор проекта и быстрый старт
- **[🇺🇸 English Documentation](README.md)** - Complete project documentation
- **[🏠 Локальная разработка](README-LOCAL.md)** - Настройка локальной PostgreSQL
- **[🚀 Развертывание в AWS](DEPLOYMENT.md)** - Руководство по продакшн развертыванию
- **[🔧 API Документация](#api-endpoints)** - Полная справка по API
- **[🏗️ Архитектура](#технологический-стек)** - Технические детали

## Обзор 

TodoApp - это веб-приложение, написанное на Go, которое предоставляет возможности управления задачами как для индивидуальных пользователей, так и для командной работы. Приложение включает аутентификацию пользователей, верификацию email, сброс паролей, управление личными задачами и совместную работу над проектами в реальном времени.


## Основные возможности 


- **🔐 Аутентификация и безопасность** 
  - 🔑 JWT-аутентификация (секреты)
  - 📧 Регистрация пользователей с верификацией email 
  - 🔄 Сброс паролей через email 
  - 🛡️ Защита от брутфорс-атак 
  - 🔒 Безопасное хеширование паролей с bcrypt 
- **📝 Управление задачами** 
  - ✨ Создание, чтение, обновление и удаление личных задач 
  - ⏰ Установка дедлайнов с визуальными индикаторами 
  - ✅ Отметка задач как выполненных 
  - 📱 Адаптивный дизайн для мобильных и десктопных устройств

- **👥 Командная работа** 
  - 🏰 Создание командных проектов с уникальными 6-значными кодами 
  - 🤝 Присоединение к существующим проектам по коду 
  - ⚔️ Совместное управление задачами в рамках проектов 
  - ⚡ Функции реального времени для совместной работы 

- **🌍 Интернационализация** (Универсальное Общение)
  - 🗣️ Поддержка нескольких языков (русский, английский, монгольский)
  - 🔄 Динамическое переключение языков (говори на любом языке!)
  - 📜 Локализованные сообщения об ошибках (понятные сообщения!)

- **🚀 Готовность к развертыванию** 
  - 🐳 Контейнеризация Docker 
  - ☁️ Конфигурация развертывания AWS ECS (твое облачное королевство!)
  - ⚖️ Балансировка нагрузки с Nginx (твой цифровой балансировщик!)
  - 🗄️ База данных PostgreSQL с миграциями (твой дворец памяти!)

## Технологический стек

### Backend
- **Язык**: Go 1.23
- **Веб-фреймворк**: Gorilla Mux
- **База данных**: PostgreSQL
- **ORM**: SQLC для типобезопасных запросов к базе данных
- **Аутентификация**: JWT токены
- **Email**: SMTP интеграция (Gmail)

### Frontend
- **HTML/CSS/JavaScript**: Vanilla реализация
- **Адаптивный дизайн**: Mobile-first подход
- **Интернационализация**: Кастомная i18n система

### Инфраструктура
- **Контейнеризация**: Docker & Docker Compose
- **Облачная платформа**: AWS (ECS, RDS, ECR)
- **Балансировщик нагрузки**: Nginx
- **База данных**: PostgreSQL с пулом соединений

## Структура проекта

```
todoapp/
├── api/                          # API обработчики и маршруты
│   ├── auth.go                   # Middleware аутентификации и сброс паролей
│   ├── routers.go                # Регистрация маршрутов
│   ├── tasks.go                  # Управление личными задачами
│   ├── team_projects.go          # Функциональность командных проектов
│   └── user.go                   # Управление пользователями
├── database/                     # Слой базы данных
│   ├── generated/                # Сгенерированный код SQLC
│   │   ├── db.go                 # Соединение с базой данных
│   │   ├── models.go             # Модели данных
│   │   ├── tasks.sql.go          # Запросы задач
│   │   ├── users.sql.go          # Запросы пользователей
│   │   ├── team_projects.sql.go  # Запросы командных проектов
│   │   ├── email_verification.sql.go
│   │   └── password_resets.sql.go
│   ├── migrations/               # Миграции базы данных
│   │   ├── 0001_create_tasks.up.sql
│   │   ├── 0002_create_users.up.sql
│   │   ├── 0003_user_id_tasks.up.sql
│   │   ├── 0004_security_users.up.sql
│   │   ├── 0005_email_password.up.sql
│   │   ├── 0005_user_blocker.up.sql
│   │   ├── 0006_email_verification.up.sql
│   │   ├── 0007_create_team_projects.up.sql
│   │   └── 0008_add_deadline_to_tasks.up.sql
│   ├── queries/                  # Определения SQL запросов
│   │   ├── tasks.sql
│   │   ├── users.sql
│   │   ├── team_projects.sql
│   │   ├── email_verification.sql
│   │   └── password_resets.sql
│   ├── database.go               # Инициализация базы данных
│   └── schema.sql                # Схема базы данных
├── i18n/                         # Интернационализация
│   ├── i18n.go                   # Конфигурация i18n
│   └── errors.go                 # Переводы сообщений об ошибках
├── settings/                     # Настройки приложения
│   └── settings.go               # Обработка переменных окружения
├── static/                       # Файлы frontend
│   ├── index.html                # Основной интерфейс приложения
│   ├── script-simple.js          # Frontend JavaScript
│   ├── i18n.js                   # Frontend интернационализация
│   ├── verify-email.html         # Страница верификации email
│   └── reset-password.html       # Страница сброса пароля
├── docker-compose.yml            # Настройка локальной разработки
├── Dockerfile                    # Контейнеризация приложения
├── ecs-task-definition.json      # Конфигурация AWS ECS
├── ecs-trust-policy.json         # IAM trust policy
├── setup_online_db.sql           # Скрипт инициализации базы данных
├── sqlc.yaml                     # Конфигурация SQLC
├── go.mod                        # Зависимости Go модуля
└── main.go                       # Точка входа приложения
```

## Схема базы данных

### Таблица Users
- `id`: Первичный ключ (SERIAL)
- `username`: Уникальное имя пользователя (TEXT)
- `email`: Уникальный email адрес (TEXT)
- `password_hash`: Хешированный пароль Bcrypt (TEXT)
- `is_blocked`: Статус блокировки аккаунта (BOOLEAN)
- `failed_attempts`: Счетчик неудачных попыток входа (INT)
- `last_failed`: Временная метка последней неудачной попытки (TIMESTAMP)
- `email_verified`: Статус верификации email (BOOLEAN)

### Таблица Tasks
- `id`: Первичный ключ (SERIAL)
- `title`: Описание задачи (TEXT)
- `done`: Статус выполнения (BOOLEAN)
- `user_id`: Внешний ключ к users (INT)
- `deadline`: Опциональный дедлайн задачи (TIMESTAMP)

### Таблица Team Projects
- `id`: Первичный ключ (SERIAL)
- `name`: Название проекта (TEXT)
- `code`: Уникальный 6-значный код проекта (TEXT)
- `created_by`: ID создателя проекта (INT)
- `created_at`: Временная метка создания (TIMESTAMP)

### Таблица Team Project Members
- `id`: Первичный ключ (SERIAL)
- `project_id`: Внешний ключ к team_projects (INT)
- `user_id`: Внешний ключ к users (INT)
- `joined_at`: Временная метка присоединения (TIMESTAMP)

### Таблица Team Tasks
- `id`: Первичный ключ (SERIAL)
- `title`: Описание задачи (TEXT)
- `done`: Статус выполнения (BOOLEAN)
- `project_id`: Внешний ключ к team_projects (INT)
- `created_by`: ID создателя задачи (INT)
- `deadline`: Опциональный дедлайн задачи (TIMESTAMP)

### Таблицы безопасности
- `password_reset_tokens`: Функциональность сброса паролей
- `email_verification_tokens`: Система верификации email

## API Endpoints

### Аутентификация
- `POST /api/auth/register` - Регистрация нового пользователя
- `POST /api/auth/login` - Вход пользователя
- `POST /api/auth/logout` - Выход пользователя
- `POST /api/auth/request-password-reset` - Запрос сброса пароля
- `POST /api/auth/reset-password` - Сброс пароля с токеном

### Личные задачи
- `GET /api/tasks` - Получить задачи пользователя
- `POST /api/tasks` - Создать новую задачу
- `PATCH /api/tasks/{id}` - Обновить задачу
- `DELETE /api/tasks/{id}` - Удалить задачу

### Командные проекты
- `POST /api/team-projects` - Создать командный проект
- `POST /api/team-projects/join` - Присоединиться к существующему проекту
- `GET /api/team-projects/my` - Получить командные проекты пользователя
- `GET /api/team-projects/{id}/members` - Получить участников проекта
- `GET /api/team-projects/{id}/tasks` - Получить задачи проекта
- `POST /api/team-projects/{id}/tasks` - Создать задачу в проекте
- `PATCH /api/team-projects/{id}/tasks/{taskId}` - Обновить задачу проекта
- `DELETE /api/team-projects/{id}/tasks/{taskId}` - Удалить задачу проекта

## Быстрый старт

### Предварительные требования
- Docker и Docker Compose
- Git

### Настройка локальной разработки

1. **Клонируйте репозиторий:**
```bash
git clone <repository-url>
cd todoapp
```

2. **Создайте файл окружения:**
```bash
cp env.example .env
# Отредактируйте .env с вашей конфигурацией
```

3. **Запустите приложение:**
```bash
docker-compose up -d
```

4. **Откройте приложение:**
```
http://localhost:3000
```

### Переменные окружения

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_PORT` | Порт приложения | `7540` |
| `APP_BASE_URL` | Базовый URL для ссылок в email | `http://localhost:7540` |
| `PG_HOST` | Хост PostgreSQL | `localhost` |
| `PG_PORT` | Порт PostgreSQL | `5432` |
| `PG_USER` | Пользователь базы данных | `postgres` |
| `PG_PASS` | Пароль базы данных | - |
| `PG_DB` | Имя базы данных | `todoapp` |
| `PG_SSLMODE` | SSL режим | `disable` |
| `JWT_SECRET` | Секрет для подписи JWT | - |
| `ADMIN_TOKEN` | Токен администратора | - |
| `SMTP_HOST` | SMTP сервер | `smtp.gmail.com` |
| `SMTP_PORT` | SMTP порт | `587` |
| `SMTP_USER` | SMTP имя пользователя | - |
| `SMTP_PASS` | SMTP пароль | - |
| `SMTP_FROM` | Email адрес отправителя | - |
| `LOGIN_MAX_FAIL` | Максимум неудачных попыток входа до блокировки | `4` |

### Настройка SMTP (Gmail)

1. Включите двухфакторную аутентификацию в вашем аккаунте Gmail
2. Создайте пароль приложения
3. Используйте пароль приложения в переменной окружения `SMTP_PASS`

## Развертывание в AWS

### Предварительные требования
- Настроенный AWS CLI
- Установленный Docker
- AWS аккаунт с соответствующими разрешениями

### Пошаговое развертывание

1. **Создайте ECR репозиторий:**
```bash
aws ecr create-repository --repository-name todoapp --region us-east-1
```

2. **Аутентифицируйте Docker в ECR:**
```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com
```

3. **Соберите и загрузите Docker образ:**
```bash
docker build -t todoapp .
docker tag todoapp:latest YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
docker push YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
```

4. **Создайте RDS базу данных:**
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

5. **Дождитесь готовности базы данных:**
```bash
aws rds wait db-instance-available --db-instance-identifier todoapp-db --region us-east-1
```

6. **Инициализируйте базу данных:**
```bash
DB_ENDPOINT=$(aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1 --query 'DBInstances[0].Endpoint.Address' --output text)
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d postgres -c "CREATE DATABASE todoapp;" --set=sslmode=require
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d todoapp -f setup_online_db.sql --set=sslmode=require
```

7. **Создайте IAM роль:**
```bash
aws iam create-role --role-name ecsTaskExecutionRole --assume-role-policy-document file://ecs-trust-policy.json
aws iam attach-role-policy --role-name ecsTaskExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
```

8. **Создайте секреты:**
```bash
aws secretsmanager create-secret --name todoapp/db-password --secret-string "YOUR_DB_PASSWORD" --region us-east-1
aws secretsmanager create-secret --name todoapp/smtp-password --secret-string "your_smtp_password" --region us-east-1
```

9. **Создайте ECS кластер:**
```bash
aws ecs create-cluster --cluster-name todoapp-cluster --region us-east-1
```

10. **Создайте CloudWatch Log Group:**
```bash
aws logs create-log-group --log-group-name /ecs/todoapp --region us-east-1
```

11. **Зарегистрируйте Task Definition:**
```bash
aws ecs register-task-definition --cli-input-json file://ecs-task-definition.json --region us-east-1
```

12. **Создайте ECS Service:**
```bash
aws ecs create-service \
  --cluster todoapp-cluster \
  --service-name todoapp-service \
  --task-definition todoapp \
  --desired-count 1 \
  --launch-type FARGATE \
  --network-configuration "awsvpcConfiguration={subnets=[YOUR_SUBNET_ID],securityGroups=[YOUR_SECURITY_GROUP_ID],assignPublicIp=ENABLED}" \
  --region us-east-1
```

### Получение URL приложения

```bash
TASK_ARN=$(aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1 --query 'taskArns[0]' --output text)
NETWORK_INTERFACE_ID=$(aws ecs describe-tasks --cluster todoapp-cluster --tasks $TASK_ARN --region us-east-1 --query 'tasks[0].attachments[0].details[?name==`networkInterfaceId`].value' --output text)
PUBLIC_IP=$(aws ec2 describe-network-interfaces --network-interface-ids $NETWORK_INTERFACE_ID --region us-east-1 --query 'NetworkInterfaces[0].Association.PublicIp' --output text)
echo "URL приложения: http://$PUBLIC_IP:8080"
```

## Функции безопасности

### Аутентификация и авторизация
- JWT токен-аутентификация
- Безопасное хеширование паролей с bcrypt
- Механизмы истечения и обновления токенов
- Контроль доступа на основе ролей

### Механизмы защиты
- Защита от брутфорс-атак с блокировкой аккаунта
- Ограничение скорости на endpoints аутентификации
- Валидация и санитизация входных данных
- Предотвращение SQL-инъекций через параметризованные запросы
- Конфигурация CORS для кросс-доменных запросов

### Безопасность данных
- Конфигурация переменных окружения для секретов
- Интеграция с AWS Secrets Manager
- SSL/TLS шифрование для соединений с базой данных
- Безопасное управление сессиями

## Мониторинг и логирование

### Логи приложения
- Структурированное логирование с деталями запросов/ответов
- Отслеживание ошибок и отладочная информация
- Логирование метрик производительности

### Интеграция с AWS CloudWatch
- Централизованная агрегация логов
- Политики хранения логов
- Потоковая передача логов в реальном времени
- Пользовательские метрики и алерты

### Проверки здоровья
- Мониторинг подключения к базе данных
- Endpoints проверки здоровья приложения
- Статус здоровья контейнеров
- Мониторинг доступности сервисов

## Устранение неполадок

### Частые проблемы

1. **Приложение не запускается:**
   - Проверьте переменные окружения
   - Убедитесь в подключении к базе данных
   - Просмотрите логи приложения

2. **Ошибки подключения к базе данных:**
   - Проверьте учетные данные базы данных
   - Проверьте сетевое подключение
   - Убедитесь, что база данных запущена

3. **Email функциональность не работает:**
   - Проверьте учетные данные SMTP
   - Проверьте настройку пароля приложения Gmail
   - Просмотрите логи email сервиса

4. **Проблемы с аутентификацией:**
   - Проверьте конфигурацию JWT секрета
   - Проверьте настройки истечения токена
   - Просмотрите логи middleware аутентификации

### Команды отладки

```bash
# Проверить логи приложения
docker-compose logs -f todoapp

# Проверить подключение к базе данных
docker-compose exec postgres psql -U postgres -d todoapp -c "\dt"

# Проверить переменные окружения
docker-compose exec todoapp env | grep -E "(PG_|JWT_|SMTP_)"

# Тестировать подключение к базе данных
docker-compose exec todoapp psql -h postgres -U postgres -d todoapp -c "SELECT version();"
```

## Разработка

### Добавление новых функций

1. **Изменения в базе данных:**
   - Создайте файлы миграций в `database/migrations/`
   - Обновите SQL запросы в `database/queries/`
   - Регенерируйте код с помощью `sqlc generate`

2. **API Endpoints:**
   - Добавьте обработчики в соответствующие файлы `api/`
   - Зарегистрируйте маршруты в `routers.go`
   - Обновите frontend JavaScript при необходимости

3. **Изменения Frontend:**
   - Измените HTML/CSS в директории `static/`
   - Обновите функциональность JavaScript
   - Добавьте новые переводы в систему i18n

### Генерация кода

```bash
# Генерировать код базы данных
sqlc generate

# Загрузить зависимости
go mod download

# Запустить тесты
go test ./...
```

## Лицензия

MIT License - см. файл LICENSE для деталей

