-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS links (
    id   UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    short_suffix TEXT,
    url TEXT,
    clicks INTEGER NOT NULL DEFAULT 0,
    expiration_date TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd
