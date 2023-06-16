CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(200) NOT NULL,
    age INT,
    created_at INT,
    updated_at INT,
    deleted_at INT,
    CONSTRAINT id_uniq UNIQUE (id),
    CONSTRAINT username_uniq UNIQUE(username),
    CONSTRAINT email_uniq UNIQUE(email),
    CONSTRAINT age_check CHECK (age > 8)
);
