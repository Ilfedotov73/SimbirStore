CREATE TABLE IF NOT EXISTS offers (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    buyer_id INT REFERENCES users(id) ON DELETE CASCADE,
    vendor_id INT REFERENCES users(id) ON DELETE CASCADE,
    offer_price FLOAT NOT NULL,
    message TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);