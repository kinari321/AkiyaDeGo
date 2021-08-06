-- CREATE DATABASE akiyadego;
USE akiyadego;

DROP TABLE IF EXISTS users, posts, sessions;

CREATE TABLE IF NOT EXISTS users
(
  id INT NOT NULL AUTO_INCREMENT,
  uuid TEXT NOT NULL,
  name TEXT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at DATETIME NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS posts
(
  id INT NOT NULL AUTO_INCREMENT,
  imagepath TEXT NULL,
  title TEXT NOT NULL,
  category TEXT NOT NULL,
  prefecture TEXT NOT NULL,
  description TEXT NULL,
  user_id INT NOT NULL,
  created_at DATETIME NULL,
  PRIMARY KEY (id)
);

-- テストデータ
INSERT INTO posts
  (imagepath, title,	category,	prefecture,	description, user_id,	created_at)
  VALUES
  ("/var/www/image/akiya.jpeg", "空き家投稿", "空き家", "北海道", "空き家の詳細説明", 1, "2021-05-22 07:58:32")
;

CREATE TABLE IF NOT EXISTS sessions(
  id INT NOT NULL AUTO_INCREMENT,
  uuid TEXT NOT NULL,
  email TEXT NOT NULL,
  user_id INT NOT NULL,
  created_at DATETIME NULL,
  PRIMARY KEY (id)
);
