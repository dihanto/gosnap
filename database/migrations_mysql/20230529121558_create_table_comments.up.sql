CREATE TABLE comments (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id int,
    Foreign Key (user_id) REFERENCES users(id),
    photo_id int,
    Foreign Key (photo_id) REFERENCES photos(id),
    message VARCHAR(300) NOT NULL,
    created_at int(8),
    updated_at int(8),
    deleted_at int(8)
) ENGINE=InnoDB;
