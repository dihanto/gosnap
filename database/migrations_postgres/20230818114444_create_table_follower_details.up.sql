CREATE TABLE
    follower_details (
        user_id uuid,
        follower_id uuid,
        Foreign Key (user_id) REFERENCES users(id),
        Foreign Key (follower_id) REFERENCES users(id)
    )