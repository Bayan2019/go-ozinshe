-- +goose Up
INSERT INTO permissions(title)
VALUES ('tıyım salınğan'),
       ('tek oqw'),
       ('redakcïyalaw');

-- +goose Down
DELETE FROM permissions;