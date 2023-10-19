 CREATE TABLE followers (
    id SERIAL PRIMARY KEY,
    follower_count INT NOT NULL DEFAULT 0,
    user_id UUID,
    username VARCHAR(100),
    Foreign Key (user_id) REFERENCES users(id)
);

CREATE TABLE follower_details (
    follow_id int,
    follower_name VARCHAR(100),
    Foreign Key (follow_id) REFERENCES followers(id)
)