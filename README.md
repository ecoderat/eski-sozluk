# eski-sözlük

A go-based sözlük clone inspired by the first version

![homepage](https://github.com/ecoderat/eski-sozluk/blob/master/ui/static/img/eski-sozluk.png?raw=true)

## Installation
### Installation MySQL
#### Macos
```bash
brew install mysql
```
#### Linux
```bash
sudo apt install mysql-server
```
### Configuration MySQL
```bash
sudo mysql
mysql>
```
to create a new eskisozluk database using UTF8 encoding:
```SQL
-- Create a new UTF-8 `eskisozluk` database.
CREATE DATABASE eskisozluk CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Switch to using the `eskisozluk` database.
USE eskisozluk;
```
to create a new sozluk table to hold the text entry for our application:
```SQL
-- Create a `sozluk` table.
CREATE TABLE sozluk (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
title VARCHAR(50) NOT NULL,
content TEXT NOT NULL,
user VARCHAR(40) NOT NULL,
created DATETIME NOT NULL,
);
-- Add an index on the created column.
CREATE INDEX idx_sozluk_created ON sozluk(created);
```
creating a new user:
```SQL
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON eskisozluk.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```

### Installation eski-sözlük
clonning of repo:
```bash
git clone https://github.com/ecoderat/eski-sozluk.git
```
open main.go and swap 'pass' value with a password of your own chosen:
```go
// main.go
// ...
dsn := flag.String("dsn", "web:pass@/eskisozluk?parseTime=true", "MySQL data source name")
// ...
```


## Usage
in the project directory
```bash
go run ./cmd
```
## Features
#### basic features
- [x] Latest entry fetching properly to homepage
- [x] Topics list on homepage is fetching properly
- [x] Topics list on topicpage is fetching properly
- [x] Adding an entry to a topic with the form (no authentication yet)
- [x] Creating a new topic 
- [x] Adding middleware

#### features related to user
- [x] Creating a users table on the database
- [x] User authentication
- [x] Creating a users model
- [ ] Log in/Sign up page
- [ ] Running a HTTPS server

#### entry search features
...
