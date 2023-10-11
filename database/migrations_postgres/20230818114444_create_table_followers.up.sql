 CREATE TABLE followers (
    id SERIAL PRIMARY KEY,
    follower_count INT NOT NULL DEFAULT 0,
    username VARCHAR(100),
    Foreign Key (username) REFERENCES users(username)
);

CREATE TABLE follower_details (
    follow_id int,
    follower_name VARCHAR(100),
    Foreign Key (follow_id) REFERENCES followers(id),
    Foreign Key (follower_name) REFERENCES users(username)
)