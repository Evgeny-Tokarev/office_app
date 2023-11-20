-- +migrate Up
ALTER TABLE offices
    ADD COLUMN img_file TEXT;

-- +migrate Down
ALTER TABLE offices
DROP COLUMN img_file;





