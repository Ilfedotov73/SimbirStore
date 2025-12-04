CREATE TABLE IF NOT EXISTS notices (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    entity_id INT, -- В контексте ServiceCustomer считаем это ID покупателя, либо ссылкой на объект
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);