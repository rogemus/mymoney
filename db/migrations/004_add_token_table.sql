-- +goose Up
CREATE TABLE IF NOT EXISTS tokens (
    id           SERIAL PRIMARY KEY,
    uuid         UUID DEFAULT gen_random_uuid(),
    token        VARCHAR(255) NOT NULL,
    userEmail    VARCHAR(255) NOT NULL,
    created      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at   TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS tokens;
