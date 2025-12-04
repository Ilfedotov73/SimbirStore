CREATE TABLE IF NOT EXISTS products_reviews (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE SET NULL,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    review TEXT,
    rating FLOAT
);