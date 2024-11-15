DROP DATABASE go_sessions;

CREATE USER IF NOT EXISTS 'session'@'localhost' IDENTIFIED BY 'password';
CREATE DATABASE IF NOT EXISTS go_sessions;
GRANT ALL ON `go_sessions`.* TO 'session'@'localhost';

use go_sessions;

CREATE TABLE user(
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	permission INT NOT NULL DEFAULT 0,
	CONSTRAINT PRIMARY KEY(id)
);

-- 1 -> customer service rep, 2 -> administrator
INSERT INTO user (username, password, permission) VALUES ("john", "password", 1), ("jane", "password", 2);
