CREATE TABLE customizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    is_user_defined BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO customizations (name, price, is_user_defined) VALUES
('No customization', 0, TRUE),
('Add bb sauce', 2.00, TRUE),
('Less oil', 0.00, TRUE);
