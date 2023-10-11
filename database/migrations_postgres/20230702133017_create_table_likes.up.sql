CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    like_count INT NOT NULL DEFAULT 0,
    photo_id INT,
    FOREIGN KEY (photo_id) REFERENCES photos(id)
);

CREATE TABLE like_details (
    like_id int,
    user_id UUID,
    liked_at INT NOT NULL DEFAULT 0,
    Foreign Key (like_id) REFERENCES likes(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
