-- +goose Up
CREATE TABLE IF NOT EXISTS budgets (
    id             SERIAL PRIMARY KEY,
    uuid           UUID DEFAULT gen_random_uuid(),
    title          VARCHAR(128) NOT NULL,
    description    VARCHAR(255),
    created        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    userId         INT NOT NULL,
    FOREIGN KEY    (userID) REFERENCES users(id)
);

INSERT INTO budgets
    (id, title, description, userid)
VALUES
    (1, 'ARCHITAX', 'Elit dolor cillum elit aute do aliquip esse. Nostrud id eu ut labore eiusmod non.', 1),
    (2, 'FUTURITY', 'In et consequat sit tempor in sit laboris. Qui amet eiusmod minim labore.', 1),
    (3, 'CYTREX', '', 1);

-- +goose Down
DROP TABLE IF EXISTS budgets;
