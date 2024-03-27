
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
    Title       VARCHAR(128) NOT NULL,
    Description VARCHAR(255),
    Uuid        VARCHAR(36) DEFAULT (uuid()),
    Created     DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
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
