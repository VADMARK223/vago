-- +goose Up

CREATE TABLE IF NOT EXISTS topics
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS questions
(
    id          SERIAL PRIMARY KEY,
    topic_id    INT NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
    text        TEXT NOT NULL,
    code        TEXT,
    explanation TEXT,
    difficulty  SMALLINT CHECK (difficulty BETWEEN 1 AND 5),
    created_at  TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_questions_topic_id
    ON questions(topic_id);

CREATE TABLE IF NOT EXISTS answers
(
    id          SERIAL PRIMARY KEY,
    question_id INT NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    text        TEXT NOT NULL,
    is_correct  BOOLEAN DEFAULT FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS uniq_correct_answer_per_question
    ON answers(question_id)
    WHERE is_correct = TRUE;

-- +goose Down
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS topics;
