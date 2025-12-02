insert into questions (id, topic_id, text) values (1, 7, 'Что такое goroutine?');
insert into answers (question_id, text) values (1,'Отдельный процесс ОС');
insert into answers (question_id, text, is_correct) values (1,'Лёгкая потокоподобная сущность, управляемая рантаймом Go',true);
insert into answers (question_id, text) values (1,'Поток ОС');
insert into answers (question_id, text) values (1,'Операция ввода-вывода');

insert into questions (id, topic_id, text) values (2, 12, 'Что делает ключевое слово defer?');
insert into answers (question_id, text) values (2,'Выполняет функцию асинхронно');
insert into answers (question_id, text, is_correct) values (2,'Выполняет функцию после выхода из текущей функции',true);
insert into answers (question_id, text) values (2,'Отменяет выполнение функции');
insert into answers (question_id, text) values (2,'Экспортирует функцию');

insert into questions (id, topic_id, text) values (3, 5, 'Какой размер у пустой структуры struct{}?');
insert into answers (question_id, text) values (3,'1 байт');
insert into answers (question_id, text, is_correct) values (3,'0 байт',true);
insert into answers (question_id, text) values (3,'4 байта');
insert into answers (question_id, text) values (3,'Зависит от архитектуры');
