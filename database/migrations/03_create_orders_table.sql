CREATE TYPE order_category AS ENUM ('IN', 'OUT', 'DELIVERY');

CREATE TABLE orders (
    id SERIAL NOT NULL PRIMARY KEY,
    subtotal DECIMAL(10, 2) NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    gst DECIMAL(10, 2) NOT NULL DEFAULT 0,
    pst DECIMAL(10, 2) NOT NULL DEFAULT 0,
    discount DECIMAL(10, 2) NOT NULL DEFAULT 0,
    timestamp TIMESTAMPTZ DEFAULT NOW(),
    category order_category NOT NULL,
    void BOOLEAN NOT NULL DEFAULT FALSE,
    paid BOOLEAN NOT NULL DEFAULT FALSE,
    customizations jsonb DEFAULT NULL,
    customer_id INT REFERENCES customers(id)
);

INSERT INTO orders
(subtotal,    total,   gst,   pst,    discount,   category,       customizations,                                                           customer_id ) VALUES
(7.50,        7.87,    0.37,  0.00,   0.00,       'IN',       NULL,                                                                     1           ),
(6.00,        6.30,    0.30,  0.00,   0.00,       'OUT',      '[{"name_eng": "add bb sauce", "name_oth": "gaseejup", "price": 1.00}]',  2           );