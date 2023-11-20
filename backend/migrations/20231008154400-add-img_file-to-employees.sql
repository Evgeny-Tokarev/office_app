-- +migrate Up
ALTER TABLE employees
    ADD COLUMN img_file text;

-- +migrate Down
ALTER TABLE employees
DROP COLUMN img_file;