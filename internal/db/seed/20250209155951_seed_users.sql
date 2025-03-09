-- +goose Up
INSERT INTO users (name)
VALUES ('admin');

-- +goose Down
DELETE FROM users WHERE name='admin';
