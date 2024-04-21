CREATE TABLE budgets (
    id             SERIAL PRIMARY KEY,
    uuid           UUID DEFAULT gen_random_uuid(),
    title          VARCHAR(128) NOT NULL,
    description    VARCHAR(255),
    created        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  	userId         INT NOT NULL,
  	FOREIGN KEY    (userID) REFERENCES users(id)
);
