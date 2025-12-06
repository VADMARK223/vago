TRUNCATE TABLE questions RESTART IDENTITY CASCADE;
BEGIN;
INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (1, 4, 'Что выведет код?', 'm := make(map[string]int)
m["a"] = 1
for k := range m {
    delete(m, k)
}
fmt.Println(len(m))
', 'Go допускает удаление элементов из map во время итерации.');
INSERT INTO answers (question_id, text, is_correct) VALUES (1, '1', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, '0', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'panic', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Неопределенное поведение', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (2, 11, 'Для чего нужен интерфейс error?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Для обработки исключений', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Для передачи ошибок как значений', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Чтобы логировать ошибки автоматически', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Чтобы завершать программу', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (3, 9, 'Какая коллекция в Go является потокобезопасной по умолчанию?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Map', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Slice', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Channel', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Array', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (4, 9, 'Какой тип передается в канал chan int?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Только указатели на int', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Только значения типа int', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Любые числа', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Любые типы', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (5, 9, 'Что произойдёт при записи в закрытый канал?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Ничего — значение игнорируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Запись блокируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Происходит panic', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Возвращается zero-value', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (6, 13, 'Когда происходит escape в heap?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Когда переменная слишком большая', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Когда значение нужно вернуть наружу и оно "живет" после выхода функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Когда в функции много переменных', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Когда программа вызывает panic', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (7, 7, 'Что такое goroutine?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Отдельный процесс ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Лёгкая потокоподобная сущность, управляемая рантаймом Go', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Поток ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Операция ввода-вывода', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (8, 12, 'Что делает ключевое слово defer?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Выполняет функцию асинхронно', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Выполняет функцию после выхода из текущей функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Отменяет выполнение функции', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Экспортирует функцию', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (9, 5, 'Какой размер у пустой структуры struct{}?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (9, '1 байт', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (9, '0 байт', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (9, '4 байта', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (9, 'Зависит от архитектуры', false);

COMMIT;
