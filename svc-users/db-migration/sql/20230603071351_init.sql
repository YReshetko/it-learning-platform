-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id          UUID                    DEFAULT uuid_generate_v4(),
    external_id UUID                    NOT NULL UNIQUE,
    first_name  VARCHAR(1024)           NOT NULL,
    last_name   VARCHAR(1024)           NOT NULL,
    email       VARCHAR(1024)           NOT NULL,
    active      BOOLEAN                 NOT NULL,
    created_at  TIMESTAMP               NOT NULL,
    updated_at  TIMESTAMP               NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX user_external_id_hash_idx ON users USING HASH (external_id);

-- +goose Down
DROP TABLE IF EXISTS users;
