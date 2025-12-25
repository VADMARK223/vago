CREATE TABLE IF NOT EXISTS topics
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
COMMENT ON TABLE topics IS 'Таблица тем';

CREATE TABLE IF NOT EXISTS questions
(
    id           SERIAL PRIMARY KEY,
    topic_id     INTEGER NOT NULL REFERENCES topics (id) ON DELETE CASCADE,
    text         TEXT    NOT NULL,
    code         TEXT,
    explanation  TEXT,
    difficulty   SMALLINT CHECK (difficulty BETWEEN 1 AND 5),
    created_at   TIMESTAMPTZ DEFAULT NOW(),
    content_hash TEXT    NOT NULL
);
COMMENT ON TABLE questions IS 'Таблица вопросов';

CREATE INDEX IF NOT EXISTS idx_questions_topic_id
    ON questions (topic_id);

CREATE UNIQUE INDEX questions_content_hash_uq
    ON questions (content_hash);


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
    ON answers (question_id)
    WHERE is_correct = true;


CREATE TABLE comments
(
    id          BIGSERIAL PRIMARY KEY,
    question_id BIGINT    NOT NULL,
    parent_id   BIGINT    NULL,
    author_id   BIGINT    NOT NULL,
    content     TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP,

    CONSTRAINT fk_question
        FOREIGN KEY (question_id) REFERENCES questions (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_parent
        FOREIGN KEY (parent_id) REFERENCES comments (id)
            ON DELETE CASCADE
);
COMMENT ON TABLE comments IS 'Таблица комментариев';
comment on column comments.parent_id is 'Если null, то комментарий к вопросу, иначе к комментарию';
CREATE INDEX idx_comments_question
    ON comments (question_id);

CREATE INDEX idx_comments_parent
    ON comments (parent_id);

CREATE INDEX idx_comments_created
    ON comments (created_at);