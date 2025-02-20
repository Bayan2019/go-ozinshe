-- +goose Up
INSERT INTO types(title)
VALUES ('Фильм'),
       ('Мультфильм'),
       ('Сериал'),
       ('Мультсериал');

-- +goose Down
DELETE FROM types;