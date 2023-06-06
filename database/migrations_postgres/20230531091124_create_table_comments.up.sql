CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    user_id INT,
    photo_id INT,
    message VARCHAR(300) NOT NULL,
    created_at INT,
    updated_at INT,
    deleted_at INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (photo_id) REFERENCES photos(id)
);
