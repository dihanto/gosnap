CREATE TABLE photos (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    caption VARCHAR(100) NOT NULL,
    photo_url VARCHAR(50) NOT NULL,
    user_id INT,
    Foreign Key (user_id) REFERENCES users(id),
    created_at int(8),
    updated_at int(8),
    deleted_at int(8)
) ENGINE=InnoDB;
