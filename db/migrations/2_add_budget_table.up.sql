DROP TABLE IF EXISTS budget;
CREATE TABLE budget (
    ID            INT AUTO_INCREMENT NOT NULL,
    Uuid          VARCHAR(36) DEFAULT (uuid()),
    Title         VARCHAR(128) NOT NULL,
    Description   VARCHAR(255),
    Created       DATETIME DEFAULT CURRENT_TIMESTAMP,
  	UserId        INT NOT NULL,
    PRIMARY KEY   (ID),
  	FOREIGN KEY   (UserID) REFERENCES user(ID)
);
