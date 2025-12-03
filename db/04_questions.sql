TRUNCATE TABLE questions RESTART IDENTITY CASCADE;
BEGIN;
INSERT INTO questions (id, topic_id, text) VALUES (1, 11, 'Для чего нужен интерфейс error?');
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Для обработки исключений', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Для передачи ошибок как значений', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Чтобы логировать ошибки автоматически', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Чтобы завершать программу', false);

INSERT INTO questions (id, topic_id, text) VALUES (2, 9, 'Какая коллекция в Go является потокобезопасной по умолчанию?');
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Map', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Slice', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Channel', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Array', false);

INSERT INTO questions (id, topic_id, text) VALUES (3, 9, 'Какой тип передается в канал chan int?');
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Только указатели на int', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Только значения типа int', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Любые числа', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Любые типы', false);

INSERT INTO questions (id, topic_id, text) VALUES (4, 9, 'Что произойдёт при записи в закрытый канал?');
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Ничего — значение игнорируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Запись блокируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Происходит panic', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Возвращается zero-value', false);

INSERT INTO questions (id, topic_id, text) VALUES (5, 13, 'Когда происходит escape в heap?');
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Когда переменная слишком большая', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Когда значение нужно вернуть наружу и оно "живет" после выхода функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Когда в функции много переменных', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Когда программа вызывает panic', false);

INSERT INTO questions (id, topic_id, text) VALUES (6, 7, 'Что такое goroutine?');
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Отдельный процесс ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Лёгкая потокоподобная сущность, управляемая рантаймом Go', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Поток ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Операция ввода-вывода', false);

INSERT INTO questions (id, topic_id, text) VALUES (7, 12, 'Что делает ключевое слово defer?');
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Выполняет функцию асинхронно', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Выполняет функцию после выхода из текущей функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Отменяет выполнение функции', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Экспортирует функцию', false);

INSERT INTO questions (id, topic_id, text) VALUES (8, 5, 'Какой размер у пустой структуры struct{}?');
INSERT INTO answers (question_id, text, is_correct) VALUES (8, '1 байт', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, '0 байт', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, '4 байта', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Зависит от архитектуры', false);

COMMIT;
