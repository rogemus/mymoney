-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    id             SERIAL PRIMARY KEY,
    uuid           UUID DEFAULT gen_random_uuid(),
    description    VARCHAR(255),
    amount         FLOAT NOT NULL,
    created        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    budgetid       INT NOT NULL,
    userid         INT NOT NULL,
    FOREIGN KEY    (budgetid) REFERENCES budgets(id),
    FOREIGN KEY    (userid) REFERENCES users(id)
);

INSERT INTO transactions
    (id, amount, description, userid, budgetid)
VALUES
    (1, 12.34, 'Elit dolor cillum elit aute do aliquip esse. Nostrud id eu ut labore eiusmod non.', 1, 1),
    (2, -6.9, 'In et consequat sit tempor in sit laboris. Qui amet eiusmod minim labore.', 1, 1),
    (3, 41.6, '', 1, 1);

-- +goose Down
DROP TABLE IF EXISTS transactions;
