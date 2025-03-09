-- +goose Up
CREATE SEQUENCE users_id_seq;
CREATE TABLE users (
      id integer NOT NULL DEFAULT nextval('users_id_seq'),
      name text,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY(id)
);
ALTER SEQUENCE users_id_seq OWNED BY users.id;

-- +goose Down
DROP TABLE users;
