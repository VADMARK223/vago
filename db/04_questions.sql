/*
insert into users (username, password, email, color, created_at) values ('1', '$2a$10$uKQGFGjps52djEN1yYUkvO5cUuELFbqZgxFOyxI6D6kjwh5Ne2W5m', 'user1@mail.ru', '#FF5733', now());
insert into users (username, password, email, color, created_at) values ('2', '$2a$10$AHiUObaB7UdeslZP2WAd.uCbXu01LspUz7KiLMPOfze67NIYJEcPy', 'user2@mail.ru', '#33FF57',now());*/

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
