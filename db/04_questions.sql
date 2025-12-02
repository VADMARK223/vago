TRUNCATE TABLE questions RESTART IDENTITY CASCADE;
BEGIN;
INSERT INTO questions (id, topic_id, text) VALUES (1, 7, 'Что такое goroutine?');
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Отдельный процесс ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Лёгкая потокоподобная сущность, управляемая рантаймом Go', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Поток ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Операция ввода-вывода', false);

INSERT INTO questions (id, topic_id, text) VALUES (2, 12, 'Что делает ключевое слово defer?');
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Выполняет функцию асинхронно', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Выполняет функцию после выхода из текущей функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Отменяет выполнение функции', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Экспортирует функцию', false);

INSERT INTO questions (id, topic_id, text) VALUES (3, 5, 'Какой размер у пустой структуры struct{}?');
INSERT INTO answers (question_id, text, is_correct) VALUES (3, '1 байт', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, '0 байт', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, '4 байта', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Зависит от архитектуры', false);

COMMIT;
