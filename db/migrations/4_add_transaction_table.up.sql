CREATE TABLE transaction (
    ID   						INT AUTO_INCREMENT NOT NULL,
    Uuid 						VARCHAR(36) DEFAULT (uuid()),
    Description     VARCHAR(255), 
    Amount          FLOAT NOT NULL,
    Created         DATETIME DEFAULT CURRENT_TIMESTAMP,
  	BudgetID        INT NOT NULL,
  	UserID 					INT NOT NULL,
    PRIMARY KEY     (ID),
    FOREIGN KEY     (BudgetID) REFERENCES budget(ID) 
);

