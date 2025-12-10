-- +goose Up
INSERT INTO topics (id, name)
VALUES (14, 'Буферизованные каналы');

UPDATE topics
SET name ='Небуферизованные каналы'
WHERE id = 9;
