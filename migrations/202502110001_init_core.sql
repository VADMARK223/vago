-- +goose Up

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    login      VARCHAR(30) UNIQUE NOT NULL,
    username   VARCHAR(30)        NOT NULL,
    password   VARCHAR(255)       NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    role       VARCHAR(20) CHECK (role IN ('user', 'moderator', 'admin')) NOT NULL,
    color      VARCHAR(7) CHECK (color ~ '^#[0-9A-Fa-f]{6}$') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    completed   BOOLEAN DEFAULT FALSE,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_id     INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    content      TEXT NOT NULL,
    message_type VARCHAR(30) NOT NULL DEFAULT 'text',
    created_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_messages_user_id ON messages (user_id);
CREATE INDEX IF NOT EXISTS idx_messages_messages_type ON messages (message_type);

-- +goose Down
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;