-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS technologies
(
    id          UUID DEFAULT uuid_generate_v4(),
    name        VARCHAR(1024) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP     NOT NULL,
    updated_at  TIMESTAMP     NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS categories
(
    id            UUID DEFAULT uuid_generate_v4(),
    technology_id UUID          NOT NULL,
    name          VARCHAR(1024) NOT NULL,
    description   TEXT,
    created_at    TIMESTAMP     NOT NULL,
    updated_at    TIMESTAMP     NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_categories_technologies_id FOREIGN KEY (technology_id) REFERENCES technologies (id)
);

CREATE TABLE IF NOT EXISTS topics
(
    id          UUID DEFAULT uuid_generate_v4(),
    category_id UUID          NOT NULL,
    seq_no      INT           NOT NULL,
    name        VARCHAR(1024) NOT NULL,
    description TEXT,
    active      BOOLEAN       NOT NULL,
    created_at  TIMESTAMP     NOT NULL,
    updated_at  TIMESTAMP     NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_topics_categories_id FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TABLE IF NOT EXISTS tasks
(
    id          UUID DEFAULT uuid_generate_v4(),
    seq_no      INT           NOT NULL,
    name        VARCHAR(1024) NOT NULL,
    description TEXT,
    active      BOOLEAN       NOT NULL,
    created_at  TIMESTAMP     NOT NULL,
    updated_at  TIMESTAMP     NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tags
(
    name       VARCHAR(128) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS topics_tags
(
    id         UUID DEFAULT uuid_generate_v4(),
    topic_id   UUID         NOT NULL,
    tag_name   VARCHAR(128) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_topics_tags_topics_id FOREIGN KEY (topic_id) REFERENCES topics (id),
    CONSTRAINT fk_topics_tags_tags_name FOREIGN KEY (tag_name) REFERENCES tags (name)
);

CREATE TABLE IF NOT EXISTS tasks_tags
(
    id         UUID DEFAULT uuid_generate_v4(),
    task_id    UUID         NOT NULL,
    tag_name   VARCHAR(128) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_tasks_tags_topics_id FOREIGN KEY (task_id) REFERENCES tasks (id),
    CONSTRAINT fk_tasks_tags_tags_name FOREIGN KEY (tag_name) REFERENCES tags (name)
);

CREATE TABLE IF NOT EXISTS courses
(
    id          UUID DEFAULT uuid_generate_v4(),
    seq_no      INT           NOT NULL,
    name        VARCHAR(1024) NOT NULL,
    description TEXT,
    active      BOOLEAN       NOT NULL,
    owner_id    UUID          NOT NULL,
    created_at  TIMESTAMP     NOT NULL,
    updated_at  TIMESTAMP     NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS courses_topics
(
    id         UUID DEFAULT uuid_generate_v4(),
    course_id  UUID      NOT NULL,
    topic_id   UUID      NOT NULL,
    seq_no     INT       NOT NULL,
    active     BOOLEAN   NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_courses_topics_courses_id FOREIGN KEY (course_id) REFERENCES courses (id),
    CONSTRAINT fk_courses_topics_topics_id FOREIGN KEY (topic_id) REFERENCES topics (id)
);

CREATE TABLE IF NOT EXISTS tasks_courses_topics
(
    id              UUID DEFAULT uuid_generate_v4(),
    course_topic_id UUID      NOT NULL,
    task_id         UUID      NOT NULL,
    weight          INT       NOT NULL,
    active          BOOLEAN   NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_tasks_courses_topics_courses_topics_id FOREIGN KEY (course_topic_id) REFERENCES courses_topics (id),
    CONSTRAINT fk_tasks_courses_topics_tasks_id FOREIGN KEY (task_id) REFERENCES tasks (id)
);

-- +goose Down
DROP TABLE IF EXISTS tasks_courses_topics;
DROP TABLE IF EXISTS courses_topics;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS tasks_tags;
DROP TABLE IF EXISTS topics_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS topics;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS technologies;
