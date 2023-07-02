CREATE TABLE like_details(
    photo_id int,
    user_id uuid,
    Foreign Key (photo_id) REFERENCES photos(id),
    Foreign Key (user_id) REFERENCES users(id)
)