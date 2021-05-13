CREATE DATABASE IF NOT EXISTS ejobs CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use ejobs;

CREATE TABLE ads (
     id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
     title VARCHAR(100) NOT NULL,
     content TEXT NOT NULL,
     created DATETIME NOT NULL,
     expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_ads_created ON ads(created);

CREATE TABLE users (
       id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
       name VARCHAR(255) NOT NULL,
       email VARCHAR(255) NOT NULL,
       hashed_password CHAR(60) NOT NULL,
       created DATETIME NOT NULL,
       active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO ads (title, content, created, expires) VALUES (
                                                              'An old silent pond',
                                                              'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo',
                                                              UTC_TIMESTAMP(),
                                                              DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
                                                          );

INSERT INTO ads (title, content, created, expires) VALUES (
                                                              'Over the wintry forest',
                                                              'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
                                                              UTC_TIMESTAMP(),
                                                              DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
                                                          );

INSERT INTO ads (title, content, created, expires) VALUES (
                                                              'First autumn morning',
                                                              'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
                                                              UTC_TIMESTAMP(),
                                                              DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
                                                          );