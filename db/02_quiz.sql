CREATE TABLE IF NOT EXISTS topics
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
COMMENT ON TABLE topics IS 'Таблица тем';

CREATE TABLE IF NOT EXISTS questions
(
    id          SERIAL PRIMARY KEY,
    topic_id    INTEGER NOT NULL REFERENCES topics (id) ON DELETE CASCADE,
    text        TEXT    NOT NULL,
    code        TEXT,
    explanation TEXT,
    difficulty  SMALLINT CHECK (difficulty BETWEEN 1 AND 5),
    created_at  TIMESTAMPTZ DEFAULT NOW()
);
COMMENT ON TABLE questions IS 'Таблица вопросов';

CREATE INDEX IF NOT EXISTS idx_questions_topic_id
    ON questions(topic_id);

CREATE TABLE IF NOT EXISTS answers
(
    id          SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL REFERENCES questions (id) ON DELETE CASCADE,
    text        TEXT    NOT NULL,
    is_correct  BOOLEAN NOT NULL DEFAULT false
);
COMMENT ON TABLE answers IS 'Таблица ответов';

-- Гарантия того, что у вопроса не более одного правильного ответа
CREATE UNIQUE INDEX uniq_correct_answer_per_question
    ON answers(question_id)
    WHERE is_correct = true;