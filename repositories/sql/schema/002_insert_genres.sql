-- +goose Up
INSERT INTO genres(title)
VALUES ('Komedïyalar'),
       ('Отбасымен көретіндер'),
       ('Ğılımï-tanımdıq'),
       ('Ойын-сауық'),
       ('Ğılımï fantastïka jäne féntezï'),
       ('Şıtırman oqïğal'),
       ('Qısqametrli'),
       ('Mwzıkalıq'),
       ('Sporttıq');

-- +goose Down
DELETE FROM genres;