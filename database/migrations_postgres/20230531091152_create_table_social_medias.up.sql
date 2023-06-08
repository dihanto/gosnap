CREATE TABLE social_medias (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    social_media_url VARCHAR(100) NOT NULL,
    user_id INT,
    created_at INT NOT NULL DEFAULT 0,
    updated_at INT NOT NULL DEFAULT 0,
    deleted_at INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
