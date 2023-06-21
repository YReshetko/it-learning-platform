-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS courses(
    id UUID                             DEFAULT uuid_generate_v4(),
    seq_no      INT                     NOT NULL,
    name        VARCHAR(1024)           NOT NULL,
    description TEXT,
    active      BOOLEAN                 NOT NULL,
    created_at  TIMESTAMP               NOT NULL,
    updated_at  TIMESTAMP               NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS topics(
    id UUID                             DEFAULT uuid_generate_v4(),
    seq_no      INT                     NOT NULL,
    name        VARCHAR(1024)           NOT NULL,
    description TEXT,
    active      BOOLEAN                 NOT NULL,
    created_at  TIMESTAMP               NOT NULL,
    updated_at  TIMESTAMP               NOT NULL,
    courses_id  UUID                    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_topics_courses_id FOREIGN KEY(courses_id) REFERENCES courses(id)
);

CREATE TABLE IF NOT EXISTS tasks(
    id UUID                             DEFAULT uuid_generate_v4(),
    seq_no      INT                     NOT NULL,
    weight      INT                     NOT NULL DEFAULT 1,
    name        VARCHAR(1024)           NOT NULL,
    description TEXT                    NOT NULL,
    active      BOOLEAN                 NOT NULL,
    created_at  TIMESTAMP               NOT NULL,
    updated_at  TIMESTAMP               NOT NULL,
    topics_id   UUID                    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_tasks_topics_id FOREIGN KEY(topics_id) REFERENCES topics(id)
);

-- +goose Down
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS topics;
DROP TABLE IF EXISTS courses;
