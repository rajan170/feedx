-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL ;

-- +goose Down
ALTER TABLE users DROP COLUMN api_key