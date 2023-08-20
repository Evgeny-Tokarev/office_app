
-- +migrate Up
create table employees
(
    id       bigserial not null
        constraint employees_pkey
            primary key,
    name text default '' not null,
    age int not null,
    office_id bigint constraint employees_office_id_fkey references offices (id),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

-- +migrate Down
drop table employees;
