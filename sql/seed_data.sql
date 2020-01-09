create database todoapp;
use todoapp;

-- table users
create table users
(
    id         int unsigned auto_increment
        primary key,
    created_at timestamp    null,
    updated_at timestamp    null,
    last_name  varchar(255) null,
    first_name varchar(255) null,
    email      varchar(255) not null,
    password   varchar(255) not null,
    constraint email
        unique (email)
);

-- table todos
create table todos
(
    id           int unsigned auto_increment
        primary key,
    created_at   timestamp            null,
    updated_at   timestamp            null,
    title        varchar(255)         null,
    completed_at timestamp            null,
    is_done      tinyint(1) default 0 null,
    user_id      int unsigned         not null,
    note         text                 null
);

create index idx_todos_completed_at
    on todos (completed_at);

create index idx_todos_user_id
    on todos (user_id);

INSERT INTO todoapp.users (id, created_at, updated_at, last_name, first_name, email, password) VALUES (1, '2020-01-09 15:17:44', '2020-01-09 16:20:05', 'Duong', 'Tai', 'user1@gmail.com', '$2a$10$WCwkEHgoxR/U0hOgVDwZ5Ov1rRJuPmshHC5P9XABg43Vn5T1ChK42');
INSERT INTO todoapp.users (id, created_at, updated_at, last_name, first_name, email, password) VALUES (2, '2020-01-09 16:40:10', '2020-01-09 16:41:54', '', '', 'user2@gmail.com', '$2a$10$lcXu7chdJVsMhb71bHeFUOE/5KdWBVOZT0jxuhPs7Gs6Dx7IK37Ye');

INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (1, '2020-01-09 15:42:34', '2020-01-09 16:03:41', 'Updated title', '2020-01-09 16:03:41', 1, 1, 'Updated note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (2, '2020-01-09 15:42:39', '2020-01-09 15:42:39', 'Second', null, 0, 1, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (3, '2020-01-09 16:54:44', '2020-01-09 16:54:44', 'Second', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (4, '2020-01-09 17:38:37', '2020-01-09 17:38:37', 'Third', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (5, '2020-01-09 17:38:44', '2020-01-09 17:38:44', 'Hello', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (6, '2020-01-09 17:38:47', '2020-01-09 17:38:47', 'Hiii', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (7, '2020-01-09 17:38:53', '2020-01-09 17:38:53', 'Do something', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (8, '2020-01-09 17:39:01', '2020-01-09 17:39:01', 'Complete go-todo app', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (9, '2020-01-09 17:39:07', '2020-01-09 17:39:07', 'Test app', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (10, '2020-01-09 17:39:15', '2020-01-09 17:39:15', 'Test 2', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (11, '2020-01-09 17:39:21', '2020-01-09 17:39:21', 'Test 3', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (12, '2020-01-09 17:39:24', '2020-01-09 17:39:24', 'Test 4', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (13, '2020-01-09 17:39:27', '2020-01-09 17:39:27', 'Test 5', null, 0, 2, 'Note');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (14, '2020-01-09 17:39:30', '2020-01-09 17:39:30', 'Test 5', null, 0, 2, 'Note 23232');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (15, '2020-01-09 17:39:34', '2020-01-09 17:39:34', 'Test 5', null, 0, 2, 'fdsafjldsf');
INSERT INTO todoapp.todos (id, created_at, updated_at, title, completed_at, is_done, user_id, note) VALUES (16, '2020-01-09 17:39:37', '2020-01-09 17:39:37', 'Test 5', null, 0, 2, 'fdsfdsafdsf');