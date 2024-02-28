-- +migrate Up
-- Create indexes on the name and email columns for search optimization
create index idx_users_name on users (name);
create index idx_users_email on users (email);

-- +migrate Down
-- Drop indexes on the name and email columns
drop index if exists idx_users_name;
drop index if exists idx_users_email;