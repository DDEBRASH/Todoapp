# 🚀 TodoApp Deployment Guide (Path to Cloud) ☁️

> *"From local to cloud kingdom - the path"*

## 🎭 Navigation 🌟

- **[🏠 Main README](../README.md)** - Project overview and quick start
- **[🇺🇸 English Documentation](README.md)** - Complete project documentation
- **[🇷🇺 Русская Документация](README-RU.md)** - Полная документация на русском
- **[🏠 Local Development](README-LOCAL.md)** - Local PostgreSQL setup
- **[🔧 API Documentation](README.md#api-endpoints)** - Complete API reference

### 🌍 Language Versions
- **[🇷🇺 Русская Версия](DEPLOYMENT.md)** - Русская версия этого руководства

## 📋 Contents
1. [Local Deployment](#local-deployment)
2. [AWS Deployment](#aws-deployment)
3. [Monitoring and Logs](#monitoring-and-logs)
4. [Troubleshooting](#troubleshooting)

## ⚠️ Important: Replace placeholders with your data

Before executing commands, replace the following placeholders with your real data:

- 🔑 `YOUR_AWS_ACCOUNT_ID` - your AWS Account ID (12 digits) - your key to cloud kingdom
- 🔒 `YOUR_DB_PASSWORD` - password for PostgreSQL database
- 🛡️ `YOUR_SECURITY_GROUP_ID` - your Security Group ID (e.g.: sg-xxxxxxxxx) - your digital shield
- 🌐 `YOUR_SUBNET_ID` - your subnet ID (e.g.: subnet-xxxxxxxxx) - your digital territory
- 📧 `your_email@gmail.com` - your email for SMTP - your mail pigeon
- 🔐 `your_smtp_password` - application password for SMTP - your pigeon's secret

## 🏠 Local Deployment

### Prerequisites
- Docker and Docker Compose
- Git

### Quick Start
```bash
# 1. Clone repository
git clone <repository-url>
cd todoapp

# 2. Create .env file
cp env.example .env
# Edit .env with your settings

# 3. Run application
docker-compose up -d

# 4. Open browser
open http://localhost:3000
```

### Verification
```bash
# Check container status
docker-compose ps

# View logs
docker-compose logs -f todoapp

# Check database connection
docker-compose exec postgres psql -U postgres -d todoapp -c "\dt"
```

## ☁️ AWS Deployment

### Prerequisites
- AWS CLI configured and authenticated
- Docker installed
- AWS account with necessary permissions

### Step 1: Prepare AWS Resources

#### 1.1 Create ECR Repository
```bash
aws ecr create-repository --repository-name todoapp --region us-east-1
```

#### 1.2 Authenticate Docker with ECR
```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com
```

### Step 2: Build and Upload Docker Image

#### 2.1 Build Image
```bash
docker build -t todoapp .
```

#### 2.2 Tag for ECR
```bash
docker tag todoapp:latest YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
```

#### 2.3 Upload to ECR
```bash
docker push YOUR_AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/todoapp:latest
```

### Step 3: Create RDS Database

#### 3.1 Create RDS Instance
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

#### 3.2 Wait for Database Ready
```bash
aws rds wait db-instance-available --db-instance-identifier todoapp-db --region us-east-1
```

#### 3.3 Create Database and Tables
```bash
# Get database endpoint
DB_ENDPOINT=$(aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1 --query 'DBInstances[0].Endpoint.Address' --output text)

# Create database
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d postgres -c "CREATE DATABASE todoapp;" --set=sslmode=require

# Create tables
PGPASSWORD="YOUR_DB_PASSWORD" psql -h $DB_ENDPOINT -U postgres -d todoapp -f setup_online_db.sql --set=sslmode=require
```

### Step 4: Configure IAM and Secrets Manager

#### 4.1 Create IAM Role for ECS
```bash
aws iam create-role --role-name ecsTaskExecutionRole --assume-role-policy-document file://ecs-trust-policy.json
aws iam attach-role-policy --role-name ecsTaskExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
```

#### 4.2 Create Secrets
```bash
aws secretsmanager create-secret --name todoapp/db-password --secret-string "YOUR_DB_PASSWORD" --region us-east-1
aws secretsmanager create-secret --name todoapp/smtp-password --secret-string "your_smtp_password" --region us-east-1
```

### Step 5: Create ECS Cluster and Service

#### 5.1 Create Cluster
```bash
aws ecs create-cluster --cluster-name todoapp-cluster --region us-east-1
```

#### 5.2 Create CloudWatch Log Group
```bash
aws logs create-log-group --log-group-name /ecs/todoapp --region us-east-1
```

#### 5.3 Register Task Definition
```bash
aws ecs register-task-definition --cli-input-json file://ecs-task-definition.json --region us-east-1
```

#### 5.4 Create ECS Service
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

### Step 6: Configure Security Groups

#### 6.1 Open Ports
```bash
# Open port 8080 for application
aws ec2 authorize-security-group-ingress --group-id YOUR_SECURITY_GROUP_ID --protocol tcp --port 8080 --cidr 0.0.0.0/0 --region us-east-1

# Open port 5432 for database (if needed)
aws ec2 authorize-security-group-ingress --group-id YOUR_SECURITY_GROUP_ID --protocol tcp --port 5432 --cidr 0.0.0.0/0 --region us-east-1
```

## 📊 Monitoring and Logs

### Check Service Status
```bash
# ECS Service Status
aws ecs describe-services --cluster todoapp-cluster --services todoapp-service --region us-east-1

# RDS Status
aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1

# List Tasks
aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1
```

### View Logs
```bash
# Get latest logs
aws logs get-log-events --log-group-name /ecs/todoapp --log-stream-name $(aws logs describe-log-streams --log-group-name /ecs/todoapp --region us-east-1 --order-by LastEventTime --descending --max-items 1 --query 'logStreams[0].logStreamName' --output text) --region us-east-1
```

### Get Public IP
```bash
# Get application IP address
TASK_ARN=$(aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1 --query 'taskArns[0]' --output text)
NETWORK_INTERFACE_ID=$(aws ecs describe-tasks --cluster todoapp-cluster --tasks $TASK_ARN --region us-east-1 --query 'tasks[0].attachments[0].details[?name==`networkInterfaceId`].value' --output text)
PUBLIC_IP=$(aws ec2 describe-network-interfaces --network-interface-ids $NETWORK_INTERFACE_ID --region us-east-1 --query 'NetworkInterfaces[0].Association.PublicIp' --output text)
echo "Application available at: http://$PUBLIC_IP:8080"
```

## 🛑 Stop Services

### Temporary Stop
```bash
# Stop application
aws ecs update-service --cluster todoapp-cluster --service todoapp-service --desired-count 0 --region us-east-1

# Stop database
aws rds stop-db-instance --db-instance-identifier todoapp-db --region us-east-1
```

### Complete Removal
```bash
# Delete ECS Service
aws ecs delete-service --cluster todoapp-cluster --service todoapp-service --region us-east-1

# Delete RDS Database
aws rds delete-db-instance --db-instance-identifier todoapp-db --skip-final-snapshot --region us-east-1

# Delete ECS Cluster
aws ecs delete-cluster --cluster todoapp-cluster --region us-east-1

# Delete ECR Repository
aws ecr delete-repository --repository-name todoapp --force --region us-east-1
```

## 🔧 Troubleshooting

### Problem: Application won't start
```bash
# Check logs
aws logs get-log-events --log-group-name /ecs/todoapp --log-stream-name $(aws logs describe-log-streams --log-group-name /ecs/todoapp --region us-east-1 --order-by LastEventTime --descending --max-items 1 --query 'logStreams[0].logStreamName' --output text) --region us-east-1

# Check task status
aws ecs describe-tasks --cluster todoapp-cluster --tasks $(aws ecs list-tasks --cluster todoapp-cluster --service-name todoapp-service --region us-east-1 --query 'taskArns[0]' --output text) --region us-east-1
```

### Problem: Can't connect to database
```bash
# Check RDS status
aws rds describe-db-instances --db-instance-identifier todoapp-db --region us-east-1 --query 'DBInstances[0].DBInstanceStatus'

# Check Security Groups
aws ec2 describe-security-groups --group-ids YOUR_SECURITY_GROUP_ID --region us-east-1
```

### Problem: Application not accessible externally
```bash
# Check Security Groups
aws ec2 describe-security-groups --group-ids YOUR_SECURITY_GROUP_ID --region us-east-1 --query 'SecurityGroups[0].IpPermissions'

# Ensure port 8080 is open
aws ec2 authorize-security-group-ingress --group-id YOUR_SECURITY_GROUP_ID --protocol tcp --port 8080 --cidr 0.0.0.0/0 --region us-east-1
```
