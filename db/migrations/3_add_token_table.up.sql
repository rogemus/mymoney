CREATE TABLE token (
    ID              INT AUTO_INCREMENT NOT NULL,
    Uuid            VARCHAR(36) DEFAULT (uuid()),
    Token           VARCHAR(255) NOT NULL,
  	UserEmail				VARCHAR(255) NOT NULL,
    Created         DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY     (ID)
);
