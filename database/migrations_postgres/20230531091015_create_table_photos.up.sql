CREATE TABLE photos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    caption VARCHAR(100) NOT NULL,
    photo_url VARCHAR(50) NOT NULL,
    user_id INT,
    created_at INT,
    updated_at INT,
    deleted_at INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
