CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    photo_id INT,
    message VARCHAR(300) NOT NULL,
    created_at INT NOT NULL DEFAULT 0,
    updated_at INT NOT NULL DEFAULT 0,
    deleted_at INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (photo_id) REFERENCES photos(id)
);
