CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    photo_url TEXT,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    characteristics TEXT,
    product_rating FLOAT DEFAULT 0
);