CREATE TABLE IF NOT EXISTS users
(
    id UUID,
    full_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    password_updated_at TIMESTAMP NOT NULL,
    email_confirmed_at TIMESTAMP DEFAULT NULL
);
