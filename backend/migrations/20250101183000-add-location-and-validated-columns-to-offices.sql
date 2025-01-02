-- +migrate Up

alter table offices
    add column location json null,
    add column is_valid_address boolean default false not null;

-- +migrate Down

alter table offices
    drop column location,
    drop column is_valid_address;