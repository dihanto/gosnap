CREATE TABLE social_medias (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    social_media_url VARCHAR(100) NOT NULL,
    user_id INT,
    Foreign Key (user_id) REFERENCES users(id),
    created_at int(8),
    updated_at int(8),
    deleted_at int(8)
) ENGINE=InnoDB;
