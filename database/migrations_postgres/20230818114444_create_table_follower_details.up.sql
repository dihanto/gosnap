CREATE TABLE
    follower_details (
        username VARCHAR(100),
        follower_id uuid,
        Foreign Key (username) REFERENCES users(username),
        Foreign Key (follower_id) REFERENCES users(id)
    )