-- Добавляем поле deadline в таблицу tasks
ALTER TABLE tasks ADD COLUMN deadline TIMESTAMP;

-- Добавляем поле deadline в таблицу team_tasks
ALTER TABLE team_tasks ADD COLUMN deadline TIMESTAMP;

-- Создаем индексы для оптимизации запросов по дедлайнам
CREATE INDEX idx_tasks_deadline ON tasks(deadline);
CREATE INDEX idx_team_tasks_deadline ON team_tasks(deadline);
