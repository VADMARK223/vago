TRUNCATE TABLE questions RESTART IDENTITY CASCADE;
BEGIN;
INSERT INTO questions (id, topic_id, text) VALUES (1, 9, 'Какой тип передается в канал chan int?');
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Только указатели на int', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Только значения типа int', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Любые числа', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Любые типы', false);

INSERT INTO questions (id, topic_id, text) VALUES (2, 9, 'Что произойдёт при записи в закрытый канал?');
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Ничего — значение игнорируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Запись блокируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Происходит panic', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Возвращается zero-value', false);

INSERT INTO questions (id, topic_id, text) VALUES (3, 13, 'Когда происходит escape в heap?');
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Когда переменная слишком большая', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Когда значение нужно вернуть наружу и оно "живет" после выхода функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Когда в функции много переменных', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Когда программа вызывает panic', false);

INSERT INTO questions (id, topic_id, text) VALUES (4, 7, 'Что такое goroutine?');
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Отдельный процесс ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Лёгкая потокоподобная сущность, управляемая рантаймом Go', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Поток ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Операция ввода-вывода', false);

INSERT INTO questions (id, topic_id, text) VALUES (5, 12, 'Что делает ключевое слово defer?');
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Выполняет функцию асинхронно', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Выполняет функцию после выхода из текущей функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Отменяет выполнение функции', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Экспортирует функцию', false);

INSERT INTO questions (id, topic_id, text) VALUES (6, 5, 'Какой размер у пустой структуры struct{}?');
INSERT INTO answers (question_id, text, is_correct) VALUES (6, '1 байт', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, '0 байт', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, '4 байта', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Зависит от архитектуры', false);

COMMIT;
