
-- +migrate Up
CREATE TABLE IF NOT EXISTS "items" (
    id INT PRIMARY KEY,
    name TEXT,
    attrs JSONB DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "items";