-- Создание таблицы командных проектов
CREATE TABLE team_projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(6) UNIQUE NOT NULL,
    created_by INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы участников командных проектов
CREATE TABLE team_project_members (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES team_projects(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, user_id)
);

-- Создание таблицы командных задач
CREATE TABLE team_tasks (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES team_projects(id) ON DELETE CASCADE,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT FALSE,
    created_by INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для оптимизации
CREATE INDEX idx_team_projects_code ON team_projects(code);
CREATE INDEX idx_team_projects_created_by ON team_projects(created_by);
CREATE INDEX idx_team_project_members_project_id ON team_project_members(project_id);
CREATE INDEX idx_team_project_members_user_id ON team_project_members(user_id);
CREATE INDEX idx_team_tasks_project_id ON team_tasks(project_id);
CREATE INDEX idx_team_tasks_created_by ON team_tasks(created_by);
