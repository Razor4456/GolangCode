CREATE TABLE IF NOT EXISTS revok_token(
    token TEXT PRIMARY KEY,
    user_id INT REFERENCES users(id),
    expires_at TIMESTAMPTZ
);