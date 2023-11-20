-- +migrate Up
create table offices
(
    id       bigserial not null
        constraint offices_pkey
            primary key,
    name text default '' not null,
    address text default '' not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null
);

insert into offices  (name, address)
values ('MEDIASOFT', 'ул. Карла Маркса, 13А, корп. 3, Ульяновск, Ульяновская обл., 432011');

-- +migrate Down
drop table offices;
