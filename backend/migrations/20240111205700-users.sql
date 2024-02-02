-- +migrate Up
create table users
(
    id       bigserial not null
            primary key,
    name varchar default '' not null,
    email varchar default '' unique not null,
    role varchar default 'user' not null CHECK (role IN ('user', 'admin', 'moderator')),
    hashed_password varchar not null,
    password_changed_at timestamp DEFAULT('0001-01-01 00:00:00Z'::timestamp) not null,
    created_at timestamp default now() not null,
    img_file text
);

-- +migrate Down
drop table if exists users;
