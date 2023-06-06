CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(200) NOT NULL,
    age INT,
    created_at INT,
    updated_at INT,
    deleted_at INT
);
