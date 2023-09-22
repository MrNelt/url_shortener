-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links (
    short_suffix TEXT PRIMARY KEY,
    link TEXT,
    secret_key TEXT UNIQUE,
    clicks INTEGER NOT NULL DEFAULT 0
);
CREATE TABLE IF NOT EXISTS vip_links (
    short_suffix TEXT PRIMARY KEY,
    link TEXT,
    secret_key TEXT UNIQUE,
    clicks INTEGER NOT NULL DEFAULT 0,
    expiration_date TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
DROP TABLE vip_links;
-- +goose StatementEnd
