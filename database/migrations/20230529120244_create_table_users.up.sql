CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(200) NOT NULL,
    age INT,
    created_at int(8),
    updated_at int(8),
    deleted_at int(8)
) ENGINE=InnoDB;
