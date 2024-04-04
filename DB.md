
# MySQL

## Env Vars
```
$ export DBUSER=root
$ export DBPASS=

```

## Login to MySQL
```
mysql -u root
```
## Create DB
```
mysql> create database tracker;
```

## Select DB
```
mysql>use tracker
```

## Budget table

### Create table 
```sql
DROP TABLE IF EXISTS budget;
CREATE TABLE budget (
    ID          INT AUTO_INCREMENT NOT NULL,
    Uuid        VARCHAR(36) DEFAULT (uuid()),
    Title       VARCHAR(128) NOT NULL,
    Description VARCHAR(255),
    Created     DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (ID)
);

```

### Populate table
```sql
INSERT INTO budget
    (Title, Description)
VALUES
    ('ARCHITAX', 'Elit dolor cillum elit aute do aliquip esse. Nostrud id eu ut labore eiusmod non.'),
    ('FUTURITY', 'In et consequat sit tempor in sit laboris. Qui amet eiusmod minim labore.'),
    ('CYTREX', '');
```

## Transaction table

### Create table
```sql
DROP TABLE IF EXISTS transaction;
CREATE TABLE transaction (
    ID              INT AUTO_INCREMENT NOT NULL,
    Uuid            VARCHAR(36) DEFAULT (uuid()),
    Description     VARCHAR(255), 
    Amount          FLOAT NOT NULL,
    Created         DATETIME DEFAULT CURRENT_TIMESTAMP,
  	BudgetID        INT NOT NULL,
    PRIMARY KEY     (ID),
    FOREIGN KEY     (BudgetID) REFERENCES budget(ID) 
);

```

### Insert Data
```sql
INSERT INTO transaction
    (Description, Amount, BudgetID)
VALUES
    ("In reprehenderit et elit aliqua officia aute sint dolor minim.", -23.67, 1),
    ("Aute anim occaecat excepteur.", 17.94, 2),
    ("", -3.41, 1);
```
