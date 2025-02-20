-- +goose Up
INSERT INTO users_roles(user_id, role_id)
VALUES(1, 1);

-- +goose Down
DELETE FROM users_roles;