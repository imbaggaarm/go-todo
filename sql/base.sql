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