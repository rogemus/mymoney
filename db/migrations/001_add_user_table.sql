-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    uuid        UUID DEFAULT gen_random_uuid(),
    email       VARCHAR(255) NOT NULL UNIQUE,
    username    VARCHAR(255) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    created     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users
    (id, username, email, password)
VALUES
    (1, 'test', 'test@test.com', '$2a$14$K4TVJJ43ddGnXZ/65J4EyOwGtgTx6UWjDyxmyhPqXWI0qhg0kGXty');

-- +goose Down
DROP TABLE IF EXISTS users;
