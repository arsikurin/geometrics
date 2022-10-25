-- schema.sql

DROP TABLE IF EXISTS courses_problems;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS submits;
DROP TABLE IF EXISTS problems;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    id          SERIAL       NOT NULL,
    login       TEXT         NOT NULL,
    password    BYTEA        NOT NULL,
    type        INT          NOT NULL DEFAULT 0,
    first_name  VARCHAR(100) NOT NULL,
    last_name   VARCHAR(100) NOT NULL,
    grade       INT,
    school      VARCHAR(100),
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    last_online TIMESTAMP    NOT NULL,
    timezone    TEXT         NOT NULL DEFAULT 'Europe/Moscow',
    PRIMARY KEY (id)
);

CREATE TABLE problems
(
    id           SERIAL NOT NULL,
    name         TEXT   NOT NULL,
    description  TEXT   NOT NULL,
    solution_raw TEXT   NOT NULL,
    toolbar      TEXT   NOT NULL DEFAULT '0 62 | 1 | 2 15 18 | 53 | 40 41 42 , 27 28 6',
    PRIMARY KEY (id)
);

CREATE TABLE submits
(
    id           SERIAL    NOT NULL,
    user_id      INT       NOT NULL,
    problem_id   INT       NOT NULL,
    status       INT       NOT NULL,
    solution_raw TEXT      NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (problem_id) REFERENCES problems (id),
    FOREIGN KEY (user_id) REFERENCES users (id)

);

create TABLE courses
(
    id          SERIAL NOT NULL,
    author_id   INT    NOT NULL,
    name        TEXT   NOT NULL,
    description TEXT,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES users (id)

);

CREATE TABLE courses_problems
(
    id         SERIAL NOT NULL,
    course_id  INT    NOT NULL,
    problem_id INT    NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (course_id) REFERENCES courses (id),
    FOREIGN KEY (problem_id) REFERENCES problems (id),
    UNIQUE (course_id, problem_id)
);


-- Getting all problems for a course:

-- SELECT * FROM problems INNER JOIN courses_problems ON courses_problems.problem_id = problems.id WHERE courses_problems.course_id = 2;


