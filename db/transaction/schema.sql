CREATE TABLE user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) DEFAULT 'kaito2'
) Engine = InnoDB DEFAULT CHARSET = utf8mb4;