-- Удаляем индексы
DROP INDEX IF EXISTS idx_tasks_deadline;
DROP INDEX IF EXISTS idx_team_tasks_deadline;

-- Удаляем поле deadline из таблиц
ALTER TABLE tasks DROP COLUMN IF EXISTS deadline;
ALTER TABLE team_tasks DROP COLUMN IF EXISTS deadline;
