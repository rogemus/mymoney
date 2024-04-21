CREATE TABLE transactions (
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
