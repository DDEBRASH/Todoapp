#!/bin/bash

# Скрипт для настройки локальной PostgreSQL базы данных

echo "🚀 Настройка локальной PostgreSQL базы данных..."

# Проверяем, установлена ли PostgreSQL
if ! command -v psql &> /dev/null; then
    echo "❌ PostgreSQL не установлена. Установите её:"
    echo "   macOS: brew install postgresql"
    echo "   Ubuntu: sudo apt install postgresql"
    exit 1
fi

# Проверяем, запущен ли PostgreSQL
if ! pg_isready -q; then
    echo "🔄 Запускаем PostgreSQL..."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        brew services start postgresql
    else
        # Linux
        sudo systemctl start postgresql
    fi
fi

# Создаем базу данных если не существует
echo "📦 Создаем базу данных todoapp..."
createdb todoapp 2>/dev/null || echo "База данных уже существует"

# Выполняем SQL скрипт
echo "🔧 Настраиваем схему базы данных..."
psql -d todoapp -f setup_online_db.sql

echo "✅ Локальная база данных настроена!"
echo "📊 Подключение: postgresql://postgres@localhost:5432/todoapp"
echo ""
echo "🚀 Теперь можно запускать:"
echo "   Локально: go run main.go"
echo "   Docker:   docker-compose -f docker-compose.dev.yml up -d"
