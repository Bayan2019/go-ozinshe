-- +goose Up
INSERT INTO types(title)
VALUES ('fïlm'),
       ('Mwltfïlm'),
       ('Serïyalıq'),
       ('Mwltserïal');

-- +goose Down
DELETE FROM types;