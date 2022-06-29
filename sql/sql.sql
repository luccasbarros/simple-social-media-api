CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users (
  id int auto_increment primary key,
  name varchar(50) NOT NULL,
  nick varchar(50) NOT NULL unique,
  email varchar(50) NOT NULL unique,
  password varchar(100) NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;

CREATE TABLE followers (
  user_id int not null, 
  FOREIGN KEY (user_id) REFERENCES users (id)
  ON DELETE CASCADE,

  follower_id int not null,
  FOREIGN KEY (follower_id) REFERENCES users (id)
  ON DELETE CASCADE,

  primary key (user_id, follower_id)
  ) ENGINE=InnoDB;