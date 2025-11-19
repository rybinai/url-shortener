CREATE TABLE IF NOT EXISTS urlshortener (
    short_code VARCHAR(6) PRIMARY KEY,
    original_url TEXT NOT NULL
);