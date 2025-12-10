CREATE TABLE IF NOT EXISTS vendors_products (
    id SERIAL PRIMARY KEY,
    vendor_id INT REFERENCES users(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE
);