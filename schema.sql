-- schema.sql

drop table if exists courses_problems;
drop table if exists courses;
drop table if exists submits;
drop table if exists problems;
drop table if exists users;

create table users
(
    id          serial       not null,
    login       text         not null,
    password    text         not null,
    type        int          not null default 0,
    first_name  varchar(100) not null,
    last_name   varchar(100) not null,
    grade       int,
    school      varchar(100),
    created_at  timestamp    not null default NOW(),
    last_online timestamp    not null,
    timezone    text         not null default 'Europe/Moscow',
    primary key (id)
);

create table problems
(
    id           serial not null,
    name         text   not null,
    description  text,
    solution_raw text   not null,
    primary key (id)
);

create table submits
(
    id           serial    not null,
    user_id      int       not null,
    problem_id   int       not null,
    status       int       not null,
    solution_raw text      not null,
    created_at   timestamp not null default NOW(),
    primary key (id),
    foreign key (problem_id) references problems (id),
    foreign key (user_id) references users (id)

);

create table courses
(
    id          serial not null,
    author_id   int    not null,
    name        text   not null,
    description text,
    primary key (id),
    foreign key (author_id) references users (id)

);

create table courses_problems
(
    id         serial not null,
    course_id  int    not null,
    problem_id int    not null,
    primary key (id),
    foreign key (course_id) references courses (id),
    foreign key (problem_id) references problems (id),
    unique (course_id, problem_id)
);


-- Getting all problems for a course:

-- SELECT * FROM problems INNER JOIN courses_problems ON courses_problems.problem_id = problems.id WHERE courses_problems.course_id = 2;


