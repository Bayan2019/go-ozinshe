-- +goose Up
INSERT INTO permissions(title)
VALUES ('тыйым салынған'),
       ('тек оқу'),
       ('редакциялау');

-- +goose Down
DELETE FROM permissions;