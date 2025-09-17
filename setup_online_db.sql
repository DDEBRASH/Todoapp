-- Полный SQL скрипт для настройки онлайн базы данных
-- Выполните этот скрипт в psql подключившись к Koyeb базе

-- 1. Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    failed_attempts INT NOT NULL DEFAULT 0,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    email TEXT,
    email_verified BOOLEAN DEFAULT FALSE
);

-- 2. Создание таблицы задач
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT false,
    user_id INT NOT NULL REFERENCES users(id),
    deadline TIMESTAMP
);

-- 3. Создание таблицы токенов сброса пароля
CREATE TABLE IF NOT EXISTS password_reset_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_password_reset_tokens_token ON password_reset_tokens(token);

-- 4. Создание таблицы токенов верификации email
CREATE TABLE IF NOT EXISTS email_verification_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_email_verification_tokens_token ON email_verification_tokens(token);

-- 5. Создание таблицы командных проектов
CREATE TABLE IF NOT EXISTS team_projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(6) UNIQUE NOT NULL,
    created_by INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 6. Создание таблицы участников командных проектов
CREATE TABLE IF NOT EXISTS team_project_members (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES team_projects(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, user_id)
);

-- 7. Создание таблицы командных задач
CREATE TABLE IF NOT EXISTS team_tasks (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES team_projects(id) ON DELETE CASCADE,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT FALSE,
    created_by INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deadline TIMESTAMP
);

-- 8. Создание индексов для оптимизации
CREATE INDEX IF NOT EXISTS idx_team_projects_code ON team_projects(code);
CREATE INDEX IF NOT EXISTS idx_team_projects_created_by ON team_projects(created_by);
CREATE INDEX IF NOT EXISTS idx_team_project_members_project_id ON team_project_members(project_id);
CREATE INDEX IF NOT EXISTS idx_team_project_members_user_id ON team_project_members(user_id);
CREATE INDEX IF NOT EXISTS idx_team_tasks_project_id ON team_tasks(project_id);
CREATE INDEX IF NOT EXISTS idx_team_tasks_created_by ON team_tasks(created_by);
CREATE INDEX IF NOT EXISTS idx_tasks_deadline ON tasks(deadline);
CREATE INDEX IF NOT EXISTS idx_team_tasks_deadline ON team_tasks(deadline);

-- Проверка создания таблиц
\dt
