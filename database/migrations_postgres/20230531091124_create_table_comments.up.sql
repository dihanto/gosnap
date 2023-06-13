CREATE TABLE comments (
    id UUID PRIMARY KEY,
    user_id UUID,
    photo_id UUID,
    message VARCHAR(300) NOT NULL,
    created_at INT NOT NULL DEFAULT 0,
    updated_at INT NOT NULL DEFAULT 0,
    deleted_at INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (photo_id) REFERENCES photos(id)
);
