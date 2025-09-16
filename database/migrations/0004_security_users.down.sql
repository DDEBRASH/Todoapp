ALTER TABLE users
    DROP COLUMN failed_attempts,
    DROP COLUMN is_blocked,
    DROP COLUMN email TEXT;