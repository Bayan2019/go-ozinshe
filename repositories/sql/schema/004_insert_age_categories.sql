-- +goose Up
INSERT INTO age_categories(title)
VALUES ('6-8 jas'),
        ('8-10 jas'),
        ('10-12 jas'),
        ('12-14 jas'),
        ('14-16 jas'),
        ('16-18 jas'),
        ('18-120 jas');

-- +goose Down
DELETE FROM age_categories;