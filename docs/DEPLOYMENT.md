# 🚀 Руководство по развертыванию TodoApp (Путь к Облачному) ☁️

> *"От локального к облачному королевству - путь"*

## 🎭 Навигация  🌟

- **[🏠 Главный README](../README.md)** - Обзор проекта и быстрый старт
- **[🇺🇸 English Documentation](README.md)** - Complete project documentation
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🏠 Local Development](README-LOCAL.md)** - Local PostgreSQL setup
- **[🔧 API Documentation](README.md#api-endpoints)** - Complete API reference

### 🌍 Language Versions
- **[🇺🇸 English Version](DEPLOYMENT-EN.md)** - English version of this guide

## 📋 Содержание
1. [Локальное развертывание](#локальное-развертывание)
2. [Развертывание в AWS](#развертывание-в-aws)
3. [Мониторинг и логи](#мониторинг-и-логи)
4. [Устранение неполадок](#устранение-неполадок)

## ⚠️ Важно: Замените placeholder'ы на ваши данные 

Перед выполнением команд замените следующие placeholder'ы на ваши реальные данные:

- 🔑 `YOUR_AWS_ACCOUNT_ID` - ваш AWS Account ID (12 цифр) - ваш ключ к облачному королевству
- 🔒 `YOUR_DB_PASSWORD` - пароль для базы данных PostgreSQL 
- 🛡️ `YOUR_SECURITY_GROUP_ID` - ID вашей Security Group (например: sg-xxxxxxxxx) - ваш цифровой щит
- 🌐 `YOUR_SUBNET_ID` - ID вашей подсети (например: subnet-xxxxxxxxx) - ваша цифровая территория
- 📧 `your_email@gmail.com` - ваш email для SMTP - ваш почтовый голубь
- 🔐 `your_smtp_password` - пароль приложения для SMTP - секрет вашего голубя

## 🏠 Локальное развертывание

### Предварительные требования
- Docker и Docker Compose
- Git

### Быстрый старт
```bash
# 1. Клонируйте репозиторий
git clone <repository-url>
cd todoapp

# 2. Создайте .env файл
cp env.example .env
# Отредактируйте .env с вашими настройками

# 3. Запустите приложение
docker-compose up -d

# 4. Откройте браузер
open http://localhost:3000
```

### Проверка работы
```bash
# Проверьте статус контейнеров
docker-compose ps

# Посмотрите логи
docker-compose logs -f todoapp

# Проверьте подключение к базе данных
docker-compose exec postgres psql -U postgres -d todoapp -c "\dt"
```

## ☁️ Развертывание в AWS

### Предварительные требования
- AWS CLI настроен и аутентифицирован
- Docker установлен
- AWS аккаунт с необходимыми правами

### Шаг 1: Подготовка AWS ресурсов

#### 1.1 Создание ECR репозитория
```bash
aws ecr create-repository --repository-name todoapp --region us-east-1
```

#### 1.2 Аутентификация Docker в ECR
```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com
```

### Шаг 2: Сборка и загрузка Docker образа

#### 2.1 Сборка образа
```bash
docker build -t todoapp .
```

#### 2.2 Тегирование для ECR
```bash
docker tag todoapp:latest YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
```

#### 2.3 Загрузка в ECR
```bash
docker push YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
```

### Шаг 3: Создание RDS базы данных

#### 3.1 Создание экземпляра RDS
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

#### 3.2 Ожидание готовности базы данных
```bash
aws rds wait db-instance-available --db-instance-identifier todoapp-db --region us-east-1
```

#### 3.3 Создание базы данных и таблиц
```bash
# Получите endpoint базы данных
DB_ENDPOINT=$(aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1 --query 'DBInstances[0].Endpoint.Address' --output text)

# Создайте базу данных
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d postgres -c "CREATE DATABASE todoapp;" --set=sslmode=require

# Создайте таблицы
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d todoapp -f setup_online_db.sql --set=sslmode=require
```

### Шаг 4: Настройка IAM и Secrets Manager

#### 4.1 Создание IAM роли для ECS
```bash
aws iam create-role --role-name ecsTaskExecutionRole --assume-role-policy-document file://ecs-trust-policy.json
aws iam attach-role-policy --role-name ecsTaskExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
```

#### 4.2 Создание секретов
```bash
aws secretsmanager create-secret --name todoapp/db-password --secret-string "YOUR_DB_PASSWORD" --region us-east-1
aws secretsmanager create-secret --name todoapp/smtp-password --secret-string "your_smtp_password" --region us-east-1
```

### Шаг 5: Создание ECS кластера и сервиса

#### 5.1 Создание кластера
```bash
aws ecs create-cluster --cluster-name todoapp-cluster --region us-east-1
```

#### 5.2 Создание CloudWatch Log Group
```bash
aws logs create-log-group --log-group-name /ecs/todoapp --region us-east-1
```

#### 5.3 Регистрация Task Definition
```bash
aws ecs register-task-definition --cli-input-json file://ecs-task-definition.json --region us-east-1
```

#### 5.4 Создание ECS Service
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

### Шаг 6: Настройка Security Groups

#### 6.1 Открытие портов
```bash
# Открыть порт 8080 для приложения
aws ec2 authorize-security-group-ingress --group-id YOUR_SECURITY_GROUP_ID --protocol tcp --port 8080 --cidr 0.0.0.0/0 --region us-east-1

# Открыть порт 5432 для базы данных (если нужно)
aws ec2 authorize-security-group-ingress --group-id YOUR_SECURITY_GROUP_ID --protocol tcp --port 5432 --cidr 0.0.0.0/0 --region us-east-1
```

## 📊 Мониторинг и логи

### Проверка статуса сервисов
```bash
# Статус ECS Service
aws ecs describe-services --cluster todoapp-cluster --services todoapp-service --region us-east-1

# Статус RDS
aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1

# Список задач
aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1
```

### Просмотр логов
```bash
# Получить последние логи
aws logs get-log-events --log-group-name /ecs/todoapp --log-stream-name $(aws logs describe-log-streams --log-group-name /ecs/todoapp --region us-east-1 --order-by LastEventTime --descending --max-items 1 --query 'logStreams[0].logStreamName' --output text) --region us-east-1
```

### Получение публичного IP
```bash
# Получить IP адрес приложения
TASK_ARN=$(aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1 --query 'taskArns[0]' --output text)
NETWORK_INTERFACE_ID=$(aws ecs describe-tasks --cluster todoapp-cluster --tasks $TASK_ARN --region us-east-1 --query 'tasks[0].attachments[0].details[?name==`networkInterfaceId`].value' --output text)
PUBLIC_IP=$(aws ec2 describe-network-interfaces --network-interface-ids $NETWORK_INTERFACE_ID --region us-east-1 --query 'NetworkInterfaces[0].Association.PublicIp' --output text)
echo "Приложение доступно по адресу: http://$PUBLIC_IP:8080"
```

## 🛑 Остановка сервисов

### Временная остановка
```bash
# Остановить приложение
aws ecs update-service --cluster todoapp-cluster --service todoapp-service --desired-count 0 --region us-east-1

# Остановить базу данных
aws rds stop-db-instance --db-instance-identifier todoapp-db --region us-east-1
```

### Полное удаление
```bash
# Удалить ECS Service
aws ecs delete-service --cluster todoapp-cluster --service todoapp-service --region us-east-1

# Удалить RDS Database
aws rds delete-db-instance --db-instance-identifier todoapp-db --skip-final-snapshot --region us-east-1

# Удалить ECS Cluster
aws ecs delete-cluster --cluster todoapp-cluster --region us-east-1

# Удалить ECR Repository
aws ecr delete-repository --repository-name todoapp --force --region us-east-1
```

## 🔧 Устранение неполадок

### Проблема: Приложение не запускается
```bash
# Проверьте логи
aws logs get-log-events --log-group-name /ecs/todoapp --log-stream-name $(aws logs describe-log-streams --log-group-name /ecs/todoapp --region us-east-1 --order-by LastEventTime --descending --max-items 1 --query 'logStreams[0].logStreamName' --output text) --region us-east-1

# Проверьте статус задачи
aws ecs describe-tasks --cluster todoapp-cluster --tasks $(aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1 --query 'taskArns[0]' --output text) --region us-east-1
```

### Проблема: Не удается подключиться к базе данных
```bash
# Проверьте статус RDS
aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1 --query 'DBInstances[0].DBInstanceStatus'

# Проверьте Security Groups
aws ec2 describe-security-groups --group-ids YOUR_SECURITY_GROUP_ID --region us-east-1
```

### Проблема: Приложение недоступно извне
```bash
# Проверьте Security Groups
aws ec2 describe-security-groups --group-ids YOUR_SECURITY_GROUP_ID --region us-east-1 --query 'SecurityGroups[0].IpPermissions'

# Убедитесь, что порт 8080 открыт
aws ec2 authorize-security-group-ingress --group-id YOUR_SECURITY_GROUP_ID --protocol tcp --port 8080 --cidr 0.0.0.0/0 --region us-east-1
```

