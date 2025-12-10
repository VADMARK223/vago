-- +goose Up
UPDATE topics
SET name ='Срезы (Slices)'
WHERE id = 2;

UPDATE topics
SET name ='Массивы (Arrays)'
WHERE id = 3;

UPDATE topics
SET name ='Карты (Maps)'
WHERE id = 4;
