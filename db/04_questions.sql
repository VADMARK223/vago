TRUNCATE TABLE questions RESTART IDENTITY CASCADE;
BEGIN;
INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (1, 4, 'Что произойдёт при исполнении?', 'ch := make(chan int)
go func() {
    ch <- 10
}()
fmt.Println(<-ch)
', 'Обычная синхронная передача данных.');
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Deadlock', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, '10', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Panic: closed channel', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (1, 'Неопределенное поведение', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (2, 4, 'Что выведет код?', 'm := make(map[string]int)
m["a"] = 1
for k := range m {
    delete(m, k)
}
fmt.Println(len(m))
', 'Go допускает удаление элементов из map во время итерации.');
INSERT INTO answers (question_id, text, is_correct) VALUES (2, '1', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, '0', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'panic', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (2, 'Неопределенное поведение', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (3, 11, 'Для чего нужен интерфейс error?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Для обработки исключений', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Для передачи ошибок как значений', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Чтобы логировать ошибки автоматически', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (3, 'Чтобы завершать программу', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (4, 9, 'Какая коллекция в Go является потокобезопасной по умолчанию?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Map', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Slice', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Channel', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (4, 'Array', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (5, 9, 'Какой тип передается в канал chan int?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Только указатели на int', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Только значения типа int', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Любые числа', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (5, 'Любые типы', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (6, 9, 'Что произойдёт при записи в закрытый канал?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Ничего — значение игнорируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Запись блокируется', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Происходит panic', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (6, 'Возвращается zero-value', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (7, 13, 'Когда происходит escape в heap?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Когда переменная слишком большая', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Когда значение нужно вернуть наружу и оно "живет" после выхода функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Когда в функции много переменных', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (7, 'Когда программа вызывает panic', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (8, 7, 'Что такое goroutine?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Отдельный процесс ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Лёгкая потокоподобная сущность, управляемая рантаймом Go', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Поток ОС', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (8, 'Операция ввода-вывода', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (9, 12, 'Что делает ключевое слово defer?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (9, 'Выполняет функцию асинхронно', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (9, 'Выполняет функцию после выхода из текущей функции', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (9, 'Отменяет выполнение функции', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (9, 'Экспортирует функцию', false);

INSERT INTO questions (id, topic_id, text, code, explanation) VALUES (10, 5, 'Какой размер у пустой структуры struct{}?', '', '');
INSERT INTO answers (question_id, text, is_correct) VALUES (10, '1 байт', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (10, '0 байт', true);
INSERT INTO answers (question_id, text, is_correct) VALUES (10, '4 байта', false);
INSERT INTO answers (question_id, text, is_correct) VALUES (10, 'Зависит от архитектуры', false);

COMMIT;
