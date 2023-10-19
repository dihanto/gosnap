CREATE TABLE photos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    caption VARCHAR(100) NOT NULL,
    photo_url VARCHAR(100) NOT NULL,
    user_id UUID,
    created_at INT NOT NULL DEFAULT 0,
    updated_at INT NOT NULL DEFAULT 0,
    deleted_at INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
