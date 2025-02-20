-- +goose Up
INSERT INTO age_categories(title)
VALUES ('6-8 жас'),
        ('8-10 жас'),
        ('10-12 жас'),
        ('12-14 жас'),
        ('14-16 жас'),
        ('16-18 жас'),
        ('18-120 жас');

-- +goose Down
DELETE FROM age_categories;